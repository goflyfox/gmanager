import request from "@/utils/request";

const BASE_URL = "/admin/generator";

const GeneratorAPI = {
  /** 代码生成表列表 */
  getPage(queryParams?: GeneratorTableQuery) {
    return request<any, PageResult<GeneratorTableVO[]>>({
      url: `${BASE_URL}/table/list`,
      method: "post",
      params: queryParams,
    });
  },

  /** 数据库表列表（导入用） */
  getDbTablePage(queryParams?: GeneratorDbTableQuery) {
    return request<any, PageResult<GeneratorTableVO[]>>({
      url: `${BASE_URL}/db/table/list`,
      method: "post",
      params: queryParams,
    });
  },

  /** 导入表 */
  importTables(names: string[]) {
    return request({
      url: `${BASE_URL}/table/import`,
      method: "post",
      data: { names },
    });
  },

  /** 获取表配置详情 */
  getTableDetail(id: number) {
    return request<any, GeneratorTableDetail>({
      url: `${BASE_URL}/table/get/${id}`,
      method: "get",
    });
  },

  /** 保存表配置 */
  saveTable(data: GeneratorTableForm) {
    return request({
      url: `${BASE_URL}/table/save`,
      method: "post",
      data: data,
    });
  },

  /** 删除表配置 */
  deleteByIds(ids: string) {
    return request({
      url: `${BASE_URL}/table/delete/${ids}`,
      method: "post",
    });
  },

  /** 预览代码 */
  preview(id: number) {
    return request<any, GeneratorPreviewVO>({
      url: `${BASE_URL}/preview/${id}`,
      method: "get",
    });
  },

  /** 下载 ZIP */
  download(id: number) {
    return request({
      url: `${BASE_URL}/download/${id}`,
      method: "get",
      responseType: "blob",
    });
  },

  /** 生成代码（写入磁盘） */
  genCode(id: number, path?: string) {
    return request({
      url: `${BASE_URL}/gen/${id}`,
      method: "post",
      data: { id, path },
    });
  },
};

export default GeneratorAPI;

/** 代码生成表查询 */
export interface GeneratorTableQuery extends PageQuery {
  keywords?: string;
}

/** 数据库表查询 */
export interface GeneratorDbTableQuery extends PageQuery {
  keywords?: string;
}

/** 代码生成表对象 */
export interface GeneratorTableVO {
  id?: number;
  tableName?: string;
  tableComment?: string;
  className?: string;
  packageName?: string;
  moduleName?: string;
  businessName?: string;
  functionName?: string;
  functionAuthor?: string;
  tplCategory?: string;
  genType?: string;
  genPath?: string;
  options?: string;
  createAt?: Date;
}

/** 代码生成表详情 */
export interface GeneratorTableDetail {
  info?: GeneratorTableVO;
  columns?: GeneratorColumnVO[];
}

/** 代码生成字段对象 */
export interface GeneratorColumnVO {
  id?: number;
  tableId?: number;
  columnName?: string;
  columnComment?: string;
  columnType?: string;
  goType?: string;
  goField?: string;
  isPk?: string;
  isIncrement?: string;
  isRequired?: string;
  isInsert?: string;
  isEdit?: string;
  isList?: string;
  isQuery?: string;
  queryType?: string;
  htmlType?: string;
  dictType?: string;
  sort?: number;
}

/** 代码生成表表单 */
export interface GeneratorTableForm {
  id?: number;
  tableName?: string;
  tableComment?: string;
  className?: string;
  packageName?: string;
  moduleName?: string;
  businessName?: string;
  functionName?: string;
  functionAuthor?: string;
  tplCategory?: string;
  genType?: string;
  genPath?: string;
  options?: string;
  columns?: GeneratorColumnVO[];
}

/** 预览代码 */
export interface GeneratorPreviewVO {
  data?: Record<string, string>;
}
