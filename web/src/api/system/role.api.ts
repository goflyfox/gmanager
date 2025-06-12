import request from "@/utils/request";

const ROLE_BASE_URL = "/admin/role";

const RoleAPI = {
  /** 获取角色分页数据 */
  getPage(queryParams?: RolePageQuery) {
    return request<any, PageResult<RolePageVO[]>>({
      url: `${ROLE_BASE_URL}/list`,
      method: "post",
      params: queryParams,
    });
  },

  /** 获取角色下拉数据源 */
  getOptions() {
    return request<any, OptionType[]>({
      url: `${ROLE_BASE_URL}/options`,
      method: "post",
    });
  },
  /**
   * 获取角色的菜单ID集合
   *
   * @param roleId 角色ID
   * @returns 角色的菜单ID集合
   */
  getRoleMenuIds(roleId: number) {
    return request<any, string[]>({
      url: `${ROLE_BASE_URL}/menuIds/${roleId}`,
      method: "post",
    });
  },

  /**
   * 分配菜单权限
   *
   * @param roleId 角色ID
   * @param menuIds 菜单ID集合
   */
  updateRoleMenus(roleId: string, menuIds: number[]) {
    return request({
      url: `${ROLE_BASE_URL}/addMenus/${roleId}`,
      method: "post",
      data: { menuIds: menuIds },
    });
  },

  /**
   * 获取角色表单数据
   *
   * @param id 角色ID
   * @returns 角色表单数据
   */
  getFormData(id: string) {
    return request<any, RoleForm>({
      url: `${ROLE_BASE_URL}/get/${id}`,
      method: "get",
    });
  },

  /**
   * 保存角色
   *
   * @param id 角色ID
   * @param data 角色表单数据
   */
  save(id: number, data: RoleForm) {
    return request({
      url: `${ROLE_BASE_URL}/save/${id}`,
      method: "post",
      data: data,
    });
  },

  /**
   * 批量删除角色，多个以英文逗号(,)分割
   *
   * @param ids 角色ID字符串，多个以英文逗号(,)分割
   */
  deleteByIds(ids: string) {
    return request({
      url: `${ROLE_BASE_URL}/delete/${ids}`,
      method: "post",
    });
  },
};

export default RoleAPI;

/** 角色分页查询参数 */
export interface RolePageQuery extends PageQuery {
  /** 搜索关键字 */
  keywords?: string;
}

/** 角色分页对象 */
export interface RolePageVO {
  /** 角色ID */
  id?: number;
  /** 角色编码 */
  code?: string;
  /** 角色名称 */
  name?: string;
  /** 排序 */
  sort?: number;
  /** 角色状态 */
  enable?: number;
  /** 创建时间 */
  createAt?: Date;
  /** 修改时间 */
  updateTime?: Date;
}

/** 角色表单对象 */
export interface RoleForm {
  /** 角色ID */
  id?: number;
  /** 角色编码 */
  code?: string;
  /** 数据权限 */
  dataScope?: number;
  /** 角色名称 */
  name?: string;
  /** 排序 */
  sort?: number;
  /** 角色状态(1-正常；2-停用) */
  enable?: number;
}
