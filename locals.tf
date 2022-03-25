locals {
  publication = (var.create_content_library && var.authentication_method != null && var.subscription_url == null) ? [
    {
      authentication_method = var.authentication_method
      username              = (var.authentication_method == "NONE") ? null : "vcsp"
      password              = (var.authentication_method == "NONE") ? null : var.password
      published             = var.publication_published
    },
  ] : []

  subscription = (var.create_content_library && length(local.publication) == 0 && var.subscription_url != null) ? [
    {
      subscription_url      = var.subscription_url
      authentication_method = var.authentication_method
      username              = (var.authentication_method == "NONE") ? null : "vcsp"
      password              = (var.authentication_method == "NONE") ? null : var.password
      automatic_sync        = var.subscription_automatic_sync
      on_demand             = var.subscription_on_demand
    },
  ] : []
}
