<!-- BEGIN_TF_DOCS -->
# Basic example

If deployed with the default values, this example will create a [content library][content\_library] named `example-content-library` backed by a [datastore] named `WorkloadDatastore`, and one ISO [item][items] named `vmware-tools-windows-11.3.0-18` in your [VMware Cloud on AWS][vmconaws] or [VMware vSphere][vsphere] on-premises environment.

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

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_vsphere_content_library"></a> [vsphere\_content\_library](#module\_vsphere\_content\_library) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_content_library_description"></a> [content\_library\_description](#input\_content\_library\_description) | The description of the vSphere content library. | `string` | `null` | no |
| <a name="input_content_library_items"></a> [content\_library\_items](#input\_content\_library\_items) | List of maps of strings defining either OVA/OVF or ISO vSphere content library items (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-D3DD122F-16A5-4F36-8467-97994A854B16.html). At this time, VM template items are not supported by this module, but can be easily added separately. Each map must have the following keys: 'name', 'description', 'file\_url', and 'type'. The value for each 'type' key must be set to either 'ovf' or 'iso'. Last, only the value for 'description' can be empty. | `list(map(string))` | <pre>[<br>  {<br>    "description": "VMware Tools for Windows.",<br>    "file_url": "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso",<br>    "name": "vmware-tools-windows-11.3.0-18",<br>    "type": "iso"<br>  }<br>]</pre> | no |
| <a name="input_content_library_name"></a> [content\_library\_name](#input\_content\_library\_name) | The name of the vSphere content library (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-254B2CE8-20A8-43F0-90E8-3F6776C2C896.html?hWord=N4IghgNiBcIMYHsB2AXApqgBBAlgIwCcwCBPEAXyA). | `string` | `"example-content-library"` | no |
| <a name="input_create_content_library"></a> [create\_content\_library](#input\_create\_content\_library) | If true, a new vSphere content library will be created; otherwise, the corresponding content library will be imported as a data source. | `bool` | `true` | no |
| <a name="input_create_content_library_items"></a> [create\_content\_library\_items](#input\_create\_content\_library\_items) | If true, new vSphere content library items will be created for each specified; otherwise, the corresponding content library items will be imported as a data source. | `bool` | `true` | no |
| <a name="input_datacenter_name"></a> [datacenter\_name](#input\_datacenter\_name) | The name of the vSphere datacenter object where the content library will be created (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-7FDFBDBE-F8AC-4D00-AE5E-3F14D7472FAF.html). | `string` | `"SDDC-Datacenter"` | no |
| <a name="input_datastore_name"></a> [datastore\_name](#input\_datastore\_name) | The name of the vSphere datastore object where the content library items will be stored (https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.storage.doc/GUID-7BED10DD-3EF2-4670-BA7F-0EEB4EC6EB85.html). | `string` | `"WorkloadDatastore"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_content_library"></a> [content\_library](#output\_content\_library) | The vSphere content library. |
| <a name="output_items"></a> [items](#output\_items) | The list of vSphere content library items. |
<!-- END_TF_DOCS -->