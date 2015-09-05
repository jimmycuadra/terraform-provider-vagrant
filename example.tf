resource "vagrant_vm" "example_box" {
  box = "coreos-stable"
}

resource "vagrant_vm" "example_networks" {
  box = "coreos-stable"

  forwarded_port {
    guest = 80
    host = 8080
  }

  private_network {
    ip = "10.0.0.10"
  }

  public_network {}
}

resource "vagrant_vm" "example_virtualbox_provider" {
  box = "coreos-stable"

  virtualbox_provider {}
}
