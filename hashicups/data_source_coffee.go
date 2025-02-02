package hashicups

import (
	"context"
	"strconv"
	"time"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
$ curl localhost:19090/coffees
[
 {
   "id": 1,
   "name": "Packer Spiced Latte",
   "teaser": "Packed with goodness to spice up your images",
   "description": "",
   "price": 350,
   "image": "/packer.png",
   "ingredients": [
     { "ingredient_id": 1 },
     { "ingredient_id": 2 },
     { "ingredient_id": 4 }
   ]
 },
 ## ...
]
*/

func dataSourceCoffees() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCoffeesRead,
		Schema: map[string]*schema.Schema{
			"coffees": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"teaser": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"price": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"image": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ingredients": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ingredient_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceCoffeesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	coffees, err := c.GetCoffees()
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Read coffees data source", map[string]interface{}{
		"amount": len(coffees),
	})

	flattenedCoffees := flattenCoffeesData(&coffees)
	if err := d.Set("coffees", flattenedCoffees); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenCoffeesData(coffees *[]hc.Coffee) []interface{} {
	if coffees != nil {
		cs := make([]interface{}, len(*coffees))

		for i, coffee := range *coffees {
			cs[i] = flattenCoffee(coffee)[0]
		}

		return cs
	}

	return make([]interface{}, 0)
}
