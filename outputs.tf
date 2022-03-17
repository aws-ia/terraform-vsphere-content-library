output "content_library" {
  description = "The vSphere content library."
  value       = (var.create_content_library) ? vsphere_content_library.content_library : data.vsphere_content_library.content_library
}

output "items" {
  description = "The list of vSphere content library items."
  value       = (var.create_content_library_items) ? vsphere_content_library_item.items : data.vsphere_content_library_item.items
}
