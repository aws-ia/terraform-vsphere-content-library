# TODO: Add tags once support for content libraries & content library items are added
# Reference: https://github.com/hashicorp/terraform-provider-vsphere/issues/1498

# module "vsphere_tags" {
#   count = length(var.vsphere_tags) > 0 ? 1 : 0

#   source = "github.com/aws-ia/terraform-vsphere-tags"
#   # version = "0.0.1"

#   vsphere_tag_category_name             = var.vsphere_tag_category_name
#   vsphere_tag_category_description      = var.vsphere_tag_category_description
#   vsphere_tag_category_associable_types = var.vsphere_tag_category_associable_types
#   vsphere_tag_category_cardinality      = var.vsphere_tag_category_cardinality
#   create_vsphere_tag_category           = var.create_vsphere_tag_category
#   vsphere_tags                          = var.vsphere_tags
#   create_vsphere_tags                   = var.create_vsphere_tags
# }

resource "vsphere_content_library" "content_library" {
  count = var.create_vsphere_content_library ? 1 : 0

  name            = var.vsphere_content_library_name
  description     = var.vsphere_content_library_description
  storage_backing = [data.vsphere_datastore.datastore.id]

  # tags = length(module.vsphere_tags) == 1 ? module.vsphere_tags[0].vsphere_tags[*].id : null
}

resource "vsphere_content_library_item" "items" {
  count = var.create_vsphere_content_library_items ? length(var.vsphere_content_library_items) : 0

  name        = var.vsphere_content_library_items[count.index]["name"]
  description = var.vsphere_content_library_items[count.index]["description"]
  file_url    = var.vsphere_content_library_items[count.index]["file_url"]
  type        = var.vsphere_content_library_items[count.index]["type"]
  library_id  = var.create_vsphere_content_library ? vsphere_content_library.content_library[0].id : data.vsphere_content_library.content_library[0].id

  # tags = length(module.vsphere_tags) == 1 ? module.vsphere_tags[0].vsphere_tags[*].id : null
}
