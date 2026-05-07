# AGENTS.md

## 项目概览
- 全栈管理后台：**GoFrame v2** 后端 (`server/`) + **Vue 3 / Vite / Element Plus** 前端 (`web/`)。
- 无 CI/CD 工作流，以本地开发为主。
- 文档和代码注释多为中文。

## 仓库结构
- `server/` — Go 后端（模块名 `gmanager`）。入口：`main.go`。
- `web/` — 前端。要求 Node >= 18，包管理器**只能用 pnpm**。
- `resources/sql/gmanager.sql` — MySQL 初始脚本（注意：README 里写的是 `resource/sql/`，实际路径是 `resources/sql/`）。

## 后端 (`server/`)
### 本地运行
```bash
cd server
go mod tidy
go run main.go
```
- 默认端口 `8000`。验证：`GET http://localhost:8000/ping` 返回 `pong`。
- 依赖 MySQL。复制 `manifest/config/config.example.yaml` → `manifest/config/config.yaml` 并修改数据库 `link`。**不要提交 `config.yaml`**（虽然已在 `.gitignore` 中，但工作树里可能存在）。
- Redis 已引入但可选；运行时缓存和 gToken 模式可配置（`1=gcache`、`2=redis`、`3=file`）。

### 代码生成与构建（GoFrame CLI `gf`）
- 安装/更新 CLI：`make cli`（Linux/Mac），Windows 需手动下载 `gf` 二进制文件。
- `make build` 或 `gf build -ew` — 构建二进制文件，将 `server/resource/` 嵌入可执行文件（配置见 `hack/config.yaml`）。
- `make dao` 或 `gf gen dao` — 根据 `hack/config.yaml` 中列出的表重新生成 DAO/DO/Entity。
- `make ctrl` 或 `gf gen ctrl` — 根据 `api/admin/v1` 和 `api/common/v1` 的请求/响应结构体生成控制器。
- `make service` 或 `gf gen service` — 生成 Service 接口。
- API 请求/响应结构体位于 `server/api/{admin,common}/v1/*.go`，控制器位于 `server/internal/{admin,common}/controller/`。

### 后端关键注意事项
- **`/admin` 路由全局开启了 CORS**，在 `internal/cmd/cmd.go` 中，注释明确说明生产环境应注释掉该行。
- `server/resource/` 下的静态资源在构建时嵌入二进制。修改资源后必须重新构建才能生效。
- Go 测试极其稀少（只有一个简单的测试），不要依赖 `go test ./...` 验证业务逻辑。

## 前端 (`web/`)
### 本地运行
```bash
cd web
pnpm install
pnpm run dev
```
- 默认端口 `3000`。
- 开发代理：`/dev-api` → `http://127.0.0.1:8000/`（来自 `.env.development`）。

### 构建与代码检查
```bash
pnpm run type-check   # vue-tsc --noEmit
pnpm run build        # vue-tsc --noEmit & vite build
pnpm run lint         # eslint + prettier + stylelint
```
- **包管理器锁定**：`npm install` 被 `only-allow pnpm` 阻止，始终使用 `pnpm`。
- **构建脚本陷阱**：`build` 命令使用单 `&`（`vue-tsc --noEmit & vite build`）。在 Unix 上后台运行类型检查；在 Windows `cmd` 上无论类型检查是否失败都会继续执行 `vite build`。实践中即使类型检查失败，构建仍可能继续。
- 已配置 Husky + lint-staged；`pnpm run prepare` 安装 husky。
- 提交辅助：`git-cz`（Commitizen + `cz-git`）。

### 前端关键注意事项
- 使用 `unplugin-auto-import` 和 `unplugin-vue-components`（`ElementPlusResolver`）。通常不需要手动导入 Vue API 或 Element Plus 组件。
- 使用 UnoCSS；全局 SCSS 变量从 `@/styles/variables.scss` 注入。
- `vite-plugin-mock-dev-server` 已存在但默认禁用（`VITE_MOCK_DEV_SERVER=false`）。

## Docker
- `resources/docker/Dockerfile` — 多阶段构建，同时构建服务端和 nginx 前端。
- `server/manifest/docker/Dockerfile` — 轻量级纯服务端镜像（要求编译后的二进制文件位于 `./temp/linux_amd64/main`）。
