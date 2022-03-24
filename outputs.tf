output "content_library" {
  description = "The vSphere content library."
  value       = (length(vsphere_content_library.content_library) == 1) ? vsphere_content_library.content_library[0] : data.vsphere_content_library.content_library[0]
}

output "items" {
  description = "The list of vSphere content library items."
  value       = (var.create_content_library_items) ? vsphere_content_library_item.items : data.vsphere_content_library_item.items
}
