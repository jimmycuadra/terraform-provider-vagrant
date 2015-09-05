package vagrant

import (
	"fmt"

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
				Computed:    true,
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
				Computed:    true,
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
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					valid_values := []string{"md5", "sha1", "sha256"}

					for _, valid_value := range valid_values {
						if value == valid_value {
							return
						}
					}

					errors = append(errors, fmt.Errorf("%q must be one of md5, sha1, or sha256", k))

					return
				},
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
			"forwarded_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_correct": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether or not to automatically resolve port collisions on the host.",
						},
						"guest": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The port on the guest you want to be exposed on the host.",
						},
						"guest_ip": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The IP the forwarded port should be bound to within the guest.",
						},
						"host": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The port on the host used to access the port on the guest.",
						},
						"host_ip": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The IP the forwarded port should be bound to on the host.",
						},
						"protocol": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The network protcol to use, one of `tcp` or `udp`.",
						},
					},
				},
				Set: func(v interface{}) int {
					return 0
				},
			},
			"private_network": resourceVagrantVmPublicOrPrivateNetwork(),
			"public_network":  resourceVagrantVmPublicOrPrivateNetwork(),
			"virtualbox_provider": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The number of CPUs to give the VM.",
						},
						"customize": &schema.Schema{
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "A call to customize the VM just before it is booted.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"arguments": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "The list of arguments for the customization command.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
							Set: func(v interface{}) int {
								return 0
							},
						},
						"gui": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether or not to create a VM with a GUI.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The amount of memory to give the VM, in MB.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A custom name for the VM within VirtualBox's GUI app.",
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

func resourceVagrantVmPublicOrPrivateNetwork() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"auto_config": &schema.Schema{
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Whether or not to automatically configure the network interface.",
				},
				"dhcp": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Description: "Whether or not to use DHCP to assign the IP on the host.",
				},
				"ip": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					Description: "The static IP to assign the interface on the host. Do not set this to use DHCP instead.",
				},
			},
		},
		Set: func(v interface{}) int {
			return 0
		},
	}
}
