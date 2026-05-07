<!-- {{.functionName}} -->
<template>
  <div class="app-container">
    <!-- 搜索区域 -->
    <div class="search-container">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
{{- range .queryColumns}}
{{- if eq .HtmlType "datetime"}}
        <el-form-item label="{{.ColumnComment}}">
          <el-date-picker
            v-model="queryParams.{{.GoField}}"
            :editable="false"
            type="daterange"
            range-separator="~"
            start-placeholder="开始时间"
            end-placeholder="截止时间"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "switch")}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-select v-model="queryParams.{{.GoField}}" placeholder="全部" clearable style="width: 100px">
            <el-option label="正常" :value="1" />
            <el-option label="禁用" :value="2" />
          </el-select>
        </el-form-item>
{{- else}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-input
            v-model="queryParams.{{.GoField}}"
            placeholder="请输入{{.ColumnComment}}"
            clearable
            @keyup.enter="handleQuery"
          />
        </el-form-item>
{{- end}}
{{- end}}
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
            v-hasPerm="['admin:{{.businessName}}:save']"
            type="success"
            icon="plus"
            @click="handleOpenDialog()"
          >
            新增
          </el-button>
          <el-button
            v-hasPerm="'admin:{{.businessName}}:delete'"
            type="danger"
            icon="delete"
            :disabled="selectIds.length === 0"
            @click="handleDelete()"
          >
            删除
          </el-button>
        </div>
      </div>

      <el-table
        v-loading="loading"
        :data="pageData"
        border
        stripe
        highlight-current-row
        class="data-table__content"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="50" align="center" />
{{- range .listColumns}}
{{- if eq .HtmlType "switch"}}
        <el-table-column label="{{.ColumnComment}}" prop="{{.GoField}}" width="100" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.{{.GoField}} == 1 ? 'success' : 'info'">
              {{ scope.row.{{.GoField}} == 1 ? "正常" : "禁用" }}
            </el-tag>
          </template>
        </el-table-column>
{{- else if eq .HtmlType "datetime"}}
        <el-table-column label="{{.ColumnComment}}" align="center" prop="{{.GoField}}" width="180" />
{{- else}}
        <el-table-column label="{{.ColumnComment}}" prop="{{.GoField}}" {{- if .Sort}}width="120"{{- end}} align="center" />
{{- end}}
{{- end}}
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <el-button
              v-hasPerm="'admin:{{.businessName}}:save'"
              type="primary"
              icon="edit"
              link
              size="small"
              @click="handleOpenDialog(scope.row.id)"
            >
              编辑
            </el-button>
            <el-button
              v-hasPerm="'admin:{{.businessName}}:delete'"
              type="danger"
              icon="delete"
              link
              size="small"
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

    <!-- 表单 -->
    <el-drawer
      v-model="dialog.visible"
      :title="dialog.title"
      append-to-body
      :size="drawerSize"
      @close="handleCloseDialog"
    >
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
{{- range .formColumns}}
{{- if eq .IsPk "1"}}
{{- else if eq .HtmlType "textarea"}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-input v-model="formData.{{.GoField}}" type="textarea" rows="3" placeholder="请输入{{.ColumnComment}}" />
        </el-form-item>
{{- else if eq .HtmlType "select"}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-select v-model="formData.{{.GoField}}" placeholder="请选择{{.ColumnComment}}">
            <el-option label="请选择" value="" />
          </el-select>
        </el-form-item>
{{- else if eq .HtmlType "radio"}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-radio-group v-model="formData.{{.GoField}}">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="2">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
{{- else if eq .HtmlType "switch"}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-switch
            v-model="formData.{{.GoField}}"
            inline-prompt
            active-text="正常"
            inactive-text="禁用"
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
{{- else if eq .HtmlType "datetime"}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-date-picker
            v-model="formData.{{.GoField}}"
            type="datetime"
            placeholder="选择{{.ColumnComment}}"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
{{- else}}
        <el-form-item label="{{.ColumnComment}}" prop="{{.GoField}}">
          <el-input v-model="formData.{{.GoField}}" placeholder="请输入{{.ColumnComment}}" />
        </el-form-item>
{{- end}}
{{- end}}
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="handleCloseDialog">取 消</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import { useAppStore } from "@/store/modules/app.store";
import { DeviceEnum } from "@/enums/settings/device.enum";

import {{.className}}API, { {{.className}}Form, {{.className}}PageQuery, {{.className}}PageVO } from "@/api/{{.moduleName}}/{{.businessName}}.api";

defineOptions({
  name: "{{.className}}",
  inheritAttrs: false,
});

const appStore = useAppStore();

const queryFormRef = ref();
const formRef = ref();

const queryParams = reactive<{{.className}}PageQuery>({
  pageNum: 1,
  pageSize: 10,
});

const pageData = ref<{{.className}}PageVO[]>();
const total = ref(0);
const loading = ref(false);

const dialog = reactive({
  visible: false,
  title: "新增{{.functionName}}",
});
const drawerSize = computed(() => (appStore.device === DeviceEnum.DESKTOP ? "600px" : "90%"));

const formData = reactive<{{.className}}Form>({
{{- range .formColumns}}
{{- if or (eq .HtmlType "switch") (eq .HtmlType "radio")}}
  {{.GoField}}: 1,
{{- end}}
{{- end}}
});

const rules = reactive({
{{- range .formColumns}}
{{- if eq .IsRequired "1"}}
  {{.GoField}}: [{ required: true, message: "{{.ColumnComment}}不能为空", trigger: "blur" }],
{{- end}}
{{- end}}
});

const selectIds = ref<number[]>([]);

async function fetchData() {
  loading.value = true;
  try {
    const data = await {{.className}}API.getPage(queryParams);
    pageData.value = data.list;
    total.value = data.total;
  } finally {
    loading.value = false;
  }
}

function handleQuery() {
  queryParams.pageNum = 1;
  fetchData();
}

function handleResetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
{{- range .queryColumns}}
{{- if or (eq .HtmlType "datetime") (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "switch")}}
  queryParams.{{.GoField}} = undefined;
{{- end}}
{{- end}}
  fetchData();
}

function handleSelectionChange(selection: any[]) {
  selectIds.value = selection.map((item) => item.id);
}

async function handleOpenDialog(id?: number) {
  dialog.visible = true;
  if (id) {
    dialog.title = "修改{{.functionName}}";
    {{.className}}API.getFormData(id).then((data) => {
      Object.assign(formData, { ...data });
    });
  } else {
    dialog.title = "新增{{.functionName}}";
  }
}

function handleCloseDialog() {
  dialog.visible = false;
  formRef.value.resetFields();
  formRef.value.clearValidate();
  formData.id = undefined;
{{- range .formColumns}}
{{- if or (eq .HtmlType "switch") (eq .HtmlType "radio")}}
  formData.{{.GoField}} = 1;
{{- else}}
  formData.{{.GoField}} = undefined;
{{- end}}
{{- end}}
}

const handleSubmit = useDebounceFn(() => {
  formRef.value.validate((valid: boolean) => {
    if (valid) {
      const id = formData.id;
      loading.value = true;
      {{.className}}API.save(id || 0, formData)
        .then(() => {
          ElMessage.success(id ? "修改成功" : "新增成功");
          handleCloseDialog();
          handleResetQuery();
        })
        .finally(() => (loading.value = false));
    }
  });
}, 1000);

function handleDelete(id?: number) {
  const ids = [id || selectIds.value].join(",");
  if (!ids) {
    ElMessage.warning("请勾选删除项");
    return;
  }

  ElMessageBox.confirm("确认删除?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    function () {
      loading.value = true;
      {{.className}}API.deleteByIds(ids)
        .then(() => {
          ElMessage.success("删除成功");
          handleResetQuery();
        })
        .finally(() => (loading.value = false));
    },
    function () {
      ElMessage.info("已取消删除");
    }
  );
}

onMounted(() => {
  handleQuery();
});
</script>
