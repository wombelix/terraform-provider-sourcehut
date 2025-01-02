<!--
    SPDX-FileCopyrightText: 2024-2025 Dominik Wombacher <dominik@wombacher.cc>
    SPDX-FileCopyrightText: 2019 The SourceHut API Contributors

    SPDX-License-Identifier: CC0-1.0
-->

# OpenTofu / Terraform Provider for sourcehut (sr.ht)

Based on the work from [SamWhited / terraform-provider-sourcehut](https://codeberg.org/SamWhited/terraform-provider-sourcehut)

---

[![REUSE status](https://api.reuse.software/badge/git.sr.ht/~wombelix/terraform-provider-sourcehut)](https://api.reuse.software/info/git.sr.ht/~wombelix/terraform-provider-sourcehut)
[![builds.sr.ht status](https://builds.sr.ht/~wombelix/terraform-provider-sourcehut.svg)](https://builds.sr.ht/~wombelix/terraform-provider-sourcehut?)

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

```
terraform {
  required_version = ">= 1.8"
  required_providers {
    sourcehut = {
      source  = "wombelix/sourcehut"
      version = "0.2.0"

      # SRHT_TOKEN env var
    }
  }
}
```

The sourcehut [oauth personal access tokens](https://meta.sr.ht/oauth)
will be read from Environment variable `SRHT_TOKEN`.

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
[git.sr.ht/~wombelix/terraform-provider-sourcehut](https://git.sr.ht/~wombelix/terraform-provider-sourcehut)

Mirrors are available on
[Codeberg](https://codeberg.org/wombelix/terraform-provider-sourcehut),
[Gitlab](https://gitlab.com/wombelix/terraform-provider-sourcehut)
and
[Github](https://github.com/wombelix/terraform-provider-sourcehut).

Publishing to
[registry.terraform.io](https://registry.terraform.io/providers/wombelix/sourcehut/latest)
and
[search.opentofu.org](https://search.opentofu.org/provider/wombelix/sourcehut/latest)
is handled by the GitHub mirror.

## Contribute

Please don't hesitate to provide Feedback,
open an Issue or create a Pull / Merge Request.

Just pick the workflow or platform you prefer and are most comfortable with.

Feedback, bug reports or patches to my sr.ht list
[~wombelix/inbox@lists.sr.ht](https://lists.sr.ht/~wombelix/inbox) or via
[Email and Instant Messaging](https://dominik.wombacher.cc/pages/contact.html)
are also always welcome.

## License

Unless otherwise stated: `BSD-2-Clause`

All files contain license information either as
`header comment` or `corresponding .license` file.

[REUSE](https://reuse.software) from the [FSFE](https://fsfe.org/)
implemented to verify license and copyright compliance.
