package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jimmycuadra/terraform-provider-vagrant/vagrant"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vagrant.Provider,
	})
}
