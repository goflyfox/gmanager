import request from "@/utils/request";

const LOG_BASE_URL = "/admin/log";

const LogAPI = {
  /**
   * 获取日志分页列表
   *
   * @param queryParams 查询参数
   */
  getPage(queryParams: LogPageQuery) {
    return request<any, PageResult<LogPageVO[]>>({
      url: `${LOG_BASE_URL}/list`,
      method: "post",
      params: queryParams,
    });
  },

  /**
   * 获取访问趋势
   *
   * @param queryParams
   * @returns
   */
  getVisitTrend(queryParams: VisitTrendQuery) {
    return request<any, VisitTrendVO>({
      url: `${LOG_BASE_URL}/visit-trend`,
      method: "post",
      params: queryParams,
    });
  },

};

export default LogAPI;

/**
 * 日志分页查询对象
 */
export interface LogPageQuery extends PageQuery {
  /** 搜索关键字 */
  keywords?: string;
  /** 日志类型 */
  logType?: number;
  /** 操作时间 */
  createAt?: [string, string];
}

/**
 * 系统日志分页VO
 */
export interface LogPageVO {
  /** 主键 */
  id: string;
  /** 类型 */
  LogType: number;
  /** 表 */
  operTable: string;
  /** 对象 */
  operObject: string;
  /** 备注 */
  operRemark: string;
  /** 类型 */
  OperType: string;
  /** 请求路径 */
  uri: string;
  /** 请求方法 */
  method: string;
  /** IP 地址 */
  ip: string;
  /** 浏览器 */
  userAgent: string;
  /** 执行时间(毫秒) */
  executionTime: number;
  /** 操作人 */
  operator: string;
}

/**  访问趋势视图对象 */
export interface VisitTrendVO {
  /** 日期列表 */
  dates: string[];
  /** 浏览量(PV) */
  pvList: number[];
  /** 访客数(UV) */
  uvList: number[];
  /** IP数 */
  ipList: number[];
}

/** 访问趋势查询参数 */
export interface VisitTrendQuery {
  /** 开始日期 */
  startDate: string;
  /** 结束日期 */
  endDate: string;
}
