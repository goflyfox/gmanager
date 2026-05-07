<!-- 代码生成编辑 -->
<template>
  <div class="app-container">
    <el-page-header @back="handleBack" title="代码生成编辑" />

    <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px" class="mt-4">
      <el-card shadow="hover" class="mb-4">
        <template #header>
          <span>基本信息</span>
        </template>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="表名称">{{ formData.tableName }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="表描述" prop="tableComment">
              <el-input v-model="formData.tableComment" placeholder="请输入表描述" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="实体类名" prop="className">
              <el-input v-model="formData.className" placeholder="请输入实体类名" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包路径" prop="packageName">
              <el-input v-model="formData.packageName" placeholder="如：gmanager" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="模块名" prop="moduleName">
              <el-input v-model="formData.moduleName" placeholder="如：system" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="业务名" prop="businessName">
              <el-input v-model="formData.businessName" placeholder="如：post" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="功能名" prop="functionName">
              <el-input v-model="formData.functionName" placeholder="如：岗位管理" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="作者" prop="functionAuthor">
              <el-input v-model="formData.functionAuthor" placeholder="作者名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="生成方式" prop="genType">
              <el-radio-group v-model="formData.genType">
                <el-radio label="0">ZIP压缩包</el-radio>
                <el-radio label="1">自定义路径</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="生成路径" prop="genPath">
              <el-input v-model="formData.genPath" placeholder="自定义路径，ZIP方式可留空" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-card>

      <el-card shadow="hover">
        <template #header>
          <span>字段配置</span>
        </template>
        <el-table :data="formData.columns" border size="small">
          <el-table-column label="列名称" prop="columnName" width="140" />
          <el-table-column label="列描述" width="160">
            <template #default="scope">
              <el-input v-model="scope.row.columnComment" size="small" />
            </template>
          </el-table-column>
          <el-table-column label="列类型" prop="columnType" width="120" />
          <el-table-column label="Go类型" width="130">
            <template #default="scope">
              <el-select v-model="scope.row.goType" size="small" style="width: 100%">
                <el-option label="string" value="string" />
                <el-option label="int" value="int" />
                <el-option label="int64" value="int64" />
                <el-option label="float64" value="float64" />
                <el-option label="*gtime.Time" value="*gtime.Time" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="Go字段名" width="140">
            <template #default="scope">
              <el-input v-model="scope.row.goField" size="small" />
            </template>
          </el-table-column>
          <el-table-column label="插入" width="60" align="center">
            <template #default="scope">
              <el-checkbox v-model="scope.row.isInsert" true-label="1" false-label="0" />
            </template>
          </el-table-column>
          <el-table-column label="编辑" width="60" align="center">
            <template #default="scope">
              <el-checkbox v-model="scope.row.isEdit" true-label="1" false-label="0" />
            </template>
          </el-table-column>
          <el-table-column label="列表" width="60" align="center">
            <template #default="scope">
              <el-checkbox v-model="scope.row.isList" true-label="1" false-label="0" />
            </template>
          </el-table-column>
          <el-table-column label="查询" width="60" align="center">
            <template #default="scope">
              <el-checkbox v-model="scope.row.isQuery" true-label="1" false-label="0" />
            </template>
          </el-table-column>
          <el-table-column label="查询方式" width="120">
            <template #default="scope">
              <el-select v-model="scope.row.queryType" size="small" style="width: 100%">
                <el-option label="=" value="EQ" />
                <el-option label="!=" value="NE" />
                <el-option label=">" value="GT" />
                <el-option label="<" value="LT" />
                <el-option label="LIKE" value="LIKE" />
                <el-option label="BETWEEN" value="BETWEEN" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="显示类型" width="120">
            <template #default="scope">
              <el-select v-model="scope.row.htmlType" size="small" style="width: 100%">
                <el-option label="输入框" value="input" />
                <el-option label="文本域" value="textarea" />
                <el-option label="下拉框" value="select" />
                <el-option label="单选框" value="radio" />
                <el-option label="复选框" value="checkbox" />
                <el-option label="日期时间" value="datetime" />
                <el-option label="开关" value="switch" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="字典类型" width="150">
            <template #default="scope">
              <el-input v-model="scope.row.dictType" size="small" placeholder="字典code" />
            </template>
          </el-table-column>
          <el-table-column label="必填" width="60" align="center">
            <template #default="scope">
              <el-checkbox v-model="scope.row.isRequired" true-label="1" false-label="0" />
            </template>
          </el-table-column>
        </el-table>
      </el-card>

      <div class="mt-4 flex justify-center">
        <el-button type="primary" @click="handleSubmit">保 存</el-button>
        <el-button @click="handleBack">返 回</el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import GeneratorAPI, {
  GeneratorTableForm,
  GeneratorColumnVO,
  GeneratorTableDetail,
} from "@/api/system/generator.api";

defineOptions({
  name: "GeneratorEdit",
  inheritAttrs: false,
});

const route = useRoute();
const router = useRouter();
const formRef = ref();

const formData = reactive<GeneratorTableForm>({
  tplCategory: "crud",
  genType: "0",
  columns: [],
});

const rules = reactive({
  className: [{ required: true, message: "实体类名不能为空", trigger: "blur" }],
  moduleName: [{ required: true, message: "模块名不能为空", trigger: "blur" }],
  businessName: [{ required: true, message: "业务名不能为空", trigger: "blur" }],
  functionName: [{ required: true, message: "功能名不能为空", trigger: "blur" }],
});

const id = ref<number | undefined>(route.query.id ? Number(route.query.id) : undefined);

async function loadDetail() {
  if (!id.value) return;
  const res: GeneratorTableDetail = await GeneratorAPI.getTableDetail(id.value);
  Object.assign(formData, res.info);
  formData.columns = res.columns || [];
}

function handleBack() {
  router.push("/system/generator");
}

function handleSubmit() {
  formRef.value.validate((valid: boolean) => {
    if (!valid) return;
    GeneratorAPI.saveTable(formData).then(() => {
      ElMessage.success("保存成功");
      router.push("/system/generator");
    });
  });
}

onMounted(() => {
  loadDetail();
});
</script>
