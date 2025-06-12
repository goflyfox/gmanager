import request from "@/utils/request";
import { Md5 } from "ts-md5";

const AUTH_BASE_URL = "/admin/auth";

const AuthAPI = {
  /** 登录接口*/
  login(data: LoginFormData) {
    const formData = new FormData();
    formData.append("username", data.username);
    formData.append("password", Md5.hashStr(data.password));
    formData.append("codeId", data.codeId);
    formData.append("code", data.code);
    return request<any, LoginResult>({
      // url: `${AUTH_BASE_URL}/login`,
      url: `/admin/login`,
      method: "post",
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  },

  /** 刷新 token 接口*/
  refreshToken(refreshToken: string) {
    return request<any, LoginResult>({
      url: `${AUTH_BASE_URL}/refresh-token`,
      method: "post",
      params: { refreshToken: refreshToken },
      headers: {
        Authorization: "no-auth",
      },
    });
  },

  /** 退出登录接口 */
  logout() {
    return request({
      // url: `${AUTH_BASE_URL}/logout`,
      url: `/admin/user/logout`,
      method: "post",
    });
  },

  /** 获取验证码接口*/
  getCaptcha() {
    return request<any, CaptchaInfo>({
      // url: `${AUTH_BASE_URL}/captcha`,
      url: `/admin/captcha/get`,
      method: "get",
    });
  },
};

export default AuthAPI;

/** 登录表单数据 */
export interface LoginFormData {
  /** 用户名 */
  username: string;
  /** 密码 */
  password: string;
  /** 验证码缓存key */
  code: string;
  /** 验证码 */
  codeId: string;
  /** 记住我 */
  rememberMe: boolean;
}

/** 登录响应 */
export interface LoginResult {
  /** 访问令牌 */
  accessToken: string;
  /** 刷新令牌 */
  refreshToken: string;
  /** 令牌类型 */
  tokenType: string;
  /** 过期时间(秒) */
  expiresIn: number;
}

/** 验证码信息 */
export interface CaptchaInfo {
  /** 验证码缓存Id */
  codeId: string;
  /** 验证码图片Base64字符串 */
  img: string;
}
