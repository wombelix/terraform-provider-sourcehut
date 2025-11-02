## v1.0.0 (2025-11-02)

### Feat

- Migrate to GraphQL API client (BREAKING CHANGE)
- **client**: add internal GraphQL client implementation

### Fix

- correct provider description and user data source
- improve GraphQL client error handling and missing fields
- complete user key resources to match documentation
- send repo name only if changed on update
- repo visibility normalization on create and update repo
- type assertions from config to *config to match the pointer type returned by configureProvider
- **resource**: StateFuc added to normalize visibility value for repos
- wrong usage of pointer to config
- **goreleaser**: DEPRECATED archives.format replaced

## v0.2.1 (2025-07-28)

### Fix

- bump sourcehut-go to v0.1.1 to fix repo desc update issues

## v0.2.0 (2024-12-27)

### Feat

- Config and Doc to publish provider

## v0.1.0 (2024-12-27)

### Feat

- Migration from TF SDK v1 to TF SDK v2
- Migration from TF Core to TF SDK v1

### Fix

- Adjust go imports to new repo
