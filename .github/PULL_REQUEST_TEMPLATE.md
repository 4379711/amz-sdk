<!--
Thanks for opening a pull request!

`amzsdk` sits on the auth + HTTP path of every consumer. Please read CONTRIBUTING.md before your first PR — the checklist below mirrors the hard rules.
-->

## Summary

<!-- What does this PR do, and why? 1-3 sentences is enough for most PRs. -->

## Changes

<!-- Bullet list of the main changes, grouped by package. -->

- [ ] `pkg/` (auth contract / cache / events / HTTP)
- [ ] `advertising/*` generated client
- [ ] `selling_partner/*` generated client
- [ ] Docs / scripts

## Test plan

```
go mod tidy
go vet ./...
go build ./...
go test ./pkg/...
```

Manual / integration tests run (with placeholder names, never real credentials):

<!-- e.g. "ran TestListSponsoredProductsCampaigns against a sandbox profile in US, got expected campaign list" -->

## Generated package changes (fill in if applicable)

- OpenAPI source URL / version: <!-- e.g. https://developer-docs.amazon.com/sp-api/openapi/sp/listingsItems_2021-08-01.json -->
- Post-processing steps applied: <!-- sed rules / rename rules / gofmt / manual fixes -->

## Checklist

- [ ] CI is green (`go build ./...` + `go vet ./...` + `go test ./pkg/...`).
- [ ] No real credentials (`ClientID`, `ClientSecret`, `refresh_token`, `Authorization` headers) committed anywhere — including test fixtures and log snippets.
- [ ] `pkg/` changes preserve existing exported API. If they don't, the PR description justifies the break and proposes a migration.
- [ ] Generated package changes are reproducible from the documented OpenAPI source + post-processing steps.
- [ ] Comments describe **current** behaviour / constraints / reasons — no change history, bug-fix narratives or "old vs new" comparisons (see [CONTRIBUTING.md](../CONTRIBUTING.md)).
- [ ] If a new dated Amazon package was added, I did **not** delete the older sibling package — downstream callers may still depend on it.

## Related issues

<!-- Closes #123 / Refs #456 -->
