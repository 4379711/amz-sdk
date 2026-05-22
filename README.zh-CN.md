<div align="center">

# amzsdk

**Amazon 广告 API 与 Selling Partner API 的统一 Go SDK —— 一套认证、一条 HTTP 管道、一套可观测能力。**

简体中文 · [English](./README.md)

[![Go](https://img.shields.io/badge/Go-1.26%2B-00ADD8?logo=go)](https://go.dev/)
[![Amazon Ads](https://img.shields.io/badge/Amazon%20Ads-SP%20%C2%B7%20SB%20%C2%B7%20SD-FF9900)](https://advertising.amazon.com/API/docs/zh-cn/)
[![Selling Partner](https://img.shields.io/badge/SP--API-v0%20%C2%B7%20v3-FF9900)](https://developer-docs.amazon.com/sp-api/)

</div>

---

## 目录

- [项目介绍](#项目介绍)
- [核心亮点](#核心亮点)
- [覆盖范围](#覆盖范围)
- [仓库结构](#仓库结构)
- [安装](#安装)
- [授权流程 (LwA OAuth)](#授权流程-lwa-oauth)
- [快速上手](#快速上手)
- [可观测性 Hook](#可观测性-hook)
- [为什么把两套 API 合在一个 module](#为什么把两套-api-合在一个-module)
- [稳定性与版本管理](#稳定性与版本管理)
- [贡献](#贡献)
- [许可协议](#许可协议)
- [免责声明](#免责声明)

---

## 项目介绍

`amzsdk` 是一个 Go SDK，把**两大类亚马逊 HTTP API** 统一收敛在一致的 client 接口下：

- **Amazon Ads** —— Sponsored Products v3、Sponsored Brands v4、Sponsored Display v1、Profiles v2、Reports v3、Portfolios、Product Metadata / Eligibility、AMS。
- **Selling Partner（SP-API）** —— Orders、Inventory、Reports、Catalog、Listings、Feeds、FBA Inbound（含 2024 新版）、Finances v0 + 2024、Pricing、Shipping、Easy Ship、A+ Content、Messaging、Notifications、Vendor 系列，共 40+ 个生成自官方 OpenAPI 的 client 包。

两套 API 在 SDK 内部共享：

- **LwA token 缓存**：5 分钟安全失效窗口、`singleflight` 合并并发刷新、按 `(refresh_token, ClientID)` 维度分键。
- **HTTP 管道**：`pkg.Configuration` + `pkg/httpx.go` + `pkg/http_trace.go`，内置 401/403 兜底重试 —— 命中时清缓存、重新换 `access_token`、原请求重放一次。
- **区域感知的 endpoint 解析**：`pkg.MarketplaceMap` 与各国 LwA URL 覆盖 NA / EU / FE / 印度 / 中东 / 澳洲 / 日本 / 新加坡 等站点。
- **可观测性 Hook**：`pkg.OnTokenRefresh` / `pkg.OnAuthRetry`，让业务程序拿到每一次真实发生的 LwA 刷新和 401 自动重试事件，SDK 自身不接管你的日志系统。

本 SDK 设计目标是可被任何需要规模化访问卖家账号的 Go 服务直接复用 —— 多租户广告运营平台、定制化报表流水线、库存 / 订单同步 worker、内部一次性工具,等等。

---

## 核心亮点

### 一套 LwA，两大 API

`advertising/auth` 与 `selling_partner/auth` 分别定义 `AdAuth` / `SpAuth`，但都实现同一个 `pkg.IAuth` 接口。意味着下游的缓存、刷新、401 重试、日志 hook **只看到 `IAuth`**，新增国家只是 auth 包里多一行 switch case。

### 面向规模化卖家

- **合并刷新**：`golang.org/x/sync/singleflight` 保证 10 个 goroutine 同时撞上过期 token 时，只会向 LwA 发起**一次**请求，其他人复用刚换出来的 token。
- **认证重试**：Amazon 返回 401 或某些"token 已轮换"的 403 时，SDK 自动 invalidate 缓存、重换 token、把原请求重放一次，业务无感知。
- **按租户分键**：缓存键是 `(refresh_token, ClientID)`，多店铺 / 多租户部署不会互相踩 token。

### 生成式 client

`advertising/sp_v3`、`sb_v4`、`sd_v1`、`report_v3` 以及大部分 `selling_partner/*` 都是基于亚马逊官方 OpenAPI spec 生成（经轻量后处理）。提供强类型的请求 builder、响应 model 与 `.Execute()` 方法，不需要手写 JSON 编解码。

### 可插拔的可观测性

SDK 不 import 你的 logger，而是暴露事件入口：

```go
pkg.OnTokenRefresh = func(ev pkg.TokenRefreshEvent) {
    // 把事件路由到 logrus / slog / zap / 自有日志系统
}
pkg.OnAuthRetry = func(ev pkg.AuthRetryEvent) {
    // 上报指标、对重试风暴做报警
}
```

典型用法:把这两个事件转发到你自己的 logger(logrus / slog / zap),按 `RefreshTokenSuffix` 聚合,既能跨 session 复盘每个店铺的 token 刷新流程,又永远不会把真实 token 写进日志。

---

## 覆盖范围

### Amazon Ads（`advertising/`）

| 包                        | 覆盖内容                                                                       |
| ------------------------- | ------------------------------------------------------------------------------ |
| `auth`                    | LwA token 交换 + 区域 endpoint 解析 + 带 Bearer 的 client 构造                  |
| `profiles_v2`             | Profiles API（advertiser context）                                              |
| `portfolios_v3`           | Portfolios CRUD                                                                 |
| `sp_v3`                   | **Sponsored Products v3** —— 活动、广告组、关键词、商品广告、Targets、否定、Budget rules、Optimization rules、Recommendations |
| `sb_v4`                   | Sponsored Brands v4                                                             |
| `sd_v1`                   | Sponsored Display v1                                                            |
| `report_v3`               | 异步广告报表 —— 申请、下载链接获取、GZIP 下载                                    |
| `product_metadata`        | Product Selector，为活动挑选可投放商品                                          |
| `product_eligibility`     | ASIN / SKU 在 SP / SB / SD 的投放资格校验                                       |
| `ams`                     | Amazon Marketing Stream 订阅                                                    |

### Selling Partner API（`selling_partner/`）

40+ 生成包，覆盖公开 SP-API。重点：

- **Catalog & Listings**：`catalog_items_20220401`、`listings_items_20210801`、`listings_restrictions_20210801`、`definitions_product_types_20200901`。
- **Orders & Fulfillment**：`orders_v0`、`fulfillment_inbound_v0`、`fulfillment_inbound_20240320`（新版）、`fulfillment_outbound_20200701`、`merchant_fulfillment_v0`、`shipping`、`shipping_v2`、`easy_ship_20220323`。
- **Inventory & FBA**：`fba_inventory`、`fba_inbound`、`awd_20240509`、`replenishment20221107`、`supply_sources_20200701`。
- **Reports & Feeds**：`reports_20210630`、`feeds_20210630`、`data_kiosk_20231115`。
- **Pricing & Fees**：`product_pricing_v0`、`product_pricing_20220501`、`product_fees_v0`。
- **Finances**：`finances_v0`、`finances_20240619`、`invoices_api_model_20240619`、`transfers_20240601`、`seller_wallet_20240301`。
- **Vendor**：`vendor_direct_fulfillment_*`、`vendor_invoices`、`vendor_orders`、`vendor_shipments`、`vendor_transaction_status`。
- **Notifications / Messaging / Tokens / Sellers / Sales / Application** 等。

每个包遵循同样的模式：`NewAPIClient(configuration *pkg.Configuration)` 返回一个 client，其 `*API` 字段暴露 builder 风格的请求方法，统一以 `.Execute()` 结束。

---

## 仓库结构

```text
amzsdk/
├── advertising/
│   ├── auth/                 # Ads API 的 LwA + Bearer client
│   ├── profiles_v2/
│   ├── portfolios_v3/
│   ├── sp_v3/                # Sponsored Products v3（最大模块）
│   ├── sb_v4/
│   ├── sd_v1/
│   ├── report_v3/
│   ├── product_metadata/
│   ├── product_eligibility/
│   └── ams/
├── selling_partner/
│   ├── auth/                 # SP-API 的 LwA + Bearer client（支持 Beta / AppID）
│   ├── orders_v0/
│   ├── fba_inventory/
│   ├── fulfillment_inbound_20240320/
│   ├── reports_20210630/
│   ├── listings_items_20210801/
│   └── ...                   # 40+ 个 OpenAPI 生成包
├── pkg/                      # 公共构件
│   ├── iauth.go              # IAuth 接口（两套 auth 共同实现）
│   ├── configuration.go      # API client 配置
│   ├── httpx.go              # HTTP client 包装(401 重试、trace 等)
│   ├── http_trace.go         # 可选的请求/响应追踪
│   ├── cache.go              # access_token 缓存(singleflight + 安全窗口)
│   ├── events.go             # OnTokenRefresh / OnAuthRetry 事件入口
│   ├── marketplace.go        # 国家码 -> marketplaceId 映射
│   ├── logger.go             # 可插拔 logger 接口
│   ├── timex.go              # 时间工具
│   └── tool.go               # pkg.Ptr[T] 等小工具
└── go.mod / go.sum
```

---

## 安装

本模块的公开 module path 是 `github.com/4379711/amz-sdk`。在你自己的 Go 项目里:

```bash
go get github.com/4379711/amz-sdk@latest
```

然后按需 `import` 子包即可:

```go
import (
    "github.com/4379711/amz-sdk/pkg"
    adAuth "github.com/4379711/amz-sdk/advertising/auth"
    "github.com/4379711/amz-sdk/advertising/sp_v3"

    spAuth "github.com/4379711/amz-sdk/selling_partner/auth"
    "github.com/4379711/amz-sdk/selling_partner/orders_v0"
)
```

> **生产部署:** 请固定到具体 tag(`go get github.com/4379711/amz-sdk@v1.0.0`)。开发期跟 `@main` 没问题,生产不推荐。

### 想边改 SDK 边联调宿主应用

如果想在开发宿主应用时同步修改本 SDK,把本仓库 clone 到宿主项目同级,然后用 `replace` 指向本地副本:

```bash
git clone https://github.com/4379711/amz-sdk.git ../amz-sdk
```

宿主应用的 `go.mod`:

```go
require github.com/4379711/amz-sdk v1.0.0

replace github.com/4379711/amz-sdk => ../amz-sdk
```

发布前把 `replace` 那行删掉,公开 module path 会通过 Go 的 module proxy 自动解析,无需任何额外配置。

工具链要求：

- **Go 1.26+**（`go.mod` 锁定）。
- 依赖刻意保持精简 —— `bytedance/sonic`、`oklog/ulid/v2`、`gopkg.in/validator.v2`、`golang.org/x/sync`。

---

## 授权流程 (LwA OAuth)

这是**每个店铺只跑一次**的接入流程。新店铺接入时跑一次,拿到 step 2 的 `refresh_token` 长期保存,后续所有 API 调用用的都是它(广告 API:只要应用一直注册着就长期有效;SP-API:Amazon 官方文档约 1 年)。

两套 API 共用同样的三步 OAuth 流程,`AdAuth` / `SpAuth` 暴露的方法名也一致,差别只在授权 URL 和应用标识符。

### Ads API (`AdAuth`)

**第一步 —— 构造 LwA 授权 URL,把卖家引导过去:**

```go
import (
    adAuth "amzsdk/advertising/auth"
)

auth := &adAuth.AdAuth{
    App: &adAuth.App{
        ClientID:     "amzn1.application-oa2-client.xxxxx",
        ClientSecret: "yyyyy",
        RedirectURL:  "https://your-app.example.com/lwa/callback",
    },
    Seller: &adAuth.Seller{
        CountryCode: "US",
    },
}

fmt.Println(auth.GetLwaURL())
// → https://www.amazon.com/ap/oa?client_id=...&scope=advertising::campaign_management
//   &response_type=code&redirect_uri=...&state=<ulid>
```

卖家授权后 Amazon 把卖家重定向回 `redirect_uri`,带 `?code=ANxxxxxxx&state=<ulid>`。先校验 `state` 与你发出去之前存的值一致,再拿 `code`。

**第二步 —— 用一次性 `code` 换长期 `refresh_token`:**

```go
if err := auth.GetRefreshToken("ANxxxxxxx"); err != nil {
    return fmt.Errorf("用 code 换 refresh_token 失败: %w", err)
}
// 落盘存好(DB 表按 seller 维度,字段加密)。
fmt.Println(auth.Token.RefreshToken) // Atzr|...
```

**第三步(可选)—— 手动换一枚 `access_token`:**

```go
auth.Token = &adAuth.Token{RefreshToken: "Atzr|..."} // 从 DB 取出
if err := auth.GetAccessTokenFromEndpoint(); err != nil {
    return fmt.Errorf("换 access_token 失败: %w", err)
}
fmt.Println(auth.Token.AccessToken)  // Bearer token, ~60 分钟过期
```

业务侧通常**不**需要手动调 `GetAccessTokenFromEndpoint` —— 任何通过 `pkg.NewConfiguration(auth)` 构造出来的 `NewAPIClient(...)` 都会用 `singleflight` 自动刷新并缓存 `access_token`。第三步只为一次性 CLI 工具和接入调试暴露。

### SP-API (`SpAuth`)

流程结构一致 —— 只有授权页 URL 不同(走 Seller Central 而不是 LwA 直连),且 `App` 多了一个 `AppID`:

```go
import (
    spAuth "amzsdk/selling_partner/auth"
)

auth := &spAuth.SpAuth{
    App: &spAuth.App{
        AppID:        "amzn1.sp.solution.<your-solution-id>",
        ClientID:     "amzn1.application-oa2-client.xxxxx",
        ClientSecret: "yyyyy",
        RedirectURL:  "https://your-app.example.com/spapi/callback",
        Beta:         true, // Seller Central 上把应用切到 Production 后改 false
    },
    Seller: &spAuth.Seller{
        CountryCode: "US",
        SellerType:  "SC", // "SC" = Seller Central,"VC" = Vendor Central
    },
}

// 第一步 —— Seller Central 上的授权页
fmt.Println(auth.GetLwaURL())

// 第二步 —— 用 callback URL 里拿到的 spapi_oauth_code 换 refresh_token
if err := auth.GetRefreshToken("ANxxxxxxx"); err != nil { /* … */ }
fmt.Println(auth.Token.RefreshToken)

// 第三步 —— 可选,API client 会自动刷新 access_token
auth.GetAccessTokenFromEndpoint()
```

卖家在 Seller Central 上授权完成后,Amazon 把卖家重定向回 `redirect_uri`,带三个 query 参数:

| 参数                    | 用途                                                       |
| ----------------------- | ---------------------------------------------------------- |
| `selling_partner_id`    | 与 seller 一起入库(是 `A1...` 形式的 Merchant ID)          |
| `mws_auth_token`        | 仅旧 MWS 迁移用 —— 纯 SP-API 应用忽略即可                  |
| **`spapi_oauth_code`** | 这串字符串传给 `auth.GetRefreshToken(code)`                |

---

## 快速上手

> 下面示例假设你已经走完上面的[授权流程](#授权流程-lwa-oauth),拿到了某个店铺的 `refresh_token` 并存好了。

### Sponsored Products v3 —— 拉活动列表

```go
import (
    "context"
    "fmt"

    adAuth "amzsdk/advertising/auth"
    "amzsdk/advertising/sp_v3"
    "amzsdk/pkg"
)

func ListCampaigns(ctx context.Context, profileID string) error {
    auth := &adAuth.AdAuth{
        App: &adAuth.App{
            ClientID:     "amzn1.application-oa2-client.xxxxx",
            ClientSecret: "yyyyy",
        },
        Seller: &adAuth.Seller{
            SellerName:  "demo-seller",
            CountryCode: "US",
            ProfileId:   profileID,
        },
        Token: &adAuth.Token{
            RefreshToken: "Atzr|...",
        },
    }

    cfg := pkg.NewConfiguration(auth)
    client := sp_v3.NewAPIClient(cfg)

    req := client.CampaignsAPI.ListSponsoredProductsCampaigns(ctx).
        SponsoredProductsListSponsoredProductsCampaignsRequestContent(
            sp_v3.SponsoredProductsListSponsoredProductsCampaignsRequestContent{
                IncludeExtendedDataFields: pkg.Ptr(true),
            },
        )

    resp, _, err := req.Execute()
    if err != nil {
        return fmt.Errorf("拉广告活动失败: %w", err)
    }
    fmt.Printf("共 %d 条活动\n", len(resp.Campaigns))
    return nil
}
```

### Selling Partner —— 查最近一天订单

```go
import (
    "context"
    "fmt"
    "time"

    spAuth "amzsdk/selling_partner/auth"
    "amzsdk/selling_partner/orders_v0"
    "amzsdk/pkg"
)

func ListOrders(ctx context.Context) error {
    auth := &spAuth.SpAuth{
        App: &spAuth.App{
            ClientID:     "amzn1.application-oa2-client.xxxxx",
            ClientSecret: "yyyyy",
        },
        Seller: &spAuth.Seller{
            SellerName:  "demo-seller",
            CountryCode: "US",
            SellerType:  "SC",
        },
        Token: &spAuth.Token{
            RefreshToken: "Atzr|...",
        },
    }

    cfg := pkg.NewConfiguration(auth)
    client := orders_v0.NewAPIClient(cfg)

    req := client.OrdersV0API.GetOrders(ctx, []string{pkg.MarketplaceMap["US"]}).
        CreatedAfter(time.Now().Add(-24 * time.Hour).UTC().Format(time.RFC3339))

    resp, _, err := req.Execute()
    if err != nil {
        return fmt.Errorf("拉订单失败: %w", err)
    }
    fmt.Printf("共 %d 条订单\n", len(resp.GetPayload().Orders))
    return nil
}
```

### 其他常用入口

上面两段示例骨架相同 —— 构造 `*Auth`、交给 `pkg.NewConfiguration`、再用对应包的 `NewAPIClient`。下表给出其他典型场景对应的入口方法,方便速查:

| 想做什么                              | 包                                          | 入口方法                                                                  |
| ------------------------------------ | ------------------------------------------- | ------------------------------------------------------------------------- |
| 拉店铺的广告 profile 列表             | `advertising/profiles_v2`                   | `client.ProfilesAPI.ListProfiles(ctx).Execute()`                          |
| 列 portfolios                        | `advertising/portfolios_v3`                 | `client.PortfoliosAPI.ListPortfolios(ctx).Execute()`                      |
| 改活动 / 广告组 / 关键词              | `advertising/sp_v3`                         | `client.{Campaigns,AdGroups,Keywords}API.{Create,Update,Delete}…`         |
| 下载 SP 广告报表(异步)               | `advertising/report_v3`                     | `CreateAsyncReport` → 轮询 `GetAsyncReport` → 拉 `Url` 指向的 GZIP        |
| ASIN 广告投放资格校验                 | `advertising/product_eligibility`           | `client.ProductEligibilityAPI.ProductEligibility(ctx)...Execute()`        |
| 列 SP-API 订单                       | `selling_partner/orders_v0`                 | `client.OrdersV0API.GetOrders(ctx, marketplaceIds)…`                      |
| 拉 FBA 库存快照                      | `selling_partner/fba_inventory`             | `client.FbaInventoryAPI.GetInventorySummaries(ctx)…`                      |
| 申请 Selling Partner 报表             | `selling_partner/reports_20210630`          | `client.ReportsAPI.CreateReport(ctx)…` → `GetReport` → `GetReportDocument` |
| 提交 Feed                            | `selling_partner/feeds_20210630`            | `client.FeedsAPI.CreateFeedDocument` → 上传 → `CreateFeed`               |
| 读取 listing item                    | `selling_partner/listings_items_20210801`   | `client.ListingsItemsAPI.GetListingsItem(ctx, …)…`                       |

所有生成包遵循统一的 **builder → `Execute()`** 模式。可空字段统一通过 `pkg.Ptr[T]`(或包内自带的 `PtrString` / `PtrBool` 等)构造。

---

## 可观测性 Hook

`amzsdk` 不带 logger。它**只在真实工作发生时**抛事件 —— 命中缓存的情况完全静默。

```go
import "amzsdk/pkg"

func init() {
    pkg.OnTokenRefresh = func(ev pkg.TokenRefreshEvent) {
        // ev.Phase: "start" | "end"
        // ev.Service: "advertising" | "selling_partner"
        // ev.Reason:  "cache-miss" | "401-retry"
        // ev.RefreshTokenSuffix / ev.AccessTokenSuffix: 脱敏到后 8 位
        // phase="end" 时还会带 ev.Duration / ev.ExpiresAt / ev.Success / ev.Err
    }

    pkg.OnAuthRetry = func(ev pkg.AuthRetryEvent) {
        // ev.FirstStatus / ev.FirstBodyHead: 触发重试的现场
        // ev.RetrySkipped: 如 "body-not-cloneable",说明重试被压制
        // ev.RetryStatus / ev.RetryErr: 重试本身的结果
    }
}
```

一个比较稳的接入做法是:

- `OnTokenRefresh` 的 `Phase=end` 走 info 级,带 `Duration` + `ExpiresAt`;`OnAuthRetry` 走 warn 级,带 `FirstStatus` + `RetryStatus`;
- 给 `Reason="401-retry"` 加一个计数器,token 流回归时能立刻报警;
- 把 LwA 相关日志单独落到一个文件/通道,接入排查时能直接 tail,不会被业务日志淹掉。

要打开逐请求 HTTP trace(很详细,**只**用于排障),从应用侧调用 `pkg/http_trace.go` 暴露的 helper,并用一个配置开关控制启停,生产默认关掉。

---

## 为什么把两套 API 合在一个 module

Amazon 的广告 API 和 SP-API 团队各自维护独立的 OpenAPI spec，但在卖家应用层面，这两套**永远是搭配出现**的：

- 同一个 **LwA 应用**为两套 API 颁发 access_token。
- 同一个 **refresh_token 仓库**给两套 API 共用 —— 多租户服务真正需要的是**一个**缓存、**一个** singleflight、**一套** 401 重试。
- 真实平台几乎都需要交叉两套 API（如 SP-API 订单数据 + 广告搜索词报表 → 盈亏平衡分析）。

如果拆成两个独立仓库，每个 consumer 都得把"共享 auth / cache / retry"层重复实现两份，还容易在双方之间产生细微行为漂移。`amzsdk` 通过把两套 API 放在不同的包树（`advertising/`、`selling_partner/`）保证 import 边界清晰，同时让 `pkg/` 共享层可以同步演进。

---

## 稳定性与版本管理

- 当前固定版本为 `v1.0.0`。
- `pkg.IAuth` 契约与 `OnTokenRefresh` / `OnAuthRetry` 事件结构视为**稳定 API**。这两处的破坏性变更需要大版本号 + 迁移说明。
- `advertising/*` 与 `selling_partner/*` 跟随 Amazon OpenAPI spec。Amazon 发新版日期版本（例如 `fulfillment_inbound_20240320` 取代 `fulfillment_inbound_v0`）时，`amzsdk` 是**新增同级包**而非原地替换，老代码不会破。
- 下游 `go.mod` 请固定 tag。开发期 track `main` 没问题，生产部署不推荐。

---

## 贡献

欢迎提 issue、修复与新 API 覆盖。流程与代码风格见 [CONTRIBUTING.md](./CONTRIBUTING.md)，社区行为准则见 [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md)。

提交新增生成包时请附带：

- 用于生成的 Amazon OpenAPI spec 准确 URL / 版本号。
- 你应用的后处理步骤（必须可重复）。
- 在 PR 描述里贴一个最小的 builder + `Execute()` 调用片段，证明新 client 能跑通你自己的账号。

---

## 许可协议

许可证条款见仓库根目录的 [LICENSE](./LICENSE) 文件。

Amazon、Amazon Ads（Sponsored Products / Sponsored Brands / Sponsored Display）、Selling Partner、FBA、Amazon Marketing Stream 等均为 Amazon.com, Inc. 及其关联公司的商标。本 SDK 由社区维护，**不**由 Amazon 背书或赞助。

---

## 免责声明

`amzsdk` 按"现状"提供，不附带任何形式的保证。两套 API 的误用都可能损害卖家账户健康度、泄露运营指标，或触发限流让正常业务卡死。你需要自行：

- 不要把 LwA 凭据 / refresh_token 提交进版本库。
- 遵守每个 API 已公布的限流规则。
- 验证生成的请求模型仍与你接入的 Amazon 现网契约一致 —— Amazon 偶尔会在同一版本上发起破坏性变更。
