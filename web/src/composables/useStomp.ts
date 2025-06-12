import { Client, IMessage, StompSubscription } from "@stomp/stompjs";
import { Auth } from "@/utils/auth";

export interface UseStompOptions {
  /** WebSocket 地址，不传时使用 VITE_APP_WS_ENDPOINT 环境变量 */
  brokerURL?: string;
  /** 用于鉴权的 token，不传时使用 getAccessToken() 的返回值 */
  token?: string;
  /** 重连延迟，单位毫秒，默认为 8000 */
  reconnectDelay?: number;
  /** 连接超时时间，单位毫秒，默认为 10000 */
  connectionTimeout?: number;
  /** 是否开启指数退避重连策略 */
  useExponentialBackoff?: boolean;
  /** 最大重连次数，默认为 5 */
  maxReconnectAttempts?: number;
  /** 最大重连延迟，单位毫秒，默认为 60000 */
  maxReconnectDelay?: number;
  /** 是否开启调试日志 */
  debug?: boolean;
}

/**
 * STOMP WebSocket连接组合式函数
 * 用于管理WebSocket连接的建立、断开、重连和消息订阅
 */
export function useStomp(options: UseStompOptions = {}) {
  // 默认值：brokerURL 从环境变量中获取，token 从 getAccessToken() 获取
  const defaultBrokerURL = import.meta.env.VITE_APP_WS_ENDPOINT || "";

  const brokerURL = ref(options.brokerURL ?? defaultBrokerURL);
  // 默认配置参数
  const reconnectDelay = options.reconnectDelay ?? 15000; // 默认15秒重连间隔
  const connectionTimeout = options.connectionTimeout ?? 10000;
  const useExponentialBackoff = options.useExponentialBackoff ?? false;
  const maxReconnectAttempts = options.maxReconnectAttempts ?? 3; // 最多重连3次
  const maxReconnectDelay = options.maxReconnectDelay ?? 60000;

  // 连接状态标记
  const isConnected = ref(false);
  // 重连尝试次数
  const reconnectCount = ref(0);
  // 重连计时器
  let reconnectTimer: any = null;
  // 连接超时计时器
  let connectionTimeoutTimer: any = null;
  // 存储所有订阅
  const subscriptions = new Map<string, StompSubscription>();

  // 用于保存 STOMP 客户端的实例
  let client = ref<Client | null>(null);

  /**
   * 初始化 STOMP 客户端
   */
  const initializeClient = () => {
    return;
  };

  /**
   * 处理重连逻辑
   */
  const handleReconnect = () => {
    if (reconnectCount.value >= maxReconnectAttempts) {
      console.error(`已达到最大重连次数(${maxReconnectAttempts})，停止重连`);
      return;
    }

    reconnectCount.value++;
    console.log(`尝试重连(${reconnectCount.value}/${maxReconnectAttempts})...`);

    // 使用指数退避策略增加重连间隔
    const delay = useExponentialBackoff
      ? Math.min(reconnectDelay * Math.pow(2, reconnectCount.value - 1), maxReconnectDelay)
      : reconnectDelay;

    // 清除之前的计时器
    if (reconnectTimer) {
      clearTimeout(reconnectTimer);
    }

    // 设置重连计时器
    reconnectTimer = setTimeout(() => {
      if (!isConnected.value && client.value) {
        client.value.activate();
      }
    }, delay);
  };

  // 监听 brokerURL 的变化，若地址改变则重新初始化
  watch(brokerURL, (newURL, oldURL) => {
    if (newURL !== oldURL) {
      console.log(`brokerURL changed from ${oldURL} to ${newURL}`);
      // 断开当前连接，重新激活客户端
      if (client.value && client.value.connected) {
        client.value.deactivate();
      }
      brokerURL.value = newURL;
      initializeClient(); // 重新初始化客户端
    }
  });

  // 初始化客户端
  initializeClient();

  /**
   * 激活连接（如果已经连接或正在激活则直接返回）
   */
  const connect = () => {
    // 检查是否有配置WebSocket端点
    if (!brokerURL.value) {
      console.error("WebSocket连接失败: 未配置WebSocket端点URL");
      return;
    }

    if (!client.value) {
      initializeClient();
    }

    if (!client.value) {
      console.error("STOMP客户端初始化失败");
      return;
    }

    // 避免重复连接:检查是否已连接或正在连接
    if (client.value.connected) {
      console.log("WebSocket已经连接,跳过重复连接");
      return;
    }

    if (client.value.active) {
      console.log("WebSocket连接正在进行中,跳过重复连接请求");
      return;
    }

    // 设置连接超时
    clearTimeout(connectionTimeoutTimer);
    connectionTimeoutTimer = setTimeout(() => {
      if (!isConnected.value) {
        console.warn("WebSocket连接超时");
        if (useExponentialBackoff) {
          handleReconnect();
        }
      }
    }, connectionTimeout);

    client.value.activate();
  };

  /**
   * 订阅指定主题
   * @param destination 目标主题地址
   * @param callback 接收到消息时的回调函数
   * @returns 返回订阅 id，用于后续取消订阅
   */
  const subscribe = (destination: string, callback: (_message: IMessage) => void): string => {
    if (!client.value || !client.value.connected) {
      console.warn(`尝试订阅 ${destination} 失败: 客户端未连接`);
      return "";
    }

    try {
      const subscription = client.value.subscribe(destination, callback);
      const subscriptionId = subscription.id;
      subscriptions.set(subscriptionId, subscription);
      console.log(`订阅成功: ${destination}, ID: ${subscriptionId}`);
      return subscriptionId;
    } catch (error) {
      console.error(`订阅 ${destination} 失败:`, error);
      return "";
    }
  };

  /**
   * 取消订阅
   * @param subscriptionId 订阅 id
   */
  const unsubscribe = (subscriptionId: string) => {
    const subscription = subscriptions.get(subscriptionId);
    if (subscription) {
      subscription.unsubscribe();
      subscriptions.delete(subscriptionId);
      console.log(`已取消订阅: ${subscriptionId}`);
    }
  };

  /**
   * 断开WebSocket连接
   */
  const disconnect = () => {
    if (client.value && client.value.connected) {
      // 清除所有订阅
      for (const [id, subscription] of subscriptions.entries()) {
        subscription.unsubscribe();
        subscriptions.delete(id);
      }

      // 断开连接
      client.value.deactivate();
      console.log("WebSocket连接已断开");
    }

    // 清除重连计时器
    if (reconnectTimer) {
      clearTimeout(reconnectTimer);
      reconnectTimer = null;
    }

    // 清除连接超时计时器
    if (connectionTimeoutTimer) {
      clearTimeout(connectionTimeoutTimer);
      connectionTimeoutTimer = null;
    }

    isConnected.value = false;
    reconnectCount.value = 0;
  };

  return {
    isConnected,
    connect,
    subscribe,
    unsubscribe,
    disconnect,
  };
}
