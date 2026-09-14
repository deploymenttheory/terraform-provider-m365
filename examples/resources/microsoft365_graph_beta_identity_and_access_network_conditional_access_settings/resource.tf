# No default is applied. Omit signaling_status to preserve the existing setting.
# Review Conditional Access, CAE and Identity Protection impact before changing it.
resource "microsoft365_graph_beta_identity_and_access_network_conditional_access_settings" "example" {
  signaling_status = "enabled"
}

# Destroy removes only Terraform state; it does not disable or reset signaling.
