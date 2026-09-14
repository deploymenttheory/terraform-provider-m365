# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v1.0.0...v1.1.0) (2026-09-14)


### Features

* add Global Secure Access cloud firewall policy and rule resources ([38e6139](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/38e613945cf15d6477f356268071a0d638a6ae92))
* add Global Secure Access cloud firewall policy and rule resources ([#3801](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3801)) ([f1fe560](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1fe560ce7686eda3fa462222b98913895da0f6a))
* add Global Secure Access custom block page resource ([9461105](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/946110522d02db186f0266820341bf43b70b19af))
* add Global Secure Access custom block page resource ([#3855](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3855)) ([4660d2b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4660d2baff27aba56b5d53e961ebae4510b5c19f))
* add Global Secure Access managed TLS certificate ([0914d34](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0914d34b116a0b3a8589555e4b8895edcf4a6db9))
* add Global Secure Access managed TLS certificate ([#3773](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3773)) ([076e8df](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/076e8df05024eef5c727563f6ab98e0077220d0d))
* add Global Secure Access MCP policy and rule resources ([142628d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/142628d1c846f1e1ba4c9b148e6ac472c0f87518))
* add Global Secure Access MCP policy and rule resources ([#3779](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3779)) ([3456fa5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3456fa5b24015b19e329d5b086aa249bbb2dcfe4))
* add Global Secure Access prompt policy and rule resources ([#3778](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3778)) ([2f196e7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2f196e7f2dc03bbaca086fa23de79743d616b757))
* add Global Secure Access threat intelligence policy and rule resources ([10dfd80](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/10dfd801a5f6494222662a33d1b112bbde68ec8b))
* add Global Secure Access threat intelligence policy and rule resources ([#3800](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3800)) ([dd27dd5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dd27dd5e1749e2c346df92e62797cb2ddfdf8155))
* add independent Global Secure Access prompt policies and rules ([dc25504](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dc2550448ca18a3a897aa2e1eb9147f0acbe7315))
* add independent TLS inspection policy and rule resources ([aa44d3e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aa44d3ec08e81533a0991a803be2063b97bc29b4))
* add independent TLS inspection policy and rule resources ([#3777](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3777)) ([2cc6309](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2cc6309627fa657cf60318de1bb0979b73a6d84c))
* add windows_autopilot_device_preparation_policy data source ([dbf13fd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dbf13fd4d7a5403b9554493706de8a826a58427e))
* add windows_autopilot_device_preparation_policy data source ([#3774](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3774)) ([b813d93](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b813d935bd6ed48f0a5a1f86a18ac505f8b5d564))
* align device preparation policy data source with device data source pattern ([ee7ea15](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ee7ea15fec44dbe7e9d4a22c655768d7bcebed85))
* ios managed app protection ([#3695](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3695)) ([9d27183](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9d27183ffd3dff99097d8a16133b2d273aeef24d))


### Bug Fixes

* align managed TLS certificate data source id handling ([f69d200](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f69d2007b2ec8f8781ec347ade64f64c1ff94d5b))
* categorize prompt policy docs and clarify error variable names ([efc8515](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/efc8515d0c232c9c48f01abcb5d6f31075182803))
* honor Graph middleware options ([3a867a4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3a867a4991a096a4b9bd9533999eadf08da15e58))
* keep prompt priority bounds visible to static analysis ([6cbbc62](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6cbbc62a6b6f096df68adb17cf72e51ece206676))
* preserve Graph request bodies when retrying compression fallback ([4fd8583](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4fd8583a52a1a1ac2eced4a7aed2dcdce4ed18dc))
* preserve Kiota retry handling for throttling ([b960ca8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b960ca88075226f3cd737599c4bc846752abf3d1))
* preserve Kiota retry handling for throttling ([0ce062e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0ce062e690b558a40dba92ebb02c4259192798be))
* preserve Kiota retry middleware in provider runtime ([486fa1b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/486fa1b8d039eebe6b79ada74db1daae1a85f19a))
* preserve request bodies on throttled retries ([34d6030](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/34d6030b0a7b7c58413d903a3de1dd1d31401877))
* retain device security group on autopilot device preparation policy update ([#3775](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3775)) ([a30944b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a30944b0814887eb8c055551efdaad18f4fdd06a))
* retain device security group on autopilot policy update ([81a6f69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/81a6f6960e12312b6da42121fa40470332fe2a77))
* retain resource state on wrapped transport errors ([4f25fdc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f25fdcb3f6e0d97d48694f9299a22fa16b41761))
* stabilize prompt rule scheme readback ([90b933e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/90b933e20ff4d7ecc8cdc597bbe0edf207260bcf))
* use Kiota retry handling for MCP creates ([e03577d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e03577df381704c4cbce63c214b8e38b61506c1f))
* windows_enrollment_status_page validate selected_mobile_app_ids by direct lookup instead of paginated list ([d816b73](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d816b73cb5d5ad4936a646aa100b9166458f181c))
* windows_enrollment_status_page validate selected_mobile_app_ids by direct lookup instead of paginated list ([#3830](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3830)) ([8da0b58](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8da0b5845240c51e51ac707b2d5ce135fd0a196f))
* wrap Graph client configuration errors ([fbe01db](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fbe01db2381e576b215537a48ec0a02acfa21998))

## [1.0.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.57.1-alpha...v1.0.0-alpha) (2026-09-01)


### ⚠ BREAKING CHANGES

* display_name, homepage, logout_url, service_principal_names and sign_in_audience are removed from microsoft365_graph_beta_applications_service_principal. Read them from the backing application resource or the service principal data source instead.

### Features

* add role_scope_tag_ids to beta device management role assignment ([#3607](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3607)) ([ba465b3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ba465b35bb89ce0c3897c2d239d2a1ce02585852))
* android managed app protection ([#3369](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3369)) ([cf01a71](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cf01a71262ef3b80684be675d818d792d4f5ee2e))
* model the service principal's own properties and drop application-derived ones ([#3492](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3492)) ([2456909](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2456909274b4201a15150440152f6de01527b063))
* support automatic mode device group assignments ([#3624](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3624)) ([56a6eaf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/56a6eafe8386b4f623b51eea117b35e419ca145b))
* support Entra ID SAML SSO token signing certificate lifecycle for service principals ([#3491](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3491)) ([1a450c7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1a450c7a3af8eed6ad4943b6998e0640924f8b65))
* support ipRange destination type for on-premises IP application segments ([#3490](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3490)) ([16eb2c0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/16eb2c0222bcddbe1a9fd30c56630010302a5de8))
* support priority and assignment filters on windows enrollment status page ([4f46495](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f464957bb61c8f4bc22abe380c179c9a5deb44d))
* support priority and assignment filters on windows enrollment status page ([#3702](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3702)) ([9aba9f4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9aba9f49d48c74c393f36ff66ab7b5f7b81f4fae))
* **windows_bios_configurations_and_other_settings_template:** add BIOS configuration and other settings template resource ([#3701](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3701)) ([44a3d7c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/44a3d7ca66507fc286208b2079a10d2f4078f4b9))
* **windows_bios_configurations_and_other_settings_template:** add resource ([734900f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/734900fa981a7dc9e77e040919b069ef0a035edb))


### Bug Fixes

* grace_period_hours to match windows_device_compliance_policy with the underlying API ([#3752](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3752)) ([a63b564](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a63b56460224cfc09aeeea6fec7f1b6694e31a9a))
* iOS managed mobile app resource fails to target apps (wrong API call) ([91891db](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/91891dba7b4a16a63629d87c72adb62eac4008c2))
* iOS managed mobile app resource fails to target apps (wrong API call) ([#3699](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3699)) ([8548c0a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8548c0a82fd3a28fbcdcc17e7a308daa33f5175e))
* **ios_store_app:** update is_featured attribute to be computed ([a9d2caa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a9d2caaa29d66295e71bb59409d415dd95c454df))
* **mobile_app_assignment:** map win32_catalog settings on read ([66b93bc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/66b93bcf076c1cf9b8dc8f986ea71b9e4e4abfc8))
* **mobile_app_assignment:** omit intent-incompatible Apple settings ([182042a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/182042a62d576fe826fa09f5d726756227fb915b)), closes [#3692](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3692)
* **mobile_app_assignment:** omit intent-incompatible Apple settings instead of defaulting them ([#3696](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3696)) ([9db2a05](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9db2a0580de650351116a141322a38daa1b66650))
* **mobile_app_assignment:** repair read refresh, update and import ([1ad7dd4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1ad7dd43b335a9626e3d9354c3fe171f0e4a7399))
* preserve Win32 app content and application identity ([d33f734](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d33f734376e12ee3b28aa27be58344d70315090e))
* preserve Win32 app content and application identity ([#3667](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3667)) ([8ea2b4d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8ea2b4d26e55b158afed8618d5119d79901f048b))
* **settings_catalog_template_json:** read assignments from /assignmen… ([#3623](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3623)) ([de8fc3b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/de8fc3b242c448efbd613dd59eff1efd8ee7faf1))
* **settings_catalog_template_json:** read assignments from /assignments endpoint ([de8fc3b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/de8fc3b242c448efbd613dd59eff1efd8ee7faf1))
* support multiple on-premises IP segment protocols ([#3449](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3449)) ([ab8220c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ab8220c5e95dbb66ca0618f7346eb3ff85c3ce70))
* support plain ZIP sources for Win32 apps ([d0ef8b1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d0ef8b18c64aef317ce6d53bdce6b5206522fd0f))
* **visionos_device_enrollment_policy:** match model field to schema attribute ([#3697](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3697)) ([1c9be30](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c9be3037ca6cf6aa84d3d138ca4a6a3fb2ad7fb))
* windows_device_compliance_policy grace_period_hours to match graph API. ([78c621e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/78c621ee3fdc1e7c22bfb255d4fc0838a3432909))
* **winget_app:** allow dots, hyphens and underscores in package_identifier ([#3694](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3694)) ([2eae4c7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2eae4c7b2d8906a5a90d50a45441ecefffe481ea))

## [0.57.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.57.0-alpha...v0.57.1-alpha) (2026-07-15)


### Bug Fixes

* **deps:** update go.sum to remove deprecated dependencies and bump versions ([#3365](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3365)) ([d398ec0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d398ec077ae1d55342f1300193516addd0c48c54))

## [0.57.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.56.0-alpha...v0.57.0-alpha) (2026-07-15)


### Features

* add Global Secure Access content policy resource ([#3317](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3317)) ([026edb4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/026edb441b908688ec5d44fa7e10bc6bb59eb438))
* add Global Secure Access content policy rule resource ([#3364](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3364)) ([67d88d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/67d88d2543e1fa7ca521f2112b7ccfa7d4021f27))
* add windows custom configuration (OMA-URI) device profile resource ([#3363](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3363)) ([3c0e37e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3c0e37ee444c418f039b1c859a785d1ec82ffa85))
* Adds a Microsoft Entra Global Secure Access Intelligent Local Access ([#3263](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3263)) ([07443c1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/07443c187751998dab5b879ee3dbb58d24fc784d))
* Adds a Microsoft Graph beta resource for Microsoft Entra Global Secure Access Security profile policy links ([#3231](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3231)) ([dfd6aad](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dfd6aadd9d197607d9b6bc4da0794aedd4ffdd8c))
* Support Global Secure Access Internet policy support ([#3265](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3265)) ([4499fb7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4499fb750d6267a021e811ab3c4c3193658cbbf8))
* windows managed app protection ([#3262](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3262)) ([4b2befa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4b2befa003a3db7efc364f35a5f68466c9084f5d))


### Bug Fixes

* handle Entra eventual consistency when creating service principal token lifetime policy assignments ([#3316](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3316)) ([7690e40](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7690e4091db36c27564a16a806cc94cc4ef2b7d8))
* OS version regex ([e037e64](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e037e64599292c432808233c0ee88863efbdf051))
* OS version regex ([#3233](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3233)) ([e037e64](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e037e64599292c432808233c0ee88863efbdf051))
* verify group and user license assignment changes are actually applied before reporting success ([#3264](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3264)) ([21ba347](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/21ba347faa00a2e3acaf6fe8f88ea1266e31d8a6))
* winget app computed state drift ([#3318](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3318)) ([58e3d96](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/58e3d96fa874eae67888f4aa1631456bf24f0bee))

## [0.56.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.55.2-alpha...v0.56.0-alpha) (2026-07-06)


### Features

* add Application Proxy connector group resources ([#3213](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3213)) ([e899000](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e899000432d83bc7d32d1182982d26a32a0d3025))
* Add Global Secure Access filtering profile resource ([#3222](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3222)) ([9ab123e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ab123ed5eaaf3e648f595548735c60df8968137))
* Add Graph beta Global Secure Access web filtering policy resources ([#3227](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3227)) ([8a8f650](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8a8f650b19554982ca9aef3b219d27e9e58d5243))
* add macOS DEP (ADE) enrollment profile resource (depMacOSEnrollmentProfile) ([#3202](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3202)) ([94f3f3d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/94f3f3d66380002660b578a5c5cffb8178b65217))
* add token lifetime policy and service principal assignment resources ([#3183](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3183)) ([e7f3a9d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e7f3a9dc987274b4ddeada02268adf916e50e685))
* add traffic_routing_method attribute to microsoft365_graph_beta_applications_application_on_premises_publishing resource ([#3184](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3184)) ([d99ce3c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d99ce3ccd4c9ae9b3e5e884d56c2e223ec30c614))
* Add VisionOS Device Enrollment Policy resource to provider ([#3230](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3230)) ([148676e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/148676e409b3755641d282a5a8c9cb78a4438ae4))
* added field 'is_default_policy_assignment' to macos_device_enrollment_policy and added resource 'identity_and_access_cross_tenant_access_partner_group_sync_settings' ([#3225](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3225)) ([4f54769](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f5476963410981ecd9497bdaf9941811eb91a06))
* added resource ios_ipados_device_enrollment_policy ([#3228](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3228)) ([1d8bada](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1d8bada9403c283617f3f7dbd1b7d29d2253ba6c))


### Bug Fixes

* Align filtering profile description validation ([#3224](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3224)) ([67c09bb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/67c09bb2fadc965e52c4ba3c7d6c38b844d41139))
* on-premises IP application segment resource ([#3210](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3210)) ([6d9d490](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6d9d490dcd5a3b318e2b9a6588ae9e2a33265568))
* update OperationApprovalPolicy resource schema ([#3181](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3181)) ([5b96c6d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5b96c6d56c2240bf6044a4b53dd6395d6786d19e))

## [0.55.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.55.1-alpha...v0.55.2-alpha) (2026-06-23)


### Bug Fixes

* application federated identity credential ([#3171](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3171)) ([732df3d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/732df3d91f887dfaa100574ef2b97a2ec6c6ab72))
* moved the identity set in create under timeouts to handle potential 404's ([#3173](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3173)) ([b9bcb22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b9bcb22fa09d212e841829d866ab2b7cb893c037))

## [0.55.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.55.0-alpha...v0.55.1-alpha) (2026-06-17)


### Bug Fixes

* add minimum resource requirements to Win32LobAppResource schema ([#3159](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3159)) ([7caaed5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7caaed599362cb1ef316522c62b906cd7bf53497))

## [0.55.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.54.0-alpha...v0.55.0-alpha) (2026-06-15)


### Features

* added Graph Beta Users datasource ([#3139](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3139)) ([d362261](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d362261eb47c58d155e62ef6143081683a1a9f5e))


### Bug Fixes

* add MinimumFreeDiskSpaceInMB to Win32LobAppResource schema ([#3155](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3155)) ([011d08f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/011d08fd8895f991851f9b9d791f5c97ecd09da2))
* deprecate detection_rules and requirement_rules in Win32LobAppResource schema ([#3154](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3154)) ([a6470c9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a6470c94567ae70215a4bc95bd3d8a00bbc5b452))
* update minimum supported Windows release in Win32LobAppResource schema ([#3153](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3153)) ([6a65e2e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6a65e2e110d790da33ef162c84c6515707e8048f))

## [0.54.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.53.1-alpha...v0.54.0-alpha) (2026-05-26)


### Features

* add settings_catalog_inventory_policy resource ([cda4557](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cda4557c5e4419cc5f1292729ae74f4d1cc6c2ea))
* add settings_catalog_inventory_policy resource ([#3105](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3105)) ([03188e3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/03188e34c2995f71c4a317a331165c0d30590d73))

## [0.53.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.53.0-alpha...v0.53.1-alpha) (2026-05-19)


### Bug Fixes

* [#3085](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3085) ([d624f30](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d624f305885b40cb980d5d11561ebb7c85681f92))
* for device_and_app_management_win32_app so handle for correctly for .exe based installers ([#3086](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3086)) ([f5f1571](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f5f1571e3fa9a0e9a0d738ed86de7548f92c830d))
* import resource handling for device_and_app_management_win32_app, device_and_app_management_macos_pkg_app and device_and_app_management_macos_dmg_app ([89cbefb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/89cbefb93098a90f8680c7d1255bd03989db83f3))
* import resource handling for device_and_app_management_win32_app, device_and_app_management_macos_pkg_app and device_and_app_management_macos_dmg_app ([#3087](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3087)) ([398a2b7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/398a2b7709e65e5578deebcebeb8e35fb5e1d1ac))
* updated release please pipeline to use ubuntu latest rather than a large runner ([#3088](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/3088)) ([5210015](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/52100157b325b5e1571d262a1f87e51521c1e25a))
* updated release please pipeline to use ubuntu latest rather than a large runner which is not needed ([626b4a1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/626b4a16d5185e3180cecbc807c3d8870fca2a70))

## [0.53.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.52.1-alpha...v0.53.0-alpha) (2026-05-02)


### Features

* Enhance Windows Device Compliance Policy with additional system security settings ([a4cb627](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a4cb6275c2d5c2be31a0c57b05496f353c4a332d))


### Bug Fixes

* Enhance Windows Device Compliance Policy with additional keys ([#2991](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2991)) ([0c031fc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0c031fc620b28b580463ad3032308a8030c30a1c))
* Temporarily disable support for 'androidFotaDeployment' assignment target ([#2987](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2987)) ([c387833](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c387833ba83f1ed11feca97f44ee5decedcb314a))
* Temporarily disable support for 'androidFotaDeployment' assignment target due to its removal in msgraph-beta-sdk-go v0.160.0. Added error handling and comments for future restoration when the type is reintroduced. Updated related mappings and schema versioning in mobile app assignment resources. ([c387833](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c387833ba83f1ed11feca97f44ee5decedcb314a))
* Temporarily disable support for 'androidFotaDeployment' assignment target due to its removal in msgraph-beta-sdk-go v0.160.0. Added error handling and comments for future restoration when the type is reintroduced. Updated related mappings and schema versioning in mobile app assignment resources. ([ea12f94](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ea12f94d58964346d40ad18281c5b427cff6a271))
* Update Intune branding profile handling ([#2988](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2988)) ([b8812d1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b8812d1bba588f2e18a14baebf28a0c626be64e6))
* Update Intune branding profile handling: Temporarily comment out 'sendDeviceOwnershipChangePushNotification' due to its removal in msgraph-beta-sdk-go v0.160.0. Added deprecation message in schema and updated versioning in resource files for future reference. ([b29e6e4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b29e6e4f7ed7613d1ab3c22a10dbe428def1f59a))

## [0.52.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.52.0-alpha...v0.52.1-alpha) (2026-03-29)


### Bug Fixes

* **macos_software_update_configuration:** persist plan state before ReadWithRetry in Update ([#2689](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2689)) ([3a105a8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3a105a821666be8f8059a582d1027e38a6369955))

## [0.52.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.51.0-alpha...v0.52.0-alpha) (2026-03-26)


### Features

* added predicates to handle eventual consistency and to move away from hard coded pauses between create -&gt; read and update -&gt; read. ([#2652](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2652)) ([564efe4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/564efe414678cad3af5756a1e3a9d33d6f927af1))
* added resource access_administrative_unit_membership ([#2648](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2648)) ([9a30bcd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9a30bcd22bdf45e0abc1abfd52623f4bbd933d60))
* added resource administrative units ([#2623](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2623)) ([0588204](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/05882046291a160a83e3871a2265b0f60c8ed1a7))


### Bug Fixes

* bi-weekly tests workflow for windows_updates and change_notifications ([#2649](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2649)) ([deaf2d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/deaf2d29da1d46ed24e150515a0a55afb504d748))
* test fixures for microsoft365_graph_identity_and_access_subscribed_skus ([#2650](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2650)) ([22e3b11](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/22e3b11c2c65f28b517698ee930926ebd1269de4))

## [0.51.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.50.0-alpha...v0.51.0-alpha) (2026-03-23)


### Features

* add utility resource for GUID list sharding ([4d742a0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4d742a0992e018db5f1a6139b64712e823af314e))
* add utility resource for GUID list sharding ([#2620](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2620)) ([7f65f48](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7f65f48281aa24ae5c728c5d5a2bfad9b15d276e))
* added new search query types for graph_identity_and_access_subscribed_skus and bug fixes for Group License Assignment and device_management_group_policy_uploaded_definition_files ([#2621](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2621)) ([4063307](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40633075977b47b5bd1d818db6972d7d38d3f9ae))
* added new search query types for microsoft365_graph_identity_and_access_subscribed_skus and bug fixes for GroupLicenseAssignment ([b9383b5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b9383b52bca12385bb807c6dc22d444a1c1dfa78))


### Bug Fixes

* added key ' force_definition_file_upload' to microsoft365_graph_beta_device_management_group_policy_uploaded_definition_files to handle upload error scenarios. ([9064d15](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9064d15abc115730ad3b46c86bb12e166b9772dc))
* graphBetaWindowsUpdatesAutopatchUpdatableAssetGroup ([dcda75d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dcda75dc142d6f8a44154923ffedbc9728234fc4))
* graphBetaWindowsUpdatesAutopatchUpdatableAssetGroup ([e9e3757](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e9e3757f4932a4bc15d74760b478bf9329edf99d))
* resource WindowsUpdatesAutopatchUpdatableAssetGroup state mapping ([#2616](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2616)) ([dcda75d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dcda75dc142d6f8a44154923ffedbc9728234fc4))
* update mobile app data source to refine filtering capabilities ([#2618](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2618)) ([6b0a712](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6b0a712f57f58a95c52153fc5234ad127a960d3e))

## [0.50.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.49.1-alpha...v0.50.0-alpha) (2026-03-20)


### Features

* added datasource graph_beta_identity_and_access_device and refactored datasource graph_beta_device_management_managed_device ([#2596](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2596)) ([e3cb8a8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e3cb8a872686adcca3e91ac84a5888a29f6c5fba))
* added windows autopatch resources: ...... ([#2522](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2522)) ([6bff051](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6bff051e38805507cf15782d96b36cc2d0b2eefb))
* added windows update autopatch resources to the provider ([#2571](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2571)) ([c4efc8e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c4efc8ed2a5ff866a605a251af7833a61d9d011b))


### Bug Fixes

* correct attribute name in autopatch_updatable_asset_group_assignment example ([1e2bf8e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1e2bf8e79f1235cf8580609b7084cbbce86a8d69))
* session_controls state mapping for global_secure_access_filtering_profile to prevent post-apply null inconsistency ([#2568](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2568)) ([e9df2c6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e9df2c6ea01e0d5fcba364fc4ecff7232d6e7482))
* update autopatch_device_registration acceptance tests for refactored managed_device API ([7ba7006](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7ba7006c33037a152c6e08d03932242f989aec6e))
* update autopatch_updatable_asset_group_assignment tests for refactored APIs ([b7b3b5a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b7b3b5a52e117e4e4fde0212ff4fbddb18d1848f))

## [0.49.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.49.0-alpha...v0.49.1-alpha) (2026-03-13)


### Bug Fixes

* migrated all dynamic errors in validate.go implementations to use sentinel errors. ([#2545](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2545)) ([32f4cf1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/32f4cf101a8404c242bca0f8cb808487ba95ff95))

## [0.49.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.48.0-alpha...v0.49.0-alpha) (2026-03-13)


### Features

* add added resource microsoft365_graph_beta_identity_and_access_cross_tenant_access_policy ([9f9b7b2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9f9b7b2664aa3d2e8b1f67feb0dcfde883203765))
* add added resource microsoft365_graph_beta_identity_and_access_cross_tenant_access_policy ([#2539](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2539)) ([22b68c0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/22b68c080c30709a375afee50538e60a116ade63))
* added resource graph_beta_identity_and_access_cross_tenant_access_default_settings ([#2540](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2540)) ([0edaa28](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0edaa285c0dc19b0d24a2e55c8015174b85a5d96))
* added resource graph_beta_identity_and_access_cross_tenant_access_default_settings.md ([5f63cd4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5f63cd44c65d6056af05f9148a6e4e32fa6fbd7b))
* added resource graph_beta_identity_and_access_cross_tenant_access_partner_settings ([#2544](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2544)) ([8ee428e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8ee428e6c2dea3155756420511a1f1c37204f837))
* added resource graph_beta_identity_and_access_cross_tenant_access_partner_user_sync_settings ([#2542](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2542)) ([4b44750](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4b44750ecf607bb2d5b98965f05c0a873768615f))
* enhance CrossTenantAccessDefaultSettingsMock with additional user and group mock registrations ([aeace3b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aeace3b3d0cb2f09fddb8cb4c6bebfc3a962185d))
* implement validation for CrossTenantAccessPartnerUserSyncSettings resource ([#2543](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2543)) ([b6a1ff4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b6a1ff450513e7a3ad727c093f1a869d9f1b558c))


### Bug Fixes

* for schema validation for device config assignments ([19d4a90](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/19d4a90d08d8ef31fda8fb6e10eb141198de07c0))
* for schema validation for resources using DeviceConfigurationWithAllGroupAssignmentsAndFilterSchema ([#2537](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2537)) ([122dc46](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/122dc4666811ed662685e7d3c92e95b42760f24a))

## [0.48.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.47.0-alpha...v0.48.0-alpha) (2026-03-06)


### Features

* Support username and password authentication for client ([2bd0e46](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2bd0e463c69f5e4e786d3338dbc07f652ced692e))
* Support username and password authentication for client ([#2517](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2517)) ([e695d69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e695d698e5724fd1e9c34ac44553422b4fea77f2))


### Bug Fixes

* update Entra ID client tests and middleware to prevent gzip compression on OAuth requests ([852d9f2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/852d9f2f970a6917ca76979b0e59e4036b3c6c34))
* update Entra ID client tests and middleware to prevent gzip compression on OAuth requests ([#2519](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2519)) ([e3feb22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e3feb22594d8aad4b0940996115658543a2baa1b))

## [0.47.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.46.0-alpha...v0.47.0-alpha) (2026-03-05)


### Features

* Support az cli for API Client ([#2493](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2493)) ([d8a50c1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d8a50c10e00f5497a59bd933012771bdef9c88e5))
* Support az cli for client ([5be2715](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5be2715fdc555e17156332c9a94464790b18a0ff))


### Bug Fixes

* added scripts to automate the definition of permissions for all tf block types so that 403 errors and tf doc templates describe the correct permissions  ([#2513](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2513)) ([540999e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/540999e5f1437f0f493cdc60abb9d4b581e144bd))
* fixed args for pipeline permissions generation ([13ac8ec](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/13ac8ec8945bc3cdf7ea11ee5c4bac644704c390))
* fixed args for pipeline permissions generation ([#2515](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2515)) ([190d0e4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/190d0e4be69a478a4ca31c166af3e5599c81c4af))
* split GO_ARGS and TEMPLATE_ARGS in update-graph-permissions workflow ([41f1c02](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/41f1c0287f37b629e0575c5098685649f7b0e1ac))
* standardised crud logic flow ([9b31864](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9b31864237850b165e010344bf99416e5f77c6ec))
* standardised crud logic flow for resource group_app_role_assignment ([#2491](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2491)) ([110fc92](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/110fc9259fe2f6afa1809e630860fbf22d4a305f))
* Update Group resource CRUD methods to utilize new types package ([a369d2b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a369d2be5ed3e318518307258302d585b6083ed1))
* Update Group resource CRUD methods to utilize new types package ([#2494](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2494)) ([d06b976](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d06b9766bb03aaa788e4d9ea510f20a6d58deed9))

## [0.46.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.45.0-alpha...v0.46.0-alpha) (2026-03-03)


### Features

* Windows Autopilot Device Preparation Policy assignments to use standard shared model schema with type/group_id fields ([8349f59](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8349f590643b16b0aefdf1457039d7ba022e3e3b))


### Bug Fixes

* resolve err113 by centralizing construct errors in sentinels and reduce Autopilot settings builder complexity to satisfy gocyclo ([6cb2d9a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6cb2d9ab650c4ec0a025b17ec1a9e55040147ad1))
* Standardize assignments schema for Windows Autopilot Device Preparation Policy ([#2474](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2474)) ([01a84f6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/01a84f6890aaaaed7f577a54a1b086d8cfc2a911))
* Use PUT instead of PATCH for windows_autopilot_device_preparation_policy updates to handle settings navigation property ([567a8b4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/567a8b4a071ac4a5b358fe011bfa8fd040dd54ce))
* Use PUT instead of PATCH for windows_autopilot_device_preparation_policy updates to handle settings navigation property ([#2489](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2489)) ([71682f4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/71682f44c092ee48beb2199bee9698f0ca9a60bc))

## [0.45.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.44.0-alpha...v0.45.0-alpha) (2026-02-15)


### Features

* add identity schema support for all resources ([fd0696e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fd0696eba7a99119369c401139b35cc25da9bc80))
* added list resource user and updated docs ([53f6229](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/53f6229e068ea00a0afae0c6214f809466bc7694))
* added list resource user and updated docs ([#2384](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2384)) ([18363c3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/18363c36f354663993784a29a5294a371a4525b4))
* Added support for `microsoft365_graph_beta_device_management_windows_platform_script` ([3333200](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3333200c0aa1f478b32294bbf5cdf168982d33c6))
* Added support for microsoft365_graph_beta_device_management_windows_platform_script ([#2397](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2397)) ([0f113b8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0f113b88fd8aeada9f44721c6391897363d755bc))
* Enhance StateContainer to support Identity management ([211b2b6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/211b2b62629c785d8d25b946e83ae3fcaaa1151e))


### Bug Fixes

* add AllowSystemGeneratedSetValues modifier for enhanced set management ([ad6a2ca](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ad6a2cae376343f04bd9472b8d23193b97ee79ca))
* add GraphToFrameworkStringSetFiltered function for tag filtering ([521db10](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/521db102c02e2c5337514d14a3d1a0bb0f55b759))
* add validation for Windows feature update policy resource constraints ([#2361](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2361)) ([3d5dca0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3d5dca09a52d5d820d4caf19d88c5443294de45f))
* added support for panicking test cases in bi-weekly test pipeline ([#2398](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2398)) ([475301c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/475301cda7a656417ba3f12a488189ac0ac83687))
* All errors from external packages are now properly wrapped with context using fmt.Errorf with the %w verb, satisfying the wrapcheck linter requirements. ([17a2a31](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/17a2a31411e6fdc19967fdf5e363ac591e8ecd87))

## [0.44.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.43.0-alpha...v0.44.0-alpha) (2026-02-03)


### Features

* enhance guid_list_sharder to support applications and improve test structure ([#2327](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2327)) ([e8d10f0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e8d10f020c2c27f3cc3b48e353818a66802dad5f))


### Bug Fixes

* enhance error handling in CRUD operations for MacOS and Win32 apps ([#2325](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2325)) ([5dc4595](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5dc4595661399c274e23e9a9799759631f6d3e47))

## [0.43.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.42.0-alpha...v0.43.0-alpha) (2026-02-02)


### Features

* added resource application, pplication_certificate_credential, application_federated_identity_credential, application_identifier_uri, application_owner, application_password_credential, service_principal, service_principal_owner ([#2321](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2321)) ([9abd25c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9abd25c58ef212c8c06eaa816053402127b10bd6))


### Bug Fixes

* update HTTP method in Delete request for Conditional Access Terms of Use ([#2238](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2238)) ([82caa60](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/82caa60370243cab63ca1809bcc9f331276e1c27))

## [0.42.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.41.0-alpha...v0.42.0-alpha) (2026-01-24)


### Features

* add comprehensive AST analysis tests for SDK usage tracking ([c7b67dd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c7b67dd78797bcac7c90d7c9ad222ba71efe7d64))
* add enum change analysis and reporting ([d75351c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d75351cea988793a136c6768714f700c1104d763))
* add enum change analysis and reporting ([#2028](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2028)) ([d53d820](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d53d820c01d613473f0cc3b08de53e890e2905ee))
* added datasource microsoft365_graph_beta_groups_group ([a62fad5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a62fad51b6e1f81245ecb34e4ec0eecff932a01e))
* added datasource microsoft365_graph_beta_groups_group ([#2012](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2012)) ([197db2c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/197db2cfe114738b394b38fcf3c17a3f27b32573))
* added datasource microsoft365_utility_guid_list_sharder ([#2032](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2032)) ([cd7b2e9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cd7b2e9bf167207b8a2e9e88b8cee0be643eb7f0))


### Bug Fixes

* enhance test coverage reporting by sanitizing package filenames in CI workflow ([48b42bf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/48b42bf8200af134383a10071f7a283a741e835d))
* improved test coverage ([6fa7600](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6fa76001cf896940bdfba3a1b1d140a111aa4a4a))
* unit and acc tests for resource microsoft365_graph_beta_device_management_settings_catalog_configuration_policy_json ([ef6b30c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ef6b30c04e3c7e7660060b0213b7eb3cedefde72))
* unit and acc tests for resource microsoft365_graph_beta_device_management_settings_catalog_configuration_policy_json ([#2024](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2024)) ([818a01c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/818a01c427ade4ac56faa55ad46ac15a06d0f574))

## [0.41.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.40.2-alpha...v0.41.0-alpha) (2026-01-20)


### Features

* added datasource microsoft365_utility_deployment_scheduler ([9ac25d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ac25d211acb26e0f75232f3c936a4a3cc4e124a))
* added datasource microsoft365_utility_deployment_scheduler ([#1977](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1977)) ([75f1f36](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/75f1f3633f18a3efd9addd7eda08e4941133a8cb))
* Update msgraph-beta-sdk-go to v0.157.0 and add conditional access template datasource ([#2011](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2011)) ([5a75ca0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5a75ca0a717f97a4abc42b6962fbc150a84ec634))


### Bug Fixes

* added group deletion handling when licenses are assigned. standardised version constraint for time provider ([c4c1f04](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c4c1f04a5984532d7af075f36c0395816e64307e))
* added group deletion handling when licenses are assigned. standardised version constraint for time provider ([#1792](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1792)) ([3de4d69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3de4d69560af1ddaf5e2de5241e4986f3105efc4))
* assertion assignment order ([633dbce](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/633dbce570ca3b5d18950a21c944d14ca124638a))
* better field extraction ([566fff1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/566fff1bead77702139ca24f301783f67f029415))
* better field extraction ([#1786](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1786)) ([120b647](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/120b64725fcb459f4cad4f8c601c24050fe85e37))
* changed applications_ip_application_segment field application_id to application_object_id in prep for full fix ([0672f8a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0672f8af7ac2f39e99ada8aa50819b86bfe7e80d))
* changed applications_ip_application_segment field application_id to application_object_id in prep for full fix and updated docs ([#2009](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2009)) ([c1b6cdc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c1b6cdcfe6b921b4d23fb4a3738d97430fd0e8a8))
* for resource macos_device_configuration_templates test harness ([8a9352c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8a9352c757722cdd622b3bfa695fd50218a577ef))
* for resource macos_device_configuration_templates test harness ([#1793](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1793)) ([41603d3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/41603d3a3bf731805664f5b807776743489618d7))
* for test harness for WindowsQualityUpdateExpeditePolicy, WindowsFeatureUpdatePolicy and WindowsEnrollmentStatusPage ([e7c2ca8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e7c2ca8e51e963255df49e43656e868f815148f7))
* for test harness for WindowsQualityUpdateExpeditePolicy, WindowsFeatureUpdatePolicy and WindowsEnrollmentStatusPage ([#1796](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1796)) ([1ce72e4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1ce72e48d2dd0821c3d76102bb6a6ed6a0b3b687))
* named_location delete workflow logic refactor ([f125ed5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f125ed525a5c7f0fbf01efffc4e3c304c909544f))
* named_location delete workflow logic refactor ([#2007](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/2007)) ([d86845a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d86845ac8d8223eb7775a8462777c94a6fa3f434))
* refactor and added docs ([71664eb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/71664ebe8507bfd368f5ee1b056812c111b2e725))
* refactor of resource group_policy_configuration test harness , standardised ExternalProviders Random with const usage and changed operation strings for HandleKiotaGraphError to consts ([3039b56](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3039b56377caedc66c26b8c0038a659670afa7aa))
* refactor of resource group_policy_configuration test harness , standardised ExternalProviders Random with const usage and changed operation strings for HandleKiotaGraphError to consts ([#1788](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1788)) ([d087094](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d087094f1cb41d467e8d9592acf7ad0131807d97))
* refactored resource device_management_app_control_for_business_built_in_controls test harness ([4d1bf5c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4d1bf5cec98ff01a4d2b62fff6f400fa3479f852))
* refactored resource device_management_app_control_for_business_built_in_controls test harness ([#1787](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1787)) ([867566a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/867566ae50bc4a01e6c4f675c274b8785a273bdf))
* refinements for overly broad codecov ignore pattens ([3042fa8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3042fa80ef70ef9e17278ea037b997faafbca37e))
* refinements for overly broad codecov ignore pattens ([#1789](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1789)) ([70d75a9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/70d75a99dd16ab27c4d32994a00269db630802ae))
* removed debugging ([13b72fb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/13b72fb3fa7b2b8deac5e0ba22d69ef567e448af))
* reverted license assignment test validation destruction helpers due to export field issues ([4797bc6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4797bc6c77c2ae525058a79c2e5423e52486d6f4))
* reverted license assignment test validation destruction helpers due to export field issues ([#1791](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1791)) ([0902ebb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0902ebbdfdd40d79d018698b2db3bcea3fabf55d))
* sec ([ff2edbe](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ff2edbe6098891235e7d5768b71eda4bc29f62fb))
* sec ([#1712](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1712)) ([342ebb5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/342ebb5cf55bc790964bf0c69828a0a7ab31c3d6))
* unit test naming convention and docs ver ([74c78bf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74c78bfe829f429bd919977e9563d72bb6c579d9))
* **windows_update_ring:** set all deadline_settings attributes as computed and optional ([fa4e6c2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fa4e6c2fa637a3f48d98f4492a51eb8ac39aa5ac))
* **windows_update_ring:** set all deadline_settings attributes as computed and optional ([#1976](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1976)) ([1aa1cc3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1aa1cc301b46a2aec464b1cc44f9dbc3317b7dd6))

## [0.40.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.40.1-alpha...v0.40.2-alpha) (2026-01-12)


### Bug Fixes

* description field removal of default value set to "" and standardisation of valid max string length ([#1644](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1644)) ([042592e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/042592e3cfecf587b50af7b15eff07ff9011b22c))
* for sku id's in acc tests for resource UserLicenseAssignmentResource ([#1670](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1670)) ([60fd5f3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/60fd5f3d976e97ebebd8306180f866d83ca77a44))
* for unit and acc tests for device_management_group_policy_uploaded_definition_files ([#1648](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1648)) ([bdd1807](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bdd18075597f3738200674fbb624786727b208dc))
* multiple fixes for unit and acc test failures and updated TestHTTPClient to accept context in resource tests  ([#1669](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1669)) ([2c8c00c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2c8c00cb75ff033bc92e28419c2b90405cb26ae3))
* pr bodies ([870d26f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/870d26f36b19aad597b98d955504d01d8164cb28))
* pr bodies ([#1708](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1708)) ([0ff9ede](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0ff9edefe49413e9dc26c184027105b62739f0f2))
* refactor of sdk_schema_change_detection ([cae2753](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cae2753f20d35fcdf08c7898418e87a19f489b3b))
* refactor of sdk_schema_change_detection ([#1704](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1704)) ([185496e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/185496ef8f34452c9ee2c93b909215bb1c5c4737))
* refactored unit and acc tests for resource device_management_macos_platform_script ([#1667](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1667)) ([5a7596e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5a7596ebd38777b5474cd5d4ab2d555d85e756d6))
* updated various docs inc. batch import, tf ops best practice and tf workspaces ([#1694](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1694)) ([430f8e4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/430f8e451ae9345c0244bbe0a3e282be88aa228c))

## [0.40.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.40.0-alpha...v0.40.1-alpha) (2026-01-07)


### Bug Fixes

* for import path ([526715e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/526715eafc977db612dc0e81c9e25f2ce85986c9))
* for import path ([#1642](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1642)) ([7a3cdd5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7a3cdd5f41aeddf15b0d53c90068fab62dc939bf))

## [0.40.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.39.2-alpha...v0.40.0-alpha) (2026-01-07)


### Features

* add list_resource microsoft365_graph_beta_device_management_settings_catalog_configuration_policy ([aa4f6c2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aa4f6c2fbfd0ea9d6125355d6fcc43abfc0a085d))
* add list_resource microsoft365_graph_beta_device_management_settings_catalog_configuration_policy ([#1625](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1625)) ([0c94be7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0c94be74d778aa33702989c18e11000c562b381c))
* added data source /group_policy_value_reference ([5aaa539](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5aaa5390f2edf59179d0c34c100db57d9bfdf659))
* added data-source group_policy_value_reference ([#1476](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1476)) ([a143543](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a143543c0899ca1509fbfb511344d4ab58ce7983))
* added resource device_management_group_policy_definition with test harness and removed old group policy resource for specific field types. ([01ceb81](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/01ceb8197cea64f1378e7112711c537abde58ea4))
* added resource device_management_group_policy_definition with test harness and removed old group policy resource for specific field types. ([#1496](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1496)) ([7604d2e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7604d2eb98e6c07532586050d2e79cdf0488bb3b))
* refactored all actions with correct examples, common syncprogress handling and support for timeout over rides ([74cde98](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74cde98787ae745aeba861abc2ca0fa2f13fcb77))
* refactored all actions with correct examples, common syncprogress handling and support for timeout over rides ([#1609](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1609)) ([54dbe59](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/54dbe5922a1cf003a3e8ef0006b81bc5be050191))


### Bug Fixes

* extended bi-weekly tests to support actions ([73f30e4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/73f30e4d8896d918e75665b97beb90c8e3598248))
* extended bi-weekly tests to support actions ([#1610](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1610)) ([8f620a3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8f620a331ab4db2b10883011e2cd553287006125))
* for actions timeouts to use terraform-plugin-framework-timeouts v0.7.0 ([b9cbbb2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b9cbbb2788a65819ec27907a7db9108d4d1a539a))
* for actions timeouts to use terraform-plugin-framework-timeouts v0.7.0 ([#1591](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1591)) ([0e903d4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0e903d433bd5755cbfa23dec583803aa0d9d823d))
* for conditional access acceptance tests ([8485bea](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8485beaea2d5ccad5557792c4250543028347d26))
* for device_management_windows_update_ring schema ([55219c1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/55219c1d88983adc46ac69a6be3b9b53e989d2f9))
* for failing conditional access acceptance tests ([#1503](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1503)) ([d5c3c69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d5c3c69dd09799d193eada584d7f8c42f902892c))
* for resource device_management_windows_update_ring schema ([#1641](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1641)) ([1defcd1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1defcd14bfe397481c79475bc7b6566040ad40b3))
* standardised schema description field logic, with default values and optional/computed atts ([db6b47f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/db6b47ffc338807622d0ee19e8be39614f73c63a))
* standardised schema description field logic, with default values and optional/computed atts ([#1627](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1627)) ([7fe73d0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7fe73d00c259f555f6c21cd359ea05106c03ab04))

## [0.39.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.39.1-alpha...v0.39.2-alpha) (2025-12-29)


### Bug Fixes

* path typos ([6886208](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/68862088c591969d0b4d2c60e8e7fe735b664502))
* path typos ([#1474](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1474)) ([835cadd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/835cadd5b1ce59cf546d4eb65007edf953ab6ba7))

## [0.39.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.39.0-alpha...v0.39.1-alpha) (2025-12-28)


### Bug Fixes

* extended pipeline timeout to 240mins ([#1463](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1463)) ([9558169](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/955816919a3b1a2f6b8a7ba00dc6eec1b3ff17cc))

## [0.39.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.38.0-alpha...v0.39.0-alpha) (2025-12-28)


### Features

* added datasource microsoft365_graph_beta_identity_and_access_tenant_information ([#1364](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1364)) ([236d177](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/236d1770ef7b706a68fc8a4d826fafc949c72e3e))
* added datasource microsoft365_graph_beta_identity_and_access_tenant_information ([#1365](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1365)) ([0d8b53c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0d8b53c83cf27f6aca5eefd5b3032cc67e11f0aa))
* new resources and refactors ([#1384](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1384)) ([ef79681](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ef796819819b8f554580bca0a96d286338b68250))
* refactored conditional_access_policy to use kiota sdk and acc tests and default values for [] ([#1361](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1361)) ([d37b8b3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d37b8b351af454e24f13690fd4164535db316065))
* refactored test harness for resource windows_remediation_script and fixed issue with schedule interval default values ([#1386](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1386)) ([b93bb03](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b93bb03c05adfc38da4eb8a185930562cd441ecf))
* various additons to test harness ([#1363](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1363)) ([6be4931](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6be4931147687d8186d66076f34f836e8f85ed56))


### Bug Fixes

* Add a runRemediationScript attribute of microsoft365_graph_beta_device_management_windows_remediation_script resource assignment ([#1385](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1385)) ([49e871f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/49e871fd8bc452b1c97da9b072c6962a67e1e8d9))
* extended bi weekly test pipeline to 180m and fixed an issue in python script where failed_step was only initialized inside the else block ([#1404](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1404)) ([5d7a473](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5d7a47302ad37f4e92b18c131ca719694b7a41b8))
* numerous fixes for test harness for resources microsoft365_graph_beta_device_management_windows_quality_update_policy and microsoft365_graph_beta_device_management_windows_quality_update_policy ([#1397](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1397)) ([05b7adc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/05b7adc7c570185a36db196eddfa26ee6291ccd4))
* refactor windows_feature_update_policy test harness ([#1402](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1402)) ([5d258f9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5d258f9c63f0af995281a63f59f88345e7c0316d))
* refactored test harness for windows_quality_update_expedite_policy ([#1398](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1398)) ([3c7507a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3c7507ab2db7fc6ee2f906fd6a0b92eecd46b113))
* refactored windows_platform_script test harness ([#1401](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1401)) ([4e11255](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4e11255e2de00fdc10bf7bb657582c30649c6650))

## [0.38.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.37.0-alpha...v0.38.0-alpha) (2025-12-11)


### Features

* added resource agent_identity_blueprint_certificate_credential ([#1294](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1294)) ([f587282](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f5872824e9a72f80dbd03b0ed13cb60efc7d2abb))
* added resource agent_identity_blueprint_certificate_credential ([#1295](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1295)) ([be4a92c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/be4a92c1c8a18c8a8354c6cb4b7781920f643e98))
* added resource agent_user and added soft and hard delete options for directory objects ([091519c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/091519c62237b7aef7d97d8a48d66c160467030b))
* added resource agent_user, service_principal_app_role_assigned_to, agents_agent_identity and added soft and hard delete options for directory objects ([#1344](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1344)) ([091519c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/091519c62237b7aef7d97d8a48d66c160467030b))
* added resource microsoft365_graph_beta_agents_agent_collection ([#1347](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1347)) ([746cb8e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/746cb8e21114d8ce5f35d30271e3de1f4544a727))
* added resource microsoft365_graph_beta_agents_agent_collection_assignment ([#1348](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1348)) ([ad38e14](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ad38e1439cc11755e9f8ad1c501f7d1573ea86bf))
* added resource microsoft365_graph_beta_agents_agent_identity and resource service_principal_app_role_assigned_to ([#1326](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1326)) ([2c6cc14](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2c6cc14af9524607766b7159cdbd81b83eeaa2eb))
* added resource microsoft365_graph_beta_agents_agent_identity_blueprint , microsoft365_graph_beta_agents_agent_identity_blueprint_federated_identity_credential , microsoft365_graph_beta_agents_agent_identity_blueprint_identifier_uri , microsoft365_graph_beta_agents_agent_identity_blueprint_password_credential and microsoft365_graph_beta_agents_agent_identity_blueprint_service_principal ([#1291](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1291)) ([0fa5ca2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0fa5ca23e8836315448b1986cee9f0996a6bb039))
* added resource microsoft365_graph_beta_agents_agent_instance ([#1346](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1346)) ([04606ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/04606acbf0e9fccceedc067fa38c9b28f1d87d1f))
* added resource microsoft365_graph_beta_users_user_manager with test harness ([#1345](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1345)) ([7ad12fe](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7ad12fecc8daf276f349f108cdc2c32546b42571))
* added resource service_principal_app_role_assigned_to ([#1314](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1314)) ([e9a85a9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e9a85a93d80780206cbff106fb908b0913702133))


### Bug Fixes

* added MEM to datasource utilityMicrosoft365EndpointReference ([#1279](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1279)) ([8b6c206](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8b6c20663ba33fcd1c5c6b6088c7d775a0d46840))
* **autopilot:** preserve hybrid_azure_ad_join_skip_connectivity_check for azure ad joined profiles ([#888](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/888)) ([1141f5c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1141f5caf7d5d57f3e15e7f9fd40a59914ee2a8c))
* **windows_esp:** resolve app validation pagination limit by querying apps individually ([#889](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/889)) ([04eda94](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/04eda94f838bfd073b61ec37255ae26667d4eb31))

## [0.37.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.36.1-alpha...v0.37.0-alpha) (2025-12-01)


### Features

* added datasource microsoft365_utility_licensing_service_plan_reference ([3c9b30a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3c9b30ac577922a3a5507a9cf2b54aca5ea536c9))
* added datasource microsoft365_utility_licensing_service_plan_reference ([#1262](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1262)) ([fba52df](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fba52dfbd9d6f39b5cbab476a3bb2f2af0c0c200))
* added new resources and refactored test harness ([#1183](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1183)) ([2b46754](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2b46754f9f2922b59112c545b72105ba8fc09e78))
* added resource user mailbox settings with unit test harness ([63b3c09](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/63b3c09d094f41918b6bc3521d78aafa22d6276a))
* added resource user mailbox settings with unit test harness ([#1277](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1277)) ([f5f5bd9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f5f5bd9982453e57e5a0422046f001457f753dd1))
* enhance run-tests.py for improved memory management and sequential test execution ([70d20e1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/70d20e109139bef763965d91d641a558947bb21b))
* enhance run-tests.py for improved memory management and sequential test execution ([#1231](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1231)) ([39bc074](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/39bc0748c78eb0040142a97dab8a26ce70fb4d40))
* enhance run-tests.py with configurable parallelism and improved test enumeration control ([e52e53e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e52e53e49b18adde1c9d0fee889ac020d56cb5af))
* enhance run-tests.py with detailed logging and timing for test discovery and execution ([883fd22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/883fd220ed79bada9e7de8c509344174675351ec))
* refactored graphBetaGroupLifecycleExpirationPolicy and added full test harness, added resource graphBetaGroupLifecycleExpirationPolicyAssignment and full test harness ([8abef15](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8abef1561c6c021910c7993e7306472331171201))
* refactored utility datasource unit tests to follow new test harness pattern ([e7de954](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e7de9542953f7f067da0ccfaa12b9952979995c3))


### Bug Fixes

* added helpers for stating and more unit tests ([#1261](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1261)) ([d67836f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d67836f60cfce6078be54fe8266eb041d071e643))
* added pipeline debug log option when bi weekly tests run in workflow dispatch ([#1215](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1215)) ([552d45f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/552d45f697974757ba55e08bd3e188532e094043))
* conditional access control set logic ([#1184](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1184)) ([1b507a5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1b507a519aa89ccf44587f7ff13ad55196506512))
* datasource rename to microsoft365_utility_microsoft_365_endpoint_reference ([1cac047](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1cac0472f035b480558c3897157eb0e7dca49740))
* datasource rename to microsoft365_utility_microsoft_365_endpoint_reference ([#1278](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1278)) ([b469e69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b469e6911d37f40161c1fbb7c98b29a873885b83))
* enhance nightly tests workflow by adding disk space cleanup step and updating runner configuration ([#1124](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1124)) ([0ada782](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0ada7823339ea526785a4f95798d0a5e7e96aee5))
* for group license assignments and added full test harness ([2980d32](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2980d325d9454ee6411de63f56849f6ba34d42fa))
* for group license assignments and added full test harness ([#1194](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1194)) ([40dc331](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40dc331ae3f1db3917171939ea9361ba27aea764))
* remove global_secure_access attribute in applications of conditional access policy ([#1182](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1182)) ([0dbbbe7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0dbbbe719433a6fb2df4ac8b421168dd988c3b25))
* resource  user license assignment tests ([eae4e70](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/eae4e701542d1a87e475bce12e176843249ecb16))
* resource  user license assignment tests ([#1106](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1106)) ([8b816fb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8b816fb8e5ab51caa46d6f759322d613804bc91b))
* trialling run-test.py in sequential mode ([aa92be8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aa92be81688ff346bfb6fb71c864e266ecbd3e0a))
* unit test refactor for datasource utilities ([#1220](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1220)) ([7c1faaf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7c1faaf1f8442f6ed4d4ca7f58ecc2467a56373c))
* update nightly tests workflow to bi-weekly schedule and changed to larger runners for memory intensive workflows ([#1213](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1213)) ([a39faae](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a39faae1b25cc0fcbe7e1bc4bccf7bd894d0655a))
* update run-tests.py to improve test enumeration and memory management ([1c7212f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c7212f9ba6ad3d5395b0ff13df4d0cf9d4dca3e))
* update run-tests.py to improve test enumeration and memory management ([#1232](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1232)) ([e8b8985](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e8b8985ecc56b20d99ccbb6161d026dfec63ee83))
* updated resource group test harness to follow new pattern ([4347f6f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4347f6f7b29d19d28389265f9dfac2593d626795))
* updated resource group test harness to follow new pattern ([#1123](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1123)) ([6f0b0d0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6f0b0d06c724945d7f98638d37c0a95fed161409))

## [0.36.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.36.0-alpha...v0.36.1-alpha) (2025-11-21)


### Bug Fixes

* go mod ([7de8336](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7de8336f02aabc8c2f5b11265285af5bd734efee))
* go mod ([#1100](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1100)) ([19fb8a6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/19fb8a6dd84e3c3f4e07a0cdfae5880581959215))

## [0.36.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.35.0-alpha...v0.36.0-alpha) (2025-11-21)


### Features

* add label management for test failures in create-test-issues.py ([5bf33c5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5bf33c57c31dcf04a45a5563511eb9fa2cd98983))
* added full unit test harness for resource microsoft365_graph_beta_identity_and_access_conditional_access_policy ([092323a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/092323ae006a66bd7718bcbd7adf4fcd92892200))
* added full unit test harness for resource microsoft365_graph_beta_identity_and_access_conditional_access_policy ([#1077](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1077)) ([32039a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/32039a72f75d5f951a67fa45f289663c2258d6b5))
* added new datasource utility_microsoft_365_endpoint_reference ([75737db](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/75737dbf90f72e23844d8fc0e5d64494d9f9901d))
* added new datasource utility_microsoft_365_endpoint_reference ([#894](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/894)) ([e7bc9d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e7bc9d267de3457cc955bfe00444b68f842b5dfd))
* added resource android_managed_device_app_configuration_policy ([0da497e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0da497ecc62aefe75f7db01c527ad71355738ea6))
* enhance acceptance tests for authentication context resource ([f166680](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1666809c49a7e1d4020a45026f8574af9e7d193))
* enhance nightly tests workflow with job-level failure detection and management ([323929a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/323929abfc323319e2f6a97709cbe7744587b21a))
* enhance resource validation and logging in acceptance tests ([cba1f19](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cba1f19527d79efbfd94beeffb3e03120be5c43f))
* enhance resource validation and logging in acceptance tests ([#1057](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1057)) ([69cde2b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/69cde2bad57fec799adefb11c034abf301e3ba60))
* implement issue management for test failures with manage-test-issues.py and remove merge-test-failures.py ([64a9d0d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/64a9d0db903ead416d0d5041c143027725083f7a))
* improve error context extraction in test results parsing for better failure diagnostics ([e85b6e5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e85b6e5bf12879d08e0a37649ceaeb865ded2617))
* improve error context extraction in test results parsing for better failure diagnostics ([#956](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/956)) ([f5b74ea](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f5b74ead684fccb781b785136d8c3bc264927122))
* multiple fixes for resource user, with correctly handling of first set of passwords. Also added full test harness ([676e860](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/676e860bee14614dc7277239f28e31685b4569de))
* multiple fixes for resource user, with correctly handling of first set of passwords. Also added full test harness ([#1099](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1099)) ([d0ff5d1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d0ff5d1ea2397591aa75d36de68b0be99aa7c556))
* refactor nightly test acceptance test pipeline ([91a92ab](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/91a92abd2a7548022d5d452de19536a037aadf6c))
* refactor nightly test acceptance test pipeline ([#870](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/870)) ([6388d61](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6388d6151feb5b50c1b7632b31c9307b5d0f8560))
* refactor of resource microsoft365_graph_beta_identity_and_access_authentication_strength_policy with fully functiioning test harness ([9ae5c6b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ae5c6b6303bcd601a5f05d6158a76b240d836ca))
* refactor of resource microsoft365_graph_beta_identity_and_access_authentication_strength_policy with fully functiioning test harness ([#1065](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1065)) ([cf252a6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cf252a6eb24541907bc14d68450f052232d7c88e))
* restructured validators by resource and att types ([2335fac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2335fac0443ebbf451f484f65d0ebd8e52d9430a))
* support filtering policy of network access ([#851](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/851)) ([a51820f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a51820fe4effd2edf556f3c837823ba93709c018))
* update issue title format in manage-test-issues.py to include 'Bug:' prefix for better clarity ([e07a38b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e07a38b23f55d52d7c72fe54dbdc7c2b7c8e79ca))


### Bug Fixes

* added docs for location-based conditional access policies ([c7342da](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c7342dac15779956e129586fe7d02374c888cd5d))
* enhance nightly tests workflow with job-level failure detection and management ([#902](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/902)) ([2271882](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2271882e567f3961cdd6c095432a5d08d33b4ff9))
* enhance unit tests for identity and access resources ([dd0d82a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dd0d82af9dba86a765cb51edfb55365db673baa1))
* enhance unit tests for identity and access resources ([#1078](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1078)) ([f430901](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f430901be5c06cdad9ab0a77d7ff2e651cf1b231))
* for datasource mobile_app_catalog_package test harness ([8c1071f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8c1071f5c392109560bf83d1ebc44b68fc521ae7))
* for datasource mobile_app_catalog_package test harness ([#885](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/885)) ([6f759db](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6f759db9399259ab2ae5fea52b0f0153ad301c9a))
* for jq: parse error: Invalid string: control characters ([5be22bb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5be22bb4c1cb0c842932cd93b24fb6418e335c46))
* for jq: parse error: Invalid string: control characters ([#896](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/896)) ([ce5a66d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ce5a66db4d2672001e1aeb38b061fbcffaf02d16))
* for multiple failing unit and acc tests ([#892](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/892)) ([26aeefb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/26aeefb365cd3ac20efc2399179bf84b60e35c8c))
* for resource user schema and refactored test harness ([3ba9c25](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3ba9c25ae956ee2a98119586376d01f0c92659ad))
* for resource user schema and refactored test harness ([#1079](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1079)) ([bf5ffb7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bf5ffb7b78fecaf6c8dd09950025667d42abdedb))
* issue management for acc test handling for nightly builds ([#901](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/901)) ([d8f7af4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d8f7af46ce6dd1f6f647fcca9395382aa3f41f3f))
* microsoft docs for conditional access. ([#1066](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/1066)) ([20c6403](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/20c64033d173c98e50d22e0a2cf9a4455ad49377))
* moved to kiota sdk. ([58d1347](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/58d1347d992e54dbb92348602d75d9a69e727e9e))
* nightly test pipeline to raise issues on gh when tests fail ([#895](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/895)) ([a7feb15](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a7feb150489a9c4df20d7abcdd1e19cb34ddbcd7))
* refactor pipeline scripts to python ([#898](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/898)) ([de2ea6e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/de2ea6eefc978756b08ae9119a14bf52cb82fbc0))
* script permissions ([#897](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/897)) ([cda6a41](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cda6a41a32b570c31a8034b48f9fdd8b7299530f))
* update M365 credentials in nightly tests workflow for improved authentication ([8242ea7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8242ea737903e0b38aa1212519ebfe1b2df162ec))
* update test directory path handling in run-tests.py for improved error messages ([087d4ed](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/087d4ed368a31d483234bf611d440b2cc48c41b3))
* updated CA policy examples ([53144ef](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/53144efed1786aa3f7180c946f67cf78fad320b4))

## [0.35.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.34.0-alpha...v0.35.0-alpha) (2025-11-12)


### Features

* add datasource Entra ID SID regex and update data source templates for consistency ([e166a4d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e166a4d179b96f7947b0dc4afc3b96505542e38e))
* add datasource Entra ID SID regex and update data source templates for consistency ([#864](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/864)) ([55f8a36](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/55f8a36dd48fd363064a0ec39d4be137c5c228a9))
* added new utility datasources entra_id_sid_converter and windows_remediation_script_registry_key_generator ([#865](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/865)) ([9e16617](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9e166177fa14fe32382d036aa22cf4cbc09e3209))
* enhance data sources with timeouts and new utility functions ([63e1f4c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/63e1f4c0ae241ed804f136d7b1645792c1476a90))
* enhance mapping of authentication strength IDs to user-friendly values in conditional access policy ([9900a49](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9900a4934abd45d2639ddacbde5d01890a3db9a1))
* refactored microsoft365_graph_identity_and_access_conditional_access_terms_of_use and added full test harness ([efdded8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/efdded8a3c1906019231d2dbbd252d8ac95c7068))


### Bug Fixes

* added test harness to microsoft365_utility_macos_pkg_app_metadata ([9a8b16a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9a8b16ab876d03fc3ecd4670e4e7c1c1877ccbce))
* changed resource named_locations to use dowithretry for reqs and updated docs ([#857](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/857)) ([b5ee323](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b5ee3239a47f6ca2f2c446c4685d7ce64365aaf0))
* changed resource to use dowithretry for reqs ([6dffc15](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6dffc150d30f6ef4d712e5cdb34d3ce656f6ccce))
* doc chores, refactored microsoft365_graph_identity_and_access_conditional_access_terms_of_use and added full test harness ([#855](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/855)) ([59f6c9e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/59f6c9edd7a25ad5a39f5826d719b947ee0f3e97))
* implement retry logic for HTTP requests in authentication strength CRUD operations ([39921d3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/39921d3f156cdc25d315050d1519311cee27e8dd))
* increase retry attempts for HTTP requests in conditional access policy CRUD operations ([ffe8a54](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ffe8a543f6a1a0e1a80c5dc2bc3b05baa4e53710))
* increase retry attempts for HTTP requests in conditional access terms of use CRUD operations ([c99cdeb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c99cdebd507446ecc6b49da547cf1b3d8f697c2c))
* naming pattern ([af54568](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/af54568132e0ca6367e056daf3bb3249431c80d3))
* numerous fixes ([b61c193](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b61c19323728aa4737fc307f11fae451ee442d0e))
* replace direct HTTP request calls with retry logic in conditional access policy validation functions ([5832bd5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5832bd57a038ea7e1f3c72ce467fa91dcb549662))
* retry logic for standard http client ([#858](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/858)) ([e513257](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e51325700b07eaffa01123e8dfe7f156e52637ba))
* update conditional access policies to include additional policies and improve timeout settings ([fd39d92](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fd39d920dd0e83ccefc77f0db58cb1991640b21f))
* update error message formatting for invalid tenant ID validation in conditional access policy ([99413c4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/99413c4883c353fa790e9510be5448a5b489d988))
* update test requirement check to exclude schema directory ([b38e1ba](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b38e1ba21235d37f58c3f1706e0eb34554614bd4))
* **windows_enrollment_status_page:** enforce false for reset/use/technician flags when blocking is enabled ([414c91b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/414c91bd862e6f17c74630b10ea2563521ec91aa))
* **windows_esp:** correct validation logic ([#854](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/854)) ([7a43f42](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7a43f427ae9088ca844cbd8fc8b2159dee0e0189))

## [0.34.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.33.2-alpha...v0.34.0-alpha) (2025-11-06)


### Features

* added retry logic for generic http client and numerous bug fixes for conditional access policies ([ad07337](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ad07337285367d919ea99048b7c0ba1eb8c8cf43))
* added retry logic for generic http client and numerous bug fixes for conditional access policies ([#849](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/849)) ([38206f1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/38206f1aac2d9e536e13388b204238ed0443b041))


### Bug Fixes

* **azure:** enforce camelCase for azure api resource ids ([1c09ee0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c09ee05064edb907407c01552abaa1dcf836533))
* **azure:** enforce camelCase for azure api resource ids ([#848](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/848)) ([7373ebe](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7373ebe87e97f3255419b07b97f6398416521449))

## [0.33.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.33.1-alpha...v0.33.2-alpha) (2025-11-04)


### Bug Fixes

* refactor of "graph_beta_groups_group" with more precise schema validation, constructor logic for creates vs. updates and full test harness ([3981df3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3981df302316b7707db4dbfc43b071716abb03b4))
* refactor of "graph_beta_groups_group" with more precise schema validation, constructor logic for creates vs. updates and full test harness ([#846](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/846)) ([87777da](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/87777dadf37553cf91730428bb5431e02c6ca030))

## [0.33.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.33.0-alpha...v0.33.1-alpha) (2025-11-04)


### Bug Fixes

* go mod tidy ([#844](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/844)) ([ffcf858](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ffcf8588678068c210d4fe27a46b5a6c8ddc67b0))

## [0.33.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.32.0-alpha...v0.33.0-alpha) (2025-11-04)


### Features

* added 30+ actions managed device. testing harness to follow ([afafe6a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/afafe6a7d145a116c2b55637fc6d73dfa35dc3a8))
* added 30+ actions managed device. testing harness to follow ([#833](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/833)) ([42cdd34](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/42cdd3420eec666be8d880be1e1e396a868b38e3))
* added action apply_cloud_pc_provisioning_policy and fixes for cloud_pc_user_setting ([8eaf974](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8eaf974bb610154e88fd88ebae5ac70e34d136fd))
* added attribute_set ([60d8dde](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/60d8ddee9a64a466c821735aab545267c5dd7b55))
* added datasource directory role definitions ([28c50a8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/28c50a872f456dd70776e5260d1f9add356f0557))
* added datasource directory role definitions and fixes for miss named datasources ([#830](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/830)) ([994a608](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/994a60804f9e91cde84f95589f51c7aa4816147b))
* added microsoft365_graph_beta_device_management_policy_set docs ([5286a1e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5286a1eb4c380b9920dce62c460ea28eb954431c))
* added microsoft365_graph_beta_identity_and_access_custom_security_attribute_allowed_value ([9061ae1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9061ae1b0f8b5d0f108ad569bd15a80ac72ac5e3))
* added microsoft365_graph_beta_identity_and_access_custom_securiy_attribute_allowed_value ([#829](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/829)) ([cef2d52](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cef2d529106b5527313b045d546871a58617a2aa))
* added resource apple_configurator_enrollment_policy ([a88130f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a88130f9843849d0a464bd602d1f519539c4919c))
* added resource apple_configurator_enrollment_policy and datasource mobile_app_catalog_package ([#817](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/817)) ([3ddecd8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3ddecd87d66e90e9d9a738729d804546a332968e))
* added resource identity_and_access_attribute_set and identity_and_access_custom_security_attribute_definition ([#828](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/828)) ([50f9864](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/50f9864eba9a99e9abe944ddbdeb1a885c583617))
* added resource intune policy_set ([6a51cc3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6a51cc3a637076d19951efe849086c3af24aa7ad))
* added resource microsoft365_graph_beta_identity_and_access_custom_security_attribute_definition ([2b144bc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2b144bc82ff40f468bbe6b561057d9c0d1b4c4f5))
* Support ip application segment ([#842](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/842)) ([ba4971c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ba4971c7ff52d800c45dc2a87003ba4178c9da39))
* Support new attribute of conditional access policy ([#841](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/841)) ([77afd41](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/77afd41e210fee01c749527f32aa67482c4e885f))


### Bug Fixes

* added additional validation and missing fields and timeouts ([402a7fa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/402a7fa6543523e97d02f9f076a7f7692eb67618))
* added cloud type for TestAccM365Provider_InvalidProxyURL ([1244edf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1244edf434f394ecbb789f120a0152e25e916dbb))
* added exclusion ([a33d569](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a33d569512be966edf2117dbf333e5cf3fdd13d1))
* added fixes for update handling ([f5d0b45](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f5d0b4572f5d63ec22abfd1580bef7bb122b9d07))
* added full odata  query support and page itteration to datasource managed_device ([3809f2c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3809f2c76bc59b698ce9ea2ce299447c1c3afb39))
* added full odata  query support and page itteration to datasource managed_device ([#831](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/831)) ([2ac60e2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2ac60e28dbd9f708c732a68faae66fb126f3ed32))
* added microsoftwindowsdesktop_windows-ent-cpc_win11-25h2-ent-cpc-m365 to resource graph_beta_windows_365_cloud_pc_provisioning_policy ([a7a8329](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a7a83297a0c9e4432e40626bbd32e331155eb835))
* added valdiation for display_name ([1b69f43](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1b69f436c4a05da154ec666d267e52e81a49e7fb))
* added validation for frontline type with provisioning_policy_id ([e85e622](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e85e622a97b412323334947ea3d9dff41d97731c))
* changed behaviour of local_admin_enabled . ([f593884](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f593884d28426a6931659163a983f4a481cc5290))
* for ca constructor field include_users ([4dc908f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4dc908f786a945152fae25afd83c05314d76e348))
* for conditional access validations ([d70db24](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d70db24cba02492ba07ff05f463590f9e4da25ab))
* numerous fixes for conditional access ([#827](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/827)) ([e42b9a9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e42b9a9e196dd43251bb9fe2f2f2e7c2883fd420))
* numerous fixes for windows_365 resources ([#819](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/819)) ([07f6699](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/07f6699366ea4ae4f330c64264210c19e93e0963))
* resolve Warning: Both go-version and go-version-file inputs are specified, only go-version will be used ([0199e04](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0199e045137755be5eaa965c494270e02d6fde90))

## [0.32.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.31.1-alpha...v0.32.0-alpha) (2025-10-05)


### Features

* added graph_beta_device_management_windows_backup_and_restore ([bd318ce](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bd318ce494f1ccb88f71b3b8a03a4e232c80c30b))
* ground work for windows_backup_and_restore ([670ce4a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/670ce4a88b12f1b289a51bc3fdcf11bb7b5699d4))
* ground work for windows_backup_and_restore and doc fixes for actions ([#815](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/815)) ([62c5340](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/62c5340fce676a1557df261e76f6b13e02090525))


### Bug Fixes

* added Windows 11, version 25H2 to windows_feature_update_profile Resource ([d60f511](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d60f51155af1aeee3d3be8cabb73000fef64317d))
* for microsoft365_graph_beta_multitenant_management_audit_event Ephemeral Resource ([c8e0529](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c8e05296c29c7c2d5aa07bc63c2fe1f86f173637))
* for microsoft365_graph_beta_multitenant_management_audit_event ephemeral Resource ([#813](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/813)) ([4ba0f13](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4ba0f13ea1af19c92ce01e5969f188ec637acd18))
* linter fix, package restructure with docs bump ([d1f527a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d1f527af4e4a3d3574a8a4f0d3b2f180ab734fd9))
* linter fix, package restructure with docs bump ([#811](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/811)) ([856856f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/856856fbbcf6aa123cc3a2a8ec6800357d4c655e))

## [0.31.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.31.0-alpha...v0.31.1-alpha) (2025-10-02)


### Bug Fixes

* go mod tidy ([63d78c1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/63d78c11f0e652666aac83f3ac0ae0ddd35d8306))
* go mod tidy ([#809](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/809)) ([429e22f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/429e22f1e88906b0bc32bd1a0bcc96eb2dc55d06))

## [0.31.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.30.4-alpha...v0.31.0-alpha) (2025-10-02)


### Features

* added actions and restructured docs ([7d0d98f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7d0d98ffe81884c14684228420e83334939245af))
* added actions and restructured docs ([#802](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/802)) ([1416d59](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1416d593967750133a0725173072cd44ea4b5d71))
* added actions for windows_autopilot_device_identity in prep for tf v1.14.0 ([e06a1b9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e06a1b995af3c031ace47cc5b5de8def29e8ea77))
* added conditional_access_terms_of_use ([363bf7d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/363bf7d471855977bdd87cbc2ebe801a2d67a77e))
* added conditional_access_terms_of_use ([#807](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/807)) ([b147bc9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b147bc9cfe82ed518e6fd3ac732070aadd2688cf))
* added data-source microsoft365_graph_beta_applications_service_principal ([08ab815](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/08ab815855dadf4941f51645973331ae983c4cb1))
* added data-source microsoft365_graph_beta_applications_service_principal ([#805](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/805)) ([0852753](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0852753e38f4a2353bf3f195c881710287f071ce))
* added resource authentication_context with test harness and updated docs ([2c139a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2c139a7119529ba729707d965e49d849a308537a))
* added resource authentication_context with test harness and updated docs ([#806](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/806)) ([0bc0452](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0bc0452cb1e7760ad543c5e83aa5fc4267b32bfc))


### Bug Fixes

* Export environment variable Client ID ([90b8d90](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/90b8d906e8e9fa1e40cbd98bece0fc000b1a7dbe))
* Export environment variable Client ID ([#808](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/808)) ([49b2c1b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/49b2c1b8fea97d8514f1c4f41bbeccf50bf09f51))
* any -&gt; any ([8074ff4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8074ff43bff9ffd6714f56fe281a418a438f9f59))
* restore main ([e3ddd31](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e3ddd3140c8e9f2729f8ead1d1ef09bdc0d38054))
* schema refinements to improve dev ux ([4e132b5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4e132b5f154d7afc349b08fccf45af573ebdbb3c))
* test for future regression ([d319a1a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d319a1a59a80874bccbdf5b6f4fc6a0244ec3abe))

## [0.30.4-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.30.3-alpha...v0.30.4-alpha) (2025-09-29)


### Bug Fixes

* map[string]any -&gt; map[string]any ([dbe67b1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dbe67b18f48b5a9643ccb7f802db55792cf49174))
* map[string]any -&gt; map[string]any ([#797](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/797)) ([7f7434f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7f7434fd60327fcbcd4e47e68bd243a2ecb69773))

## [0.30.3-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.30.2-alpha...v0.30.3-alpha) (2025-09-29)


### Bug Fixes

* added more logging ([7927d4e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7927d4e95374fc057e4e26aaff39fbf03a508935))
* markdown docs update ([824449a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/824449a956b009827022bcd404f2580e0fca2ed5))
* more debug logging for oidc ([#795](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/795)) ([13613b8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/13613b8520a34660a44411b1a3e3ad051c06aa99))

## [0.30.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.30.1-alpha...v0.30.2-alpha) (2025-09-29)


### Bug Fixes

* Support GitHub OIDC authentication in Terraform Cloud remote execution ([52b6ed9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/52b6ed99da9cbc3f546d1470ed0dfed82712c4bf))
* Support GitHub OIDC authentication in Terraform Cloud remote execution ([#793](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/793)) ([25ae985](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/25ae985cbbf5ff9fa4d024eca89dfdc7da07623a))

## [0.30.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.30.0-alpha...v0.30.1-alpha) (2025-09-29)


### Bug Fixes

* comments ([5ac269a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ac269a8ee22a424d6f88633f78ab894e7784b87))
* docs ([c938a25](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c938a25746a155aed05ce1290b10f3a665704b84))
* docs and comments ([#791](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/791)) ([97f135a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/97f135aaaf61a9788cbbc18a3c55ed9faf8d31ff))
* tidy up comments ([48b18c7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/48b18c7dfa5834a72b4a50d11afaeae16d47d521))

## [0.30.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.29.3-alpha...v0.30.0-alpha) (2025-09-26)


### Features

* 1st itteration of testing ([e36b6a6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e36b6a63ae9277b165399f5934d0ca1bf573b541))
* added office suite app with examples ([331664a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/331664aa4bc630726016d2b60eaed7a21ff458f5))
* added office_suite_app with examples ([#770](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/770)) ([6f5d1ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6f5d1acc80b413554362480ea5d14be6b65b2953))
* added unit and acc tests for windows_autopilot_deployment_profile ([cb04b34](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cb04b344f853f47fc1f9e8c53a14f9543b2bda71))
* refactored windows_autopilot_deployment_profile ([cc98227](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cc98227365c97d92bf94318580c5ee6fe4ecce80))
* refactored windows_autopilot_deployment_profile and added unit and acc test suite ([#784](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/784)) ([096b8a5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/096b8a5c6977a35d8dee10371c9bed8b58eb024d))
* updated docs ([7ca77fa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7ca77faa2260687e5d807a906a894eeec1708d27))
* updated docs ([#787](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/787)) ([1e80022](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1e80022ddd92e0d376332fe74a60af57e7b2c6b4))


### Bug Fixes

* added gp multi text value examples ([018ab4c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/018ab4caef143593585ea7be271949ce533dcec0))
* added mising repo badges ([f31025e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f31025e002d28ee13b8eb2d78f9a4d760eceb4f8))
* added mising repo badges ([#786](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/786)) ([1e7ff4e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1e7ff4ef19629d18a466ed3944b8c93730fd63a0))
* docs ([3ad48dd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3ad48dd13f794e1e43db92f056ca8f487f615f8d))
* docs ([#788](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/788)) ([8cd9c41](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8cd9c41c3d284a25eb336ca96cec333bc8a8a488))
* for github_oidc auth ([a894899](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a89489981e907ab6ed56ba12ae8e0ea8cde4a349))
* for github_oidc auth and updated relevant docs ([#785](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/785)) ([02c1d92](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/02c1d92dae824b9fab74158a42d00b23af842e43))
* group policy descriptions ([7480155](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74801555cdbaee022c7ef71fe3a98e2be8e08b95))
* pipeline permissions and go lint version ([7b5256e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7b5256ed44a845447b908bd6cd33485ab3129f7b))
* removed trailing white space in file name ([f21c1d4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f21c1d4b94d2f0b0e6f591921519ea18cacb5f0e))
* standardised usage of setup go steps ([4b02dcc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4b02dcc7e8b42b535264132f342604e4e382cb25))

## [0.29.3-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.29.2-alpha...v0.29.3-alpha) (2025-09-13)


### Bug Fixes

* alignment of go.mod version ([e3ad76c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e3ad76c1f80e91f1bd1cf9624f1af44747192df8))
* alignment of go.mod version ([#768](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/768)) ([e1eb436](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e1eb4366f26bd850bfc45f2f4b51067ce3c5f73c))

## [0.29.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.29.1-alpha...v0.29.2-alpha) (2025-09-13)


### Bug Fixes

* goreleaser setup go ver ([a35f8ed](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a35f8edc4b997ce5967312295367ba6ce764c5f8))
* goreleaser setup go ver ([#765](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/765)) ([80ec02b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80ec02bb489a265ab7f5aa466efa0152e65ea017))

## [0.29.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.29.0-alpha...v0.29.1-alpha) (2025-09-12)


### Bug Fixes

* go mod tidy ([8f33e23](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8f33e23743fedfed264039c8e267dc38d2331dc2))
* go mod tidy ([#763](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/763)) ([0f2d37a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0f2d37a3bf023f81690ee161e1ef8ec2d187c248))

## [0.29.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.28.1-alpha...v0.29.0-alpha) (2025-09-12)


### Features

* added 1st itteration of group_policy_configurations ([b2b8078](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b2b8078492e8b1cea411c905124d1ce322dc3799))
* added centralised resolver logic and added docs ([46c4004](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/46c4004aa9e6ef48b9fd52e35757de9cbc6660a3))
* added group_policy_boolean_value ([88b6b77](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/88b6b77a0ac6118d12d6676182570eeb42dbc307))
* added group_policy_configuration, group_policy_boolean_value and group_policy_test_value as new resources ([#762](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/762)) ([33d3fc6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/33d3fc61229fb93e41e0c8d71dcad5319fa19ca1))
* added group_policy_text_value ([b1900d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b1900d2e6605450aa4a80bb53bad6d78d3b94baa))
* added group_policy_uploaded_definition_files ([5ecf420](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ecf4202f434aef1ef47f539adad33724401b5d2))
* added resource group_policy_uploaded_definition_files and datasource group_policy_category ([#753](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/753)) ([9ce3d0d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ce3d0d7097046a0f41f6c75505bf94275523e21))


### Bug Fixes

* map[string]any -&gt; map[string]any ([fcd9a51](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fcd9a51dc63fd02f9d001caab20d2fa515374209))
* numerous ([f74e41e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f74e41eb964e98eb66870f71d9f7e2cd0b3d3ace))

## [0.28.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.28.0-alpha...v0.28.1-alpha) (2025-09-06)


### Bug Fixes

* docs for authentication strength and named locations ([2fabdf8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2fabdf82c7c28d6e3f66653b8bee1f7eb9e88f44))
* docs for authentication strength and named locations ([#751](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/751)) ([759a5fd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/759a5fde3f8173b56714b276c9a99907b4a3b618))

## [0.28.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.6-alpha...v0.28.0-alpha) (2025-09-06)


### Features

* added app_control_for_business_built_in_controls and app_control_for_business_policy with examples and test suite ([#748](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/748)) ([08effe6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/08effe61ad2df9c96f421838a2a811f548c9bc80))
* added app_control_for_business_managed_installer ([c7761b5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c7761b5a4ccda439565f778cee5e32b73c1fa486))
* added graph error handling for non kiota http clients ([e41c7b8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e41c7b82c87f49abc6fdf64972b67614f2d23d44))
* added graph error handling for non kiota http clients ([#734](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/734)) ([c09ea43](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c09ea4363b4b861ee91bbe427a1c168079b85a67))
* added microsoft365_graph_beta_device_management_app_control_for_business_built_in_controls with examples and test suite ([e1b5086](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e1b50861bd44d30ee3631fab1fbfda2d10722de6))
* added microsoft365_graph_beta_device_management_app_control_for_business_policy with test suite and examples ([27df46b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/27df46b7d24b40bc0f52615b162d9c4c63438743))
* Enrollment Status Page support for installing Windows security updates during Windows OOBE ([8b5d6a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8b5d6a7bc08db4287291583373a7d44baaa2c8b0))
* Enrollment Status Page support for installing Windows security updates during Windows OOBE ([#749](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/749)) ([1c2ec22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c2ec22407fce59b93e43a273119159322c2a093))
* windows_device_compliance_policy ([5894296](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/58942963c1233ac047872b7b5f22ece85cbda471))


### Bug Fixes

* added docs and unit tests for app-control-for-business ([#746](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/746)) ([6195f3e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6195f3eca23229a2df23ce513701ec93fb6797b6))
* app-control-for-business unit test fixes and test harness ([ebcee04](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ebcee04b05da0ed453201633adf286a3df6df793))
* for conditional access ([fd2c0a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fd2c0a7fac16431aa2466bf7b0cd5085554e92f3))
* for conditional access ([#743](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/743)) ([90080b4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/90080b44570f28d29ba04060ee6aa110ba9de411))
* for set value handling for role scope tag id sets ([21840d7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/21840d78c917ba9fc19f19d1ffcc45110fd26301))
* for state logic for conditional access policies ([5698dd3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5698dd395a998e5d3352d4ed39d9ab7527ff89c3))
* missnamed resource and docs ([#731](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/731)) ([5f92828](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5f92828186f5440b2150a901ed4535c29317631a))
* tidy up docs and removed old resources ([#729](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/729)) ([4f6dfd1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f6dfd1f6dfda2f6284ee33337fda2227a1c8032))
* unit and acc tests for windows_compliance_policy ([83e1cb5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/83e1cb58b07264fa605e9197796e9a9e808fea33))
* windows_device_compliance_policy refactor ([#732](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/732)) ([a3082b3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a3082b3dd708ca89ecb89e205938de3f592866f0))

## [0.27.6-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.5-alpha...v0.27.6-alpha) (2025-08-21)


### Bug Fixes

* trying parallelism 1 with ubuntu-22.04-arm-4c-l ([#727](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/727)) ([57f40cc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/57f40cc73ecdfb916be1fe7c5223b5dcfab75eec))

## [0.27.5-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.4-alpha...v0.27.5-alpha) (2025-08-21)


### Bug Fixes

* 1st itteration of autopatch_groups ([#725](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/725)) ([a9ef4e2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a9ef4e27a9547204dda60a014e3fb4625522514c))

## [0.27.4-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.3-alpha...v0.27.4-alpha) (2025-08-21)


### Bug Fixes

* runs-on: ubuntu-24.04-arm-4c-l ([#723](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/723)) ([22f431f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/22f431fe331e792efb7d62f5614b2d973c6bf460))

## [0.27.3-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.2-alpha...v0.27.3-alpha) (2025-08-21)


### Bug Fixes

* added unit and acc tests for windows_device_compliance_notifications ([#721](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/721)) ([0546e4f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0546e4f637b822eb6c44fc142661cc2331229196))

## [0.27.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.1-alpha...v0.27.2-alpha) (2025-08-20)


### Bug Fixes

* refactored WindowsDeviceComplianceNotifications ([c63f934](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c63f934ed694ab291c66ec72457d12a249902961))
* refactored WindowsDeviceComplianceNotifications ([#718](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/718)) ([2bcac69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2bcac697ed9349ab456da0914e713911decfb98a))

## [0.27.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.27.0-alpha...v0.27.1-alpha) (2025-08-20)


### Bug Fixes

* for unit and acc tests ([a9ecc32](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a9ecc32f0c462bacbfba35426f0074b8afa9c3c3))
* for unit and acc tests ([#716](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/716)) ([c9859f2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c9859f2fd18669cf8d9b72ffaf9bd860877243df))

## [0.27.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.26.3-alpha...v0.27.0-alpha) (2025-08-19)


### Features

* added android_enrollment_notifications with unit and acc tests ([#708](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/708)) ([6c67540](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6c675408d743e204b1b3449148ec0264256c6829))
* added dynamic dcv2 settings catalogue build depth ([1a5e24a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1a5e24ac3a51c6496ca62081ec19bfc498c18b32))
* added dynamic dcv2 settings catalogue build depth ([#697](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/697)) ([ad3db91](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ad3db9182376c8f2ca758dd17b7db44cd672a583))
* added notifcation messages ([3d0f75b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3d0f75bb9df75eb0d97b264abd31a2a6b546368a))
* added unit and acc tests for windows_driver_update_profile ([db51ffc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/db51ffc7f13aad931732c91e8573dcab97326645))
* added unit and acc tests for windows_feature_update_profile ([cb16457](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cb164573b68d20d0cd8af602f21390faa6969443))
* added unit and acc tests for windows_feature_update_profile ([#700](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/700)) ([782efa5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/782efa517847d8abc15e9f1db0285bfc2e6f5828))
* added unit and acc tests for windows_quality_update_expedite_policy ([76b4ccf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/76b4ccf9228160b1ceb6b419439c20a4ec8cb030))
* added unit and acc tests for windows_quality_update_expedite_policy ([#699](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/699)) ([a8fd98d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a8fd98d8b8e6536c29c0d483e1a7405a985486cd))
* added unit tests for string validators ([ce6e13b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ce6e13b18d65bacf7dd04d5b81b4033b8ce68f44))
* added unit tests for string validators ([#702](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/702)) ([1c9f3ce](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c9f3ced4edc7b7c1eb0bb5881eac263a920a696))
* device clean up rule validators and refactored acc and unit tests ([#704](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/704)) ([3de342d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3de342d59b62c9492b3d14f561d3116e2521c48a))
* implemented macos_device_configuration_templates with examples ([eec3d36](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/eec3d36febabd24c15eac93afa6edce9759e4423))
* implemented macos_device_configuration_templates with examples ([#715](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/715)) ([25fc4e6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/25fc4e66a1de7f12ec800afbedd6b101add098c3))
* implemented notification messages with fully passing test suite ([faaa4a0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/faaa4a0ff93829e8838a8584493f5468c6886e05))
* implemented notification messages with fully passing test suite ([#711](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/711)) ([43be9ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/43be9ac5d41a51da4658b9d08375b978e033fe7d))
* new acc and unit tests for provider config ([fa62672](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fa626726850d749f8991c29ff4d25322f8cdd058))
* new acc and unit tests for provider config ([#694](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/694)) ([2071cc1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2071cc1ad85bb9b8cacab88f14a9923d0e0b31dc))
* refactored unit and acc tests for macos_custom script atts ([a413905](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a41390594edfd085b7687942f72f0ef752725d80))
* refactored unit and acc tests for macos_custom script atts ([#703](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/703)) ([ad05c51](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ad05c511103a9987d2ad0d1022c7be02c6eebd9c))
* settings catalog json unit and acc tests ([8d6a6e5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8d6a6e5b2612d6633325c1c88dc2a9cb34a6704c))


### Bug Fixes

* added new pattern unit and acc tests for various windows script types ([#706](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/706)) ([095ed32](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/095ed32d99a0a6381dee555b9f59e7e839c0a546))
* added unit and acc tests for device compliance scripts ([20f0990](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/20f099087c570f5ea094090af93e25f425d85f14))
* added unit and acc tests for windows expedite policy ([109836c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/109836ca7c88b9ff36939a51fb1f516f91bc6b3a))
* added unit tests for utilities ([c991a4d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c991a4d63b3d75330f2c2e20490a8c7994af5d36))
* added validator for singular resource uniqueness for managed_device_cleanup_rule ([6591dfc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6591dfc371c33313e67eba0c5999a0530e5c0768))
* device clean up rule acc and unit tests ([e6ca94f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e6ca94f524d15d469d83649fa7c92c8b36fea564))
* exclusions ([a0d0aa4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a0d0aa46de70af27651aedbcbd603aa0209cb332))
* for acc tests for win quality updates ([5b746b4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5b746b43482b01a15f5125d21e0a0ba8f3dbeaf3))
* for locales to lower case ([5f99f57](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5f99f57461cf76383d3fdd9b72afcb72039c7e8e))
* for notification messages and testing ([5266ba0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5266ba0c90f138fcd76ec7ff51a7c3b0e29490b0))
* for update logic ([19815a3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/19815a34382d8cfb4840823d6063df9e35df6972))
* for windows_feature_update acc and unit tests ([ca83357](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ca833579f6b0036992a182d2108738894f539cf3))
* for windows_feature_update acc and unit tests ([#701](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/701)) ([75549e0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/75549e0a74351305d9edc754ed607ed2870668b1))
* migrated to new unit test pattern ([776faab](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/776faab6119f55d8bf756a73be754b57bc69fceb))
* migrated to new unit test pattern ([#698](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/698)) ([dcf641d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dcf641d8267470ef71da8b0853570e34fc3c7199))
* numerous unit test fixes for settings catalog for schema recursion depth ([80ef819](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80ef8198c3d762540c9fc5b0eacc407c5392852e))

## [0.26.3-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.26.2-alpha...v0.26.3-alpha) (2025-08-12)


### Bug Fixes

* trying with 1 ([58d6768](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/58d6768ba0f0dec2ef49e174f81b987cf6b666e3))
* trying with 1 ([#692](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/692)) ([439e5e3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/439e5e3e39655238c60925c5fd304d9b4f2cadf9))

## [0.26.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.26.1-alpha...v0.26.2-alpha) (2025-08-12)


### Bug Fixes

* release pipelines ([6d7444b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6d7444b5dd26358b5e6164bb06cac404806ca485))
* release pipelines ([#690](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/690)) ([09fb398](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/09fb398d928d463eebaee8a5421a1afc49a2dab3))

## [0.26.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.26.0-alpha...v0.26.1-alpha) (2025-08-12)


### Bug Fixes

* deps ([#688](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/688)) ([7a59f87](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7a59f872e2bdeda8c7b6f3219bde9957ef251fa6))
* deps bump ([2f43ef9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2f43ef933a5da665773b363798311d8453030a84))
* pipeline schedule ([8650af6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8650af66c9e0e2b831b84db3418b71b7a4005d59))

## [0.26.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.25.1-alpha...v0.26.0-alpha) (2025-08-12)


### Features

* delete/read with retry based on status codes ([347f874](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/347f874aa64a139fdf04f18066bb6238bb8a5886))
* export settings catalog templates via powershell with an export all option added too. ([e28ec4d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e28ec4dee789bd1a5e549f7c8d7fdbb94fe34135))
* moved goreleaser to ubuntu runner and pipeline fixes ([420e342](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/420e3420c6e96194d69c6033bcc3db39d23f6ca6))


### Bug Fixes

* added unit tests for mobile app helpers ([664aa35](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/664aa350fef960595e41979e05949aef2155e962))
* bug fixing assignments ([9b71c8f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9b71c8f47977202f0e8406b313433b96a2eada54))
* for http methods ([c99cd96](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c99cd9676d2f1f1bca3acd25b76ab55557100692))
* for settings catalog unit test timeouts ([#682](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/682)) ([8a8cbc4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8a8cbc4ba15d651e2a25a4323778f5820d24c555))
* various bug fixes for pipelines and settings catalog, with improved exporter. ([#673](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/673)) ([7b101ff](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7b101ff8cf42424b3ae7ab10beccc5d85fe24797))

## [0.25.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.25.0-alpha...v0.25.1-alpha) (2025-08-08)


### Bug Fixes

* for pipeline ([adc888e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/adc888e6e082f9b28a439a58b4b78f5128f016b0))
* for pipeline runner ([#671](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/671)) ([3987281](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/39872811eb2b35a1362599ec5389434ea5e468ef))

## [0.25.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.24.0-alpha...v0.25.0-alpha) (2025-08-08)


### Features

* added cloud_pc_role_definition with acc and unit tests ([62df1c6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/62df1c6adb9262f815d885c4ef85651212c0daf6))
* added more unit tests and updated testing pattern ([b4e6a6d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b4e6a6dc3dea4735bd9e8ab6ca811131293f3a7e))
* added unit and acc tests ([f7dd14f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f7dd14faec64e3d7641e42d10528120f79e46d89))
* added unit and acc tests for assignment filters , role scope tags and device categories ([#653](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/653)) ([affc09a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/affc09a166792006bdc34e3dcf2adbeba069a6e8))
* added unit tests and bug fixes ([5146459](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5146459e08e24f80145fd0db8931c0a5a99730c6))
* migrated settings catalog resources to the new pattern for assignments ([80693ed](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80693ed0a29f49e86b50b982232692934cac8179))
* migrated settings catalog resources to the new pattern for assignments ([#665](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/665)) ([f08713e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f08713ead36e297ffd9698e1b96835a8e93ef338))
* refined test harness for windows_remediation_scripts and windows_update_ring ([d01f3fa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d01f3fab38affcceb131fa5832490132ab4c5dd6))
* unit and acc tests for terms and conditions ([1cff2ee](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1cff2ee8e98a2c14a32bafe915081738201dc70c))
* updated testing pattern with centralised tf resources and external tf provider dependancies ([#664](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/664)) ([b5cab67](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b5cab67bcc1be9e5e108f1dd2d5ba6af2d8d5f88))


### Bug Fixes

* added unit and acc tests for role_assignments ([#661](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/661)) ([46bcbf5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/46bcbf5e70558ab7650563945afe221956b3949a))
* changed file path. cant be bothered ([befa30c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/befa30cf8b72c3c0a0f958f11c67e0fbe25ddbee))
* for role definitions ([eb7f4e6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/eb7f4e6e450c46760c1a82f770040b98d9488cee))
* for role_definitions ([83dcd6e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/83dcd6ef99f20e51114698e5785f7e268f938124))
* for role_definitions and tests ([#660](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/660)) ([ad8c5fc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ad8c5fc8c59077928c4e4b3cae671590ff53e512))
* for unit tests ([b7ad6ea](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b7ad6ea3da16f41fa44cd2bd8d85a15cd8b37699))
* initital unit and acc tests for settings catalog ([c53c4a3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c53c4a3ff3908198719a83c340eacb9d7507b57f))
* restructured file paths ([a000b35](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a000b3560663c1ef003e2e8acf107752afcf414b))
* settings catalog unit tests ([#667](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/667)) ([9acd789](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9acd7892809bb5b98ab62921b9cdeaf93e1da4f7))

## [0.24.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.23.0-alpha...v0.24.0-alpha) (2025-08-01)


### Features

* trialing code-cov ([#641](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/641)) ([1a1a8e0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1a1a8e0586063ddd7d71ccca6c197870245683da))


### Bug Fixes

* for macos windows script unit and acc tests ([5b6d120](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5b6d12092a1acfb7ed8513a5d46604db727b13a6))
* for testing strategy ([593d1b5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/593d1b53ececd5846c4a8961b4afe1d748128367))
* for unit tests ([da2533f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/da2533ff9f41b72247535d2fc0c21122422f32eb))
* implemented unit tests for macos scripts ([89d5532](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/89d55323871f2c624d9fbf9c4a67648538398fb5))
* implemented unit tests for macos scripts ([#647](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/647)) ([4fa0ec7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4fa0ec7bd75706870eeee65b96e27879ee3cf07e))
* permissions on pipeline ([e5e46b6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e5e46b62c9aaf5f38e23a5fa5199dc8ce52368d0))
* pipeline test strategy ([#646](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/646)) ([57b0fe3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/57b0fe380d467245ca0fd6e5cf1285d7545e610e))
* removed redundant assignment resources ([49c34a3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/49c34a3311c3a75d51aea06d5e401347805d7097))
* removed redundant assignment resources ([#645](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/645)) ([0b7accc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0b7accc24c3beabaccc8fd2014c55efd10ebcb78))
* token for auto merge dependabot ([836a205](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/836a2051072cfd36ae1096e9d0515486b5f18963))
* trials with code-cov ([630db22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/630db2227784f410ada2f23ec6f7d2e091021558))
* windows remediation scripts ([7d1157b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7d1157beb95f0fd17ad56f10c0314886e45593be))

## [0.23.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.22.0-alpha...v0.23.0-alpha) (2025-07-31)


### Features

* app protection policy assignments (mam) , reinstated windows update rings ([#627](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/627)) ([949f6f0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/949f6f04bc1907533d3aa8869b776f9f786c332d))


### Bug Fixes

* branch merge ([#634](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/634)) ([82ee18e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/82ee18e3c3a609d43df78bbe7173b29430b02922))
* unit tests ([#635](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/635)) ([8f1b315](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8f1b31508633b5ba9c0cc82ff33fc561626a80b2))
* unit tests for windows 365 and assignment logic ([#638](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/638)) ([47c26af](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/47c26afa78a634986f6eac891cee53ee17449581))

## [0.22.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.21.3-alpha...v0.22.0-alpha) (2025-07-21)


### Features

* added settings catalog to hcl generator script ([#615](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/615)) ([762ab6f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/762ab6f349d9c9ce111deecd4ba953c42c164a84))


### Bug Fixes

* final set of fixes to handle a list assignment response as a set ([#623](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/623)) ([47fb051](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/47fb051ea448c5d6da025d9aaf5f80a02594d643))
* for settings catalog simple settings with secret values and added schema validator to only accept "notEncrypted" for requests ([#617](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/617)) ([f40a8dd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f40a8ddd8753cdeba4d755d6c1972a78c4456959))
* for state handling of assignments for macos_platform_script , macos_custom_attribute_script , windows_remediation_script ([#620](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/620)) ([decfee9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/decfee95177344aa29c5aa58e863a22a0c166f1d))
* specify provider name in tfplugindocs go:generate directive ([#621](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/621)) ([5cffa73](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5cffa737074d85cc0f7f03716b58ffd98912d6c7))
* updated powershell hcl exporter for settings catalog to support complex types ([#619](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/619)) ([9d20c88](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9d20c88e5aba1a2adc7e238affb1e2f2e865af40))

## [0.21.3-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.21.2-alpha...v0.21.3-alpha) (2025-07-16)


### Bug Fixes

* added assignment removal to deletion of CloudPcProvisioningPolicy resource ([#610](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/610)) ([aeb93c7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aeb93c78439f676010ef1d2cdb3cf1ca71ff830c))
* multiple assignment fixes for macos scripts and cloudpc provisioning policies ([#612](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/612)) ([769bb85](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/769bb85cc8d573418652b3e21646f1488886445b))

## [0.21.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.21.1-alpha...v0.21.2-alpha) (2025-07-16)


### Bug Fixes

* doc gen for device_management_windows_driver_update_profile and device_management_windows_driver_update_inventory datasources ([#607](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/607)) ([bacccc7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bacccc7cc7d6305753c15085e0724d215f3c889f))

## [0.21.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.21.0-alpha...v0.21.1-alpha) (2025-07-16)


### Bug Fixes

* added page itteration to settings_catalog to handle &gt;25 config items within hcl. ([e9bfa49](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e9bfa4956acc56e650b1d4d845f5b902e69ed2fe))
* added page itteration to settings_catalog to handle &gt;25 config items within hcl. ([#603](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/603)) ([50947df](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/50947df300c70b48cc20ff59424cc382ab98fb06))
* added template_reference block to device_management_settings_catalog_configuration_policy and updated examples ([0460890](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/046089081783545b364328eed01d2d535075f9f8))
* for macos_custom_attribute_script resource updates ([c175a29](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c175a29db1dd7d5461a525d834fed444cfbe096c))
* for macos_custom_attribute_script resource updates and added template_reference block to device_management_settings_catalog_configuration_policy and updated examples ([#601](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/601)) ([38e1a4c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/38e1a4c1e87455be5bd33a1a1d50dafc56c87eba))
* setValueTemplateReference unreachable case matching ([45709ca](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/45709ca36b7a74f4f10f007e6f0612ca0403a731))

## [0.21.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.20.0-alpha...v0.21.0-alpha) (2025-07-14)


### Features

* added 1st itteration of enterprise app catalog ([65bcd68](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/65bcd68cb69394888fe00c71a470e30c388f634b))
* added datasource utility_windows_msi_app_metadata with examples ([d547d5c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d547d5c2121ec4e5cf04d72ea023df1cd9c6f29a))
* added datasource utility_windows_msi_app_metadata with examples ([#593](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/593)) ([370093b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/370093bd1359d99983d5b74d2e64bedd91ee85f9))
* added ios_store_app with examples ([e6aa49f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e6aa49f6297a50afbbb0829f087c5ec4708b4cd6))
* added ios_store_app with examples ([#586](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/586)) ([0f9342d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0f9342dd220c2e09582edaf78d7c0b11434cb49d))
* added mobile_app_relationship datasource and mobile_app_supersedence resource with example ([b5bdc2d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b5bdc2d80edca31c8d55b8a9b9ae9309cf72ac4c))
* added mobile_app_relationship datasource and mobile_app_supersedence resource with examples ([#591](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/591)) ([4531c1b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4531c1b5b992fe74b379a6dffc07a4f592c83662))
* added utility_microsoft_store_package_manifest_metadata with examples ([0fb7804](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0fb78041580686f9a4935d79e0da705e31d920b1))
* added utility_microsoft_store_package_manifest_metadata with examples ([#592](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/592)) ([28187fa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/28187fa4b8c8b8e0d96e1f2775a5907312e57203))
* added win32_lob_app_with_examples ([ba42395](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ba423951160590485e02af4f9bd1582733aa0410))
* added win32_lob_app_with_examples ([#594](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/594)) ([ef9ee54](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ef9ee54e11b586aab828df896668625962258659))
* added windows_autopilot_device_identity with example ([6eb785a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6eb785a71fd73c0dbf87f4e2aed7e0aea7499432))
* added windows_autopilot_device_identity with example ([#596](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/596)) ([c3f2425](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c3f2425566e7300693b47cff3fb418ca6bf83f00))
* added windows_web_app resource and ios_ipados_web_clip with examples ([#590](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/590)) ([2b34fe1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2b34fe1d1ce469226c2e7b55aff4b7b76b2058f7))
* added windows_web_app resource with examples ([d59e19c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d59e19c7e6b685c43b5a83cbd900fbc1da4a662d))
* initial commit of vpp apps ([58b00f3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/58b00f3169719f3d003fe523577a16f404ba0d09))
* initial implementation of microsoft_teams resources ([4ee45cf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4ee45cfc924e1f6de39d426b779e22090e6b049a))
* initial implementation of microsoft_teams resources ([#584](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/584)) ([912c14a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/912c14a61fb0f346d1ac55ad0cbc1a6c1eef140b))


### Bug Fixes

* for macos platform script assignments ([32a67de](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/32a67ded912795e02aebd4c107ba821fa6c44088))
* for windows 32 log app ([69d45c1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/69d45c1c757b7d406933798c91fb87739f99bd24))
* macos script assignments ([#581](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/581)) ([2635444](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2635444fd9ac8beb68aeac0c0f4eacd1152a9528))
* updated all mobile apps to support all image types and added http[s] validators with constants consistently ([89d6277](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/89d6277b081e1973e392a6b4cd08d5fb8aaa7443))

## [0.20.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.19.0-alpha...v0.20.0-alpha) (2025-07-03)


### Features

* added cloudpc alert rule ([3f04efd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3f04efdb66497d71c94f2a60bbe2d83fc35b7d13))


### Bug Fixes

* updated docs ([#576](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/576)) ([4340d62](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4340d62e2a1e26ee7faf478b1abd73773a1df0eb))

## [0.19.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.18.2-alpha...v0.19.0-alpha) (2025-07-02)


### Features

* added graph_beta_windows_365_azure_network_connection, graph_beta_windows_365_cloud_pc_organization_settings, graph_beta_windows_365_cloud_pc_provisioning_policy and graph_beta_windows_365_user_setting with examples ([#567](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/567)) ([005dbf8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/005dbf8585f2bbe06f9b263751373f563693f8d2))
* added windows 365 examples and bugies for move to graph beta ([5da3eaa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5da3eaa34501da056131aa27982b95d5c67e79af))


### Bug Fixes

* added managed devices 1st itteration ([2355602](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/235560283c5b4344b676d23cc35bfaec6a4231c8))
* added numerous cloud pc datasources ([d1eef9d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d1eef9d883ea610b5c213bef5f65d6f118d00407))
* added numerous cloud pc datasources ([#557](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/557)) ([3e576e1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3e576e1e0c39052d70702b31fb42e7b580192a0f))
* added python venv to gitignore ([be5d039](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/be5d039574d85cf84432e8ea7a6e4a4d186040f4))
* added windows_365 datasources and added odata support ([08b3b8f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/08b3b8f76c46f1bb82ac7bca9554d6a0942aaf43))
* added windows_365 datasources and added odata support ([#559](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/559)) ([74c69e0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74c69e0203b0aa2793b50562d112b59775a4c321))
* example ([fd0bae5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fd0bae56cb1c529e158cebe7b981b608fe192aa3))
* for assignments with cloud provisoning polciies and tester ps script ([bde2160](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bde2160635e9e6d3d0adcfa19a9ced1b394bfb2e))
* for branch merge ([07f4a5e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/07f4a5e0b0a5d6aeda02c64b6518914887a75eb7))
* for branch merge ([#560](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/560)) ([2481027](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2481027cda4a473df85e2102ad765fe829c66d37))
* for cloud pc provisioning profiles graph beta ([da6a74a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/da6a74af856acf168f4bca4a583a21ef467ad595))
* for conditional access policies ([8d7f5f0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8d7f5f01304079faf2c6b567378432caf10f3e8c))
* for GroupSettings unit tests ([ffb8f71](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ffb8f719dedf3155f2d13242dda63dcb1dfaa580))
* for plan and state definitions within update funcs ([1cfce7a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1cfce7a1a2d04f5eee804b14cbae506d3d8acb67))
* for pr merge ([bb7b5a2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bb7b5a2e9de1026da924fc22c2d31125dca764e1))
* for pr merge ([#568](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/568)) ([a4cf05a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a4cf05a3d8aaf1a1098e5c68895b6a7d0a2c81da))
* for read with retry loop. correct ctx handling ([d2bf94b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d2bf94bc304006bc7b01333f1c59ca4c7f58e089))
* graph beta additions for windows 365 ([de472bb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/de472bbaeb29dbbaf63fd18732ba99076540d152))
* implemented win365 provisoning policy fixes with frontline support and assignments. ([#561](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/561)) ([a0d66fe](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a0d66fe8c314d9b27d13598ba6c2bed493613059))
* implemented win365 provisoning policy fixes with frontline support. ([fe1ff4f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fe1ff4f7b2d8b5c5041b1aabd8c205d7b08661c9))
* more fixes for conditional access unit tests ([a916947](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a9169476781401a45e6a540e7348b0ec21e671d5))
* numerous fixes for conditional access policies. uses a new non kiota based sdk http client ([f62c5cc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f62c5cc95be1e96192929098e70d910ad55e96f8))
* numerous fixes for conditional access policies. uses a new non kiota based sdk http client ([#553](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/553)) ([ab62461](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ab624617b6f329c42cb17df17866970b9412f372))
* numerous fixes for read with retry logic to now correctly handle operation type ([#562](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/562)) ([16275b8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/16275b84a72254328a0813e19630fd03c87baf87))
* partial fix for conditional access introduced with sdk v0.137.0 ([#550](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/550)) ([82452cc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/82452cc393627f84402e5e8181bc91584d1eb421))
* re-up ([9fa33bf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9fa33bf3a68a935a28208942de6b319742a24430))
* re-up ([#569](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/569)) ([e30b87c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e30b87cb931cddc4d4848515a49bd706452c913e))
* set unit test timeout to 120 minutes ([e39a0af](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e39a0afa4d436c9b528a8fe31bd0039ea2e7f6d3))
* state removal logic and logging consistency across all resources ([b77d94a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b77d94ac10c980e5daf860c3c40bdd14fdc4e9f9))
* tflog logs for update func ([109e2cc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/109e2ccadaaf6617ecdf0c5d93f82ccb24f3c446))
* tidy up ([a378df2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a378df2d4f1e4579ebe3ef2ae3cb47a42d5cea73))
* unit tests for conditional access policies ([2fcabab](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2fcababe1f4cfa2924b0530f5ade95df5a2222ac))
* unit tests only run now on changes to go files ([47ff1f5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/47ff1f54d7ddd1b799da8ae44679cb8e2b52896a))
* updated docs for groups resources ([f733e4c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f733e4c3fe4ae71a0e99237d4a6e983e689a2c6d))

## [0.18.2-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.18.1-alpha...v0.18.2-alpha) (2025-06-25)


### Bug Fixes

* tf-registry-goreleaser.yml permission fix ([#548](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/548)) ([ef2e54e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ef2e54e157242ef55afe5e9aac1e141932bead17))

## [0.18.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.18.0-alpha...v0.18.1-alpha) (2025-06-25)


### Bug Fixes

* added group acceptance and unit tests ([6b0bc50](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6b0bc5029011509f60c167911672f34712ff28d9))
* added unit and acc tests for groups and groups member assignment ([aef90f6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aef90f6cc2397a7f0ee066083ba7fbee638e14ab))
* added unit and acc tests for groups and groups member assignment ([#545](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/545)) ([99759a4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/99759a48b342313f9d6e26a7d2eda40f769282e5))

## [0.18.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.17.1-alpha...v0.18.0-alpha) (2025-06-24)


### Features

* added numerous acc and unit tests ([#541](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/541)) ([5be07af](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5be07af61c225192c817139f90c36a0bfb71d986))
* added unit and acc testing for users, m365_admin_install_options , license assignment ([a558b64](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a558b6458aa0c1f2a2192ddc2d60f064c777dbcc))
* added unit and acc tests for resource user and resource license assignment ([76058c0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/76058c034b568a2cdc081aafc06e432c3cc9a647))


### Bug Fixes

* for true to false ([369f494](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/369f49448fb9ed595c7c4736c8728ad19c93174f))
* for unit tests for microsoft365_graph_beta_device_management_macos_software_update_configuration ([ced892f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ced892f18331710ee05be492c1788c9d47788675))

## [0.17.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.17.0-alpha...v0.17.1-alpha) (2025-06-23)


### Bug Fixes

* added docs for device config assignment graph v1.0 ([ef66aab](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ef66aab7ae4c7720e18016d2380cf87bd6bad5d0))
* added docs for device config assignment graph v1.0 ([#536](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/536)) ([788f7f1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/788f7f1d38fa728ab864987d2bb6c06460decb54))
* increased timeout for goreleaser ([709cbb2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/709cbb29a33a7062d699c1e82427a3c5f9ea269a))
* increased timeout for goreleaser to 90 minutes ([#538](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/538)) ([a528317](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a5283177fe997a903a4f072d159d118d32f8044a))

## [0.17.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.16.1-alpha...v0.17.0-alpha) (2025-06-23)


### Features

* Add macOS LOB app publisher validation and update gitignore ([c9cd922](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c9cd922080ec768367bc79f1fc73632c9248d2f4))
* added 1st itteration of oidc support for github actions and azdo ([3a2c14b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3a2c14bdcc51b039fc9269e9febf8edae89fa866))
* added additional extraction points for graph metadata ([3893926](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3893926ae66dca2ed555653745049b530d0fbc32))
* added ai instructions set ([af64371](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/af64371837294cfc9e11d78ed2bf2395c661fedd))
* added application category resource and datasource with examples ([55a099a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/55a099a087f713f4e4c01558e4aec5fdfaf303f8))
* added assignments for windows updates as a seperate set of resource types ([837ee80](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/837ee809bc587496f2e8be44996db786cbd0d10f))
* added configuration policy unit tests ([cc5a938](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cc5a93863fa5112832619bdf1f897e0ba6966b66))
* added datasource graphBetaDeviceAndAppManagementWindowsQualityUpdateProfile ([9f6f79d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9f6f79d7472eef5d7f76973c8195046b65c4462b))
* added datasource microsoft365_utility_macos_pkg_app_metadata ([2c2cbca](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2c2cbca7feb8d1336c9327e487767ecec1e4bc35))
* added dev docs ([d31361f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d31361fa979172bed9709402f76d8fbc3616b32d))
* added developer docs ([6b1717a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6b1717a81351b40460835e8f9da819d32a7cf7c2))
* added developer docs for resource development ([#521](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/521)) ([2b5178c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2b5178c7cfd28ef3df52ae9fe9672a4d0623d999))
* added device_and_app_management_windows_update_catalog_item datasource with examples ([655e804](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/655e804d7ea7989c549fa86a8925ea82b0b33d46))
* added graph metadata extraction script ([cef8591](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cef8591dc8904959fb4844c06aa72d1770f1fe6b))
* added graph_beta_device_and_app_management_windows_driver_update_profile_assignment ([73dc481](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/73dc481446cb3c88284ef6b5fcac719da08bfb81))
* added graph_beta_device_and_app_management_windows_feature_update_profile with example ([050cec1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/050cec1724781fa5bdf81b14c9c5890f5a0386e4))
* added graph_beta_device_management_apple_user_initiated_enrollment_profile_assignment with examples ([5ae2eca](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ae2eca5809e0083d525cd9695002aab18ed886e))
* added graph_beta_device_management_apple_user_initiated_enrollment_profile_assignment with examples ([#503](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/503)) ([e4cade6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e4cade6e0009400e6ec09ffe9b8935fe02b5cadf))
* added graph_beta_device_management_macos_software_update_configuration with examples ([33247a9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/33247a91e23b12381eaa41c2364448c26d5e502a))
* added graph_beta_device_management_macos_software_update_configuration with examples ([#506](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/506)) ([d82a919](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d82a9197b67a1b1625a32f800d153dd738daf860))
* added group license assignment, user license assignment resources and datasource license subscribed SKUs ([bede181](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bede1812ddd65df921206baf67602e02693e1f2a))
* added group ownership assignment with examples ([18810ea](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/18810ea49955060d9bfb87947a2b25a4e560c3ba))
* added group ownership assignment with examples / renamed user license assignment to follow resource naming pattern ([#497](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/497)) ([6e2fdf3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e2fdf3e6411537f679f5201abf9db426f82fd5b))
* added GroupMemberAssignment resource with examples ([21dde58](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/21dde582a0048b75e462dd881831d08706ef8069))
* added GroupOwnerAssignmentResource ([4f9bfd6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f9bfd60e1013b7030ac31187b42fdd5224d78a7))
* added GroupOwnerAssignmentResource ([#499](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/499)) ([9bc5dee](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9bc5dee1a016aa39f78acbee5b91236d4c95d267))
* added groups with examples ([ffc38e3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ffc38e3db9b488cbfb86b062d40bc946c9ebe22a))
* added groups, and group membership assignment, updates to doc descriptions ([#495](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/495)) ([0601507](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0601507bf998791c5f74ae433005a492d0ecf1b0))
* added helper SetObjectsFromStringSet for object construction ([e2b110e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e2b110e4f7b4aa73aa2e823a5ae9b58f5b597e8c))
* added intune device categories resource and data source with examples ([#389](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/389)) ([b8d33fb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b8d33fba7b11121bf61514164b77d01858a9a74c))
* added macos pkg app datasource ([a65ca03](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a65ca038103e064225b7d136205179b7fbb04fad))
* added macos pkg datasource ([f0f6144](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f0f61447b307f221cdf6f0b0d65f74d8e4f6beb6))
* added macos_dmg_app fixes and updated resource examples ([1da2c8d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1da2c8d26bd0963fdd48d72ff384cb0245a62be2))
* added makefile for tf development ([36b7a0a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/36b7a0a57f687dbd96bd1c521c745078cbed2977))
* added makefile, ai instruction set, refactored data type conversion helpers, scaffolding for unit and acc tests with mocking ([#528](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/528)) ([6788408](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6788408661ed58258ee10589b6415f7ddab18800))
* added managed device clean up rule with examples ([fc33266](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fc332665c817df572c0d349ec4d68ccf1b037570))
* added managed device clean up rule with examples ([#463](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/463)) ([f1a8777](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1a87770fa8e61a0edbc8370876cf86b4cb07850))
* added microsoft365_graph_beta_device_and_app_management_windows_update_profile_assignment with example ([44e3539](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/44e3539f943013c5a0414e2053f56a1c02757f04))
* added microsoft365_graph_beta_device_management_terms_and_conditions ([335e4a1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/335e4a1ec8441e7f49b980af60de4d07e2afea8f))
* added microsoft365_graph_beta_device_management_terms_and_conditions ([#466](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/466)) ([16849a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/16849a77e9262aeddbae84bc43b55552da62c146))
* added microsoft365_graph_beta_device_management_terms_and_conditions_assignment with examples ([c3db9ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c3db9acb7c2fb11122658bdf817c2ff57af01fec))
* added microsoft365_graph_beta_device_management_terms_and_conditions_assignment with examples ([#468](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/468)) ([6e407c0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e407c067f71cbc2f01cef185adaff9f214774db))
* added microsoft365_graph_beta_device_management_windows_feature_update_policy_assignment with examples ([74f286e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74f286eb668a342f700b1ce3c5da9671695ffdcb))
* added microsoft365_graph_beta/microsoft365_graph_beta_device_management_windows_driver_update_profile_assignment with examples ([863774c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/863774cbb6a0b9944c1d79b2e87ac2f4ad66a6b5))
* added mobile app data source with odata queries ([3a1c0f4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3a1c0f414e70ea9c55f36b31c1c48e85feb2aa02))
* added multiple new resources for license assignment and fixes for macos dmg app ([#491](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/491)) ([8d09b67](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8d09b67ec036a30817aecb7745d050adc96c880c))
* added multiple onboarding resources ([602135e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/602135e85aeab9ad2f8923c437004be647357d58))
* added operation approval policy ([40b307c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40b307cfd8c1ec89fd342c7a29f20254d2cacaef))
* added ps graph metadata extraction script ([#522](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/522)) ([a3ca194](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a3ca1944e6bbd89c00dca90a327a1188e4074aa1))
* added refactored unit tests for provider build and credential factory ([bd7691c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bd7691c76d6403dadb4eb68d76ba37093f506b9b))
* added refactored unit tests for provider build and credential factory ([#439](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/439)) ([c191a93](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c191a9345c962f5f19a1bfd76dcafb3fb1deec01))
* added resource operation for intune rbac with examples ([4f36c5f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f36c5f5b50c9d19f3cdcdcc2507d8d8929a5931))
* added resource operations with examples for intune rbac ([#465](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/465)) ([1c31571](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c315714f220e136300bf4da2f0c029dd39e2c11))
* added setting catalog in native hcl ([596d6c9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/596d6c9f0454810004ff30a3177a5c4ec2bc4ebc))
* added settings catalog in native hcl ([#461](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/461)) ([6e5e693](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e5e6937324d83df7fe5f96445dc45f851bf771f))
* added timeouts to datasources ([#397](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/397)) ([100840f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/100840f1bff363a598fd9a76e26183c16c086150))
* added unit tests for configuration policy settings go and json builders ([#511](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/511)) ([3293255](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/32932553d3ce83264e3c0d618f5ca0db5681bc09))
* added validation for mobile app assignments for 8 invalid scenarios ([cc88fd1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cc88fd104a35cf071e4a5693a84420f57d1151d8))
* added windows autopilot and assignments thereof ([#488](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/488)) ([5ffe081](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ffe08175062fa490c4b7babc5e8ecf658cc3ba3))
* added windows driver update profile and inventory resources and datasource ([#416](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/416)) ([8e3272d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8e3272dbc77296cdd874a11e7c7693103173419b))
* added windows driver update profile and inventory resources and datasources ([85ff58f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/85ff58f605f939be963c3aba58ebd44c0ac41f0f))
* added windows quality update profile assignment resource with examples ([a49a109](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a49a109323d23f0c8e0ed2be4afea53a77b1b05d))
* device_enrollment_configuration with examples ([#420](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/420)) ([88516db](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/88516db1ff74964069d9965f29c21972bb147144))
* finalised device preparation policies until msft fix their api ([b9949b1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b9949b141afcd3822012fbaf850c417ed4b51b93))
* graph_beta_device_management_device_enrollment_notification_configuration with examples ([ff7fecb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ff7fecbdb1f77970dabf1994fc47269cd2cc58c1))
* graph_beta_device_management_device_enrollment_notification_configuration with examples ([#501](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/501)) ([145625d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/145625d2510a28544419c69beab4eb777da6e3cc))
* half implementation of cloudPC as a datasource ([68e7a43](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/68e7a43dc062a7808d2265f7b76c7a354fff5d85))
* implemented device categories resource and datasource with examples ([acf2626](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/acf262635ca9aeb5175f0744fc108dfd0694b027))
* implemented refactor state and constructor funcs ([307fb91](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/307fb91f0966100ac6517915ac263a6db1dc4b96))
* initial implementation of mac_dmg_app ([35b6fe6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/35b6fe60dd97db4c75b442387d26c26680f61afd))
* mobile app assignments ([80c892e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80c892ec95186025afc2c6f7edbd61f7ed1863e4))
* mobile app assignments can now be set in any order with modify plan (diff supression) ([c8f97ef](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c8f97ef7201c1a79efc3850add74eac0e78468be))
* mobile app assignments resource with examples ([#453](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/453)) ([498ff50](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/498ff5087f942dffa701be9da736c8d3e1a5dde2))
* mocking scaffolding ([f1be6d1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1be6d1f6fddfca091d34df19750da364545c1b1))
* mocks scaffolding ([#440](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/440)) ([9ee57df](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ee57dfc4eecebb95c35c0e60686925d3228c2c6))
* moved all client initization to package client and refined client init logic with provider configure func ([45de96e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/45de96e0ef0052544f38055b483c138db8698d84))
* moved all client related logic for building to the client package from the provider package ([9e2561c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9e2561cbcaa78b5458eae138ca025609a3416f87))
* moved custom requests to it's own package within common out of the client package ([356f090](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/356f09017278074550b35a4b4f5e4a3aac175356))
* refactored data source for maospkgapp ([#442](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/442)) ([fec56b7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fec56b72934dc5c9fbd1dad05e8d690fab0793ab))
* refactored directory hirachy so that all client related logic is within the client package. moved all custom request logic to it's own package within custom from client package ([#529](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/529)) ([8754a07](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8754a071caed0fd816f7080b840e4f5003c07981))
* refactored macospkg app datasource with example ([cbc1681](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cbc1681feca36a096ea63b69dc61540178b21447))
* refactored retry mechanism for create update func calls to read and fixed resource name issue in tflog ([6f3842b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6f3842b684eab3bcde146a90a9afd4b7b1c1cb34))
* refactored retry mechanism for create update func calls to read and fixed resource name issue in tflog ([#498](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/498)) ([5257e5e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5257e5ef5653b0ed0ca2ba5fe4ff1c70c3a5cf29))
* refactored unit and acc test strategy to use local responders and centralised factories for common opperations ([c878bda](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c878bdaaf8fa7c65ce69323323a27dec6ff340bd))
* refactored unit and acc test strategy to use local responders and centralised factories for common opperations ([#534](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/534)) ([8179a7d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8179a7d00b43e34f21e7b941fea16b90ca0d4d76))
* repo structure to move datasoources and resources within a common services folder. moved common to services folder in folder hierarchy ([0799a5b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0799a5b52ca34d78f186c7c6068b7d0abf47313c))
* repo structure to move datasources and resources within a common ./internal/services folder. Moved the common folder with shared assets to ./internal/services from /resources ([#524](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/524)) ([22520ec](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/22520ec6ed4492addcc0a0ac1d205a2fc1720da1))


### Bug Fixes

* added close stale issues maintainence pipeline ([1d90efd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1d90efd55aab386abff032eda94948987f69d50f))
* added close stale issues maintainence pipeline ([#475](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/475)) ([40affe2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40affe22d27c45128e8b4f7887419f10d41efc58))
* added discord info to readme ([88c0a53](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/88c0a53746578ceb5479c71a5c9e61bb83485f7a))
* added discord info to readme ([#519](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/519)) ([aded08c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aded08ca82a4a35e3ba2e111b96f9ed5e140571b))
* added examples and more fixes for windows_remediation_script_assignments ([9827467](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9827467ac9972380540ccb29490670d5db960c87))
* added mutex to all api calls to defender against Kiota's middleware (e.g., HeadersInspectionHandler) when modifing shared header maps during HTTP request processing. appears to be a bug in the latest version ([976ad22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/976ad22dea6b1ddf6e3027714a04c449b73ac11c))
* added mutex to all api calls to defender against Kiota's middleware (e.g., HeadersInspectionHandler) when modifing shared header maps during HTTP request processing. appears to be a bug in the latest version ([#450](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/450)) ([10327c8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/10327c856db27c79ae6883677a48049c6caf2a1e))
* added new regex for ISO8601Duration and comment fix ups ([1e11d71](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1e11d71d4890d1adaf6f3901c6ed2268264b597a))
* added planmodifiers.UseStateForUnknownString(), for id ([dccde7e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dccde7e78d92bf90a23bbdf7ad5673391baafc65))
* added read timeouts for datasources ([243e6b2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/243e6b2fdc92441d54949d0d1bd9544b4b479c1e))
* base resource updates for macos_pkg_app doesnt use redundant file size checks ([6162a1d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6162a1d4dc9860de2aa812ff386289bd6fc387f7))
* categories fix for macos app categories and stating issues for infured included app values ([defa936](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/defa93679fa6bed0bc6065678987c76429383133))
* changed all actions to use commit sha's ([cd48ede](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cd48ede424a0363694107d44fb6fe64823bce198))
* content version ([74c417b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74c417b8312e6353bbbdbf78fd012d9fa96b39e1))
* defined default values and plan modifier for role_scope_tag_ids ([e8909e0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e8909e0fc876dbe03422bfb26943ac872cfab432))
* determined that the api for AutopilotDevicePreparationPolicy is broken.sigh ([481dcdd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/481dcddaab9b8510b0740960919ff1c7725d98e4))
* docs ([0a676f8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0a676f87825c2d75b2c3ed5175d48474e23ec28d))
* examples so that device_management examples are reflective ([d6856b4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d6856b480ef5a0c9f7277592bc391eb47487b443))
* examples so that device_management examples are reflective ([#441](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/441)) ([6062154](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6062154d0eedb88177fec252a4c5fb5f925ca24a))
* extended go lint timeout to 20minutes ([e5f9b3b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e5f9b3bbf42b420ba1568a9042930c281d82523b))
* final fixes for macos assignments ([8d998cf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8d998cfc27ccb4aa1b4bab3be60823f479b7cf85))
* final fixes for macos assignments ([#451](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/451)) ([7e3b853](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7e3b853731d22cc832840ceab3ca9c7a256418c4))
* finalised device preparation policies until msft fix their api ([#505](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/505)) ([089d398](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/089d398d18efd2a065d5dcdd4ed191bda458fab5))
* for deps bump of go sdk version ([54a119e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/54a119edcca117fa697b21c59f2ba35cdda12ae2))
* for doc gen pipeline ([f959ac8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f959ac87f1d6ea3e0d66b9979502a5c1f914466f))
* for doc gen pipeline ([d3fecdf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d3fecdf388ed6ad8e82fffffce38cbdd0126553a))
* for doc gen pipeline ([#471](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/471)) ([c93301e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c93301e7f95b32b56ed812ba5f373d2c2d3b0d80))
* for doc gen pipeline ([#473](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/473)) ([3c58522](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3c58522f6f1522eca9da2b4ee613b32b96319368))
* for docs path ([2fce669](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2fce66923162574af62abe0c7236cf9da8bd76ab))
* for enum handling within GetDeviceEnrollmentConfigurationType ([#437](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/437)) ([896c291](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/896c2918aa799ec20127a79538a06cfe1cff892e))
* for example path ([771bb1c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/771bb1c2792a16955e05f124cbe7a4e6d88cd3bb))
* for examples ([244a04b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/244a04b7ef7521d33ffcf42e0af0015d4d2791c8))
* for go lint failure handling ([e3e52be](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e3e52be832b1ca02cb214c3c0bbb7d6d5b8c57e9))
* for GraphToFrameworkISODuration helpers to resolve unit tests ([e8338a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e8338a764a864067ad78aef8ed9e14be86f4c592))
* for handling enum ([ce13f69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ce13f693c89723f9766c02519f17596ea328d7e2))
* for handling enum ([6e2f1ff](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e2f1ff38ef16ddbcd405eff88c380371e315856))
* for macos pkg apps ([#395](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/395)) ([deb9309](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/deb9309edc8a8aff4d33a022102f30f9c90eb679))
* for mobile app assignment id's and content version stating ([a57a269](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a57a26945a184c294e4c37c41b6b3457dbdad180))
* for mutex handling for SetCommittedContentVersion ([058bf32](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/058bf32f301bc9b5082f668427aaf4ca3cff3878))
* for parrallism issues with concurrent header map settings in client mobile apps ([49561f6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/49561f64cea87bd23eab9dda180824ddfd8934de))
* for pipeline permissions ([a8b6990](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a8b69906db84a5a90db4b9f9373a64e4dabf9a08))
* for role_scope_tag_ids ([#444](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/444)) ([b1e0fd8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b1e0fd819d4956ee065e96d4300fb24785719dc2))
* for successful stating logging with resource name and resource id ([8c81932](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8c819320ff467a553640c8ff5ce9c78c5f7cf76a))
* for winget assignments ([8a2bdd0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8a2bdd03111f263f47db1614ba1e0fb8deb5a650))
* Go Lint configuration yml for provider project types ([#424](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/424)) ([b92b8ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b92b8ac28100a40611b0cca6e191750db9b8d277))
* go unit test pipeline trigger ([e972c36](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e972c3658f3f1a4c74f29fa55657ccef63117d5b))
* go-lint ([#423](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/423)) ([57ba99d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/57ba99dd57bf8b26e9785767cf05449df1b2fca2))
* improved markdown descriptions for documentation generation ([16bc1d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/16bc1d2d7324c2a5c0a224a274e73100648b4caa))
* m365 app installl options ([e2db1a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e2db1a7982f5cced9b3f346b7f2df8a30c92695a))
* migrated microsoft365_graph_device_and_app_management_m365_apps_installation_options to graph v1.0 ([a3c9682](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a3c968250d67787711d46b2589facd62cf3ab8b3))
* migrating winget app assignments to sets ([2272cc6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2272cc64166db8c02accd2a243d79a646fc3dfe4))
* minor fix for sorting order of assignments ([fb50e0f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fb50e0fbac88e72d4d784f400131449be1c4ff07))
* misc naming fix ups ([502ba24](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/502ba24503a3473bd3351a9ae31cb667ee5b409c))
* moved mobile app assignment validation to create and update funcs to fail early if any hcl mistakes before any rresources get deployed ([8c79958](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8c79958155e1445d959487a27cd66ddca38947b3))
* numerous bug fixes and unit test additions ([44bb11a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/44bb11a80d63c5d3f2b1a7a726d32f39387591a7))
* numerous edits and tweaks ([dccd487](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dccd4873fc4dec676b2a0e3859c7a2a42d595b4f))
* numerous fixes and scaffolding for mock tests ([f951cda](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f951cdaeec2b4e5401f447731cd383011963f9b7))
* numerous schema fixes ([ddb6e87](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ddb6e8780360a63c31354f0899314fd4ed51c111))
* pipeline permissions and tidy up ([#532](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/532)) ([475c686](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/475c686ffce21927d8b72d3419d5fada6e79766c))
* pipeline step naming ([3fade5e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3fade5e96ab87097c3b8b18dc3e2f0ea5850a7a3))
* pipeline step naming ([#425](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/425)) ([8067017](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80670170abf8aa817432805c0d4a786859655951))
* python test scripts ([591c402](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/591c40260095757c86431a7888715f3b3f4b50c3))
* read with retry test ([3053bd5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3053bd530f85440aa1ac41911e6b94d4a5347ed7))
* readme's and schema docs for tf registry ([afbab65](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/afbab654a7ad3e83868e0b7f52fb7d0849444a28))
* readme's and schema docs for tf registry ([#520](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/520)) ([bc16e69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bc16e692736e681bf81d45f348156e3e18763c1f))
* rearraging provider docs ([a0aab65](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a0aab656a60e99ef2554e053d01fdad5c4539016))
* refactored docs for repo restructure ([546ebb4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/546ebb4ff5dfe06b9b6c9febad4223a4acf71abf))
* refactored docs for repo restructure ([#435](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/435)) ([3b14e78](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3b14e78f1db1659b176983abced03e04b32195d6))
* refining logic for handling app metadata ([9ecf223](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ecf2236482cffbe5ef3cc67cf9e10e94068f345))
* regex const naming patterns ([54b5af1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/54b5af124f7a2b97901be6c7b3144707f13de508))
* removed conditional access. complete refactor required with latest sdk version ([4cb550d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4cb550d56f376cb30bde71c71b4699e63653e07c))
* removed mutex, going with parallism of 1 instead ([dd0e6d8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dd0e6d822dcf3eb6fd43bbdea105d62020dc488f))
* renamed user license assignment to follow resource naming pattern ([29f4c8b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/29f4c8b163dd7e1d95dfe045831fc64e125db76f))
* restructured repo to follow api endpoint paths more precisely ([1b2c110](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1b2c1108add87c7395eaed7cfce79887c7899f12))
* restructured repo to follow api endpoint paths more precisely ([#433](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/433)) ([6aed990](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6aed9906d8f76430c646909e6111c6d91b29f1b3))
* tidy up pipeline step naming ([102e356](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/102e35695ac0629381b740e08aec981cfdaf84a1))
* tidy up repo restructure ([b5b35cf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b5b35cf4b001d4ba9ebc0cfed8913bc8a7054e12))
* typo fix for datasource ([5c60f13](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5c60f13a9ce6d8701d9310ea41f07e1104524967))
* unit test trigger ([2bf41c4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2bf41c4a70e4a47e268bb9d68950dda48d24bfe0))
* unit test trigger ([#474](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/474)) ([e426630](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e42663025764a290342fde1a4f489a19931e94a1))
* unit tests ([e0dd719](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e0dd7194b4dd58cf18a241f927119e8dd185c630))
* unit tests ([5ffcd30](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ffcd30efd1beb34be608df2095a0eda5e632c9f))
* updated create and delete to use the crud.ReadWithRetry pattern ([83003fd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/83003fd8275e537495cf3a94865b425e9712f7d8))
* win_get_app now skips property 'InstallExperience' and 'PackageIdentifier' as it cannot be patched during update operations. and added support for app categories ([8b5a4a5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8b5a4a59a43e6bddb6e82a3763713a85bcfe4047))
* windows remediation script assignments ([04243b6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/04243b62ecb12ed869edec6caf334e6671c8fd90))

## [0.16.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.16.0-alpha...v0.16.1-alpha) (2025-06-21)


### Bug Fixes

* for pipeline permissions ([a8b6990](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a8b69906db84a5a90db4b9f9373a64e4dabf9a08))
* pipeline permissions and tidy up ([#532](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/532)) ([475c686](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/475c686ffce21927d8b72d3419d5fada6e79766c))

## [0.16.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.15.0-alpha...v0.16.0-alpha) (2025-06-21)


### Features

* added additional extraction points for graph metadata ([3893926](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3893926ae66dca2ed555653745049b530d0fbc32))
* added ai instructions set ([af64371](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/af64371837294cfc9e11d78ed2bf2395c661fedd))
* added configuration policy unit tests ([cc5a938](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cc5a93863fa5112832619bdf1f897e0ba6966b66))
* added dev docs ([d31361f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d31361fa979172bed9709402f76d8fbc3616b32d))
* added developer docs ([6b1717a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6b1717a81351b40460835e8f9da819d32a7cf7c2))
* added developer docs for resource development ([#521](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/521)) ([2b5178c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2b5178c7cfd28ef3df52ae9fe9672a4d0623d999))
* added graph metadata extraction script ([cef8591](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cef8591dc8904959fb4844c06aa72d1770f1fe6b))
* added graph_beta_device_management_apple_user_initiated_enrollment_profile_assignment with examples ([5ae2eca](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ae2eca5809e0083d525cd9695002aab18ed886e))
* added graph_beta_device_management_apple_user_initiated_enrollment_profile_assignment with examples ([#503](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/503)) ([e4cade6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e4cade6e0009400e6ec09ffe9b8935fe02b5cadf))
* added graph_beta_device_management_macos_software_update_configuration with examples ([33247a9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/33247a91e23b12381eaa41c2364448c26d5e502a))
* added graph_beta_device_management_macos_software_update_configuration with examples ([#506](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/506)) ([d82a919](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d82a9197b67a1b1625a32f800d153dd738daf860))
* added makefile for tf development ([36b7a0a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/36b7a0a57f687dbd96bd1c521c745078cbed2977))
* added makefile, ai instruction set, refactored data type conversion helpers, scaffolding for unit and acc tests with mocking ([#528](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/528)) ([6788408](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6788408661ed58258ee10589b6415f7ddab18800))
* added ps graph metadata extraction script ([#522](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/522)) ([a3ca194](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a3ca1944e6bbd89c00dca90a327a1188e4074aa1))
* added unit tests for configuration policy settings go and json builders ([#511](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/511)) ([3293255](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/32932553d3ce83264e3c0d618f5ca0db5681bc09))
* finalised device preparation policies until msft fix their api ([b9949b1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b9949b141afcd3822012fbaf850c417ed4b51b93))
* graph_beta_device_management_device_enrollment_notification_configuration with examples ([ff7fecb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ff7fecbdb1f77970dabf1994fc47269cd2cc58c1))
* graph_beta_device_management_device_enrollment_notification_configuration with examples ([#501](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/501)) ([145625d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/145625d2510a28544419c69beab4eb777da6e3cc))
* implemented refactor state and constructor funcs ([307fb91](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/307fb91f0966100ac6517915ac263a6db1dc4b96))
* moved all client initization to package client and refined client init logic with provider configure func ([45de96e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/45de96e0ef0052544f38055b483c138db8698d84))
* moved all client related logic for building to the client package from the provider package ([9e2561c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9e2561cbcaa78b5458eae138ca025609a3416f87))
* moved custom requests to it's own package within common out of the client package ([356f090](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/356f09017278074550b35a4b4f5e4a3aac175356))
* refactored directory hirachy so that all client related logic is within the client package. moved all custom request logic to it's own package within custom from client package ([#529](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/529)) ([8754a07](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8754a071caed0fd816f7080b840e4f5003c07981))
* repo structure to move datasoources and resources within a common services folder. moved common to services folder in folder hierarchy ([0799a5b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0799a5b52ca34d78f186c7c6068b7d0abf47313c))
* repo structure to move datasources and resources within a common ./internal/services folder. Moved the common folder with shared assets to ./internal/services from /resources ([#524](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/524)) ([22520ec](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/22520ec6ed4492addcc0a0ac1d205a2fc1720da1))


### Bug Fixes

* added discord info to readme ([88c0a53](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/88c0a53746578ceb5479c71a5c9e61bb83485f7a))
* added discord info to readme ([#519](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/519)) ([aded08c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/aded08ca82a4a35e3ba2e111b96f9ed5e140571b))
* added new regex for ISO8601Duration and comment fix ups ([1e11d71](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1e11d71d4890d1adaf6f3901c6ed2268264b597a))
* changed all actions to use commit sha's ([cd48ede](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cd48ede424a0363694107d44fb6fe64823bce198))
* determined that the api for AutopilotDevicePreparationPolicy is broken.sigh ([481dcdd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/481dcddaab9b8510b0740960919ff1c7725d98e4))
* extended go lint timeout to 20minutes ([e5f9b3b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e5f9b3bbf42b420ba1568a9042930c281d82523b))
* finalised device preparation policies until msft fix their api ([#505](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/505)) ([089d398](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/089d398d18efd2a065d5dcdd4ed191bda458fab5))
* for go lint failure handling ([e3e52be](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e3e52be832b1ca02cb214c3c0bbb7d6d5b8c57e9))
* for GraphToFrameworkISODuration helpers to resolve unit tests ([e8338a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e8338a764a864067ad78aef8ed9e14be86f4c592))
* go unit test pipeline trigger ([e972c36](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e972c3658f3f1a4c74f29fa55657ccef63117d5b))
* misc naming fix ups ([502ba24](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/502ba24503a3473bd3351a9ae31cb667ee5b409c))
* numerous bug fixes and unit test additions ([44bb11a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/44bb11a80d63c5d3f2b1a7a726d32f39387591a7))
* numerous fixes and scaffolding for mock tests ([f951cda](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f951cdaeec2b4e5401f447731cd383011963f9b7))
* python test scripts ([591c402](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/591c40260095757c86431a7888715f3b3f4b50c3))
* read with retry test ([3053bd5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3053bd530f85440aa1ac41911e6b94d4a5347ed7))
* readme's and schema docs for tf registry ([afbab65](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/afbab654a7ad3e83868e0b7f52fb7d0849444a28))
* readme's and schema docs for tf registry ([#520](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/520)) ([bc16e69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bc16e692736e681bf81d45f348156e3e18763c1f))
* regex const naming patterns ([54b5af1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/54b5af124f7a2b97901be6c7b3144707f13de508))
* unit tests ([5ffcd30](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ffcd30efd1beb34be608df2095a0eda5e632c9f))

## [0.15.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.14.1-alpha...v0.15.0-alpha) (2025-06-04)


### Features

* Add macOS LOB app publisher validation and update gitignore ([c9cd922](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c9cd922080ec768367bc79f1fc73632c9248d2f4))
* added group license assignment, user license assignment resources and datasource license subscribed SKUs ([bede181](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bede1812ddd65df921206baf67602e02693e1f2a))
* added group ownership assignment with examples ([18810ea](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/18810ea49955060d9bfb87947a2b25a4e560c3ba))
* added group ownership assignment with examples / renamed user license assignment to follow resource naming pattern ([#497](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/497)) ([6e2fdf3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e2fdf3e6411537f679f5201abf9db426f82fd5b))
* added GroupMemberAssignment resource with examples ([21dde58](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/21dde582a0048b75e462dd881831d08706ef8069))
* added GroupOwnerAssignmentResource ([4f9bfd6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f9bfd60e1013b7030ac31187b42fdd5224d78a7))
* added GroupOwnerAssignmentResource ([#499](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/499)) ([9bc5dee](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9bc5dee1a016aa39f78acbee5b91236d4c95d267))
* added groups with examples ([ffc38e3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ffc38e3db9b488cbfb86b062d40bc946c9ebe22a))
* added groups, and group membership assignment, updates to doc descriptions ([#495](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/495)) ([0601507](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0601507bf998791c5f74ae433005a492d0ecf1b0))
* added macos_dmg_app fixes and updated resource examples ([1da2c8d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1da2c8d26bd0963fdd48d72ff384cb0245a62be2))
* added multiple new resources for license assignment and fixes for macos dmg app ([#491](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/491)) ([8d09b67](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8d09b67ec036a30817aecb7745d050adc96c880c))
* added multiple onboarding resources ([602135e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/602135e85aeab9ad2f8923c437004be647357d58))
* added windows autopilot and assignments thereof ([#488](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/488)) ([5ffe081](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ffe08175062fa490c4b7babc5e8ecf658cc3ba3))
* half implementation of cloudPC as a datasource ([68e7a43](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/68e7a43dc062a7808d2265f7b76c7a354fff5d85))
* initial implementation of mac_dmg_app ([35b6fe6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/35b6fe60dd97db4c75b442387d26c26680f61afd))
* refactored retry mechanism for create update func calls to read and fixed resource name issue in tflog ([6f3842b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6f3842b684eab3bcde146a90a9afd4b7b1c1cb34))
* refactored retry mechanism for create update func calls to read and fixed resource name issue in tflog ([#498](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/498)) ([5257e5e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5257e5ef5653b0ed0ca2ba5fe4ff1c70c3a5cf29))


### Bug Fixes

* for examples ([244a04b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/244a04b7ef7521d33ffcf42e0af0015d4d2791c8))
* improved markdown descriptions for documentation generation ([16bc1d2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/16bc1d2d7324c2a5c0a224a274e73100648b4caa))
* rearraging provider docs ([a0aab65](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a0aab656a60e99ef2554e053d01fdad5c4539016))
* renamed user license assignment to follow resource naming pattern ([29f4c8b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/29f4c8b163dd7e1d95dfe045831fc64e125db76f))

## [0.14.1-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.14.0-alpha...v0.14.1-alpha) (2025-05-26)


### Bug Fixes

* added close stale issues maintainence pipeline ([1d90efd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1d90efd55aab386abff032eda94948987f69d50f))
* added close stale issues maintainence pipeline ([#475](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/475)) ([40affe2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40affe22d27c45128e8b4f7887419f10d41efc58))
* docs ([0a676f8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0a676f87825c2d75b2c3ed5175d48474e23ec28d))
* for doc gen pipeline ([f959ac8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f959ac87f1d6ea3e0d66b9979502a5c1f914466f))
* for doc gen pipeline ([d3fecdf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d3fecdf388ed6ad8e82fffffce38cbdd0126553a))
* for doc gen pipeline ([#471](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/471)) ([c93301e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c93301e7f95b32b56ed812ba5f373d2c2d3b0d80))
* for doc gen pipeline ([#473](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/473)) ([3c58522](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3c58522f6f1522eca9da2b4ee613b32b96319368))
* unit test trigger ([2bf41c4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2bf41c4a70e4a47e268bb9d68950dda48d24bfe0))
* unit test trigger ([#474](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/474)) ([e426630](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e42663025764a290342fde1a4f489a19931e94a1))

## [0.14.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.13.0-alpha...v0.14.0-alpha) (2025-05-25)


### Features

* added assignments for windows updates as a seperate set of resource types ([837ee80](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/837ee809bc587496f2e8be44996db786cbd0d10f))
* added datasource microsoft365_utility_macos_pkg_app_metadata ([2c2cbca](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2c2cbca7feb8d1336c9327e487767ecec1e4bc35))
* added managed device clean up rule with examples ([fc33266](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fc332665c817df572c0d349ec4d68ccf1b037570))
* added managed device clean up rule with examples ([#463](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/463)) ([f1a8777](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1a87770fa8e61a0edbc8370876cf86b4cb07850))
* added microsoft365_graph_beta_device_management_terms_and_conditions ([335e4a1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/335e4a1ec8441e7f49b980af60de4d07e2afea8f))
* added microsoft365_graph_beta_device_management_terms_and_conditions ([#466](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/466)) ([16849a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/16849a77e9262aeddbae84bc43b55552da62c146))
* added microsoft365_graph_beta_device_management_terms_and_conditions_assignment with examples ([c3db9ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c3db9acb7c2fb11122658bdf817c2ff57af01fec))
* added microsoft365_graph_beta_device_management_terms_and_conditions_assignment with examples ([#468](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/468)) ([6e407c0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e407c067f71cbc2f01cef185adaff9f214774db))
* added microsoft365_graph_beta_device_management_windows_feature_update_policy_assignment with examples ([74f286e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74f286eb668a342f700b1ce3c5da9671695ffdcb))
* added microsoft365_graph_beta/microsoft365_graph_beta_device_management_windows_driver_update_profile_assignment with examples ([863774c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/863774cbb6a0b9944c1d79b2e87ac2f4ad66a6b5))
* added mobile app data source with odata queries ([3a1c0f4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3a1c0f414e70ea9c55f36b31c1c48e85feb2aa02))
* added operation approval policy ([40b307c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40b307cfd8c1ec89fd342c7a29f20254d2cacaef))
* added resource operation for intune rbac with examples ([4f36c5f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4f36c5f5b50c9d19f3cdcdcc2507d8d8929a5931))
* added resource operations with examples for intune rbac ([#465](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/465)) ([1c31571](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1c315714f220e136300bf4da2f0c029dd39e2c11))
* added setting catalog in native hcl ([596d6c9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/596d6c9f0454810004ff30a3177a5c4ec2bc4ebc))
* added settings catalog in native hcl ([#461](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/461)) ([6e5e693](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e5e6937324d83df7fe5f96445dc45f851bf771f))
* added validation for mobile app assignments for 8 invalid scenarios ([cc88fd1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cc88fd104a35cf071e4a5693a84420f57d1151d8))
* added windows quality update profile assignment resource with examples ([a49a109](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a49a109323d23f0c8e0ed2be4afea53a77b1b05d))
* mobile app assignments ([80c892e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80c892ec95186025afc2c6f7edbd61f7ed1863e4))
* mobile app assignments resource with examples ([#453](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/453)) ([498ff50](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/498ff5087f942dffa701be9da736c8d3e1a5dde2))


### Bug Fixes

* final fixes for macos assignments ([8d998cf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8d998cfc27ccb4aa1b4bab3be60823f479b7cf85))
* final fixes for macos assignments ([#451](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/451)) ([7e3b853](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7e3b853731d22cc832840ceab3ca9c7a256418c4))
* for docs path ([2fce669](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2fce66923162574af62abe0c7236cf9da8bd76ab))
* for mutex handling for SetCommittedContentVersion ([058bf32](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/058bf32f301bc9b5082f668427aaf4ca3cff3878))
* for successful stating logging with resource name and resource id ([8c81932](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8c819320ff467a553640c8ff5ce9c78c5f7cf76a))
* for winget assignments ([8a2bdd0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8a2bdd03111f263f47db1614ba1e0fb8deb5a650))
* migrating winget app assignments to sets ([2272cc6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2272cc64166db8c02accd2a243d79a646fc3dfe4))
* minor fix for sorting order of assignments ([fb50e0f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fb50e0fbac88e72d4d784f400131449be1c4ff07))
* moved mobile app assignment validation to create and update funcs to fail early if any hcl mistakes before any rresources get deployed ([8c79958](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8c79958155e1445d959487a27cd66ddca38947b3))
* removed mutex, going with parallism of 1 instead ([dd0e6d8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dd0e6d822dcf3eb6fd43bbdea105d62020dc488f))
* typo fix for datasource ([5c60f13](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5c60f13a9ce6d8701d9310ea41f07e1104524967))
* win_get_app now skips property 'InstallExperience' and 'PackageIdentifier' as it cannot be patched during update operations. and added support for app categories ([8b5a4a5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8b5a4a59a43e6bddb6e82a3763713a85bcfe4047))

## [0.13.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.12.0-alpha...v0.13.0-alpha) (2025-05-12)


### Features

* mobile app assignments can now be set in any order with modify plan (diff supression) ([c8f97ef](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c8f97ef7201c1a79efc3850add74eac0e78468be))


### Bug Fixes

* added mutex to all api calls to defender against Kiota's middleware (e.g., HeadersInspectionHandler) when modifing shared header maps during HTTP request processing. appears to be a bug in the latest version ([976ad22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/976ad22dea6b1ddf6e3027714a04c449b73ac11c))
* added mutex to all api calls to defender against Kiota's middleware (e.g., HeadersInspectionHandler) when modifing shared header maps during HTTP request processing. appears to be a bug in the latest version ([#450](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/450)) ([10327c8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/10327c856db27c79ae6883677a48049c6caf2a1e))
* base resource updates for macos_pkg_app doesnt use redundant file size checks ([6162a1d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6162a1d4dc9860de2aa812ff386289bd6fc387f7))
* defined default values and plan modifier for role_scope_tag_ids ([e8909e0](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e8909e0fc876dbe03422bfb26943ac872cfab432))
* for parrallism issues with concurrent header map settings in client mobile apps ([49561f6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/49561f64cea87bd23eab9dda180824ddfd8934de))
* for role_scope_tag_ids ([#444](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/444)) ([b1e0fd8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b1e0fd819d4956ee065e96d4300fb24785719dc2))
* removed conditional access. complete refactor required with latest sdk version ([4cb550d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4cb550d56f376cb30bde71c71b4699e63653e07c))

## [0.12.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.11.0-alpha...v0.12.0-alpha) (2025-05-11)


### Features

* added refactored unit tests for provider build and credential factory ([bd7691c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bd7691c76d6403dadb4eb68d76ba37093f506b9b))
* added refactored unit tests for provider build and credential factory ([#439](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/439)) ([c191a93](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c191a9345c962f5f19a1bfd76dcafb3fb1deec01))
* mocking scaffolding ([f1be6d1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1be6d1f6fddfca091d34df19750da364545c1b1))
* mocks scaffolding ([#440](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/440)) ([9ee57df](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ee57dfc4eecebb95c35c0e60686925d3228c2c6))
* refactored data source for maospkgapp ([#442](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/442)) ([fec56b7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fec56b72934dc5c9fbd1dad05e8d690fab0793ab))
* refactored macospkg app datasource with example ([cbc1681](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cbc1681feca36a096ea63b69dc61540178b21447))


### Bug Fixes

* examples so that device_management examples are reflective ([d6856b4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d6856b480ef5a0c9f7277592bc391eb47487b443))
* examples so that device_management examples are reflective ([#441](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/441)) ([6062154](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6062154d0eedb88177fec252a4c5fb5f925ca24a))
* for enum handling within GetDeviceEnrollmentConfigurationType ([#437](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/437)) ([896c291](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/896c2918aa799ec20127a79538a06cfe1cff892e))
* for handling enum ([ce13f69](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ce13f693c89723f9766c02519f17596ea328d7e2))
* for handling enum ([6e2f1ff](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6e2f1ff38ef16ddbcd405eff88c380371e315856))
* refactored docs for repo restructure ([546ebb4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/546ebb4ff5dfe06b9b6c9febad4223a4acf71abf))
* refactored docs for repo restructure ([#435](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/435)) ([3b14e78](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3b14e78f1db1659b176983abced03e04b32195d6))

## [0.11.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.10.0-alpha...v0.11.0-alpha) (2025-05-07)


### Features

* added 1st itteration of oidc support for github actions and azdo ([3a2c14b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3a2c14bdcc51b039fc9269e9febf8edae89fa866))
* added application category resource and datasource with examples ([55a099a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/55a099a087f713f4e4c01558e4aec5fdfaf303f8))
* added datasource graphBetaDeviceAndAppManagementWindowsQualityUpdateProfile ([9f6f79d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9f6f79d7472eef5d7f76973c8195046b65c4462b))
* added device_and_app_management_windows_update_catalog_item datasource with examples ([655e804](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/655e804d7ea7989c549fa86a8925ea82b0b33d46))
* added graph_beta_device_and_app_management_windows_driver_update_profile_assignment ([73dc481](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/73dc481446cb3c88284ef6b5fcac719da08bfb81))
* added graph_beta_device_and_app_management_windows_feature_update_profile with example ([050cec1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/050cec1724781fa5bdf81b14c9c5890f5a0386e4))
* added helper SetObjectsFromStringSet for object construction ([e2b110e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e2b110e4f7b4aa73aa2e823a5ae9b58f5b597e8c))
* added macos pkg app datasource ([a65ca03](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a65ca038103e064225b7d136205179b7fbb04fad))
* added macos pkg datasource ([f0f6144](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f0f61447b307f221cdf6f0b0d65f74d8e4f6beb6))
* added microsoft365_graph_beta_device_and_app_management_windows_update_profile_assignment with example ([44e3539](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/44e3539f943013c5a0414e2053f56a1c02757f04))
* added timeouts to datasources ([#397](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/397)) ([100840f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/100840f1bff363a598fd9a76e26183c16c086150))
* added windows driver update profile and inventory resources and datasource ([#416](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/416)) ([8e3272d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8e3272dbc77296cdd874a11e7c7693103173419b))
* added windows driver update profile and inventory resources and datasources ([85ff58f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/85ff58f605f939be963c3aba58ebd44c0ac41f0f))
* device_enrollment_configuration with examples ([#420](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/420)) ([88516db](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/88516db1ff74964069d9965f29c21972bb147144))


### Bug Fixes

* added examples and more fixes for windows_remediation_script_assignments ([9827467](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9827467ac9972380540ccb29490670d5db960c87))
* added read timeouts for datasources ([243e6b2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/243e6b2fdc92441d54949d0d1bd9544b4b479c1e))
* categories fix for macos app categories and stating issues for infured included app values ([defa936](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/defa93679fa6bed0bc6065678987c76429383133))
* content version ([74c417b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/74c417b8312e6353bbbdbf78fd012d9fa96b39e1))
* for deps bump of go sdk version ([54a119e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/54a119edcca117fa697b21c59f2ba35cdda12ae2))
* for example path ([771bb1c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/771bb1c2792a16955e05f124cbe7a4e6d88cd3bb))
* for macos pkg apps ([#395](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/395)) ([deb9309](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/deb9309edc8a8aff4d33a022102f30f9c90eb679))
* for mobile app assignment id's and content version stating ([a57a269](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a57a26945a184c294e4c37c41b6b3457dbdad180))
* Go Lint configuration yml for provider project types ([#424](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/424)) ([b92b8ac](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b92b8ac28100a40611b0cca6e191750db9b8d277))
* go-lint ([#423](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/423)) ([57ba99d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/57ba99dd57bf8b26e9785767cf05449df1b2fca2))
* m365 app installl options ([e2db1a7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e2db1a7982f5cced9b3f346b7f2df8a30c92695a))
* migrated microsoft365_graph_device_and_app_management_m365_apps_installation_options to graph v1.0 ([a3c9682](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a3c968250d67787711d46b2589facd62cf3ab8b3))
* numerous edits and tweaks ([dccd487](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dccd4873fc4dec676b2a0e3859c7a2a42d595b4f))
* numerous schema fixes ([ddb6e87](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ddb6e8780360a63c31354f0899314fd4ed51c111))
* pipeline step naming ([3fade5e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3fade5e96ab87097c3b8b18dc3e2f0ea5850a7a3))
* pipeline step naming ([#425](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/425)) ([8067017](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/80670170abf8aa817432805c0d4a786859655951))
* refining logic for handling app metadata ([9ecf223](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9ecf2236482cffbe5ef3cc67cf9e10e94068f345))
* restructured repo to follow api endpoint paths more precisely ([1b2c110](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1b2c1108add87c7395eaed7cfce79887c7899f12))
* restructured repo to follow api endpoint paths more precisely ([#433](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/433)) ([6aed990](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6aed9906d8f76430c646909e6111c6d91b29f1b3))
* tidy up pipeline step naming ([102e356](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/102e35695ac0629381b740e08aec981cfdaf84a1))
* tidy up repo restructure ([b5b35cf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b5b35cf4b001d4ba9ebc0cfed8913bc8a7054e12))
* windows remediation script assignments ([04243b6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/04243b62ecb12ed869edec6caf334e6671c8fd90))

## [0.10.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.9.0-alpha...v0.10.0-alpha) (2025-04-08)


### Features

* 1st itteration of intune applications. ([1a83d27](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1a83d277be16d0e08fd61d32919d07e39d776313))
* add debug logging for Graph API request bodies in constructResource function ([3fdbba4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3fdbba463403254f7dbd8e3473d48f7a3d71e3a3))
* add macOS PKG metadata extraction functionality and update dependencies ([bedd812](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bedd8121ef7b3c1bca5e5022ec2db27c88bb2f37))
* add package installer file source and extract metadata for MacOS PKG applications ([7cfd15a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7cfd15a8b37a61d5a5cc18abb4f6c9c86739798c))
* add support for extracting files from CPIO archives and enhance payload extraction logic ([609fc22](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/609fc22d130a2480dbd533d737a22c08192138f1))
* add validation for MacOSPkgApp publisher field and implement requiredWith validator ([6000864](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6000864632b107c490e6dbe0d899766af23aafce))
* added 3 secuirty baseline settings catalog templates ([ff7b7de](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ff7b7de4f4d984bce801a182cd6394b5bcbcbd05))
* added debug logging to uploadToAzureStorage func ([ec57f66](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ec57f66dce52a0153ed2f8ead56154913a9710a5))
* added Encrypted file analysis for troubleshooting for macos packages ([10886cb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/10886cb97abb6b533e22f7ef38eca7681e86a06c))
* added first itteration of schema builder for settings catalog ([7590e6f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/7590e6fabb8fd49a3b7f5111cd476ddcf5a6fff2))
* added intune device categories resource and data source with examples ([#389](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/389)) ([b8d33fb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b8d33fba7b11121bf61514164b77d01858a9a74c))
* added macospkg resource to provider ([#347](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/347)) ([10001df](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/10001df8134fe806c661b94850531f98b245643f))
* added new type conversion helpers that handle null values ([8aaa14e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8aaa14ed15e4805bd63309d8de32db5b1266b284))
* added ps scripts for testing dmb, pkg and macoslob apps. also added some getters to help with data model definitions ([37877de](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/37877defb6ae270df8059273e31322bace8708e6))
* added schema Mutually Exclusive Attribute validator ([e2c7deb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e2c7deb3b6b008c0556c296bf2f31eb8b7c3c2a9))
* added secuirty baseline templates support to settings catalog template resource ([8e4649b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/8e4649b21793674f196d27c11608c73ffb87be38))
* added support for url based and local file based icons and pkg installers ([c75ca2c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c75ca2c77727f7ce35e3f5990eb6b90e488ab71f))
* added web sources for icon uploads for macos app ([dbc36d6](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dbc36d6f5642ad011adaedf0729ab0fb87b88407))
* completed macos pkg app creation flow ([99a96fb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/99a96fb43efc1b09483c5f98c29f21138027a338))
* enhance schema for package installer with validation and improved descriptions ([5488a54](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5488a54f7c3fbacf79ac359ab14c5109f3c38ab8))
* implement XAR file extraction and field parsing functionality ([c6dab71](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c6dab71a99f18b5a08708c60a3771cbe4914c8af))
* implement XAR file parsing and metadata extraction for macOS installers ([b91e24c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b91e24c2030639be1052b4b97292e821018b0b8c))
* implement XAR package reader and metadata extraction for MacOS installers ([5d70348](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5d703488093e25721b0ffae8b797ec6ad70b5aa7))
* implemented device categories resource and datasource with examples ([acf2626](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/acf262635ca9aeb5175f0744fc108dfd0694b027))
* updated macos pkg example ([6ecc553](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6ecc553ce50049b5afd46e72a4671dd27211c342))
* updated read and stating logic ([82a2cfd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/82a2cfd8f2d39eb1ce758e39428170e77a94e91e))


### Bug Fixes

* added content version assignment ([0418ecb](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0418ecb57f547d216f7121e1aa187a12cf3758be))
* added initial stating logic to mascos create func ([e99d915](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e99d915beef8bf56424eab4197663a712c0fe92f))
* added planmodifiers.UseStateForUnknownString(), for id ([dccde7e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dccde7e78d92bf90a23bbdf7ad5673391baafc65))
* changed assignment optional settings to a pointer ([b3bb2d7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b3bb2d74af6c3732af8b3eba83cd95c42b255d36))
* extended file upload chunk time ([6f7ba33](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6f7ba3328934b9d7256c887034dc3f102b16d59e))
* for package imports ([95c9a15](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/95c9a156c5fbd9c53e6dc467c0150c011b84eeeb))
* for stating icons to use stringunknwon logic ([0006aec](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0006aec1068ea0ceb599998b9183096525c585e8))
* improve comments and formatting in constants and assignment filter read functions ([e48f5c7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e48f5c7a4ca32e3e5db93abe0a243a47e81e159d))
* nil pointer dereference in func validateInstallTimeSettings ([b52d393](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b52d39323ffb12a8b76d2fee02a9994fb571be43))
* numerous fixes for constructors and stating logic ([4d09816](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4d09816644f027bdbac3de734ff2984f6c0624a3))
* removed redundant struct field ([a7a8094](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a7a80945b91f07d4b5fb87871a388ef39e90c6df))
* testing new encryption strategy for pkg's with intune ([203c69e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/203c69e877f275ba56dce3e8673db86bdac26994))
* timeout usage of consts ([d94aec2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d94aec2d5e2da25d9de54f256477671590ea6582))

## [0.9.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.8.0-alpha...v0.9.0-alpha) (2025-01-18)


### Features

* add device management template type attribute with validation options ([d07cbb9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d07cbb92532ebead04ed8f956b53017ef0b7e798))
* add new Windows Defender Antivirus policy templates to device management configuration ([45834c3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/45834c33cb439eb5166976e58abba68f3bfb5f76))
* add settings catalog template and update related configurations ([a3b9a33](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a3b9a33347942f8545b4878d7de88dc195ab831e))
* add template for graph beta device and app management reusable policy setting documentation ([01334e2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/01334e26b1fa9092bc9bd141ca161f7bdc0f0bef))
* add validation for settings templates and extend policy configuration map with new templates ([b240513](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b24051395001714144e98e3c9d149164bca61db9))
* add Windows Firewall rules template and remove deprecated script ([4b264f7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4b264f781e65bf776da40f17037a2a683b3860b1))
* added additional endpoint security templates to settings catalog templates resource + numerous doc edits ([#318](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/318)) ([d9f3ff2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d9f3ff22346ccffdc635070d316cd7c6323b03b4))
* added all remaining settings catalog templates for Intune Endpoint security ([#319](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/319)) ([e9a9320](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e9a9320ae6e7d7ba7eb68ebe86a60e7452948b0a))
* added settings catalog template to provider with examples ([#314](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/314)) ([3f062f1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/3f062f113536da431e3abd606174cb33768e3698))
* enhanced setting catalog construction logic ([eec7669](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/eec7669ad0cb63f0561f4617ed0509ef2eb21f35))
* extend policy configuration map with new settings_catalog_templates for endpoint security ([1d76a96](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1d76a96e6bcbed86574d59c4378406d9bbca8fd6))


### Bug Fixes

* update example file path for reusable policy setting documentation ([a944199](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a94419961b97490b0b187dc3b29c41b16ce48405))
* update settings key in resource configuration for Endpoint Privilege Management ([763c8ae](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/763c8aed4b7b9dedbc32d7a83397b3d08f1cc547))
* update Terraform version requirements in documentation and configuration ([97a200c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/97a200cbe4e9805c166e3690322e885d507eabd3))
* updated schema in settings catalog templates to reflect all secuirty templates supported with descriptions ([d9c1e7a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d9c1e7a255a0adb8483981fc149ea3c9d5dd61a5))

## [0.8.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.7.0-alpha...v0.8.0-alpha) (2025-01-15)


### Features

* add endpoint privilege management resource and example usage ([b7b4b3b](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b7b4b3b0ff5d3fed63cd4aa52dce7f173690baa7))
* add Linux platform script resource and example ([bcd6530](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bcd65300c18bed23339a10c1ac8eaf90ec148f3d))
* add plan modifiers to use state values for unknown attributes in reusable policy settings ([f8f2c14](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f8f2c1423a437ccc300857436af72cd6c6b16c37))
* add reusable policy settings data source and example usage for Endpoint Privilege Management ([d5a5d10](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d5a5d107b5fe54d3f9a35a8f9551577ba95e6898))
* add reusable policy settings models and modify plan handling ([5029495](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5029495db1a3cbd5465235d103aa40b9488f7acd))
* added endpoint privilege management resource and example usage ([#303](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/303)) ([48a2514](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/48a2514f4cccca52df60431c9699d2eda3d80b76))
* added linux platform script with examples ([#302](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/302)) ([6324ebe](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6324ebeafdb411527b0801664df04b20c85efe11))
* enhance reusable policy settings API calls with additional select parameters and improved debug logging ([fc03249](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/fc0324948032fe79fd83742eed43a7a530dafcf3))
* enhance setting instance handling in reusable policy settings resource ([40f98e7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/40f98e7e99025e64998ae645acf0cd23108001ec))
* enhance state handling for reusable policy settings and normalize JSON responses ([6c90118](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6c901185f05926f5e2a99f8bd9921ce6d7417ec0))
* implement custom DELETE request handling and refactor URL template configuration ([ec7c0a2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ec7c0a24f2a4fc16d6d996006c172a0e90fcbf51))
* implemented data source for reuseable policies for epm ([#312](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/312)) ([6b0bbbe](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6b0bbbed729938659e8193a82021bf1a65a20697))
* improve error handling and logging in StateReusablePolicySettings function ([0361f31](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/0361f314aedf74f34261ad94788995d39bacd3ec))
* refined logic for reuseable policies and updates to schema ([#305](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/305)) ([00bd476](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/00bd4769d1d44d85bcd7b1344492e11675ef2cb7))
* rename settingsDetails to settings for consistency across models and resources and to align with intune gui exports ([937cdbd](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/937cdbdbbef05e8f64ae25ea122916b485d1be8b))
* update reusable policy settings documentation and add example resource for Endpoint Privilege Management ([a0aa693](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a0aa6936c2098b84ab0700ce90e67441b0cef877))
* update reusable policy settings model and integrate into provider resources ([f1e528d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f1e528df99b148ac7d3887ff8d71ae59cd242715))


### Bug Fixes

* add comment to clarify ConfigurationPolicyTemplates requirement in state_base_resource.go ([f38ebe7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f38ebe7c314a50359b1fc3830b68b31df1b23976))
* correct resource naming for reusable policy settings to singular form ([5ccd617](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ccd617bb4ce09e98c0079656afd4192ec908b15))
* correct resource naming for reusable policy settings to singular form ([#313](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/313)) ([db7ab1e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/db7ab1e53e04586f0a89b161243cec18abbc62c3))
* remove unused plan modifiers for created and last modified date attributes in reusable policy settings ([878f800](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/878f8000a328967da6ffe20f0ba9195729ee3dd0))
* Update Graph Metadata - 2025-01-05_00-07-29 ([#298](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/298)) ([43a831d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/43a831d447bda7122e00c8a7b31ec18ad8cba93d))
* update PowerShell script links in Markdown descriptions for reusable policy settings and settings catalog ([a6c0292](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a6c02923de3fdfa33bd412746c52109e30544a5a))

## [0.7.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.6.0-alpha...v0.7.0-alpha) (2025-01-04)


### Features

* add mapping functions for various remote assignment settings to Terraform ([e24fb0f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e24fb0f0851a662a562a210a1902d74343641ed9))
* add resource documentation and modify plan handling for various device management scripts ([644dbc9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/644dbc996981837466b0d30d9bf00e6f2721a1e9))
* add StringListToTypeList function for converting string slices ([#293](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/293)) ([cd62438](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/cd62438c70bdee68a2b1ec30184d098c1c72d17a))
* add StringListToTypeList function for converting string slices to types.List ([dd28488](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/dd28488c7fd29d58c8331a9e904155acaa94f5aa))
* add validation for mobile app assignment ordering ([c0fdeb7](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/c0fdeb744e9ac2218a7fc54d07c40310cc2ead84))
* added mobile app assignment schema for all app types ([ec8edc1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ec8edc1c22f91bab217489985541be0fb14410c7))
* added the option to manually define winget app metadata along side auto generation + plan modifers ([#281](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/281)) ([6049e8e](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/6049e8ec1b32d42c1449f263b9723dc4f8d6c372))
* enhance mobile app assignment configuration with new settings and sorting logic ([81dd048](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/81dd0486edf1bda8e3b59490faf464059b4834ec))
* implement mobile app assignment validation and update related constructors ([5bbc084](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5bbc084ac6f42644c94b8bd073c4866c04fdacbb))
* implement validation for mobile app assignment settings and restart timing relationships ([28aa0b3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/28aa0b34a27577a132e6c4067efe88da3fbd2ee8))
* implement validation for mobile app assignment settings and restart timing relationships ([#292](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/292)) ([ed375cc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/ed375cce3138b1b1bd35844ff59d0842904843c2))
* refactored stating and constructor func patterns to be more concise and leverage correct lib and project helpers ([#277](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/277)) ([694ade9](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/694ade95bf2d9d761e0c187f868836db8054c96d))
* standardized stating structure and constructor approach ([1db44c8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/1db44c86b3fd9afc5c8fa72a916c461c3af0bf4f))
* standardized stating structure and constructor approach ([#286](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/286)) ([e798fd5](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e798fd567519686f46aec1bbd738edc303af456b))


### Bug Fixes

* add target type field to AssignmentTargetResourceModel for improved clarity ([41cbea2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/41cbea24ae15fcb023281076179bdb14529c601d))
* added function comments for mobile app assignments ([#294](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/294)) ([d742bd4](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d742bd4efd35b70be286acaadf94e0e35df36a2b))
* bug fixes for conditional access policies and fixes for crud permissions ([#267](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/267)) ([4d7f8d8](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/4d7f8d86e44e9d1ee88723f25d8279fc4de514ee))
* centralised settings catalog assignments ([#287](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/287)) ([751e2f3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/751e2f3277a1adfc4f19a9b428a3eb9b1938400a))
* for constructAssignmentTarget within mobile app assignments ([#289](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/289)) ([5ee4531](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/5ee45311814e9a2bc56810be63c4f13f78dba663))
* numerous fixes in docs and pipelines ([#268](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/268)) ([b243425](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/b243425bb0bbc2ac00a69d0bbce03c27f2f51f3f))
* refined gorelease pipeline validation ([#269](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/269)) ([18ff5aa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/18ff5aadafd3449983fa655deb5fbab6cd767427))
* remove obsolete mobile app assignment resource from provider ([#297](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/297)) ([af9b4cf](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/af9b4cf84eebb03e25c84868501bacd1c2be5d73))
* removed icons plan from plans ([47bc42a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/47bc42a81d109f0a411456b22a4c92645219d1b5))
* rename struct for clarity and add configuration policy assignment constructor ([32b5b65](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/32b5b65b48aa167e4ae24ddd2778e928fe6f1d6e))
* reorganised repo to use graph_beta and and graph_api consistently for all package naming ([#290](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/290)) ([54fc3aa](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/54fc3aa0e9091bb56af3142ab2bd1a13a92e442f))
* replace constructAssignment function calls with specific constructors and remove obsolete construct_assignment.go file ([#288](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/288)) ([2475c9f](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2475c9f954dc9ab06b29aa18bb04a29d00c4ec3c))
* sorting mobile app assignment stating logic ([#295](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/295)) ([9de00c2](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/9de00c2792a64a55b8c0d7cfd4541a1e8242048b))
* standardised use of object throughout crud functions ([#259](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/259)) ([bbc3dc3](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/bbc3dc3c0a8b3a3b26d16b86b65d0300912995db))
* streamline property handling in WinGetApp resource mapping ([#278](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/278)) ([025a5b1](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/025a5b144d8b0255e38e7d430fa751e911c5c7eb))
* tidied up repo and add validation for mobile app assignment order ([#296](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/296)) ([036e95a](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/036e95a1b0b4364c5b27fcbe02ba26b291e0d954))
* update import paths for device and app management resource to use graph_beta and graph_v1.0 ([#291](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/291)) ([62aa968](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/62aa9686904eca813e6e921b34265a8354a34074))
* update import paths from 'construct' to 'constructors' for consistency ([d878700](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d8787001898571f07e5a1d65fabcef6c679e00ef))
* update import paths from 'construct' to 'constructors' for consistency ([#284](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/284)) ([d878700](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d8787001898571f07e5a1d65fabcef6c679e00ef))
* updated docs ([#283](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/283)) ([e76a368](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e76a368632c7a8e077c016f89e35978c5d9299aa))
* updated win_get examples ([#282](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/282)) ([e54f1fc](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/e54f1fc540368183230afed9dfe0e214ba99b353))
* various small fixes to docs and pipelines ([#271](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/271)) ([d244e16](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/d244e1633fcbe824f19a8d3e82ffe5ff8f3d8a0c))

## [0.6.0-alpha](https://github.com/deploymenttheory/terraform-provider-microsoft365/compare/v0.5.0-alpha...v0.6.0-alpha) (2024-12-16)


### Features

* added tf-docs auto generation pipeline for merge into main ([#257](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/257)) ([2b4e9ad](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/2b4e9add53640a4c950248b23deec223c60e181b))
* intune role scope tags + added release please ([#246](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/246)) ([f09c60c](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/f09c60cbfac38b0bf78e83aaf1b5d18903714d34))
* refactored datasource examples to support search by name or by resource id ([#255](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/255)) ([a07bd03](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/a07bd035ac0e6b0a2324673adc59ee31eff7136e))


### Bug Fixes

* added retry logic from sdkv2 to settings catalog resource types ([#250](https://github.com/deploymenttheory/terraform-provider-microsoft365/issues/250)) ([deb384d](https://github.com/deploymenttheory/terraform-provider-microsoft365/commit/deb384d335dba7cf219f099676213d957c7663c5))

## [Unreleased]

### Added

- Added xyz [@your_username](https://github.com/your_username)

### Fixed

- Fixed zyx [@your_username](https://github.com/your_username)

## [1.1.0] - 2021-06-23

### Added

- Added x [@your_username](https://github.com/your_username)

### Changed

- Changed y [@your_username](https://github.com/your_username)

## [1.0.0] - 2021-06-20

### Added

- Inititated y [@your_username](https://github.com/your_username)
- Inititated z [@your_username](https://github.com/your_username)
