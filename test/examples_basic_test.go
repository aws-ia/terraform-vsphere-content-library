package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

const (
	user     = "VSPHERE_USER"
	password = "VSPHERE_PASSWORD"
	server   = "VSPHERE_SERVER"
	// allow_unverified_ssl = "VSPHERE_ALLOW_UNVERIFIED_SSL"

	existing_content_library_name         = "VSPHERE_EXISTING_CONTENT_LIBRARY_NAME"
	existing_content_library_item_name    = "VSPHERE_EXISTING_CONTENT_LIBRARY_ITEM_NAME"
	existing_content_library_item_fileurl = "VSPHERE_EXISTING_CONTENT_LIBRARY_ITEM_FILEURL"
	existing_content_library_item_type    = "VSPHERE_EXISTING_CONTENT_LIBRARY_ITEM_TYPE"

	input_validation_test_failed_message = "Invalid '%s' value input validation test failed."
)

func GetEnvOrExit(env_var string) {
	if os.Getenv(env_var) == "" {
		fmt.Println("Environment variable: '" + env_var + "' is not set.")
		os.Exit(1)
	}
}

func GetEnvsOrExit() {
	GetEnvOrExit(user)
	GetEnvOrExit(password)
	GetEnvOrExit(server)
}

func TestExamplesNewContentLibraryNoItems(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		// Vars: map[string]interface{}{
		// 	"datacenter_name":              "SDDC-Datacenter",
		// 	"datastore_name":               "WorkloadDatastore",
		// 	"content_library_name":         "example-content-library",
		// 	"content_library_description":  "",
		// 	"create_content_library":       true,
		// 	"create_content_library_items": true,
		// 	// "tag_category_name":            "terraform",
		// 	// "tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
		// 	// "create_tag_category":          false,
		// 	// "create_tags":                  false,
		// 	"content_library_items": []interface{}{map[string]interface{}{
		// 		"name":        "vmware-tools-windows-11.3.0-18",
		// 		"description": "VMware Tools for Windows.",
		// 		"file_url":    "https://packages.vmware.com/tools/esx/7.0u3/windows/VMware-tools-windows-11.3.0-18090558.iso",
		// 		"type":        "iso",
		// 	}},
		// },
		Vars: map[string]interface{}{
			"create_content_library": true,
			"content_library_items":  []interface{}{},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesNewContentLibraryNewItem(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"create_content_library":       true,
			"create_content_library_items": true,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportContentLibraryImportItem(t *testing.T) {
	GetEnvsOrExit()
	GetEnvOrExit(existing_content_library_name)
	GetEnvOrExit(existing_content_library_item_name)
	GetEnvOrExit(existing_content_library_item_fileurl)
	GetEnvOrExit(existing_content_library_item_type)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_name":         os.Getenv(existing_content_library_name),
			"create_content_library":       false,
			"create_content_library_items": false,
			"content_library_items": []interface{}{map[string]interface{}{
				"name":        os.Getenv(existing_content_library_item_name),
				"description": "",
				"file_url":    os.Getenv(existing_content_library_item_fileurl),
				"type":        os.Getenv(existing_content_library_item_type),
			}},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesImportContentLibraryNewItem(t *testing.T) {
	GetEnvsOrExit()
	GetEnvOrExit(existing_content_library_name)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_name":         os.Getenv(existing_content_library_name),
			"create_content_library":       false,
			"create_content_library_items": true,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestExamplesNewContentLibraryImportItem(t *testing.T) {
	// Create library and skip importing library items since they don't exist. Intentionally passes.
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"datacenter_name":              "SDDC-Datacenter",
			"datastore_name":               "WorkloadDatastore",
			"content_library_name":         "example-content-library",
			"content_library_description":  "",
			"create_content_library":       true,
			"create_content_library_items": false,
			// "tag_category_name":            "terraform",
			// "tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
			// "create_tag_category":          false,
			// "create_tags":                  false,
			"content_library_items": []interface{}{map[string]interface{}{
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
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"datacenter_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "datacenter_name"))
	}
}

func TestExamplesInvalidDatastoreName(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"datastore_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "datastore_name"))
	}
}

func TestExamplesInvalidContentLibraryName(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_name": "",
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "content_library_name"))
	}
}

func TestExamplesInvalidContentLibraryItemsNumKeys(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_items": []interface{}{map[string]interface{}{
				"name": "test",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsNameKey(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_items": []interface{}{map[string]interface{}{
				"name":        "",
				"description": "",
				"file_url":    "a",
				"type":        "iso",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsFileUrlKey(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_items": []interface{}{map[string]interface{}{
				"name":        "a",
				"description": "",
				"file_url":    "",
				"type":        "iso",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsTypeKey1(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_items": []interface{}{map[string]interface{}{
				"name":        "a",
				"description": "",
				"file_url":    "a",
				"type":        "",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "content_library_items"))
	}
}

func TestExamplesInvalidContentLibraryItemsTypeKey2(t *testing.T) {
	GetEnvsOrExit()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic",
		Vars: map[string]interface{}{
			"content_library_items": []interface{}{map[string]interface{}{
				"name":        "a",
				"description": "",
				"file_url":    "a",
				"type":        "ova",
			}},
		},
	}

	if _, err := terraform.ApplyE(t, terraformOptions); err == nil {
		defer terraform.Destroy(t, terraformOptions)
		require.Error(t, err, fmt.Sprintf(input_validation_test_failed_message, "content_library_items"))
	}
}
