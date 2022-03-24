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

variable "password" {
  type        = string
  description = "Password if creating a publication/subscription relationship between content libraries with authentication."
  sensitive   = true
}
