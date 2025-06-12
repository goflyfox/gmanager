<template>
  <div class="app-container">
    <!-- 搜索区域 -->
    <div class="search-container">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item prop="keywords" label="关键字">
          <el-input
            v-model="queryParams.keywords"
            placeholder="日志内容"
            clearable
            @keyup.enter="handleQuery"
          />
        </el-form-item>

        <el-form-item label="日志类型" prop="logType">
          <el-select
            v-model="queryParams.logType"
            placeholder="全部"
            clearable
            style="width: 100px"
          >
            <el-option label="全部" :value="0" />
            <el-option label="系统类型" :value="1" />
            <el-option label="操作类型" :value="2" />
          </el-select>
        </el-form-item>

        <el-form-item prop="createAt" label="操作时间">
          <el-date-picker
            v-model="queryParams.createAt"
            :editable="false"
            type="daterange"
            range-separator="~"
            start-placeholder="开始时间"
            end-placeholder="截止时间"
            value-format="YYYY-MM-DD"
            style="width: 300px"
          />
        </el-form-item>

        <el-form-item class="search-buttons">
          <el-button type="primary" icon="search" @click="handleQuery">搜索</el-button>
          <el-button icon="refresh" @click="handleResetQuery">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="hover" class="data-table">
      <el-table
        v-loading="loading"
        :data="pageData"
        highlight-current-row
        border
        class="data-table__content"
      >
        <el-table-column label="操作时间" prop="createAt" width="180" />
        <el-table-column label="操作人" prop="operator" width="120" />
        <el-table-column prop="logType" label="日志类型" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.logType == 1" type="warning">系统类型</el-tag>
            <el-tag v-else-if="scope.row.logType == 2" type="success">操作类型</el-tag>
            <el-tag v-else type="info">默认</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="日志表" prop="operTable" width="100" />
        <el-table-column label="日志内容" prop="operObject" min-width="200" show-overflow-tooltip />
        <el-table-column label="日志备注" prop="operRemark" min-width="200" show-overflow-tooltip />
        <!-- <el-table-column label="操作类型" prop="operType" min-width="200" /> -->
        <el-table-column prop="operType" label="操作类型" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.operType == '登录'" type="success">
              {{ scope.row.operType }}
            </el-tag>
            <el-tag v-else-if="scope.row.operType == '删除'" type="warning">
              {{ scope.row.operType }}
            </el-tag>
            <el-tag v-else type="info">{{ scope.row.operType }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="URL地址" prop="url" width="150" show-overflow-tooltip />
        <el-table-column label="请求方法" prop="method" width="80" />
        <el-table-column label="IP 地址" prop="ip" width="100" />
        <el-table-column label="浏览器UA" prop="userAgent" width="200" show-overflow-tooltip />
        <el-table-column label="执行时间(ms)" prop="executionTime" width="120" />
      </el-table>

      <pagination
        v-if="total > 0"
        v-model:total="total"
        v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize"
        @pagination="fetchData"
      />
    </el-card>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "Log",
  inheritAttrs: false,
});

import LogAPI, { LogPageVO, LogPageQuery } from "@/api/system/log.api";

const queryFormRef = ref();

const loading = ref(false);
const total = ref(0);

const queryParams = reactive<LogPageQuery>({
  pageNum: 1,
  pageSize: 10,
  keywords: "",
  logType: 0,
  createAt: ["", ""],
});

// 日志表格数据
const pageData = ref<LogPageVO[]>();

/** 获取数据 */
function fetchData() {
  loading.value = true;
  LogAPI.getPage(queryParams)
    .then((data) => {
      pageData.value = data.list;
      total.value = data.total;
    })
    .finally(() => {
      loading.value = false;
    });
}

/** 查询（重置页码后获取数据） */
function handleQuery() {
  queryParams.pageNum = 1;
  fetchData();
}

/** 重置查询 */
function handleResetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  queryParams.createAt = undefined;
  fetchData();
}

onMounted(() => {
  handleQuery();
});
</script>
