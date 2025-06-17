<!-- 系统配置 -->
<template>
  <div class="app-container">
    <!-- 搜索区域 -->
    <div class="search-container">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="关键字" prop="keywords">
          <el-input
            v-model="queryParams.keywords"
            placeholder="请输入配置键\配置名称"
            clearable
            @keyup.enter="handleQuery"
          />
        </el-form-item>

        <el-form-item label="配置类型" prop="dataType">
          <el-select
            v-model="queryParams.dataType"
            placeholder="全部"
            clearable
            style="width: 100px"
          >
            <el-option :value="1" label="配置" />
            <el-option :value="2" label="字典" />
            <el-option :value="3" label="字典数据" />
          </el-select>
        </el-form-item>

        <el-form-item label="部门状态" prop="enable">
          <el-select v-model="queryParams.enable" placeholder="全部" clearable style="width: 100px">
            <el-option :value="1" label="正常" />
            <el-option :value="2" label="禁用" />
          </el-select>
        </el-form-item>

        <el-form-item class="search-buttons">
          <el-button type="primary" icon="search" @click="handleQuery">搜索</el-button>
          <el-button icon="refresh" @click="handleResetQuery">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="hover" class="data-table">
      <div class="data-table__toolbar">
        <div class="data-table__toolbar--actions">
          <el-button
            v-hasPerm="['admin:config:save']"
            type="success"
            icon="plus"
            @click="handleOpenDialog()"
          >
            新增
          </el-button>
          <el-button
            v-hasPerm="['admin:config:refresh']"
            color="#626aef"
            icon="RefreshLeft"
            @click="handleRefreshCache"
          >
            刷新缓存
          </el-button>
        </div>
      </div>

      <el-table
        ref="dataTableRef"
        v-loading="loading"
        :data="pageData"
        highlight-current-row
        class="data-table__content"
        border
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="index" label="序号" width="60" />
        <el-table-column key="name" label="配置名称" prop="name" min-width="100" />
        <el-table-column key="key" label="配置键" prop="key" min-width="100" />
        <el-table-column key="value" label="配置值" prop="value" min-width="100" />
        <el-table-column key="code" label="配置编码" prop="code" min-width="100" />
        <el-table-column prop="dataType" label="配置类型" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.dataType == 1" type="success">配置</el-tag>
            <el-tag v-if="scope.row.dataType == 2" type="info">字典</el-tag>
            <el-tag v-if="scope.row.dataType == 3" type="success">字典数据</el-tag>
          </template>
        </el-table-column>
        <el-table-column key="parentKey" label="所属字典" prop="parentKey" min-width="100" />
        <el-table-column prop="enable" label="状态" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.enable == 1" type="success">正常</el-tag>
            <el-tag v-else type="info">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="排序" width="100" />
        <el-table-column key="remark" label="描述" prop="remark" min-width="100" />
        <el-table-column fixed="right" label="操作" width="220">
          <template #default="scope">
            <el-button
              v-hasPerm="['admin:config:save']"
              type="primary"
              size="small"
              link
              icon="edit"
              @click="handleOpenDialog(scope.row.id)"
            >
              编辑
            </el-button>
            <el-button
              v-hasPerm="['admin:config:delete']"
              type="danger"
              size="small"
              link
              icon="delete"
              @click="handleDelete(scope.row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination
        v-if="total > 0"
        v-model:total="total"
        v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize"
        @pagination="fetchData"
      />
    </el-card>

    <!-- 系统配置表单弹窗 -->
    <el-dialog
      v-model="dialog.visible"
      :title="dialog.title"
      width="500px"
      @close="handleCloseDialog"
    >
      <el-form
        ref="dataFormRef"
        :model="formData"
        :rules="rules"
        label-suffix=":"
        label-width="100px"
      >
        <el-form-item label="配置名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入配置名称" :maxlength="50" />
        </el-form-item>
        <el-form-item label="配置键" prop="key" >
          <el-input v-model="formData.key" placeholder="请输入配置键" :maxlength="50" />
        </el-form-item>
        <el-form-item label="配置类型" prop="dataType">
          <el-radio-group v-model="formData.dataType">
            <el-radio :value="1">配置</el-radio>
            <el-radio :value="2">字典</el-radio>
            <el-radio :value="3">字典数据</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="formData.dataType == 1 || formData.dataType == 3" label="配置值" prop="value">
          <el-input v-model="formData.value" placeholder="请输入配置值" :maxlength="100" />
        </el-form-item>
        <el-form-item  v-if="formData.dataType == 1 || formData.dataType == 3" label="配置编码" prop="code">
          <el-input v-model="formData.code" placeholder="请输入配置编码" :maxlength="100" />
        </el-form-item>
        <el-form-item  v-if="formData.dataType == 3" label="所属字典" prop="parentId">
          <el-select v-model="formData.parentId" placeholder="请选择">
            <el-option
              v-for="item in dictOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="显示排序" prop="sort">
          <el-input-number
            v-model="formData.sort"
            controls-position="right"
            style="width: 100px"
            :min="0"
          />
        </el-form-item>
        <el-form-item label="部门状态">
          <el-radio-group v-model="formData.enable">
            <el-radio :value="1">正常</el-radio>
            <el-radio :value="2">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="描述" prop="remark">
          <el-input
            v-model="formData.remark"
            :rows="4"
            :maxlength="100"
            show-word-limit
            type="textarea"
            placeholder="请输入描述"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确定</el-button>
          <el-button @click="handleCloseDialog">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "Config",
  inheritAttrs: false,
});

import ConfigAPI, { ConfigPageVO, ConfigForm, ConfigPageQuery } from "@/api/system/config.api";
import { ElMessage, ElMessageBox } from "element-plus";
import { useDebounceFn } from "@vueuse/core";

const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const selectIds = ref<number[]>([]);
const total = ref(0);

const queryParams = reactive<ConfigPageQuery>({
  pageNum: 1,
  pageSize: 10,
  keywords: "",
});

// 系统配置表格数据
const pageData = ref<ConfigPageVO[]>([]);

const dialog = reactive({
  title: "",
  visible: false,
});

const formData = reactive<ConfigForm>({
  id: undefined,
  name: "",
  key: "",
  value: "",
  code: "",
  dataType: 1,
  parentId: undefined,
  enable: 1,
  sort: 10,
  remark: "",
});

const rules = reactive({
  name: [{ required: true, message: "请输入系统配置名称", trigger: "blur" }],
  key: [{ required: true, message: "请输入系统配置编码", trigger: "blur" }],
  value: [{ trigger: "blur" , validator: (rule: any, value: any, callback: any) => {
          if (formData.dataType == 1 || formData.dataType == 3) {
            if (!value) {
              callback(new Error("请输入系统配置值"));
              return
            }
          }
          callback();
  }}],
  parentId: [{ trigger: "blur" , validator: (rule: any, value: any, callback: any) => {
          if (formData.dataType == 3) {
            if (!value) {
              callback(new Error("请选择所属字典"));
              return
            }
          }
          callback();
  }}],
});

// 字典下拉数据源
const dictOptions = ref<OptionType[]>();

// 获取数据
function fetchData() {
  loading.value = true;
  ConfigAPI.getPage(queryParams)
    .then((data) => {
      pageData.value = data.list;
      total.value = data.total;
    })
    .finally(() => {
      loading.value = false;
    });
}

// 查询（重置页码后获取数据）
function handleQuery() {
  queryParams.pageNum = 1;
  fetchData();
}

// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  fetchData();
}

// 行复选框选中项变化
function handleSelectionChange(selection: any) {
  selectIds.value = selection.map((item: any) => item.id);
}

// 打开系统配置弹窗
async function handleOpenDialog(id?: number) {
  // 加载字典下拉数据源
  dictOptions.value = await ConfigAPI.getOptions();
    
  dialog.visible = true;
  if (id) {
    dialog.title = "修改系统配置";
    ConfigAPI.getFormData(id).then((data) => {
      Object.assign(formData, data);
    });
  } else {
    dialog.title = "新增系统配置";
    formData.id = undefined;
  }
}

// 刷新缓存(防抖)
const handleRefreshCache = useDebounceFn(() => {
  ConfigAPI.refreshCache().then(() => {
    ElMessage.success("刷新成功");
  });
}, 1000);

// 系统配置表单提交
function handleSubmit() {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      loading.value = true;
      const id = formData.id;

      if (formData.dataType == 1 || formData.dataType == 2) {
        formData.parentId = 0;
      }
      if (formData.dataType == 2) {
        formData.value = "";
        formData.code = "";
      }
      if (id) {
        ConfigAPI.save(id, formData)
          .then(() => {
            ElMessage.success("修改成功");
            handleCloseDialog();
            handleResetQuery();
          })
          .finally(() => (loading.value = false));
      } else {
        ConfigAPI.save(0, formData)
          .then(() => {
            ElMessage.success("新增成功");
            handleCloseDialog();
            handleResetQuery();
          })
          .finally(() => (loading.value = false));
      }
    }
  });
}

// 重置表单
function resetForm() {
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();
  formData.id = undefined;
}

// 关闭系统配置弹窗
function handleCloseDialog() {
  dialog.visible = false;
  resetForm();
}

// 删除系统配置
function handleDelete(id: string) {
  ElMessageBox.confirm("确认删除该项配置?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    loading.value = true;
    ConfigAPI.deleteById(id)
      .then(() => {
        ElMessage.success("删除成功");
        handleResetQuery();
      })
      .finally(() => (loading.value = false));
  });
}

onMounted(() => {
  handleQuery();
});
</script>
