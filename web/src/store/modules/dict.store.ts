import { store } from "@/store";
import DictAPI, { type DictItemOption } from "@/api/system/dict.api";

export const useDictStore = defineStore("dict", () => {
  // 字典数据缓存
  const dictCache = useStorage<Record<string, DictItemOption[]>>("dict_cache", {});

  // 请求队列（防止重复请求）
  const requestQueue: Record<string, Promise<void>> = {};

  /**
   * 缓存字典数据
   * @param parentKey 字典Key
   * @param data 字典项列表
   */
  const cacheDictItems = (parentKey: string, data: DictItemOption[]) => {
    dictCache.value[parentKey] = data;
  };

  /**
   * 加载字典数据（如果缓存中没有则请求）
   * @param parentKey 字典Key
   */
  const loadDictItems = async (parentKey: string) => {
    if (dictCache.value[parentKey]) return;
    // 防止重复请求
    if (!requestQueue[parentKey]) {
      requestQueue[parentKey] = DictAPI.getDictItems(parentKey).then((data) => {
        cacheDictItems(parentKey, data);
        Reflect.deleteProperty(requestQueue, parentKey);
      });
    }
    await requestQueue[parentKey];
  };

  /**
   * 获取字典项列表
   * @param parentKey 字典Key
   * @returns 字典项列表
   */
  const getDictItems = (parentKey: string): DictItemOption[] => {
    return dictCache.value[parentKey] || [];
  };

  /**
   * 移除指定字典项
   * @param parentKey 字典Key
   */
  const removeDictItem = (parentKey: string) => {
    if (dictCache.value[parentKey]) {
      Reflect.deleteProperty(dictCache.value, parentKey);
    }
  };

  /**
   * 清空字典缓存
   */
  const clearDictCache = () => {
    dictCache.value = {};
  };

  return {
    loadDictItems,
    getDictItems,
    removeDictItem,
    clearDictCache,
  };
});

export function useDictStoreHook() {
  return useDictStore(store);
}
