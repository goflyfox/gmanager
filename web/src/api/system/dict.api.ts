import request from "@/utils/request";

const DICT_BASE_URL = "/admin/config";

const DictAPI = {
  //---------------------------------------------------
  // 字典相关接口
  //---------------------------------------------------

  /**
   * 获取字典项列表
   */
  getDictItems(parentKey: string) {
    return request<any, DictItemOption[]>({
      url: `${DICT_BASE_URL}/dict/items/${parentKey}`,
      method: "post", 
    });
  },

};

export default DictAPI;

/**
 * 字典项下拉选项
 */
export interface DictItemOption {
  value: number | string;
  label: string;
  tagType?: "" | "success" | "info" | "warning" | "danger";
  [key: string]: any;
}
