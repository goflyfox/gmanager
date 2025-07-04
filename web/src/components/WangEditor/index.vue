
<template>
  <div style="z-index: 999; border: 1px solid var(--el-border-color)">
    <!-- 工具栏 -->
    <Toolbar
      :editor="editorRef"
      mode="simple"
      :default-config="toolbarConfig"
      style="border-bottom: 1px solid var(--el-border-color)"
    />
    <!-- 编辑器 -->
    <Editor
      v-model="modelValue"
      :style="{ height: height, overflowY: 'hidden' }"
      :default-config="editorConfig"
      mode="simple"
      @on-created="handleCreated"
    />
  </div>
</template>

<script setup lang="ts">
import "@wangeditor-next/editor/dist/css/style.css";
import { Toolbar, Editor } from "@wangeditor-next/editor-for-vue";
import { IToolbarConfig, IEditorConfig } from "@wangeditor-next/editor";

// 文件上传 API
import FileAPI from "@/api/file.api";

// 上传图片回调函数类型
type InsertFnType = (_url: string, _alt: string, _href: string) => void;

defineProps({
  height: {
    type: String,
    default: "500px",
  },
});
// 双向绑定
const modelValue = defineModel("modelValue", {
  type: String,
  required: false,
});

// 编辑器实例，必须用 shallowRef，重要！
const editorRef = shallowRef();

// 工具栏配置
const toolbarConfig = ref<Partial<IToolbarConfig>>({});

// 编辑器配置
const editorConfig = ref<Partial<IEditorConfig>>({
  placeholder: "请输入内容...",
  MENU_CONF: {
    uploadImage: {
      customUpload(file: File, insertFn: InsertFnType) {
        // 上传图片
        FileAPI.uploadFile(file).then((res) => {
          // 插入图片
          insertFn(res.url, res.name, res.url);
        });
      },
    } as any,
  },
});

// 记录 editor 实例，重要！
const handleCreated = (editor: any) => {
  editorRef.value = editor;
};

// 组件销毁时，也及时销毁编辑器，重要！
onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});
</script>
