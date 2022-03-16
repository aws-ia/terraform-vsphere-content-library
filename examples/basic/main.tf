module "vsphere_content_library" {
  source = "../.."

  vsphere_datacenter_name = var.vsphere_datacenter_name
  vsphere_datastore_name  = var.vsphere_datastore_name

  vsphere_content_library_name        = var.vsphere_content_library_name
  vsphere_content_library_description = var.vsphere_content_library_description
  create_vsphere_content_library      = var.create_vsphere_content_library

  vsphere_content_library_items        = var.vsphere_content_library_items
  create_vsphere_content_library_items = var.create_vsphere_content_library_items

  # TODO: Add tags once support for content libraries & content library items are added
  # Reference: https://github.com/hashicorp/terraform-provider-vsphere/issues/1498

  # vsphere_tag_category_name             = var.vsphere_tag_category_name
  # vsphere_tag_category_description      = var.vsphere_tag_category_description
  # vsphere_tag_category_associable_types = var.vsphere_tag_category_associable_types
  # vsphere_tag_category_cardinality      = var.vsphere_tag_category_cardinality
  # create_vsphere_tag_category           = var.create_vsphere_tag_category
  # vsphere_tags                          = var.vsphere_tags
  # create_vsphere_tags                   = var.create_vsphere_tags
}
