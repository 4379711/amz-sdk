# Security Policy

中文版本见 [#安全政策中文版](#安全政策中文版)。

## Supported Versions

`amzsdk` is currently distributed only via the `main` branch and the tagged releases that originate from it. Security fixes are landed on `main` and back-ported only to the **latest** tagged minor line. If you are running a fork or an older snapshot, please rebase onto the latest `main` before reporting.

## Reporting a Vulnerability

**Please do not file public GitHub issues for security bugs.**

Use one of the following private channels instead:

1. **GitHub Security Advisory (preferred)** — open a [private security advisory](https://docs.github.com/code-security/security-advisories) on this repository. Maintainers will be notified privately.
2. **Email** — send the report to the maintainer email listed on the repository profile. Please use a descriptive subject line such as `[security] amzsdk: <short summary>`.

When reporting, please include:

- A clear description of the issue and the impact you observed.
- The affected sub-package (e.g. `pkg/cache.go`, `advertising/auth`, `selling_partner/auth`).
- Steps to reproduce — ideally a self-contained Go function that demonstrates the issue against a fresh LwA app (placeholder credentials are fine; we don't want your real ones).
- Affected commit / tag / build, plus your environment (OS, Go version).
- Any Amazon-side request / response captures. **Scrub `Authorization` headers, `client_secret`, `refresh_token` and `access_token` values before sharing — last 8 characters at most.**

## Response Process

We aim to:

- **Acknowledge** the report within **3 business days**.
- **Triage and classify** the issue within **7 business days** of acknowledgement.
- **Fix and ship** a patch release as soon as a verified fix is ready, then coordinate disclosure with the reporter.

We will keep you updated on progress and credit you in the release notes / advisory unless you ask to stay anonymous.

## Scope

The following are explicitly **in scope**:

- **Authentication path** — anything in `advertising/auth`, `selling_partner/auth` or `pkg/cache.go` / `pkg/events.go` that could leak credentials, mis-route tokens between sellers, or skip the safety-window expiry.
- **HTTP pipeline** — `pkg/httpx.go` / `pkg/http_trace.go`, especially the 401/403 auto-retry path which could in principle replay a request with the wrong identity.
- **Cache key collisions** — any condition under which two different `(refresh_token, ClientID)` pairs could share a cache slot.
- **Logger hook misuse** — `OnTokenRefresh` / `OnAuthRetry` events leaking full credentials (they MUST only carry suffixes).
- **Build / supply-chain issues** in the published module artifacts.

The following are **out of scope**:

- Vulnerabilities that require the attacker to already control your application process — they can read your refresh tokens directly anyway.
- Issues in the **generated** request/response models that boil down to "Amazon's spec is wrong" — we'll happily review the report but the upstream fix lives with Amazon.
- Reports against downstream applications that embed `amzsdk` — please file those on the consumer's own repository, not here.
- Brute-force or rate-limit attacks that do not exceed Amazon's own documented protections.

## Consumer Responsibilities

`amzsdk` mints and caches access tokens that can move advertising budget and read seller PII. As a consumer you must:

- **Never** commit LwA `ClientID` / `ClientSecret` / `refresh_token` / `access_token` to source control — including in test scaffolding, log dumps or screenshots. The repo `.gitignore` excludes `test/` for exactly this reason; keep credentials there or outside the repo entirely.
- Treat the cache (`pkg.AccessTokenCache`) as **process-local**. Do not serialise it or share it across mutually-untrusted services.
- When subscribing to `OnTokenRefresh` / `OnAuthRetry`, log only the **suffixes** the SDK already exposes — do **not** introspect the raw `*Auth` struct to log the full token.
- Rotate any credential immediately if you suspect it was leaked, and rewrite git history with `git filter-repo` before pushing fixes.

---

# 安全政策中文版

## 支持的版本

`amzsdk` 通过 `main` 分支以及在它之上打的 tag 发布。安全修复合入 `main`,并 back-port 到**最新**的小版本线。如果你跑的是 fork 或更老的快照,请先 rebase 到最新 `main` 再报告。

## 漏洞报告渠道

**请勿**在公开的 GitHub issue 中披露安全漏洞。

请改用下面任意一个私密渠道:

1. **GitHub Security Advisory(推荐)** —— 在本仓库提交 [private security advisory](https://docs.github.com/zh/code-security/security-advisories),维护者会收到私密通知。
2. **邮件** —— 发送到仓库 profile 上公布的维护者邮箱,建议主题写成 `[security] amzsdk: <简要描述>`。

报告中请尽量包含:

- 问题的清晰描述与你观察到的影响。
- 受影响的子包(例如 `pkg/cache.go`、`advertising/auth`、`selling_partner/auth`)。
- 复现步骤 —— 一个能在新建 LwA 应用上复现的独立 Go 函数最佳(凭据用占位即可,我们不要你的真实凭据)。
- 受影响的 commit / tag / 构建产物,以及你的环境(操作系统、Go 版本)。
- 必要的亚马逊侧请求 / 响应抓包。**贴出来之前先脱敏 `Authorization` header、`client_secret`、`refresh_token`、`access_token` —— 最多保留后 8 位。**

## 响应流程

我们承诺:

- **3 个工作日**内确认收到。
- **7 个工作日**内完成初步分级与定性。
- 验证可行的修复后立即发布 patch release,并与报告者协商披露节奏。

我们会持续同步进展,并在发布说明 / advisory 中署名感谢,除非你明确希望保持匿名。

## 适用范围

明确**在范围内**:

- **认证路径** —— `advertising/auth`、`selling_partner/auth`、`pkg/cache.go`、`pkg/events.go` 中任何可能泄漏凭据、跨店错路 token、绕过安全窗口的问题。
- **HTTP 管道** —— `pkg/httpx.go` / `pkg/http_trace.go`,尤其是 401/403 自动重试链路,理论上可能把请求重放给错误的身份。
- **缓存键碰撞** —— 任何使两个不同 `(refresh_token, ClientID)` 命中同一缓存槽的条件。
- **Hook 滥用** —— `OnTokenRefresh` / `OnAuthRetry` 事件泄漏完整凭据(它们**只**应携带 suffix)。
- **构建 / 供应链** 在发布的 module 制品上的问题。

明确**不在范围内**:

- 需要攻击者已经控制你的进程才能利用的问题 —— 那种情况下他直接读你的 refresh token 就行了。
- 生成模型与 Amazon 现网契约不符 —— 我们会受理,但根因得 Amazon 自己修。
- 针对引用 `amzsdk` 的下游应用本身的报告 —— 请到那个仓库的安全政策提交,不要发到这里。
- 未超过 Amazon 自身文档化保护阈值的暴力 / 限流攻击。

## 使用者责任

`amzsdk` 缓存的 access token 可以撬动广告预算与读取卖家 PII。作为使用方,你必须:

- **永远不要**把 LwA `ClientID` / `ClientSecret` / `refresh_token` / `access_token` 提交进版本库 —— 包括测试脚手架、日志 dump、截图。仓库 `.gitignore` 已经整体屏蔽 `test/` 就是为了这件事:凭据要么放进 `test/`、要么放到仓库外。
- 把缓存(`pkg.AccessTokenCache`)视为**进程内**结构,不要序列化,也不要跨互不信任的服务共享。
- 订阅 `OnTokenRefresh` / `OnAuthRetry` 时只记录 SDK 已经暴露的**后缀** —— **不要**反射原始 `*Auth` 把完整 token 写进日志。
- 一旦怀疑凭据泄漏,立刻轮换,并用 `git filter-repo` 清理历史后再推送修复。
