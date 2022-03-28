output "content_library" {
  description = "The vSphere content library."
  value       = module.vsphere_content_library.content_library
  sensitive   = true
}

output "items" {
  description = "The list of vSphere content library items."
  value       = module.vsphere_content_library.items
}
