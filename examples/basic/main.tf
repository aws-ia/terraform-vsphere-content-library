module "vsphere_content_library" {
  # source = "../.."
  source  = "aws-ia/content-library/vsphere"
  version = ">= 0.0.1"

  datacenter_name = var.datacenter_name
  datastore_name  = var.datastore_name

  content_library_name        = var.content_library_name
  content_library_description = var.content_library_description
  create_content_library      = var.create_content_library

  content_library_items        = var.content_library_items
  create_content_library_items = var.create_content_library_items

  # TODO: Add tags once support for content libraries & content library items are added
  # Reference: https://github.com/hashicorp/terraform-provider-vsphere/issues/1498

  # tag_category_name             = var.tag_category_name
  # tag_category_description      = var.tag_category_description
  # tag_category_associable_types = var.tag_category_associable_types
  # tag_category_cardinality      = var.tag_category_cardinality
  # create_tag_category           = var.create_tag_category
  # tags                          = var.tags
  # create_tags                   = var.create_tags
}
