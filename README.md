<!--
    SPDX-FileCopyrightText: 2024-2025 Dominik Wombacher <dominik@wombacher.cc>
    SPDX-FileCopyrightText: 2019 The SourceHut API Contributors

    SPDX-License-Identifier: CC0-1.0
-->

# OpenTofu / Terraform Provider for sourcehut (sr.ht)

Based on the work from [SamWhited / terraform-provider-sourcehut](https://codeberg.org/SamWhited/terraform-provider-sourcehut)

**!!! IMPORTANT !!!**: The sr.ht legacy REST API was deprecated a while
ago and is now in its
[final phase of removal](https://sourcehut.org/blog/2025-09-01-whats-cooking-q3-2025/).
From version v1.0.0 of the provider uses the GraphQL API directly and
doesn't rely on the project
[sourcehut-go](https://github.com/wombelix/sourcehut-go) and the legacy
REST API anymore. This is a breaking change, you have to update to the
latest version and switch over to a new Oauth2 personal access token (PAT),
see [Usage](#usage) for more details.

---

<!-- markdownlint-disable MD013 -->
[![REUSE status](https://api.reuse.software/badge/github.com/wombelix/terraform-provider-sourcehut)](https://api.reuse.software/info/github.com/wombelix/terraform-provider-sourcehut)
[![Mirror](https://github.com/wombelix/terraform-provider-sourcehut/actions/workflows/mirror.yml/badge.svg)](https://github.com/wombelix/terraform-provider-sourcehut/actions/workflows/mirror.yml)
<!-- markdownlint-enable MD013 -->

## Table of Contents

* [Usage](#usage)
* [Source](#source)
* [Contribute](#contribute)
* [License](#license)

## Usage

The provider is available in the
[OpenTofu](https://search.opentofu.org/provider/wombelix/sourcehut/latest)
and
[Terraform](https://registry.terraform.io/providers/wombelix/sourcehut/latest)
registry.

Example usage in a `provider.tf` file:

```hcl
terraform {
  required_version = ">= 1.8"
  required_providers {
    sourcehut = {
      source  = "wombelix/sourcehut"
      version = "1.0.0"

      # SRHT_TOKEN env var
    }
  }
}
```

The sourcehut [oauth2 personal access tokens](https://meta.sr.ht/oauth2)
will be read from Environment variable `SRHT_TOKEN`.

The recommended scope is:

```text
git.sr.ht/PROFILE:RO git.sr.ht/REPOSITORIES:RW
paste.sr.ht/PROFILE:RO paste.sr.ht/PASTES:RW
meta.sr.ht/PGP_KEYS:RW meta.sr.ht/SSH_KEYS:RW meta.sr.ht/PROFILE:RO
```

You also have the option to build the provider and install it manually.

After the build is complete (`make`), copy the `terraform-provider-sourcehut`
binary into the third party plugins directory (e.g. `~/.terraform.d/plugins`)
and re-run `terraform init`. For more information, see the documentation about
[third party plugins](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

The documentation can be found in the`docs/` sub-folder. The online version is
available in the
[OpenTofu](https://search.opentofu.org/provider/wombelix/sourcehut/latest)
and
[Terraform](https://registry.terraform.io/providers/wombelix/sourcehut/latest)
registry.

## Source

The primary location is:
[github.com/wombelix/terraform-provider-sourcehut](https://github.com/wombelix/terraform-provider-sourcehut)

Mirrors are available on
[Codeberg](https://codeberg.org/wombelix/terraform-provider-sourcehut) and
[Gitlab](https://gitlab.com/wombelix/terraform-provider-sourcehut).

Publishing to
[registry.terraform.io](https://registry.terraform.io/providers/wombelix/sourcehut/latest)
and
[search.opentofu.org](https://search.opentofu.org/provider/wombelix/sourcehut/latest)
is handled from GitHub.

## Contribute

Pick the platform you prefer and are most comfortable with.

Provide feedback, open an issue or create a pull / merge request.

## License

Unless otherwise stated: `BSD-2-Clause`

All files contain license information either as a
`header comment` or a `corresponding .license` file.

[REUSE](https://reuse.software) from the [FSFE](https://fsfe.org/)
is implemented to verify license and copyright compliance.
