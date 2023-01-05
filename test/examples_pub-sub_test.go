package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

const (
	content_library_password = "VSPHERE_CONTENT_LIBRARY_PASSWORD"
)

func TestExamplesPubSub(t *testing.T) {
	GetEnvsOrExit()
	GetEnvOrExit(content_library_password)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/pub-sub",
		// Vars: map[string]interface{}{
		// 	"datacenter_name":              "SDDC-Datacenter",
		// 	"datastore_name":               "WorkloadDatastore",
		// 	"content_library_name":         "example-content-library",
		// 	"content_library_description":  "",
		// 	"create_content_library":       true,
		// 	"create_content_library_items": true,
		// 	"authentication_method":        "BASIC",
		// 	"password":                     os.Getenv(content_library_password),
		// 	"publication_published":        true,
		// 	"subscription_url":             "http://localhost/lib.json",
		// 	"subscription_automatic_sync":  true,
		// 	"subscription_on_demand":       false,
		// 	// "tag_category_name":            "terraform",
		// 	// "tags":                         []interface{}{map[string]interface{}{"name": "terraform", "description": ""}},
		// 	// "create_tag_category":          false,
		// 	// "create_tags":                  false,
		// 	"content_library_items": []interface{}{map[string]interface{}{
		// 		"name":        "vmware-tools-windows-12.0.6-20",
		// 		"description": "VMware Tools for Windows.",
		// 		"file_url":    "https://packages.vmware.com/tools/esx/8.0/windows/VMware-tools-windows-12.0.6-20104755.iso",
		// 		"type":        "iso",
		// 	}},
		// },
		Vars: map[string]interface{}{
			"password": os.Getenv(content_library_password),
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}
