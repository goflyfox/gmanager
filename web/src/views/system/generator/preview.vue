<!-- 代码预览 -->
<template>
  <div class="app-container">
    <el-page-header @back="handleBack" title="代码预览" />
    <el-row :gutter="20" class="mt-4" style="height: calc(100vh - 200px)">
      <el-col :span="6" style="height: 100%">
        <el-card shadow="hover" style="height: 100%">
          <template #header>
            <span>文件列表</span>
          </template>
          <el-scrollbar>
            <div
              v-for="(content, fileName) in previewData"
              :key="fileName"
              class="file-item"
              :class="{ active: currentFile === fileName }"
              @click="handleSelectFile(fileName)"
            >
              <el-icon class="mr-1"><Document /></el-icon>
              <span class="file-name">{{ fileName }}</span>
            </div>
          </el-scrollbar>
        </el-card>
      </el-col>
      <el-col :span="18" style="height: 100%">
        <el-card shadow="hover" style="height: 100%">
          <template #header>
            <span>{{ currentFile }}</span>
            <el-button
              type="primary"
              link
              size="small"
              class="float-right"
              @click="handleCopy"
            >
              复制
            </el-button>
          </template>
          <el-scrollbar>
            <pre class="code-block"><code>{{ currentContent }}</code></pre>
          </el-scrollbar>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import GeneratorAPI, { GeneratorPreviewVO } from "@/api/system/generator.api";

defineOptions({
  name: "GeneratorPreview",
  inheritAttrs: false,
});

const route = useRoute();
const router = useRouter();

const previewData = ref<Record<string, string>>({});
const currentFile = ref<string>("");
const currentContent = computed(() => {
  return previewData.value[currentFile.value] || "";
});

async function loadPreview() {
  const id = route.query.id ? Number(route.query.id) : undefined;
  if (!id) return;
  const res: GeneratorPreviewVO = await GeneratorAPI.preview(id);
  previewData.value = res.data || {};
  const keys = Object.keys(previewData.value);
  if (keys.length > 0) {
    currentFile.value = keys[0];
  }
}

function handleSelectFile(fileName: string) {
  currentFile.value = fileName;
}

function handleBack() {
  router.push("/system/generator");
}

function handleCopy() {
  if (!currentContent.value) return;
  navigator.clipboard.writeText(currentContent.value).then(() => {
    ElMessage.success("复制成功");
  });
}

onMounted(() => {
  loadPreview();
});
</script>

<style scoped>
.file-item {
  display: flex;
  align-items: center;
  padding: 8px 12px;
  cursor: pointer;
  border-radius: 4px;
  margin-bottom: 4px;
}
.file-item:hover {
  background-color: var(--el-fill-color-light);
}
.file-item.active {
  background-color: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
}
.file-name {
  font-size: 13px;
  word-break: break-all;
}
.code-block {
  margin: 0;
  padding: 16px;
  background-color: #f6f8fa;
  border-radius: 6px;
  font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
  font-size: 13px;
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-all;
}
.float-right {
  float: right;
}
</style>
