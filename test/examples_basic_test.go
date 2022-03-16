package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

// You must set these environment variables for this test
const (
	vsphere_user     = "VSPHERE_USER"
	vsphere_password = "VSPHERE_PASSWORD"
	vsphere_server   = "VSPHERE_SERVER"
	// vsphere_allow_unverified_ssl = "VSPHERE_ALLOW_UNVERIFIED_SSL"

	input_validation_test_failed_message = "Invalid '%s' value input validation test failed."
)

func GetEnvOrExit(env_var string) {
	if os.Getenv(env_var) == "" {
		fmt.Println("Environment variable: '" + env_var + "' is not set.")
		os.Exit(1)
	}
}

func TestExamplesNewContentLibraryNewItem(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_datacenter_name":              "SDDC-Datacenter",
			"vsphere_datastore_name":               "WorkloadDatastore",
			"vsphere_content_library_name":         "example-content-library",
			"vsphere_content_library_description":  "",
			"create_vsphere_content_library":       true,
			"create_vsphere_content_library_items": true,
			// "vsphere_tag_category_name":            "terraform",
			// "vsphere_tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
			// "create_vsphere_tag_category":          false,
			// "create_vsphere_tags":                  false,
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "vmware-tools-windows-11.3.0-18",
				"description": "VMware Tools for Windows.",
				"file_url":    "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso",
				"type":        "iso",
			}},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportContentLibraryImportItem(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_datacenter_name":              "SDDC-Datacenter",
			"vsphere_datastore_name":               "WorkloadDatastore",
			"vsphere_content_library_name":         "Content library",
			"vsphere_content_library_description":  "",
			"create_vsphere_content_library":       false,
			"create_vsphere_content_library_items": false,
			// "vsphere_tag_category_name":            "terraform",
			// "vsphere_tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
			// "create_vsphere_tag_category":          false,
			// "create_vsphere_tags":                  false,
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "aws-appliance-latest.ova",
				"description": "",
				"file_url":    "https://d28e23pnuuv0hr.cloudfront.net/aws-appliance-latest.ova",
				"type":        "ovf",
			}},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportContentLibraryNewItem(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_datacenter_name":              "SDDC-Datacenter",
			"vsphere_datastore_name":               "WorkloadDatastore",
			"vsphere_content_library_name":         "Content library",
			"vsphere_content_library_description":  "",
			"create_vsphere_content_library":       false,
			"create_vsphere_content_library_items": true,
			// "vsphere_tag_category_name":            "terraform",
			// "vsphere_tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
			// "create_vsphere_tag_category":          false,
			// "create_vsphere_tags":                  false,
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "vmware-tools-windows-11.3.0-18",
				"description": "VMware Tools for Windows.",
				"file_url":    "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso",
				"type":        "iso",
			}},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesNewContentLibraryImportItem(t *testing.T) {
	// Create library and skip importing library items since they don't exist. Intentionally passes.
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_datacenter_name":              "SDDC-Datacenter",
			"vsphere_datastore_name":               "WorkloadDatastore",
			"vsphere_content_library_name":         "example-content-library",
			"vsphere_content_library_description":  "",
			"create_vsphere_content_library":       true,
			"create_vsphere_content_library_items": false,
			// "vsphere_tag_category_name":            "terraform",
			// "vsphere_tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
			// "create_vsphere_tag_category":          false,
			// "create_vsphere_tags":                  false,
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "vmware-tools-windows-11.3.0-18",
				"description": "VMware Tools for Windows.",
				"file_url":    "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso",
				"type":        "iso",
			}},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesInvalidDatacenterName(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_datacenter_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_datacenter_name"))
	}
}

func TestExamplesInvalidDatastoreName(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_datastore_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_datastore_name"))
	}
}

func TestExamplesInvalidContentLibraryName(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_content_library_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_content_library_name"))
	}
}

func TestExamplesInvalidContentLibraryItemsNumKeys(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name": "test",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsNameKey(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "",
				"description": "",
				"file_url":    "a",
				"type":        "iso",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsFileUrlKey(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "a",
				"description": "",
				"file_url":    "",
				"type":        "iso",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsTypeKey1(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "a",
				"description": "",
				"file_url":    "a",
				"type":        "",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsTypeKey2(t *testing.T) {
	GetEnvOrExit(vsphere_user)
	GetEnvOrExit(vsphere_password)
	GetEnvOrExit(vsphere_server)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"vsphere_content_library_items": []interface{}{map[string]interface{}{
				"name":        "a",
				"description": "",
				"file_url":    "a",
				"type":        "ova",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "vsphere_content_library_items"))
	}
}
