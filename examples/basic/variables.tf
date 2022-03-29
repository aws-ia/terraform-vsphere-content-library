variable "datacenter_name" {
  type        = string
  description = "The name of the vSphere datacenter object where the content library will be created (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-7FDFBDBE-F8AC-4D00-AE5E-3F14D7472FAF.html)."
  default     = "SDDC-Datacenter"
}

variable "datastore_name" {
  type        = string
  description = "The name of the vSphere datastore object where the content library items will be stored (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.storage.doc/GUID-7BED10DD-3EF2-4670-BA7F-0EEB4EC6EB85.html)."
  default     = "WorkloadDatastore"
}

variable "content_library_name" {
  type        = string
  description = "The name of the vSphere content library (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-254B2CE8-20A8-43F0-90E8-3F6776C2C896.html?hWord=N4IghgNiBcIMYHsB2AXApqgBBAlgIwCcwCBPEAXyA)."
  default     = "example-content-library"
}

variable "content_library_description" {
  type        = string
  description = "The description of the vSphere content library."
  default     = null
}

variable "create_content_library" {
  type        = bool
  description = "If true, a new vSphere content library will be created; otherwise, the corresponding content library will be imported as a data source."
  default     = true
}

variable "content_library_items" {
  type        = list(map(string))
  description = "List of maps of strings defining either OVA/OVF or ISO vSphere content library items (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-D3DD122F-16A5-4F36-8467-97994A854B16.html). At this time, VM template items are not supported by this module, but can be easily added separately. Each map must have the following keys: 'name', 'description', 'file_url', and 'type'. The value for each 'type' key must be set to either 'ovf' or 'iso'. Last, only the value for 'description' can be empty."
  default = [
    {
      name        = "vmware-tools-windows-11.3.0-18"
      description = "VMware Tools for Windows."
      file_url    = "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso"
      type        = "iso"
    },
  ]
}

variable "create_content_library_items" {
  type        = bool
  description = "If true, new vSphere content library items will be created for each specified; otherwise, the corresponding content library items will be imported as a data source."
  default     = true
}

# TODO: Add tags once support for content libraries & content library items are added
# Reference: https://github.com/hashicorp/terraform-provider-vsphere/issues/1498

# variable "tag_category_name" {
#   type        = string
#   description = "The name of the vSphere tag category."
#   default     = "example-category"
# }

# variable "tag_category_description" {
#   type        = string
#   description = "The description of the vSphere tag category."
#   default     = null
# }

# variable "tag_category_cardinality" {
#   type        = string
#   description = "The number of tags that can be assigned from this category to a single object at once."
#   default     = "MULTIPLE"
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
# }

# variable "create_tag_category" {
#   type        = bool
#   description = "If true, a new vSphere tag category will be created."
#   default     = false
#   nullable    = false
# }

# variable "tags" {
#   type        = list(map(string))
#   description = "List of one or more maps of strings defining vSphere tags. Each map must only have 'name' & 'description' keys, and the value for 'name' cannot be empty."
#   default = {
#     terraform = "Managed by Terraform"
#     project = "terraform-vsphere-tags"
#   }
# }

# variable "create_tags" {
#   type        = bool
#   description = "If true, new vSphere tags will be created for each entry."
#   default     = false
# }
