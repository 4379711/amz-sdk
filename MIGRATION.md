# Upgrading to v1.3.0

The Go module path remains `github.com/4379711/amz-sdk`. Existing v1 import paths
can be retained:

```bash
go get github.com/4379711/amz-sdk@v1.3.0
```

```go
import (
    "github.com/4379711/amz-sdk/advertising/sp_v3"
    "github.com/4379711/amz-sdk/pkg"
)
```

Applications whose campaign IDs are strings can keep those fields as strings.
This release retains the module path but includes the source-incompatible
generated-type corrections listed below. Consumers of those APIs must update the
affected arguments and fields; keeping the import path does not remove those
type changes.

For local SDK integration tests, use one dependency and replacement:

```go
require github.com/4379711/amz-sdk v1.3.0

replace github.com/4379711/amz-sdk => ../
```

# Numeric entity IDs

Numeric Amazon entity IDs use `int64` in the Go API and JSON numbers on the wire.
The following public types require callers that use `float32` variables or pointers
to update those declarations. Existing string IDs retain their string types.

| Package | Affected API or model | ID fields |
| --- | --- | --- |
| `advertising/sp_v3` | CreateAssociatedBudgetRulesForSPCampaigns, DisassociateAssociatedBudgetRuleForSPCampaigns, ListAssociatedBudgetRulesForSPCampaigns | campaignId parameter and request field |
| `advertising/sb_v4` | CreateAssociatedBudgetRulesForSBCampaigns, DisassociateAssociatedBudgetRuleForSBCampaigns, ListAssociatedBudgetRulesForSBCampaigns | campaignId parameter and request field |
| `advertising/sd_v1` | CampaignResponseEx | CampaignId |
| `advertising/sd_v1` | AdGroupResponseEx | CampaignId, AdGroupId |
| `advertising/sd_v1` | ProductAdResponse | AdId |
| `advertising/sd_v1` | ProductAdResponseEx | CampaignId, AdGroupId, AdId |
| `advertising/sd_v1` | TargetingClauseEx | CampaignId, AdGroupId, TargetId |
| `advertising/sd_v1` | NegativeTargetingClauseEx | AdGroupId, TargetId |
| `advertising/sd_v1` | CreateCreative | AdGroupId |
| `advertising/sd_v1` | Creative, CreativeUpdate, CreativeResponse, CreativeModeration | CreativeId |

Constructors, getters, `Get…Ok` methods and setters use the matching `int64` or
`*int64` type. JSON property names and optional-field behavior are preserved.
Prices, bids, budgets and existing string identifiers are outside this mapping.

```go
campaignID := int64(36027908138365)
request := client.BudgetRulesAPI.ListAssociatedBudgetRulesForSPCampaigns(ctx, campaignID)
```

When an application stores an ID as a string, convert only at an API boundary that
requires an integer:

```go
campaignID, err := strconv.ParseInt(rawID, 10, 64)
if err != nil {
    return err
}
```

Converting an ID that has already been rounded through `float32` cannot recover
the original value. Obtain it from the original string or Amazon response.

These are source-incompatible public Go type corrections included in v1.3.0.
The `pkg.IAuth` interface, authentication hooks and SP campaign string IDs retain
their existing public signatures.

# SD targeting expressions

`GetActualInstance()` selects a nested predicate for an array `value`, a legacy
predicate when `eventType` is present, a content predicate for
`contentCategorySameAs`, and a regular predicate for other scalar types. Callers
that type-assert the result should handle the selected type rather than assuming
every scalar expression is a content predicate.

`TargetingPredicateBase` and `SDTargetingPredicateBaseV31` carry private JSON state
to keep extension fields attached when nested elements are copied or reordered.
Use their constructors or keyed struct literals rather than positional literals.
Setting more than one variant in either targeting union returns a marshal error.

Edits to known fields are reflected in outgoing JSON. Unknown fields remain with
the parsed object; replacing an element with a newly constructed object uses the
new object's fields. Setting an optional field from a non-nil value to nil removes
that field. Assigning nil to a field already decoded from null preserves null,
because it is indistinguishable from leaving the field untouched.

# TLS and authentication

Both SDK HTTP transports verify server certificates by default. Applications
using a private certificate authority can set `TLSClientConfig.RootCAs` before
making requests. Report download clients that reuse `LongTimeHttpClient.Transport`
also use this verification.

Requests waiting for token refresh honor their own context cancellation. The
shared refresh has an independent 180-second limit and can complete for other
waiters or populate the cache after one waiter cancels.

Pass the original callback URL to `RedirectURL`; `GetLwaURL` performs query
encoding. Do not pre-encode the URL before assigning it.

# Enum decoding and exception codes

String enums preserve unknown response values. Use `IsValid()` or the corresponding
`New…FromValue()` constructor for explicit checks against the known values.
Non-string JSON values, except null, still produce decoding errors.

The following exported `sp_v3` exception-code constants use uppercase values from
v1.3.0. Code comparing stored earlier values against these constants should account for the
value changes. Decoding preserves the incoming spelling; it does not normalize it.

| Exception code | Earlier value | Value from v1.3.0 |
| --- | --- | --- |
| AccessDenied | `accessDenied` | `ACCESS_DENIED` |
| InternalServerException | `internalServerException` | `INTERNAL_SERVER_EXCEPTION` |
| NotImplemented | `notImplemented` | `NOT_IMPLEMENTED` |
| SchemaValidation | `invalidSchema` | `INVALID_SCHEMA` |
| ServiceUnavailable | `serviceUnavailable` | `SERVICE_UNAVAILABLE` |
| Throttling | `throttled` | `THROTTLED` |
| Unauthenticated | `unauthenticated` | `UNAUTHENTICATED` |
