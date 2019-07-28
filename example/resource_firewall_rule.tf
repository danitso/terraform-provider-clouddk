resource "clouddk_firewall_rule" "example" {
  network_interface_id = "${element(flatten(clouddk_server.example.network_interface_ids), 0)}"
  server_id            = "${clouddk_server.example.id}"

  command  = "ACCEPT"
  protocol = "TCP"
  address  = "8.8.8.8/32"
  port     = 8080
}
