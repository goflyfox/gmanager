# gmanager 代码生成管理功能实施计划

> 文档版本：v1.0  
> 创建日期：2026-05-07  
> 关联需求：引入 RuoYi 风格的在线代码生成管理  
> 状态：待评审 / 待实施

---

## 一、背景与调研结论

### 1.1 项目现状
- **技术栈**：GoFrame v2 后端 + Vue3/Vite/Element Plus 前端 + MySQL。
- **现有表**：仅 8 张核心系统表（`sys_config`、`sys_dept`、`sys_log`、`sys_menu`、`sys_role`、`sys_role_menu`、`sys_user`、`sys_user_role`），**无代码生成相关表**。
- **已有代码生成能力**：仅依赖 GoFrame CLI (`gf gen dao`) 在开发时生成 DAO/Entity/DO，**无运行时/在线代码生成能力**。
- **前端 CURD 组件**：`web/src/components/CURD/` 下存在一套通用配置化组件，但存量页面（`system/user`、`system/role` 等）**并未使用**，均采用手写 Element Plus 方式。
- **前端残留**：`web/src/enums/codegen/` 下存在 `form.enum.ts` 和 `query.enum.ts`，但无实际页面支撑。

### 1.2 RuoYi 参考结论
- RuoYi 的代码生成器核心依赖两张配置表（`gen_table`、`gen_table_column`）+ Velocity 模板引擎。
- 支持三种模板：`crud`（单表）、`tree`（树表）、`sub`（主子表）。
- 前端 Vue3 版本生成手写式 Element Plus 页面，与 gmanager 存量风格一致。
- 关键特性：从 `information_schema` 自动导入表结构、字段属性可编辑、代码预览、ZIP 下载。

---

## 二、核心设计决策

| 决策项 | 方案 | 理由 |
|--------|------|------|
| **生成目标** | 支持 **ZIP 下载** + **直接写入源码目录** | 默认 ZIP，高级配置可指定路径直接写入 |
| **前端风格** | 手写 Element Plus（对齐 `system/user/index.vue`） | 与存量风格一致，不引入 CURD 组件的学习成本 |
| **模板范围（一期）** | **仅单表 CRUD** | 覆盖 80% 业务场景；树表、主子表放到二期 |
| **与 `gf gen dao` 的关系** | **复用已有 DAO**，运行时不重新生成 entity/do/dao | 避免与 CLI 能力重复，保持架构简洁 |
| **模板引擎** | Go 标准库 `text/template` | 足够轻量，无需引入 Velocity 等外部引擎 |
| **权限前缀** | `admin:generator:*` | 与现有权限体系一致 |

---

## 三、数据库设计

### 3.1 sys_gen_table — 代码生成业务表

```sql
CREATE TABLE `sys_gen_table` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) DEFAULT '' COMMENT '实体类名称（首字母大写）',
  `package_name` varchar(100) DEFAULT '' COMMENT '生成包路径',
  `module_name` varchar(30) DEFAULT '' COMMENT '生成模块名（如 system）',
  `business_name` varchar(30) DEFAULT '' COMMENT '生成业务名（如 post）',
  `function_name` varchar(50) DEFAULT '' COMMENT '生成功能名（如 岗位管理）',
  `function_author` varchar(50) DEFAULT '' COMMENT '生成作者',
  `tpl_category` varchar(10) DEFAULT 'crud' COMMENT '模板类型（crud/tree/sub，一期仅 crud）',
  `gen_type` char(1) DEFAULT '0' COMMENT '生成方式（0=ZIP压缩包 1=自定义路径）',
  `gen_path` varchar(200) DEFAULT '' COMMENT '自定义生成路径',
  `options` varchar(1000) DEFAULT '' COMMENT '其它生成选项（JSON）',
  `create_by` bigint(20) DEFAULT '0' COMMENT '创建人',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` bigint(20) DEFAULT '0' COMMENT '更新人',
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代码生成业务表';
```

### 3.2 sys_gen_table_column — 代码生成业务字段表

```sql
CREATE TABLE `sys_gen_table_column` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) NOT NULL COMMENT '归属表编号',
  `column_name` varchar(200) DEFAULT '' COMMENT '列名称',
  `column_comment` varchar(500) DEFAULT '' COMMENT '列描述',
  `column_type` varchar(100) DEFAULT '' COMMENT '列类型（如 varchar(64)）',
  `go_type` varchar(50) DEFAULT '' COMMENT 'Go 类型（string/int64/time.Time 等）',
  `go_field` varchar(100) DEFAULT '' COMMENT 'Go 字段名（驼峰）',
  `is_pk` char(1) DEFAULT '0' COMMENT '是否主键（1=是）',
  `is_increment` char(1) DEFAULT '0' COMMENT '是否自增（1=是）',
  `is_required` char(1) DEFAULT '0' COMMENT '是否必填（1=是）',
  `is_insert` char(1) DEFAULT '1' COMMENT '是否为插入字段（1=是）',
  `is_edit` char(1) DEFAULT '1' COMMENT '是否编辑字段（1=是）',
  `is_list` char(1) DEFAULT '1' COMMENT '是否列表字段（1=是）',
  `is_query` char(1) DEFAULT '0' COMMENT '是否查询字段（1=是）',
  `query_type` varchar(20) DEFAULT 'EQ' COMMENT '查询方式（EQ/NE/GT/LT/LIKE/BETWEEN）',
  `html_type` varchar(20) DEFAULT 'input' COMMENT '显示类型（input/textarea/select/radio/checkbox/datetime/switch）',
  `dict_type` varchar(200) DEFAULT '' COMMENT '字典类型（绑定 sys_config 的 config_key）',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `create_by` bigint(20) DEFAULT '0',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_by` bigint(20) DEFAULT '0',
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_table_id` (`table_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代码生成业务字段表';
```

---

## 四、后端实现方案

### 4.1 模块位置

在现有 `admin` 模块之外新增 `generator` 模块用于管理，但**生成目标**落在 `admin` 模块内：

```
server/internal/
  generator/
    controller/        # 代码生成管理接口（/admin/generator/*）
    logic/             # 导入表、预览、下载、写入逻辑
    model/entity/      # sys_gen_table, sys_gen_table_column
    model/do/
    model/input/
    dao/
  admin/
    api/v1/            # 生成的 api 文件落在这里
    controller/        # 生成的 controller 落在这里
    logic/             # 生成的 logic 落在这里
```

### 4.2 模板清单

模板存放于 `server/resource/template/generator/`，构建时随二进制嵌入。

| 模板文件 | 生成目标路径 |
|----------|-------------|
| `go/api.go.tpl` | `server/api/admin/v1/{businessName}.go` |
| `go/controller.go.tpl` | `server/internal/admin/controller/{businessName}.go` |
| `go/logic.go.tpl` | `server/internal/admin/logic/{businessName}.go` |
| `vue/index.vue.tpl` | `web/src/views/{moduleName}/{businessName}/index.vue` |
| `vue/api.ts.tpl` | `web/src/api/{moduleName}/{businessName}.api.ts` |

> **说明**：`model/entity`、`model/do`、`dao` 由 `gf gen dao` 生成，不在运行时生成器范围内。

### 4.3 API 设计

```go
// 代码生成表管理
type GeneratorTableListReq struct {
    g.Meta   `path:"/generator/table/list" method:"post" perms:"admin:generator:query"`
    Keywords string
    input.PageReq
}

type GeneratorTableImportReq struct {
    g.Meta `path:"/generator/table/import" method:"post" perms:"admin:generator:save"`
    Names  []string `json:"names" v:"required#至少选择一张表"`
}

type GeneratorTableSaveReq struct {
    g.Meta `path:"/generator/table/save" method:"post" perms:"admin:generator:save"`
    // 包含 GenTable + []GenTableColumn
}

type GeneratorTableDeleteReq struct {
    g.Meta `path:"/generator/table/delete/:ids" method:"post" perms:"admin:generator:delete"`
    Ids string
}

// 代码生成
type GeneratorPreviewReq struct {
    g.Meta `path:"/generator/preview/:id" method:"get" perms:"admin:generator:query"`
    Id int64
}

type GeneratorDownloadReq struct {
    g.Meta `path:"/generator/download/:id" method:"get" perms:"admin:generator:query"`
    Id int64
}

type GeneratorGenCodeReq struct {
    g.Meta `path:"/generator/gen/:id" method:"post" perms:"admin:generator:save"`
    Id   int64  `json:"id"`
    Path string `json:"path"` // 仅 gen_type=1 时生效
}
```

### 4.4 核心逻辑说明

1. **导入表**
   - 查询 `information_schema.tables` / `columns`，根据表名批量导入。
   - 自动映射 MySQL 类型 → Go 类型（如 `varchar`→`string`、`bigint`→`int64`、`datetime`→`*gtime.Time`）。
   - 自动推断 `html_type`：字段名含 `_status` → `radio`，含 `_type` → `select`，`datetime` 类型 → `datetime`，`tinyint(1)` → `switch`。
   - 默认规则：主键不插入/不编辑，审计字段（`create_at`、`update_at`、`create_by` 等）不插入/不编辑，不在列表中显示。

2. **预览**
   - 读取 `sys_gen_table` + `sys_gen_table_column` 数据渲染模板，返回 `map[文件名]代码内容`。

3. **下载 ZIP**
   - 渲染所有模板，按目录结构打包 ZIP，通过 `Content-Disposition: attachment` 下载。

4. **直接写入**
   - 根据 `gen_path` 将文件写入磁盘。
   - **安全校验**：写入路径必须落在项目根目录下，禁止 `../` 等路径遍历；若目标文件已存在则给出警告或覆盖确认。

---

## 五、前端实现方案

### 5.1 新增页面

| 页面 | 说明 |
|------|------|
| `web/src/views/system/generator/index.vue` | 代码生成表列表（导入、删除、编辑入口） |
| `web/src/views/system/generator/edit.vue` | **核心页面**：编辑表配置 + 字段配置表格 |
| `web/src/views/system/generator/preview.vue` | 代码预览（文件树 + 代码高亮） |

### 5.2 编辑页交互设计

**基本信息区**
- 表描述、实体类名、包路径、模块名、业务名、功能名、作者、生成方式（ZIP/自定义路径）、生成路径。

**字段配置表格区**
- 以 `el-table` 展示所有字段，可编辑：
  - 列描述
  - Go 类型（`string`/`int`/`int64`/`float64`/`time.Time`）
  - Go 字段名（自动驼峰，可手动改）
  - 插入 / 编辑 / 列表 / 查询（`el-checkbox`）
  - 查询方式（`EQ`/`NE`/`GT`/`LT`/`LIKE`/`BETWEEN`）
  - 显示类型（`input`/`textarea`/`select`/`radio`/`checkbox`/`datetime`/`switch`）
  - 字典类型（下拉，从 `sys_config` 拉取所有 `config_key`）

### 5.3 预览页交互设计

- 左侧文件树（`api.go`、`controller.go`、`logic.go`、`index.vue`、`api.ts`）。
- 右侧代码展示（使用 `codemirror` 组件或 `<pre>` 展示，只读）。

---

## 六、标准开发流程（用户侧）

```
1. 在数据库创建业务表（如 sys_post）
2. 进入 server/ 执行 gf gen dao，生成 entity/do/dao
3. 登录管理后台 → 系统工具 → 代码生成
4. 点击【导入表】，选择 sys_post，导入
5. 点击【编辑】，检查/调整字段属性（描述、类型、是否查询、显示类型等）
6. 点击【预览】，确认生成代码符合预期
7. 点击【生成代码】，选择 ZIP 下载或直接写入
8. 若 ZIP 下载：解压到对应目录；若直接写入：重启前后端即可验证
```

---

## 七、实施步骤（甘特图式分解）

### Phase 1：基础数据层（预估 1 天）
- [ ] 在 `resources/sql/gmanager.sql` 追加 `sys_gen_table`、`sys_gen_table_column` 建表语句
- [ ] 执行 `gf gen dao` 生成 generator 模块的 entity/do/dao
- [ ] 新增 `server/api/admin/v1/generator.go` 定义 API 结构体
- [ ] 在 `sys_menu` 中预置代码生成菜单（系统工具 → 代码生成）

### Phase 2：后端核心逻辑（预估 2~3 天）
- [ ] 实现 `logic/generator.go` — 导入表（information_schema 读取）
- [ ] 实现 `logic/generator.go` — 预览、下载 ZIP、直接写入
- [ ] 开发 `go/api.go.tpl` — 生成 Req/Res 结构体（带 g.Meta、权限、校验）
- [ ] 开发 `go/controller.go.tpl` — 薄控制器
- [ ] 开发 `go/logic.go.tpl` — 列表查询、Get、Save、Delete
- [ ] 开发 `vue/index.vue.tpl` — 手写 Element Plus 页面（搜索+表格+Drawer表单）
- [ ] 开发 `vue/api.ts.tpl` — API 封装 + TypeScript 接口
- [ ] 实现 `controller/generator.go` 暴露接口

### Phase 3：前端页面（预估 2 天）
- [ ] `system/generator/index.vue` — 列表、导入弹窗
- [ ] `system/generator/edit.vue` — 基本信息 + 字段配置表格
- [ ] `system/generator/preview.vue` — 文件树 + 代码预览
- [ ] 配置菜单路由和权限标识

### Phase 4：联调与边界处理（预估 1 天）
- [ ] 端到端流程验证：建表 → gf gen dao → 导入 → 编辑 → 预览 → 下载 → 写入 → 运行
- [ ] 边界处理：下划线转驼峰、BETWEEN 双日期选择器、字典绑定、路径遍历防护、文件覆盖提示
- [ ] 更新 `AGENTS.md` 补充代码生成使用说明

---

## 八、风险与缓解措施

| 风险 | 影响 | 缓解措施 |
|------|------|----------|
| 模板维护成本高 | 中 | 模板独立在 `server/resource/template/generator/`，后续改通用样式只需改模板文件 |
| 与 `gf gen dao` 职责重叠导致困惑 | 高 | 文档明确：运行时不生成 entity/do/dao，开发者仍需手动执行 `gf gen dao`；在 UI 和文档中强提示此步骤 |
| 直接写入覆盖已有代码 | 高 | 写入前检测文件是否存在并给出警告；生产环境建议禁用直接写入（通过配置开关 `generator.allowWrite = false`） |
| MySQL → Go 类型映射不完整 | 中 | 只覆盖常用类型，复杂类型（`json`、`text`、`blob`）默认 `string`，允许在编辑页手动修正 |
| 前端 `Dict` 组件依赖字典配置 | 低 | 生成模板中若 `dict_type` 不为空则使用 `Dict` 组件；若字典不存在则回退为普通 `el-select`，并在控制台打印警告 |
| 生成代码风格不统一 | 低 | 模板严格对标 `system/user/index.vue`、`api/admin/v1/user.go`、`logic/user.go` 的现有风格 |

---

## 九、验收标准

1. **功能验收**：
   - 在管理后台能完整完成：选择数据库表 → 导入 → 配置字段属性 → 预览代码 → 下载 ZIP → 解压后可直接编译运行。
   - 生成的后端代码能通过 `go build`（需配合 `gf gen dao` 生成的 entity/do/dao）。
   - 生成的前端页面能正常渲染列表、搜索、新增、编辑、删除。

2. **代码风格验收**：
   - 生成的 Go 文件风格与现有 `user`、`role` 模块一致（Req/Res 结构体命名、权限标签、校验规则、logic 查询模式）。
   - 生成的 Vue 文件风格与现有 `system/user/index.vue` 一致（搜索区、表格区、Drawer 表单、分页、权限指令）。

3. **安全验收**：
   - 直接写入路径必须限制在项目根目录内，无法通过 `../` 等方式写入系统任意目录。
   - 覆盖已有文件前必须有明确提示或需要二次确认。

---

## 十、附录

### 10.1 参考资源
- RuoYi 代码生成模块（GitHub）：`https://github.com/yangzongzhuan/RuoYi/tree/master/ruoyi-generator`
- RuoYi-Vue 代码生成模块（GitHub）：`https://github.com/yangzongzhuan/RuoYi-Vue/tree/master/ruoyi-generator`
- gmanager 现有典型模块：`server/internal/admin/logic/user.go`、`web/src/views/system/user/index.vue`

### 10.2 变更日志

| 版本 | 日期 | 变更人 | 变更内容 |
|------|------|--------|----------|
| v1.0 | 2026-05-07 | OpenCode | 初始版本，完成方案设计 |
