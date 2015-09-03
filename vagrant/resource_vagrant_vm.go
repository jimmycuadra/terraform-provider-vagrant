package vagrant

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVagrantVm() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"boot_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     300,
				Description: "The time in seconds that Vagrant will wait for the machine to boot.",
			},
			"box": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Installed box name or Atlas box shorthand name to use.",
			},
			"box_check_update": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not to check for updates to the box on each `vagrant up`.",
			},
			"box_download_checksum": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The expected checksum of the box file.",
			},
			"box_download_checksum_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of checksum specified by `box_download_checksum`. Should be `md5`, `sha1`, or `sha256`.",
			},
			"box_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URL to download the box.",
			},
			"box_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the box to use.",
			},
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hostname the machine should have.",
			},
		},
	}
}
