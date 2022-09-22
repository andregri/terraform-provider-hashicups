package hashicups

import (
	"fmt"
	"testing"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccHashicupsOrder_basic(t *testing.T) {
	coffeeID := "4"
	quantity := "7"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHashicupsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckHashicupsOrderConfig_basic(coffeeID, quantity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHashicupsOrderExists("hashicups_order.new"),
					resource.TestCheckResourceAttr("hashicups_order.new", "items.0.coffee.0.id", "4"),
					resource.TestCheckResourceAttr("hashicups_order.new", "items.0.quantity", "7"),
				),
			},
		},
	})
}

// testAccCheckHashicupsResourceDestroy verifies the Order
// has been destroyed
func testAccCheckHashicupsResourceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	c := testAccProvider.Meta().(*hc.Client)

	// loop through the resources in state, verifying each order
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "hashicups_order" {
			continue
		}

		orderID := rs.Primary.ID

		_, err := c.GetOrder(orderID, &c.Token)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}

func testAccCheckHashicupsOrderConfig_basic(coffeeID, quantity string) string {
	return fmt.Sprintf(`
	resource "hashicups_order" "new" {
		items {
			coffee {
				id = %s
			}
    		quantity = %s
  		}
	}
	`, coffeeID, quantity)
}

func testAccCheckHashicupsOrderExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Order ID is not set")
		}

		return nil
	}
}
