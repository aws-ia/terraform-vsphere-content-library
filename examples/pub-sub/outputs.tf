output "publisher" {
  description = "The publisher vSphere content library."
  value       = module.publisher.content_library
  sensitive   = true
}

output "subscriber" {
  description = "The subscriber vSphere content library."
  value       = module.subscriber.content_library
  sensitive   = true
}
