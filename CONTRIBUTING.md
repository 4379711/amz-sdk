# Contributing to amzsdk

中文版本见 [#贡献指南中文版](#贡献指南中文版)。

Thanks for taking the time to contribute! `amzsdk` is a small but high-leverage library — it sits on the auth + HTTP path of every API call our consumers make. The rules below exist to keep that path predictable.

## Code of conduct

By participating, you agree to abide by our [Code of Conduct](./CODE_OF_CONDUCT.md). Be respectful, focus on the work.

## Project layout reminder

```text
amzsdk/
├── advertising/           # Amazon Ads API client packages
├── selling_partner/       # SP-API client packages
└── pkg/                   # Shared building blocks (auth contract, cache, HTTP, events)
```

- `pkg/` is the **contract layer** — every change here can ripple through every consumer. Stable APIs (`IAuth`, `TokenRefreshEvent`, `AuthRetryEvent`) must not break without a major-version bump.
- `advertising/<pkg>` and `selling_partner/<pkg>` are **mostly generated** from official OpenAPI specs. Manual edits inside these directories are strongly discouraged — see ["Generated packages"](#generated-packages) below.

Live integration testing against your own LwA credentials is expected to happen in a **local-only** `test/` directory; the repo's `.gitignore` excludes `test/` entirely so real `ClientID` / `ClientSecret` / `refresh_token` values cannot accidentally land in source control.

## Reporting bugs / requesting features

Before opening an issue, please:

1. Search [existing issues](../../issues?q=is%3Aissue) to avoid duplicates.
2. Reproduce on the latest `main` (or specify the commit/tag you tested against).
3. Provide:
   - Which sub-package you hit it in (e.g. `advertising/sp_v3`).
   - The smallest possible call sequence that reproduces it — a self-contained Go function you can paste into the issue body is ideal (placeholder credentials are fine).
   - The Amazon-side response (status, body head, headers) if relevant. **Scrub LwA tokens, `Authorization` headers and refresh tokens before sharing.**
   - Your environment — Go version, OS.

For security issues, **do not** open a public issue — see [SECURITY.md](./SECURITY.md) for the private channel.

## Development workflow

### Branching & commits

- Branch off `main`. Use a short, kebab-case name: `fix/sb-bid-recommendation-pagination`, `feat/fba-inbound-2024-step2`, `docs/readme-zh-cn`.
- Keep commits focused. Conventional Commits are appreciated but not enforced (`feat:`, `fix:`, `docs:`, `refactor:`, `chore:`).
- Commit messages in English or 中文 are both accepted.

### Local checks

```bash
# from repo root
go mod tidy
go vet ./...
go build ./...
go test ./pkg/...            # pkg/ has unit tests that don't need LwA
```

If you want to exercise live API calls, write the integration code under a local `test/` directory — it is `.gitignore`-d, so real `ClientID` / `ClientSecret` / `refresh_token` values stay on your machine. **Never** commit credentials — see [SECURITY.md](./SECURITY.md).

### Pull request checklist

- [ ] CI is green (`go build ./...` + `go vet ./...` + `go test ./pkg/...`).
- [ ] No real credentials (`ClientID`, `ClientSecret`, `refresh_token`, `Authorization` headers) committed anywhere — including test fixtures and log snippets.
- [ ] `pkg/` changes preserve the existing exported API. If they don't, the PR description must justify the break and propose a migration.
- [ ] Generated package changes are reproducible (see below).
- [ ] Comments describe **current** behaviour / constraints / reasons — **not** change history, bug-fix narratives or "old vs new" comparisons.

## Generated packages

Most files under `advertising/*` and `selling_partner/*` were generated from Amazon's OpenAPI specs. To keep regenerations sane:

- **Do not** hand-tweak `model_*.go` / `api_*.go` files unless absolutely necessary — your patch will be wiped out the next time we regenerate.
- If a model is wrong, prefer **fixing the input** (the spec or the post-processing script). Document the post-processing step in your PR so future regenerations stay correct.
- If you need to add a brand-new generated package (e.g. tracking a new dated SP-API version), include in the PR:
  - The exact OpenAPI source URL and version you regenerated from.
  - Any post-processing steps applied (`sed` / `gofmt` / manual rename rules).
  - A minimal builder + `Execute()` snippet in the PR description proving the new client works against your account.
- **Never** delete an existing dated package just because a newer one shipped — Amazon keeps the old endpoint live for months, and downstream callers still depend on it.

## Code style

### Go

- `go fmt ./...` and `goimports` are mandatory. CI will reject misformatted files.
- The `pkg/` layer must stay **logger-free**. We expose events (`OnTokenRefresh`, `OnAuthRetry`) and an `IAuth` interface so consumers wire their own logging; do not push logger dependencies down.
- Public functions in `pkg/` need a doc comment explaining the **why**, not the **what** — the docs sit on the contract surface and need to age well.
- Use `pkg.Ptr[T]` for `*T` literals. Avoid hand-rolled inline helpers like `func boolPtr(b bool) *bool`.
- Prefer `errors.Is` / `errors.As` over string matching when wrapping errors.

### Comments

Comments describe **current** behaviour / constraints / reasons only. Do **not** write change history, bug-fix narratives ("fixed the 401 retry bug"), or "old vs new" comparisons inside source comments — that information belongs in commit messages and the changelog.

### Tests

- Unit tests live next to the code (`pkg/cache_test.go`, etc.). They must not require network access and they are what CI runs.
- Live integration tests are expected to live in a **local-only** `test/` directory (excluded by `.gitignore`). They need real LwA credentials, so they are intentionally not part of the published source tree.

## Security

If you discover a security vulnerability, **do not** open a public issue — see [SECURITY.md](./SECURITY.md) for the private reporting channel.

---

# 贡献指南中文版

感谢愿意为 `amzsdk` 贡献代码与文档！本 SDK 体量不大,但坐落在每一次 API 调用的认证 + HTTP 路径上,影响面很广,以下规则是为了让这条路径保持可预测。

## 行为准则

参与即视为接受我们的 [Code of Conduct](./CODE_OF_CONDUCT.md)。请相互尊重,对事不对人。

## 仓库结构提醒

```text
amzsdk/
├── advertising/           # Amazon 广告 API client 包
├── selling_partner/       # SP-API client 包
└── pkg/                   # 公共构件(IAuth 契约、cache、HTTP、events)
```

- `pkg/` 是**契约层** —— 任何变更都会传导到所有使用者。`IAuth`、`TokenRefreshEvent`、`AuthRetryEvent` 这些稳定 API 不可在 minor / patch 版本破坏。
- `advertising/<pkg>` 与 `selling_partner/<pkg>` 大多**由 OpenAPI spec 生成**,强烈不建议手改 —— 详见下面的「生成式包」。

需要拿真实 LwA 凭据跑集成测试时,请把代码放在**本地的** `test/` 目录;仓库 `.gitignore` 已经整体屏蔽 `test/`,真实的 `ClientID` / `ClientSecret` / `refresh_token` 不会意外入库。

## 提 issue / feature request

提 issue 前请:

1. 先搜一下是否已有重复 issue。
2. 在最新 `main` 上复现(若用历史 commit 请明确说明)。
3. 提供:
   - 你撞到问题的子包(例如 `advertising/sp_v3`)。
   - 最小可复现的调用序列 —— 一个能直接贴进 issue 的独立 Go 函数最佳(凭据用占位即可)。
   - 必要时附 Amazon 的响应(状态码、body 前若干字节、header)。**贴出来前先脱敏 LwA token / `Authorization` header / refresh token。**
   - 环境信息(Go 版本、操作系统)。

安全漏洞请勿在 issue 中公开,流程见 [SECURITY.md](./SECURITY.md)。

## 开发流程

### 分支与提交

- 从 `main` 切分支。命名 kebab-case 短描述,例如 `fix/sb-bid-recommendation-pagination`、`feat/fba-inbound-2024-step2`、`docs/readme-zh-cn`。
- 单 commit 聚焦单一职责。Conventional Commits 风格推荐但不强制。
- 中英文提交信息均可。

### 本地校验

```bash
# 仓库根目录
go mod tidy
go vet ./...
go build ./...
go test ./pkg/...            # pkg/ 的单测不依赖 LwA
```

需要跑真实 API 的集成测试时,请把代码写在**本地的** `test/` 目录 —— 它被 `.gitignore` 屏蔽,真实 `ClientID` / `ClientSecret` / `refresh_token` 不会入库。**严禁**提交真实凭据,详见 [SECURITY.md](./SECURITY.md)。

### PR 自查清单

- [ ] CI 绿(`go build ./...` + `go vet ./...` + `go test ./pkg/...`)。
- [ ] **没有**真实凭据落进 diff(包括测试 fixture 与日志片段)。
- [ ] `pkg/` 改动保持现有导出 API。如确需破坏,PR 描述里要说明破坏原因并给迁移方案。
- [ ] 生成式包的改动可重复执行(见下)。
- [ ] 注释只说**当前**行为 / 约束 / 原因,**不**写改动史、bug 修复历史、"旧 vs 新"对比。

## 生成式包

`advertising/*` 与 `selling_partner/*` 下大部分文件由 Amazon OpenAPI spec 生成。为了保证重新生成不打架:

- **不要**手改 `model_*.go` / `api_*.go`,除非别无选择 —— 下一次重新生成会覆盖你的补丁。
- 模型错了优先**修上游输入**(spec 本身或 post-processing 脚本)。PR 里写清楚 post-processing 步骤,保证后续生成仍然正确。
- 新增生成包(例如跟进 SP-API 新日期版本)请在 PR 里附:
  - 你重新生成所用的 OpenAPI 源 URL 与版本号。
  - 应用的 post-processing 步骤(`sed` / `gofmt` / 重命名规则)。
  - PR 描述里贴一个最小的 builder + `Execute()` 调用片段,证明新 client 能跑通你自己的账号。
- **不要**因为 Amazon 出了新日期版本就把老包删掉 —— 老 endpoint 通常会再活几个月,下游可能仍在用。

## 代码风格

### Go

- `go fmt ./...` + `goimports` 必须过,CI 会卡格式。
- `pkg/` 层保持**无 logger 依赖**。事件入口(`OnTokenRefresh`、`OnAuthRetry`)和 `IAuth` 已经覆盖外部接入需要,不要把 logger 依赖下沉到 SDK。
- `pkg/` 中的导出函数注释要说**为什么**而不是**做什么** —— 注释长在契约表面,需要长期可读。
- `*T` 字面量用 `pkg.Ptr[T]`,不要写 `func boolPtr(b bool) *bool` 这种 inline 工具。
- 包错误优先用 `errors.Is` / `errors.As`,不要用字符串匹配。

### 注释

注释**只描述当前行为 / 约束 / 原因**。**不要**写改动史、bug 修复叙事("修复某 401 重试 bug")、"旧 vs 新"对比 —— 这些信息属于 commit message 与 changelog。

### 测试

- 单元测试与代码同目录(`pkg/cache_test.go` 等),不能依赖网络,这是 CI 跑的部分。
- 拿真实 LwA 凭据跑的集成测试请放在**本地的** `test/` 目录(被 `.gitignore` 屏蔽),不会进入公开源码树。

## 安全

如发现安全漏洞,**请勿**在公开 issue 中披露,流程见 [SECURITY.md](./SECURITY.md)。
