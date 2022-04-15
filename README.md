<!-- BEGIN_TF_DOCS -->
# vSphere Content Library Terraform Module

This Terraform module creates or imports a [datastore]-backed [content library][content\_library] in your [VMware Cloud on AWS][vmconaws] or [VMware vSphere][vsphere] on&#8209;premises environment. You can configure the content library as one of three types: standalone, publisher, or subscriber.

You can optionally specify a list of new [items], such as OVA and ISO images, to deploy into the new or existing content library or a list of items to import from an existing content library.

## Usage

### Create a new content library

```hcl
module "vsphere_content_library" {
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  datacenter_name              = "SDDC-Datacenter"
  datastore_name               = "WorkloadDatastore"
  content_library_name         = "example-content-library"
  content_library_description  = "Example content library."
  create_content_library       = true
}
```

### Create a new content library with an item

```hcl
module "vsphere_content_library" {
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  datacenter_name              = "SDDC-Datacenter"
  datastore_name               = "WorkloadDatastore"
  content_library_name         = "example-content-library"
  content_library_description  = "Example content library."
  create_content_library       = true
  create_content_library_items = true
  content_library_items        = [
    {
      name        = "vmware-tools-windows-11.3.0-18"
      description = "VMware Tools for Windows."
      file_url    = "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso"
      type        = "iso"
    },
  ]
}
```

### Create a new item in an existing content library

```hcl
module "vsphere_content_library" {
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  content_library_name         = "example-content-library"
  create_content_library       = false
  create_content_library_items = true
  content_library_items        = [
    {
      name        = "vmware-tools-windows-11.3.0-18"
      description = "VMware Tools for Windows."
      file_url    = "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso"
      type        = "iso"
    },
  ]
}
```

### Import existing items into an existing content library

```hcl
module "vsphere_content_library" {
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  content_library_name         = "example-content-library"
  create_content_library       = false
  create_content_library_items = false
  content_library_items        = [
    {
      name        = "vmware-tools-windows-11.3.0-18"
      description = ""
      file_url    = "n/a"
      type        = "iso"
    },
  ]
}
```

### Create a publisher content library

```hcl
module "vsphere_content_library" {
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  content_library_name   = "example-publisher"
  create_content_library = true

  authentication_method = "BASIC"
  password              = var.password
  publication_published = true
}
```

### Create a subscriber content library

```hcl
module "vsphere_content_library" {
  source  = "aws-ia/vsphere-content-library/vsphere"
  version = ">= 0.0.1"

  content_library_name   = "example-subscriber"
  create_content_library = true

  subscription_url            = "https://vcenter.sddc-127-0-0-1.vmwarevmc.com:443/cls/vcsp/lib/60d8b31e-c6d0-47d7-9758-ff0f4ff5af14/lib.json"
  authentication_method       = "BASIC"
  password                    = var.password
  subscription_automatic_sync = true
  subscription_on_demand      = false
}
```

[content\_library]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-254B2CE8-20A8-43F0-90E8-3F6776C2C896.html?hWord=N4IghgNiBcIMYHsB2AXApqgBBAlgIwCcwCBPEAXyA
[datastore]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.storage.doc/GUID-7BED10DD-3EF2-4670-BA7F-0EEB4EC6EB85.html
[items]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-D3DD122F-16A5-4F36-8467-97994A854B16.html
[vsphere]: https://docs.vmware.com/en/VMware-vSphere/index.html
[vmconaws]: https://aws.amazon.com/vmware/

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.1.0 |
| <a name="requirement_vsphere"></a> [vsphere](#requirement\_vsphere) | >= 2.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_vsphere"></a> [vsphere](#provider\_vsphere) | >= 2.1.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [vsphere_content_library.content_library](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/content_library) | resource |
| [vsphere_content_library_item.items](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/content_library_item) | resource |
| [vsphere_content_library.content_library](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/data-sources/content_library) | data source |
| [vsphere_content_library_item.items](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/data-sources/content_library_item) | data source |
| [vsphere_datacenter.datacenter](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/data-sources/datacenter) | data source |
| [vsphere_datastore.datastore](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/data-sources/datastore) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_authentication_method"></a> [authentication\_method](#input\_authentication\_method) | Method for authenticating users if creating a publication/subscription relationship between content libraries. | `string` | `null` | no |
| <a name="input_content_library_description"></a> [content\_library\_description](#input\_content\_library\_description) | The description of the vSphere content library. | `string` | `null` | no |
| <a name="input_content_library_items"></a> [content\_library\_items](#input\_content\_library\_items) | List of maps of strings defining either OVA/OVF or ISO vSphere content library items (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-D3DD122F-16A5-4F36-8467-97994A854B16.html). At this time, VM template items are not supported by this module, but can be easily added separately. Each map must have the following keys: 'name', 'description', 'file\_url', and 'type'. The value for each 'type' key must be set to either 'ovf' or 'iso'. Last, only the value for 'description' can be empty. | `list(map(string))` | `[]` | no |
| <a name="input_content_library_name"></a> [content\_library\_name](#input\_content\_library\_name) | The name of the vSphere content library (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-254B2CE8-20A8-43F0-90E8-3F6776C2C896.html?hWord=N4IghgNiBcIMYHsB2AXApqgBBAlgIwCcwCBPEAXyA). | `string` | `"Content library"` | no |
| <a name="input_create_content_library"></a> [create\_content\_library](#input\_create\_content\_library) | If true, a new vSphere content library will be created; otherwise, the corresponding content library will be imported as a data source. | `bool` | `false` | no |
| <a name="input_create_content_library_items"></a> [create\_content\_library\_items](#input\_create\_content\_library\_items) | If true, new vSphere content library items will be created for each specified; otherwise, the corresponding content library items will be imported as a data source. | `bool` | `false` | no |
| <a name="input_datacenter_name"></a> [datacenter\_name](#input\_datacenter\_name) | The name of the vSphere datacenter object where the content library will be created (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-7FDFBDBE-F8AC-4D00-AE5E-3F14D7472FAF.html). | `string` | `"SDDC-Datacenter"` | no |
| <a name="input_datastore_name"></a> [datastore\_name](#input\_datastore\_name) | The name of the vSphere datastore object where the content library items will be stored (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.storage.doc/GUID-7BED10DD-3EF2-4670-BA7F-0EEB4EC6EB85.html). | `string` | `"WorkloadDatastore"` | no |
| <a name="input_password"></a> [password](#input\_password) | Password if creating a publication/subscription relationship between content libraries with authentication. Password length and complexity requirements are determined by the configuration in vCenter. | `string` | `null` | no |
| <a name="input_publication_published"></a> [publication\_published](#input\_publication\_published) | If true, publish the content library if creating a publication/subscription relationship between content libraries. | `bool` | `null` | no |
| <a name="input_subscription_automatic_sync"></a> [subscription\_automatic\_sync](#input\_subscription\_automatic\_sync) | If true, enable automatic synchronization with the published library if creating a publication/subscription relationship between content libraries. | `bool` | `null` | no |
| <a name="input_subscription_on_demand"></a> [subscription\_on\_demand](#input\_subscription\_on\_demand) | If true, download the published content library items only when needed if creating a publication/subscription relationship between content libraries. | `bool` | `null` | no |
| <a name="input_subscription_url"></a> [subscription\_url](#input\_subscription\_url) | URL of the published content library if creating a publication/subscription relationship between content libraries. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_content_library"></a> [content\_library](#output\_content\_library) | The vSphere content library. |
| <a name="output_items"></a> [items](#output\_items) | The list of vSphere content library items. |
<!-- END_TF_DOCS -->