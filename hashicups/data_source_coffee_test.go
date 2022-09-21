package hashicups

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHashicupsCoffee_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckNoAuth(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckHashicupsCoffeeConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.hashicups_coffees.all", "coffees.#", "6"),
				),
			},
		},
	})
}

func testAccCheckHashicupsCoffeeConfig_basic() string {
	return "data \"hashicups_coffees\" \"all\" {}"
}
