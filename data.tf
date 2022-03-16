data "vsphere_datacenter" "datacenter" {
  name = var.vsphere_datacenter_name
}

data "vsphere_datastore" "datastore" {
  name          = var.vsphere_datastore_name
  datacenter_id = data.vsphere_datacenter.datacenter.id
}

data "vsphere_content_library" "content_library" {
  count = var.create_vsphere_content_library ? 0 : 1

  name = var.vsphere_content_library_name
}

data "vsphere_content_library_item" "items" {
  count = length(data.vsphere_content_library.content_library) == 1 && !var.create_vsphere_content_library_items ? length(var.vsphere_content_library_items) : 0

  name       = var.vsphere_content_library_items[count.index]["name"]
  type       = var.vsphere_content_library_items[count.index]["type"]
  library_id = data.vsphere_content_library.content_library[0].id
}
