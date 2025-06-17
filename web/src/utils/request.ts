import axios, { type InternalAxiosRequestConfig, type AxiosResponse } from "axios";
import qs from "qs";
import { useUserStoreHook } from "@/store/modules/user.store";
import { ResultEnum } from "@/enums/api/result.enum";
import { Auth } from "@/utils/auth";
import router from "@/router";

// 创建 axios 实例
const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
  timeout: 50000,
  headers: { "Content-Type": "application/json;charset=utf-8" },
  paramsSerializer: (params) => qs.stringify(params),
});

// 请求拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const accessToken = Auth.getAccessToken();
    // 如果 Authorization 设置为 no-auth，则不携带 Token
    if (config.headers.Authorization !== "no-auth" && accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    } else {
      delete config.headers.Authorization;
    }
    return config;
  },
  (error) => Promise.reject(error)
);
// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    // 如果响应是二进制流，则直接返回，用于下载文件、Excel 导出等
    if (response.config.responseType === "blob") {
      return response;
    }
    const { code, data, message } = response.data;
    if (code === ResultEnum.SUCCESS) {
      return data;
    }
    ElMessage.error(message || "系统出错");
    return Promise.reject(new Error(message || "Error"));
  },
  async (error) => {
    console.error("request error", error); // for debug
    const { response } = error;
    if (response) {
      const { code, message } = response.data;
      if (code === ResultEnum.AUTH_INVALID) {
        // 刷新 Token 过期，跳转登录页
        await handleLogin();
        return Promise.reject(new Error(message || "Error"));
      } else {
        ElMessage.error(message || "系统出错");
      }
    }
    return Promise.reject(error.message);
  }
);
export default service;

// 处理会话过期
async function handleLogin() {
  ElNotification({
    title: "提示",
    message: "您的会话已过期，请重新登录",
    type: "info",
  });
  await useUserStoreHook().resetAllState();
  router.push("/login");
}
