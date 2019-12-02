# Terraform SourceHut Provider

[![Issue Tracker][badge]][issues]
[![Patches][listbadge]][mailing list]

[badge]: https://img.shields.io/badge/style-todo.sr.ht-green.svg?longCache=true&style=popout-square&label=issues
[listbadge]: https://img.shields.io/badge/style-lists.sr.ht-blue.svg?longCache=true&style=popout-square&label=patches
[issues]: https://todo.sr.ht/~samwhited/terraform-provider-sourcehut


This is the repository for the Terraform SourceHut (srht) Provider, which one
can use with Terraform to manage resources such as Git repos and issue trackers
on [SourceHut].

[SourceHut]: https://sourcehut.org/

For general information about Terraform, visit the [official
website] and the [GitHub project page].

[official website]: https://www.terraform.io/
[GitHub project page]: https://github.com/hashicorp/terraform


## Using the Provider

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


## Submitting Patches

To submit a patch, first read the [mailing list etiquette] and [contribution]
guides and then send patches to the [mailing list].
Please prefix the subject with `[PATCH terraform-provider-sourcehut]`.
To configure your checkout of this repo to always use the correct prefix and
send to the correct list cd into the repo and run:

    git config sendemail.to ~samwhited/terraform-provider-sourcehut@lists.sr.ht
    git config format.subjectPrefix 'PATCH terraform-provider-sourcehut'

[mailing list etiquette]: https://man.sr.ht/lists.sr.ht/etiquette.md
[contribution]: https://man.sr.ht/git.sr.ht/send-email.md
[mailing list]: https://lists.sr.ht/~samwhited/terraform-provider-sourcehut
