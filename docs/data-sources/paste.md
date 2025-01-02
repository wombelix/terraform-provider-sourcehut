---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sourcehut_paste Data Source - sourcehut"
subcategory: ""
description: |-

---

# sourcehut_paste (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The SHA1 hash of the paste.

### Read-Only

- `canonical_user` (String) The canonical name of the user that owns the paste (eg. '~example').
- `created` (String) The date on which the paste was created in RFC3339 format.
- `created_unix` (Number) The date on which the paste was created as a unix timestamp.
- `user` (String) The name of the user that owns the paste (eg. 'example').