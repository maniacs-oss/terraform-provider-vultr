package vultr

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVultrReservedIP(t *testing.T) {
	rLabel := acctest.RandomWithPrefix("tf-rip-ds")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVultrReservedIPRead(rLabel),
			},
			{
				Config: testAccVultrReservedIPRead(rLabel),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vultr_reserved_ip.foo", "id"),
					resource.TestCheckResourceAttrSet("data.vultr_reserved_ip.foo", "subnet"),
					resource.TestCheckResourceAttrSet("data.vultr_reserved_ip.foo", "region"),
					resource.TestCheckResourceAttrSet("data.vultr_reserved_ip.foo", "label"),
					resource.TestCheckResourceAttrSet("data.vultr_reserved_ip.foo", "ip_type"),
				),
			},
		},
	})
}

func testAccVultrReservedIPRead(label string) string {
	return fmt.Sprintf(`
		resource "vultr_reserved_ip" "bar" {
		label = "%s"
		region = "sea"
		ip_type = "v4"
	}

		data "vultr_reserved_ip" "foo" {
			filter {
				name = "label"
				values = ["${vultr_reserved_ip.bar.label}"]
			}
		}
		`, label)
}
