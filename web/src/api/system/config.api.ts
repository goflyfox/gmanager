import request from "@/utils/request";

const CONFIG_BASE_URL = "/admin/config";

const ConfigAPI = {
  /** 系统配置分页 */
  getPage(queryParams?: ConfigPageQuery) {
    return request<any, PageResult<ConfigPageVO[]>>({
      url: `${CONFIG_BASE_URL}/list`,
      method: "post",
      params: queryParams,
    });
  },

  getOptions() {
    return request<any, OptionType[]>({
      url: `${CONFIG_BASE_URL}/dict/options`,
      method: "post",
    });
  },

  /** 系统配置表单数据 */
  getFormData(id: number) {
    return request<any, ConfigForm>({
      url: `${CONFIG_BASE_URL}/get/${id}`,
      method: "get",
    });
  },

  /** 更新系统配置 */
  save(id: number, data: ConfigForm) {
    return request({
      url: `${CONFIG_BASE_URL}/save/${id}`,
      method: "post",
      data: data,
    });
  },

  /**
   * 删除系统配置
   *
   * @param ids 系统配置ID
   */
  deleteById(ids: string) {
    return request({
      url: `${CONFIG_BASE_URL}/delete/${ids}`,
      method: "post",
    });
  },

  refreshCache() {
    return request({
      url: `${CONFIG_BASE_URL}/refresh`,
      method: "post",
    });
  },
};

export default ConfigAPI;

/** $系统配置分页查询参数 */
export interface ConfigPageQuery extends PageQuery {
  /** 搜索关键字 */
  keywords?: string;
  /** 配置类型 */
  dataType?: number;
  /** 状态 */
  enable?: number;
}

/** 系统配置表单对象 */
export interface ConfigForm {
  /** 主键 */
  id?: number;
  /** 配置名称 */
  name?: string;
  /** 配置键 */
  key?: string;
  /** 配置值 */
  value?: string;
  /** 配置编码 */
  code?: string;
  /** 配置类型 */
  dataType?: number;
  /** 所属字典Id */
  parentId?: number;
  /** 排序 */
  sort?: number;
  /** 描述、备注 */
  remark?: string;
  /** 状态 */
  enable?: number;
}

/** 系统配置分页对象 */
export interface ConfigPageVO {
  /** 主键 */
  id?: number;
  /** 配置名称 */
  name?: string;
  /** 配置键 */
  key?: string;
  /** 配置值 */
  value?: string;
  /** 配置编码 */
  code?: string;
  /** 配置类型 */
  dataType?: number;
  /** 所属字典Id */
  parentId?: number;
  /** 所属字典 */
  parentKey?: string;
  /** 排序 */
  sort?: number;
  /** 描述、备注 */
  remark?: string;
  /** 状态 */
  enable?: number;
}
