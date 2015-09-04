package vagrant

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVagrantVm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVagrantVmCreate,
		Read:   resourceVagrantVmRead,
		Update: resourceVagrantVmUpdate,
		Delete: resourceVagrantVmDelete,

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
			"networks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of network, one of `forwarded_port`, `private_network`, or `public_network`.",
						},

						// forwarded_port
						"guest": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The port on the guest you want to be exposed on the host.",
						},
						"guest_ip": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The IP the forwarded port should be bound to within the guest.",
						},
						"host": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The port on the host used to access the port on the guest.",
						},
						"host_ip": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The IP the forwarded port should be bound to on the host.",
						},
						"protocol": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The network protcol to use, one of `tcp` or `udp`.",
						},
						"auto_correct": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to automatically resolve port collisions on the host.",
						},

						// private_network and public_network
						"ip": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The static IP to assign the interface on the host. Do not set this to use DHCP instead.",
						},
					},
				},
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func resourceVagrantVmCreate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVagrantVmRead(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVagrantVmUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVagrantVmDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}
