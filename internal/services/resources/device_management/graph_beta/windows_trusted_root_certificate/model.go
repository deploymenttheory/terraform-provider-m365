package graphBetaWindowsTrustedRootCertificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WindowsTrustedRootCertificateResourceModel struct {
	ID                     types.String   `tfsdk:"id"`
	DisplayName            types.String   `tfsdk:"display_name"`
	Description            types.String   `tfsdk:"description"`
	CertFileName           types.String   `tfsdk:"cert_file_name"`
	TrustedRootCertificate types.String   `tfsdk:"trusted_root_certificate"`
	DestinationStore       types.String   `tfsdk:"destination_store"`
	RoleScopeTagIds        types.Set      `tfsdk:"role_scope_tag_ids"`
	Assignments            types.Set      `tfsdk:"assignments"`
	Timeouts               timeouts.Value `tfsdk:"timeouts"`
}
