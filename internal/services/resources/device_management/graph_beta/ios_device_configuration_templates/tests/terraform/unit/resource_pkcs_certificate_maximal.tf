resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "pkcs_cert_example" {
  display_name = "unit-test-iOS-pkcs-certificate-example"
  description  = "PKCS certificate profile for user authentication"

  # iOS PKCS profiles have no key_size, key_usage, extended_key_usages or scep_server_urls,
  # unlike their macOS counterparts.
  pkcs_certificate = {
    renewal_threshold_percentage      = 30
    certificate_store                 = "user"
    certificate_validity_period_scale = "months"
    certificate_validity_period_value = 12
    subject_name_format               = "commonNameIncludingEmail"
    subject_name_format_string        = "CN={{UserName}},E={{EmailAddress}},O=Example Corp"

    subject_alternative_name_type          = ["emailAddress"]
    subject_alternative_name_format_string = "{{EmailAddress}}"

    certification_authority      = "ExampleCA.example.com"
    certification_authority_name = "ExampleCA-CA"
    certificate_template_name    = "UserAuthentication"

    custom_subject_alternative_names = [
      {
        san_type = "emailAddress"
        name     = "{{EmailAddress}}"
      }
    ]
  }

  role_scope_tag_ids = ["00000000-0000-0000-0000-000000000001"]

  assignments = [
    {
      type        = "groupAssignmentTarget"
      group_id    = "00000000-0000-0000-0000-000000000002"
      filter_id   = "00000000-0000-0000-0000-000000000003"
      filter_type = "include"
    },
    {
      type        = "groupAssignmentTarget"
      group_id    = "00000000-0000-0000-0000-000000000002"
      filter_id   = "00000000-0000-0000-0000-000000000003"
      filter_type = "exclude"
    },
    {
      type     = "exclusionGroupAssignmentTarget"
      group_id = "00000000-0000-0000-0000-000000000002"
    },
    {
      type     = "exclusionGroupAssignmentTarget"
      group_id = "00000000-0000-0000-0000-000000000004"
    }
  ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
