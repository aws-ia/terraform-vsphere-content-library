module "publisher" {
  # source = "../.."
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  datacenter_name = var.datacenter_name
  datastore_name  = var.datastore_name

  content_library_name   = "example-publisher"
  create_content_library = true

  authentication_method = "BASIC"
  password              = var.password
  publication_published = true
}

module "subscriber" {
  # source = "../.."
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  datacenter_name = var.datacenter_name
  datastore_name  = var.datastore_name

  content_library_name   = "example-subscriber"
  create_content_library = true

  subscription_url            = module.publisher.content_library.publication[0].publish_url
  authentication_method       = module.publisher.content_library.publication[0].authentication_method
  password                    = module.publisher.content_library.publication[0].password
  subscription_automatic_sync = true
  subscription_on_demand      = true
}
