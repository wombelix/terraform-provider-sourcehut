<!--
    SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
    SPDX-FileCopyrightText: 2019 The SourceHut API Contributors

    SPDX-License-Identifier: BSD-2-Clause
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

Until the provider finds its way into the Terraform repository or your favorite
operating systems package repository, you will need to build the provider and
install it manually.

After the build is complete (try running `make`), copy the
`terraform-provider-sourcehut` binary into the third party plugins directory
(eg. `~/.terraform.d/plugins`) and re-run `terraform init`.
For more information, see the documentation about [third party plugins].

The documentation is not being built yet, so for an example of the plugins use
see the `example/` tree.

[third party plugins]: https://www.terraform.io/docs/configuration/providers.html#third-party-plugins

## Source

The primary location is:
[git.sr.ht/~wombelix/terraform-provider-sourcehut](https://git.sr.ht/~wombelix/terraform-provider-sourcehut)

Mirrors are available on
[Codeberg](https://codeberg.org/wombelix/terraform-provider-sourcehut),
[Gitlab](https://gitlab.com/wombelix/terraform-provider-sourcehut)
and
[Github](https://github.com/wombelix/terraform-provider-sourcehut).

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