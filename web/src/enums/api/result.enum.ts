/**
 * 响应码枚举
 */
export const enum ResultEnum {
  /**
   * 成功
   */
  SUCCESS = 0,
  /**
   * 错误
   */
  ERROR = -1,
  /**
   * 权限不足
   */
  SECURITY_INVALID = 62,
  /**
   * 认证失败
   */
  AUTH_INVALID = 300,
}
