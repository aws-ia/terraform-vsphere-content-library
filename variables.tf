variable "datacenter_name" {
  type        = string
  description = "The name of the vSphere datacenter object where the content library will be created (https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-7FDFBDBE-F8AC-4D00-AE5E-3F14D7472FAF.html)."
  default     = "SDDC-Datacenter"
  nullable    = false

  validation {
    condition     = length(var.datacenter_name) > 0
    error_message = "Must be a string of one or more characters."
  }
}

variable "datastore_name" {
  type        = string
  description = "The name of the vSphere datastore object where the content library items will be stored (https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-storage/GUID-7BED10DD-3EF2-4670-BA7F-0EEB4EC6EB85.html)."
  default     = "WorkloadDatastore"
  nullable    = false

  validation {
    condition     = length(var.datastore_name) > 0
    error_message = "Must be a string of one or more characters."
  }
}

variable "content_library_name" {
  type        = string
  description = "The name of the vSphere content library (https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-254B2CE8-20A8-43F0-90E8-3F6776C2C896.html)."
  default     = "Content library"
  nullable    = false

  validation {
    condition     = length(var.content_library_name) > 0
    error_message = "Must be a string of one or more characters."
  }
}

variable "content_library_description" {
  type        = string
  description = "The description of the vSphere content library."
  default     = null
  nullable    = true
}

variable "create_content_library" {
  type        = bool
  description = "If true, a new vSphere content library will be created; otherwise, the corresponding content library will be imported as a data source."
  default     = false
  nullable    = false
}

variable "content_library_items" {
  type        = list(map(string))
  description = "List of maps of strings defining either OVA/OVF or ISO vSphere content library items (https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vm-administration/GUID-D3DD122F-16A5-4F36-8467-97994A854B16.html). At this time, VM template items are not supported by this module, but can be easily added separately. Each map must have the following keys: 'name', 'description', 'file_url', and 'type'. The value for each 'type' key must be set to either 'ovf' or 'iso'. Last, only the value for 'description' can be empty."
  default     = []
  nullable    = false

  /*
  example = [
    {
      name        = "aws-backup-gateway-virtual-appliance"
      description = "AWS Backup Gateway virtual appliance."
      file_url    = "https://d28e23pnuuv0hr.cloudfront.net/aws-appliance-latest.ova"
      type        = "ovf"
    },
  ]
  */

  validation {
    condition     = alltrue([for t in var.content_library_items : length(keys(t)) == 4])
    error_message = "Must be a list of one or more maps of strings with exactly 4 keys."
  }

  validation {
    condition = alltrue([for t in var.content_library_items : alltrue([for k, v in t : (
      contains(["name", "file_url"], k) &&
      length(v) > 0) ||
      (k == "type" && contains(["ovf", "iso"], v) ||
      k == "description")
    ])])
    error_message = "Must be a list of one or more maps of strings with 'name', 'description', 'file_url', and 'type' keys, the value for the 'type' key must be either 'ovf' or 'iso', and only the value for 'description' is permitted to be empty."
  }
}

variable "create_content_library_items" {
  type        = bool
  description = "If true, new vSphere content library items will be created for each specified; otherwise, the corresponding content library items will be imported as a data source."
  default     = false
  nullable    = false
}

variable "authentication_method" {
  type        = string
  description = "Method for authenticating users if creating a publication/subscription relationship between content libraries."
  default     = null
  nullable    = true

  validation {
    condition     = var.authentication_method == null || try(contains(["NONE", "BASIC"], var.authentication_method), false)
    error_message = "Accepted values: null, 'NONE', or 'BASIC'."
  }
}

variable "password" {
  type        = string
  description = "Password if creating a publication/subscription relationship between content libraries with authentication. Password length and complexity requirements are determined by the configuration in vCenter."
  default     = null
  sensitive   = true
  nullable    = true

  validation {
    condition     = var.password == null || try(length(var.password) >= 0, false)
    error_message = "Must be null or a string of one or more characters."
  }
}

variable "publication_published" {
  type        = bool
  description = "If true, publish the content library if creating a publication/subscription relationship between content libraries."
  default     = null
  nullable    = true
}

variable "subscription_url" {
  type        = string
  description = "URL of the published content library if creating a publication/subscription relationship between content libraries."
  default     = null
  nullable    = true

  validation {
    condition     = var.subscription_url == null || try(length(var.subscription_url) > 0, false)
    error_message = "Must be null or a string of one or more characters."
  }
}

variable "subscription_automatic_sync" {
  type        = bool
  description = "If true, enable automatic synchronization with the published library if creating a publication/subscription relationship between content libraries."
  default     = null
  nullable    = true
}

variable "subscription_on_demand" {
  type        = bool
  description = "If true, download the published content library items only when needed if creating a publication/subscription relationship between content libraries."
  default     = null
  nullable    = true
}

# TODO: Add tags once support for content libraries & content library items are added
# Reference: https://github.com/hashicorp/terraform-provider-vsphere/issues/1498

# variable "tag_category_name" {
#   type        = string
#   description = "The name of the vSphere tag category."
#   default     = "category"
#   nullable    = false

#   validation {
#     condition     = length(var.tag_category_name) > 0
#     error_message = "Must be a string of one or more characters."
#   }
# }

# variable "tag_category_description" {
#   type        = string
#   description = "The description of the vSphere tag category."
#   default     = null
#   nullable    = true
# }

# variable "tag_category_cardinality" {
#   type        = string
#   description = "The number of tags that can be assigned from this category to a single object at once."
#   default     = "MULTIPLE"
#   nullable    = false

#   validation {
#     condition     = contains(["SINGLE", "MULTIPLE"], var.tag_category_cardinality)
#     error_message = "Accepted values: 'SINGLE', 'MULTIPLE'."
#   }
# }

# variable "tag_category_associable_types" {
#   type        = list(string)
#   description = "A list object types that this category to which this category can be assigned (https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag_category#associable-object-types)."
#   default = [
#     "Folder",
#     "ClusterComputeResource",
#     "Datacenter",
#     "Datastore",
#     "StoragePod",
#     "DistributedVirtualPortgroup",
#     "DistributedVirtualSwitch",
#     "VmwareDistributedVirtualSwitch",
#     "HostSystem",
#     "com.vmware.content.Library",
#     "com.vmware.content.library.Item",
#     "HostNetwork",
#     "Network",
#     "OpaqueNetwork",
#     "ResourcePool",
#     "VirtualApp",
#     "VirtualMachine",
#   ]
#   nullable = false

#   validation {
#     condition = alltrue([
#       for t in var.tag_category_associable_types : contains([
#         "Folder",
#         "ClusterComputeResource",
#         "Datacenter",
#         "Datastore",
#         "StoragePod",
#         "DistributedVirtualPortgroup",
#         "DistributedVirtualSwitch",
#         "VmwareDistributedVirtualSwitch",
#         "HostSystem",
#         "com.vmware.content.Library",
#         "com.vmware.content.library.Item",
#         "HostNetwork",
#         "Network",
#         "OpaqueNetwork",
#         "ResourcePool",
#         "VirtualApp",
#         "VirtualMachine",
#       ], t)
#     ])
#     error_message = "Accepted values: 'Folder', 'ClusterComputeResource', 'Datacenter', 'Datastore', 'StoragePod', 'DistributedVirtualPortgroup', 'DistributedVirtualSwitch', 'VmwareDistributedVirtualSwitch', 'HostSystem', 'com.vmware.content.Library', 'com.vmware.content.library.Item', 'HostNetwork', 'Network', 'OpaqueNetwork', 'ResourcePool', 'VirtualApp', 'VirtualMachine'."
#   }
# }

# variable "create_tag_category" {
#   type        = bool
#   description = "If true, a new vSphere tag category will be created."
#   default     = false
#   nullable    = false
# }

# variable "tags" {
#   type        = map(string)
#   description = "Map of strings defining vSphere tag names and descriptions."
#   default     = {}
#   nullable    = false

#   /*
#   example = {
#     terraform = "Managed by Terraform"
#     project   = "terraform-vsphere-tags"
#   }
#   */

#   validation {
#     condition     = length(keys(var.tags)) > 0
#     error_message = "Must be a maps of strings with one or more key/value pairs."
#   }
# }

# variable "create_tags" {
#   type        = bool
#   description = "If true, new vSphere tags will be created for each entry."
#   default     = false
#   nullable    = false
# }
