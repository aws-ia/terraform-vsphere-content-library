# TODO: Add tags once support for content libraries & content library items are added
# Reference: https://github.com/hashicorp/terraform-provider-vsphere/issues/1498

# module "vsphere_tags" {
#   count = (length(var.tags) > 0) ? 1 : 0

#   source = "github.com/aws-ia/terraform-vsphere-tags"
#   # version = "0.0.1"

#   tag_category_name             = var.tag_category_name
#   tag_category_description      = var.tag_category_description
#   tag_category_associable_types = var.tag_category_associable_types
#   tag_category_cardinality      = var.tag_category_cardinality
#   create_tag_category           = var.create_tag_category
#   tags                          = var.tags
#   create_tags                   = var.create_tags
# }

resource "vsphere_content_library" "content_library" {
  count = (var.create_content_library) ? 1 : 0

  name            = var.content_library_name
  description     = var.content_library_description
  storage_backing = [data.vsphere_datastore.datastore.id]

  dynamic "publication" {
    for_each = local.publication

    content {
      authentication_method = publication.value.authentication_method
      username              = publication.value.username
      password              = publication.value.password
      published             = publication.value.published
    }
  }

  dynamic "subscription" {
    for_each = local.subscription

    content {
      subscription_url      = subscription.value.subscription_url
      authentication_method = subscription.value.authentication_method
      username              = subscription.value.username
      password              = subscription.value.password
      automatic_sync        = subscription.value.automatic_sync
      on_demand             = subscription.value.on_demand
    }
  }

  # tags = (length(module.vsphere_tags) == 1) ? module.vsphere_tags[0].tags[*].id : null
}

resource "vsphere_content_library_item" "items" {
  count = (var.create_content_library_items) ? length(var.content_library_items) : 0

  name        = var.content_library_items[count.index]["name"]
  description = var.content_library_items[count.index]["description"]
  file_url    = var.content_library_items[count.index]["file_url"]
  type        = var.content_library_items[count.index]["type"]
  library_id  = (length(vsphere_content_library.content_library) == 1) ? vsphere_content_library.content_library[0].id : data.vsphere_content_library.content_library[0].id

  # tags = (length(module.vsphere_tags) == 1) ? module.vsphere_tags[0].tags[*].id : null
}
