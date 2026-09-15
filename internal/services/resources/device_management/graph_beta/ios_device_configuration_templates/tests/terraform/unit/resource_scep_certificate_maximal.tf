resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "scep_cert_example" {
  display_name = "unit-test-iOS-scep-certificate-example"
  description  = "SCEP certificate profile for device authentication"

  scep_certificate = {
    renewal_threshold_percentage      = 20
    certificate_store                 = "machine"
    certificate_validity_period_scale = "years"
    certificate_validity_period_value = 2
    subject_name_format               = "custom"
    subject_name_format_string        = "CN={{AAD_Device_ID}},O=Example Corp,C=US"

    # Bitmask on the wire: Graph joins these with commas into a single value.
    subject_alternative_name_type          = ["emailAddress", "universalResourceIdentifier"]
    subject_alternative_name_format_string = "{{EmailAddress}}"

    # Supplied as a bare GUID; the provider expands it to the full @odata.bind URL.
    root_certificate_odata_bind = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

    key_size  = "size2048"
    key_usage = ["keyEncipherment", "digitalSignature"]

    scep_server_urls = [
      "https://scep.example.com/certsrv/mscep/mscep.dll",
      "https://scep2.example.com/certsrv/mscep/mscep.dll"
    ]

    custom_subject_alternative_names = [
      {
        san_type = "emailAddress"
        name     = "{{EmailAddress}}"
      },
      {
        san_type = "userPrincipalName"
        name     = "{{UserPrincipalName}}"
      }
    ]

    extended_key_usages = [
      {
        name              = "Client Authentication"
        object_identifier = "1.3.6.1.5.5.7.3.2"
      },
      {
        name              = "Any Purpose"
        object_identifier = "2.5.29.37.0"
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
