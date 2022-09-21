package hashicups

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"hashicups": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if ip := os.Getenv("HASHICUPS_IP"); ip == "" {
		t.Fatal("HASHICUPS_IP must be set for acceptance tests")
	}
	if err := os.Getenv("HASHICUPS_USERNAME"); err == "" {
		t.Fatal("HASHICUPS_USERNAME must be set for acceptance tests")
	}
	if err := os.Getenv("HASHICUPS_PASSWORD"); err == "" {
		t.Fatal("HASHICUPS_PASSWORD must be set for acceptance tests")
	}
}
