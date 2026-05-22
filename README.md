<div align="center">

# amzsdk

**A unified Go SDK for the Amazon Ads API and the Amazon Selling Partner API — one auth stack, one HTTP pipeline, one observability surface.**

[简体中文](./README.zh-CN.md) · English

[![Go](https://img.shields.io/badge/Go-1.26%2B-00ADD8?logo=go)](https://go.dev/)
[![Amazon Ads](https://img.shields.io/badge/Amazon%20Ads-SP%20%C2%B7%20SB%20%C2%B7%20SD-FF9900)](https://advertising.amazon.com/API/docs/en-us/)
[![Selling Partner](https://img.shields.io/badge/SP--API-v0%20%C2%B7%20v3-FF9900)](https://developer-docs.amazon.com/sp-api/)

</div>

---

## Table of Contents

- [What is amzsdk](#what-is-amzsdk)
- [Highlights](#highlights)
- [Coverage](#coverage)
- [Repository layout](#repository-layout)
- [Install](#install)
- [Authorization (LwA OAuth)](#authorization-lwa-oauth)
- [Quick start](#quick-start)
- [Observability hooks](#observability-hooks)
- [Why a single module for both APIs](#why-a-single-module-for-both-apis)
- [Stability & versioning](#stability--versioning)
- [Contributing](#contributing)
- [License](#license)
- [Disclaimer](#disclaimer)

---

## What is amzsdk

`amzsdk` is a Go SDK that ships **two large families of Amazon HTTP APIs** behind one consistent client surface:

- **Amazon Ads** — Sponsored Products v3, Sponsored Brands v4, Sponsored Display v1, Profiles v2, Reports v3, Portfolios, Product Metadata / Eligibility, AMS.
- **Selling Partner (SP-API)** — Orders, Inventory, Reports, Catalog, Listings, Feeds, FBA Inbound (incl. 2024 redesign), Finances v0 + 2024, Pricing, Shipping, Easy Ship, A+ Content, Messaging, Notifications, Vendor APIs, and more — 40+ generated client packages.

Internally both stacks share the same:

- **LwA token cache** with a 5-minute safety window, `singleflight`-coalesced refresh, and per-`(refresh_token, ClientID)` keying.
- **HTTP pipeline** — `pkg.Configuration` + `pkg/httpx.go` + `pkg/http_trace.go`, with a built-in 401/403 retry that transparently re-mints `access_token` and replays the original request.
- **Region-aware endpoint resolution** — `pkg.MarketplaceMap` and per-country LwA URLs cover NA / EU / FE / India / MENA / AU / JP / SG marketplaces.
- **Observability hooks** — `pkg.OnTokenRefresh` / `pkg.OnAuthRetry` let your application receive structured events for every actual LwA refresh and every auto-retried 401, without the SDK itself touching your logging system.

The SDK is designed to be embedded in any Go service that needs to talk to seller accounts at scale — multi-tenant ad-ops platforms, custom report pipelines, inventory / order ingestion workers, or one-off internal tools.

---

## Highlights

### One LwA, two API families

`advertising/auth` and `selling_partner/auth` produce different `*Auth` structs (`AdAuth` vs `SpAuth`) but both implement the same `pkg.IAuth` contract. That means downstream code — caching, refresh, 401 retry, logging hooks — only ever sees `IAuth`. Adding a new marketplace endpoint is a one-line switch case in the auth package.

### Built for high-throughput sellers

- **Coalesced refresh** — `golang.org/x/sync/singleflight` guarantees that ten goroutines simultaneously hitting an expired token trigger **exactly one** LwA call; the rest reuse the freshly minted token.
- **Auth retry** — when Amazon returns 401 (or one of the documented 403 variants that really mean "token rotated"), the SDK invalidates the cache, re-mints, and replays the original request once, transparently.
- **Per-seller cache keys** — the cache is keyed on `(refresh_token, ClientID)` so multi-tenant deployments don't blow each other's tokens away.

### Generated client surface

The `advertising/sp_v3`, `sb_v4`, `sd_v1`, `report_v3`, and most `selling_partner/*` packages are generated from the official Amazon OpenAPI specs (lightly post-processed). You get strongly-typed request builders, response models, and `Execute()` methods — no hand-rolled marshalling.

### Drop-in observability

The SDK never imports your logger. Instead it exposes:

```go
pkg.OnTokenRefresh = func(ev pkg.TokenRefreshEvent) {
    // route to logrus / slog / zap / your bespoke logger
}
pkg.OnAuthRetry = func(ev pkg.AuthRetryEvent) {
    // record metrics, alert on retry storms
}
```

Typical use: forward both events into your own logger (logrus / slog / zap), keyed by `RefreshTokenSuffix`, so you can correlate every seller's token-refresh flow per session without ever logging the raw token.

---

## Coverage

### Amazon Ads (`advertising/`)

| Package                  | What it covers                                                                                 |
| ------------------------ | ---------------------------------------------------------------------------------------------- |
| `auth`                   | LwA token exchange + per-region endpoint resolution + Bearer client builder.                    |
| `profiles_v2`            | Profiles API (entity / advertiser context).                                                     |
| `portfolios_v3`          | Portfolios CRUD.                                                                                |
| `sp_v3`                  | **Sponsored Products v3** — Campaigns, Ad groups, Keywords, Product Ads, Targets, Negatives, Budget rules, Optimization rules, Recommendations. |
| `sb_v4`                  | Sponsored Brands v4.                                                                            |
| `sd_v1`                  | Sponsored Display v1.                                                                           |
| `report_v3`              | Async ads reports — request, download URL, presigned GZIP fetch.                                |
| `product_metadata`       | Product Selector — picks ad-eligible products for a campaign.                                   |
| `product_eligibility`    | ASIN / SKU eligibility checks for SP / SB / SD.                                                 |
| `ams`                    | Amazon Marketing Stream subscriptions.                                                          |

### Selling Partner API (`selling_partner/`)

40+ generated packages spanning the public SP-API. Highlights:

- **Catalog & Listings** — `catalog_items_20220401`, `listings_items_20210801`, `listings_restrictions_20210801`, `definitions_product_types_20200901`.
- **Orders & Fulfillment** — `orders_v0`, `fulfillment_inbound_v0`, `fulfillment_inbound_20240320` (new design), `fulfillment_outbound_20200701`, `merchant_fulfillment_v0`, `shipping`, `shipping_v2`, `easy_ship_20220323`.
- **Inventory & FBA** — `fba_inventory`, `fba_inbound`, `awd_20240509`, `replenishment20221107`, `supply_sources_20200701`.
- **Reports & Feeds** — `reports_20210630`, `feeds_20210630`, `data_kiosk_20231115`.
- **Pricing & Fees** — `product_pricing_v0`, `product_pricing_20220501`, `product_fees_v0`.
- **Finances** — `finances_v0`, `finances_20240619`, `invoices_api_model_20240619`, `transfers_20240601`, `seller_wallet_20240301`.
- **Vendor APIs** — `vendor_direct_fulfillment_*`, `vendor_invoices`, `vendor_orders`, `vendor_shipments`, `vendor_transaction_status`.
- **Notifications, Messaging, Tokens, Sellers, Sales, Application** — the small but essential helpers.

Each package follows the same pattern: `NewAPIClient(configuration *pkg.Configuration)` returns a client whose `*API` fields expose builder-style request methods, all ending in `.Execute()`.

---

## Repository layout

```text
amzsdk/
├── advertising/
│   ├── auth/                 # LwA + Bearer client for Ads API
│   ├── profiles_v2/
│   ├── portfolios_v3/
│   ├── sp_v3/                # Sponsored Products v3 (largest module)
│   ├── sb_v4/
│   ├── sd_v1/
│   ├── report_v3/
│   ├── product_metadata/
│   ├── product_eligibility/
│   └── ams/
├── selling_partner/
│   ├── auth/                 # LwA + Bearer client for SP-API (Beta / AppID-aware)
│   ├── orders_v0/
│   ├── fba_inventory/
│   ├── fulfillment_inbound_20240320/
│   ├── reports_20210630/
│   ├── listings_items_20210801/
│   └── ...                   # 40+ packages, one per official API
├── pkg/                      # Shared building blocks
│   ├── iauth.go              # IAuth interface (the contract both auths satisfy)
│   ├── configuration.go      # API client configuration
│   ├── httpx.go              # HTTP client wrapping (401 retry, trace, etc.)
│   ├── http_trace.go         # Optional request/response tracing
│   ├── cache.go              # Access-token cache (singleflight + safety window)
│   ├── events.go             # OnTokenRefresh / OnAuthRetry hooks
│   ├── marketplace.go        # Country -> marketplaceId map
│   ├── logger.go             # Pluggable logger interface
│   ├── timex.go              # Time helpers
│   └── tool.go               # pkg.Ptr[T] and friends
└── go.mod / go.sum
```

---

## Install

The module is published as `github.com/4379711/amz-sdk`. In your own Go project:

```bash
go get github.com/4379711/amz-sdk@latest
```

Then `import` whichever sub-package you need:

```go
import (
    "github.com/4379711/amz-sdk/pkg"
    adAuth "github.com/4379711/amz-sdk/advertising/auth"
    "github.com/4379711/amz-sdk/advertising/sp_v3"

    spAuth "github.com/4379711/amz-sdk/selling_partner/auth"
    "github.com/4379711/amz-sdk/selling_partner/orders_v0"
)
```

> **Production tip:** pin a specific tag (`go get github.com/4379711/amz-sdk@v1.0.0`). Tracking `@main` is fine for development but not recommended for production deploys.

### Hacking on the SDK from a host application

If you need to patch the SDK while developing your host app, clone the repo next to your project and point `go.mod` at the local copy:

```bash
git clone https://github.com/4379711/amz-sdk.git ../amz-sdk
```

In your host app's `go.mod`:

```go
require github.com/4379711/amz-sdk v1.0.0

replace github.com/4379711/amz-sdk => ../amz-sdk
```

Drop the `replace` line before shipping; the public module path resolves through Go's proxy without any extra configuration.

Toolchain prerequisites:

- **Go 1.26+** (pinned in `go.mod`).
- The dependency set is intentionally small — `bytedance/sonic`, `oklog/ulid/v2`, `gopkg.in/validator.v2`, `golang.org/x/sync`.

---

## Authorization (LwA OAuth)

This is the **one-time per-seller** onboarding flow. Run it when you connect a new seller account; the `refresh_token` you get out of step 2 is long-lived (Ads: indefinite while your app stays registered; SP-API: ≈ 1 year per Amazon's docs) and is what every subsequent API call needs.

Both API families use the same three-step OAuth flow, exposed through identical method names on `AdAuth` / `SpAuth`. Only the consent URL and the app identifier differ.

### Ads API (`AdAuth`)

**Step 1 — build the LwA consent URL and redirect the seller to it:**

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

Amazon redirects the seller back to your `redirect_uri` with `?code=ANxxxxxxx&state=<ulid>`. Verify the `state` against the value you stored before sending the user out, then take the `code`.

**Step 2 — exchange the one-time `code` for a long-lived `refresh_token`:**

```go
if err := auth.GetRefreshToken("ANxxxxxxx"); err != nil {
    return fmt.Errorf("exchange auth code: %w", err)
}
// Persist this somewhere safe (DB row keyed by seller, encrypted at rest).
fmt.Println(auth.Token.RefreshToken) // Atzr|...
```

**Step 3 (optional) — manually mint an `access_token`:**

```go
auth.Token = &adAuth.Token{RefreshToken: "Atzr|..."} // hydrate from your DB
if err := auth.GetAccessTokenFromEndpoint(); err != nil {
    return fmt.Errorf("mint access token: %w", err)
}
fmt.Println(auth.Token.AccessToken)  // Bearer token, expires after ~60 min
```

You normally **don't** need to call `GetAccessTokenFromEndpoint` yourself — every `NewAPIClient(...)` built from `pkg.NewConfiguration(auth)` automatically refreshes (and caches) the `access_token` for you using `singleflight`. Step 3 is exposed for one-off CLI tooling and integration debugging.

### SP-API (`SpAuth`)

The flow is structurally identical — only the consent URL is different (Seller Central, not LwA directly) and the `App` carries an additional `AppID`:

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
        Beta:         true, // set to false once your app is in Production mode in Seller Central
    },
    Seller: &spAuth.Seller{
        CountryCode: "US",
        SellerType:  "SC", // "SC" = Seller Central, "VC" = Vendor Central
    },
}

// Step 1 — consent URL on Seller Central.
fmt.Println(auth.GetLwaURL())

// Step 2 — exchange the spapi_oauth_code from your callback URL.
if err := auth.GetRefreshToken("ANxxxxxxx"); err != nil { /* … */ }
fmt.Println(auth.Token.RefreshToken)

// Step 3 — optional, the API clients refresh access_token automatically.
auth.GetAccessTokenFromEndpoint()
```

When the seller authorizes on Seller Central, Amazon redirects to `redirect_uri` with three query parameters:

| Param                  | What you do with it                                                   |
| ---------------------- | --------------------------------------------------------------------- |
| `selling_partner_id`   | Save with the seller record (it's the `A1...` Merchant ID).            |
| `mws_auth_token`       | Legacy MWS migration only — ignore for SP-API-only apps.              |
| **`spapi_oauth_code`** | Pass this string to `auth.GetRefreshToken(code)`.                     |

---

## Quick start

> The snippets below assume you have already completed the [Authorization flow](#authorization-lwa-oauth) above and have a per-seller `refresh_token` stored somewhere.

### Sponsored Products v3 — list campaigns

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
        return fmt.Errorf("list campaigns: %w", err)
    }
    fmt.Printf("got %d campaigns\n", len(resp.Campaigns))
    return nil
}
```

### Selling Partner — list orders since yesterday

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
        return fmt.Errorf("list orders: %w", err)
    }
    fmt.Printf("got %d orders\n", len(resp.GetPayload().Orders))
    return nil
}
```

### Other common entry points

The two snippets above share the same template — build an `*Auth`, hand it to `pkg.NewConfiguration`, then call the relevant package's `NewAPIClient`. The table below points at the entry point you most likely want for other typical scenarios:

| Want to                             | Package                                       | Entry point                                                                |
| ----------------------------------- | --------------------------------------------- | -------------------------------------------------------------------------- |
| List ad profiles for the seller     | `advertising/profiles_v2`                     | `client.ProfilesAPI.ListProfiles(ctx).Execute()`                           |
| List portfolios                     | `advertising/portfolios_v3`                   | `client.PortfoliosAPI.ListPortfolios(ctx).Execute()`                       |
| Mutate campaigns / ad groups / keywords | `advertising/sp_v3`                       | `client.{Campaigns,AdGroups,Keywords}API.{Create,Update,Delete}…`          |
| Pull a Sponsored Products report (async) | `advertising/report_v3`                  | `CreateAsyncReport` → poll `GetAsyncReport` → fetch `Url` (GZIP)           |
| Check ASIN ad-eligibility           | `advertising/product_eligibility`             | `client.ProductEligibilityAPI.ProductEligibility(ctx)...Execute()`         |
| List SP-API orders                  | `selling_partner/orders_v0`                   | `client.OrdersV0API.GetOrders(ctx, marketplaceIds)…`                       |
| Pull an FBA inventory snapshot      | `selling_partner/fba_inventory`               | `client.FbaInventoryAPI.GetInventorySummaries(ctx)…`                       |
| Request a Selling Partner report    | `selling_partner/reports_20210630`            | `client.ReportsAPI.CreateReport(ctx)…` → `GetReport` → `GetReportDocument` |
| Submit a feed                       | `selling_partner/feeds_20210630`              | `client.FeedsAPI.CreateFeedDocument` → upload → `CreateFeed`               |
| Read a listing item                 | `selling_partner/listings_items_20210801`     | `client.ListingsItemsAPI.GetListingsItem(ctx, …)…`                         |

Every generated package follows the **builder → `Execute()`** pattern. Optional fields go through `pkg.Ptr[T]` (or the per-package `PtrString` / `PtrBool` helpers).

---

## Observability hooks

`amzsdk` doesn't ship a logger. It instead emits structured events when **real** work happens — cache hits stay silent.

```go
import "amzsdk/pkg"

func init() {
    pkg.OnTokenRefresh = func(ev pkg.TokenRefreshEvent) {
        // ev.Phase: "start" | "end"
        // ev.Service: "advertising" | "selling_partner"
        // ev.Reason:  "cache-miss" | "401-retry"
        // ev.RefreshTokenSuffix / ev.AccessTokenSuffix: redacted (last 8 chars)
        // On phase="end" you also get ev.Duration / ev.ExpiresAt / ev.Success / ev.Err.
    }

    pkg.OnAuthRetry = func(ev pkg.AuthRetryEvent) {
        // ev.FirstStatus / ev.FirstBodyHead: what triggered the retry
        // ev.RetrySkipped: e.g. "body-not-cloneable" — retry was suppressed
        // ev.RetryStatus / ev.RetryErr: outcome of the replayed request
    }
}
```

A robust integration usually:

- routes `OnTokenRefresh` (`Phase=end`) into an info-level log with `Duration` + `ExpiresAt`, and `OnAuthRetry` into a warn-level log with `FirstStatus` + `RetryStatus`;
- exports a counter for `Reason="401-retry"` events to alert on token-flow regressions;
- writes the LwA-related logs into a dedicated file or sink so they can be tailed during onboarding without drowning your business logs.

To enable per-request HTTP tracing (verbose; only for debugging), wire up the helpers in `pkg/http_trace.go` from your application and gate it behind a config flag so it stays off in production.

---

## Why a single module for both APIs

Amazon's Ads and SP-API teams ship independent OpenAPI specs, but at the seller's app level the two stacks **always** show up together:

- The same **LwA app** issues access tokens for both.
- The same **refresh token vault** powers both — so a multi-tenant service really wants one cache, one singleflight group, one 401-retry policy.
- A real platform almost always cross-references the two APIs (e.g. SP-API order data + Ads search-term reports → break-even analysis).

Splitting these into two repos forces every consumer to re-implement the shared auth / cache / retry layer twice and risks subtle drift between them. `amzsdk` keeps the two API surfaces under separate package trees (`advertising/`, `selling_partner/`) so the import graph is unambiguous, while still letting the shared `pkg/` layer evolve in lockstep.

---

## Stability & versioning

- The current pinned version is `v1.0.0`.
- The `pkg.IAuth` contract and the `OnTokenRefresh` / `OnAuthRetry` event shapes are considered **stable**. Breaking changes here will get a major-version bump and a migration note in the release.
- The generated `advertising/*` and `selling_partner/*` packages track Amazon's OpenAPI specs. When Amazon releases a new dated version (e.g. `fulfillment_inbound_20240320` superseding `fulfillment_inbound_v0`), `amzsdk` adds the new package as a sibling rather than replacing the old one — your code keeps compiling.
- Always pin a specific tag in your downstream `go.mod`. Tracking `main` is fine for development but not recommended for production deploys.

---

## Contributing

Issues, fixes, and additional API coverage are welcome. See [CONTRIBUTING.md](./CONTRIBUTING.md) for the workflow and house style, and [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md) for community guidelines.

When proposing a new generated package, please include:

- The exact Amazon OpenAPI spec URL / version you generated from.
- Any post-processing steps you applied (these need to be reproducible).
- A minimal builder + `Execute()` snippet in the PR description showing the new client works against your account.

---

## License

See the [LICENSE](./LICENSE) file in the repository root for licensing terms.

Amazon, the Amazon Ads marks (Sponsored Products, Sponsored Brands, Sponsored Display), Selling Partner, FBA, and Amazon Marketing Stream are trademarks of Amazon.com, Inc. or its affiliates. This SDK is community-maintained and **not** endorsed or sponsored by Amazon.

---

## Disclaimer

`amzsdk` is provided **as-is**, with no warranty of any kind. Misuse of either API family can damage a seller's account standing, leak operational metrics, or trigger throttling that blocks legitimate operations. You are responsible for:

- Keeping LwA credentials and refresh tokens out of source control.
- Respecting each API's documented rate limits.
- Validating that the generated request models still match the live Amazon contracts you are integrating against — Amazon occasionally ships breaking changes on the same API version.
