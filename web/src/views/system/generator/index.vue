<!-- 代码生成 -->
<template>
  <div class="app-container">
    <!-- 搜索区域 -->
    <div class="search-container">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="关键字" prop="keywords">
          <el-input
            v-model="queryParams.keywords"
            placeholder="表名称/表描述"
            clearable
            @keyup.enter="handleQuery"
          />
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
            v-hasPerm="['admin:generator:save']"
            type="success"
            icon="download"
            @click="handleOpenImportDialog()"
          >
            导入表
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
      >
        <el-table-column type="index" label="序号" width="60" align="center" />
        <el-table-column label="表名称" prop="tableName" min-width="150" />
        <el-table-column label="表描述" prop="tableComment" min-width="150" />
        <el-table-column label="实体类名" prop="className" width="150" />
        <el-table-column label="模块名" prop="moduleName" width="100" align="center" />
        <el-table-column label="业务名" prop="businessName" width="100" align="center" />
        <el-table-column label="创建时间" prop="createAt" width="180" align="center" />
        <el-table-column label="操作" fixed="right" width="280" align="center">
          <template #default="scope">
            <el-button
              v-hasPerm="'admin:generator:query'"
              type="primary"
              icon="View"
              link
              size="small"
              @click="handlePreview(scope.row.id)"
            >
              预览
            </el-button>
            <el-button
              v-hasPerm="'admin:generator:save'"
              type="primary"
              icon="edit"
              link
              size="small"
              @click="handleEdit(scope.row.id)"
            >
              编辑
            </el-button>
            <el-button
              v-hasPerm="'admin:generator:save'"
              type="success"
              icon="Download"
              link
              size="small"
              @click="handleDownload(scope.row.id)"
            >
              下载
            </el-button>
            <el-button
              v-hasPerm="'admin:generator:save'"
              type="warning"
              icon="Promotion"
              link
              size="small"
              @click="handleGenCode(scope.row)"
            >
              生成
            </el-button>
            <el-button
              v-hasPerm="'admin:generator:delete'"
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

    <!-- 导入表弹窗 -->
    <el-dialog v-model="importDialog.visible" title="导入表" width="800px" append-to-body>
      <div class="search-container mb-2">
        <el-form :inline="true">
          <el-form-item label="表名称">
            <el-input v-model="importQuery.keywords" placeholder="请输入表名称" clearable />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="fetchDbTables">搜索</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table
        v-loading="importDialog.loading"
        :data="importDialog.tableData"
        border
        @selection-change="handleImportSelectionChange"
      >
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="表名称" prop="tableName" />
        <el-table-column label="表描述" prop="tableComment" />
      </el-table>
      <pagination
        v-if="importDialog.total > 0"
        v-model:total="importDialog.total"
        v-model:page="importQuery.pageNum"
        v-model:limit="importQuery.pageSize"
        @pagination="fetchDbTables"
      />
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleImportSubmit">确 定</el-button>
          <el-button @click="importDialog.visible = false">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import GeneratorAPI, {
  GeneratorTableQuery,
  GeneratorTableVO,
  GeneratorDbTableQuery,
} from "@/api/system/generator.api";

defineOptions({
  name: "Generator",
  inheritAttrs: false,
});

const router = useRouter();
const queryFormRef = ref();

const queryParams = reactive<GeneratorTableQuery>({
  pageNum: 1,
  pageSize: 10,
});

const pageData = ref<GeneratorTableVO[]>();
const total = ref(0);
const loading = ref(false);

// 导入弹窗
const importDialog = reactive({
  visible: false,
  loading: false,
  tableData: [] as GeneratorTableVO[],
  total: 0,
  selectedNames: [] as string[],
});

const importQuery = reactive<GeneratorDbTableQuery>({
  pageNum: 1,
  pageSize: 10,
});

async function fetchData() {
  loading.value = true;
  try {
    const data = await GeneratorAPI.getPage(queryParams);
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
  fetchData();
}

function handleOpenImportDialog() {
  importDialog.visible = true;
  importQuery.pageNum = 1;
  importQuery.keywords = "";
  fetchDbTables();
}

async function fetchDbTables() {
  importDialog.loading = true;
  try {
    const data = await GeneratorAPI.getDbTablePage(importQuery);
    importDialog.tableData = data.list;
    importDialog.total = data.total;
  } finally {
    importDialog.loading = false;
  }
}

function handleImportSelectionChange(selection: GeneratorTableVO[]) {
  importDialog.selectedNames = selection.map((item) => item.tableName!).filter(Boolean);
}

function handleImportSubmit() {
  if (importDialog.selectedNames.length === 0) {
    ElMessage.warning("请至少选择一张表");
    return;
  }
  GeneratorAPI.importTables(importDialog.selectedNames).then(() => {
    ElMessage.success("导入成功");
    importDialog.visible = false;
    handleQuery();
  });
}

function handleEdit(id?: number) {
  router.push({
    path: "/system/generator/edit",
    query: { id },
  });
}

function handlePreview(id?: number) {
  router.push({
    path: "/system/generator/preview",
    query: { id },
  });
}

function handleDownload(id?: number) {
  if (!id) return;
  GeneratorAPI.download(id).then((response: any) => {
    const blob = new Blob([response.data], { type: "application/zip" });
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.href = url;
    link.download = "generator.zip";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  });
}

function handleGenCode(row: GeneratorTableVO) {
  if (!row.id) return;
  let path = row.genPath || "";
  ElMessageBox.prompt("请输入生成路径（留空使用默认配置）", "生成代码", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    inputValue: path,
  }).then(({ value }) => {
    GeneratorAPI.genCode(row.id!, value).then(() => {
      ElMessage.success("生成成功");
    });
  });
}

function handleDelete(id?: number) {
  if (!id) return;
  ElMessageBox.confirm("确认删除该表配置?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    loading.value = true;
    GeneratorAPI.deleteByIds(String(id))
      .then(() => {
        ElMessage.success("删除成功");
        handleQuery();
      })
      .finally(() => (loading.value = false));
  });
}

onMounted(() => {
  handleQuery();
});
</script>
