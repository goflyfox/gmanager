import request from "@/utils/request";

const BASE_URL = "/admin/{{.businessName}}";

const {{.className}}API = {
  /** 分页列表 */
  getPage(queryParams?: {{.className}}PageQuery) {
    return request<any, PageResult<{{.className}}PageVO[]>>({
      url: `${BASE_URL}/list`,
      method: "post",
      params: queryParams,
    });
  },

  /** 获取详情 */
  getFormData(id: number) {
    return request<any, {{.className}}Form>({
      url: `${BASE_URL}/get/${id}`,
      method: "get",
    });
  },

  /** 保存 */
  save(id: number, data: {{.className}}Form) {
    return request({
      url: `${BASE_URL}/save/${id}`,
      method: "post",
      data: data,
    });
  },

  /** 删除 */
  deleteByIds(ids: string) {
    return request({
      url: `${BASE_URL}/delete/${ids}`,
      method: "post",
    });
  },
};

export default {{.className}}API;

/** 分页查询参数 */
export interface {{.className}}PageQuery extends PageQuery {
  /** 关键字 */
  keywords?: string;
{{- range .queryColumns}}
  /** {{.ColumnComment}} */
  {{.GoField}}?: {{if eq .GoType "int"}}number{{else if eq .GoType "int64"}}number{{else}}string{{end}};
{{- end}}
}

/** 分页对象 */
export interface {{.className}}PageVO {
  /** ID */
  id: number;
{{- range .listColumns}}
  /** {{.ColumnComment}} */
  {{.GoField}}?: {{if eq .GoType "int"}}number{{else if eq .GoType "int64"}}number{{else if eq .GoType "*gtime.Time"}}Date{{else}}string{{end}};
{{- end}}
}

/** 表单类型 */
export interface {{.className}}Form {
  /** ID */
  id?: number;
{{- range .formColumns}}
  /** {{.ColumnComment}} */
  {{.GoField}}?: {{if eq .GoType "int"}}number{{else if eq .GoType "int64"}}number{{else if eq .GoType "*gtime.Time"}}string{{else}}string{{end}};
{{- end}}
}
