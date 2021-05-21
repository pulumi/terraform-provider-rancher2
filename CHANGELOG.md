## 1.15.1 (May 21, 2021)

FEATURES:



ENHANCEMENTS:



BUG FIXES:

* Added timeout to `CatalogV2Client` function when getting new catalog v2 client

## 1.15.0 (May 20, 2021)

FEATURES:

* **Deprecated Argument:** `rancher2_cluster.aks_config.tag` - (Deprecated) Use `tags` argument instead as []string
* **New Argument:** `rancher2_cluster.aks_config.tags` - (Optional/Computed) Tags for Kubernetes cluster. For example, `["foo=bar","bar=foo"]` (list)
* **New Argument:** `rancher2_cluster.agent_env_vars` - (Optional) Optional Agent Env Vars for Rancher agent. Just for Rancher v2.5.6 and above (list)
* **Deprecated provider Argument:** `retries` - (Deprecated) Use timeout instead
* **New provider Argument:** `timeout` - (Optional) Timeout duration to retry for Rancher connectivity and resource operations. Default: `"120s"`
* **New Argument:** `rancher2_cluster.oke_config.pod_cidr` - (Optional) A CIDR IP range from which to assign Kubernetes Pod IPs (string)
* **New Argument:** `rancher2_cluster.oke_config.service_cidr` - (Optional) A CIDR IP range from which to assign Kubernetes Service IPs (string)

ENHANCEMENTS:

* Added timeout to `CatalogV2Client` function when getting new catalog v2 client

BUG FIXES:

* Fixed `rancher2_cluster.hetzner_config.UsePrivateNetwork` with proper json field name

## 1.14.0 (May 7, 2021)

FEATURES:

* **New Argument:** `rancher2_cluster.oke_config.limit_node_count` - (Optional) The maximum number of worker nodes. Can limit `quantity_per_subnet`. Default `0` (no limit) (int)
* **New Argument:** `rancher2_cluster.rke_config.ingress.default_backend` - (Optional) Enable ingress default backend. Default: `true` (bool)
* **New Argument:** `rancher2_cluster.rke_config.ingress.http_port` - (Optional/Computed) HTTP port for RKE Ingress (int)
* **New Argument:** `rancher2_cluster.rke_config.ingress.https_port` - (Optional/Computed) HTTPS port for RKE Ingress (int)
* **New Argument:** `rancher2_cluster.rke_config.ingress.network_mode` - (Optional/Computed) Network mode for RKE Ingress (string)
* **New Argument:** `rancher2_cluster.rke_config.ingress.update_strategy` - (Optional) RKE ingress update strategy (list Maxitems: 1)
* **New Argument:** `rancher2_cluster.rke2_config` - (Optional/Computed) The RKE2 configuration for `rke2` Clusters. Conflicts with `aks_config`, `eks_config`, `gke_config`, `oke_config`, `k3s_config` and `rke_config` (list maxitems:1)
* **New Argument:** `rancher2_cluster_sync.wait_alerting` - (Optional) Wait until alerting is up and running. Default: `false` (bool)
* **New Argument:** `rancher2_cluster.gke_config_v2` - (Optional) The Google GKE V2 configuration for `gke` Clusters. Conflicts with `aks_config`, `eks_config`, `eks_config_v2`, `gke_config`, `oke_config`, `k3s_config` and `rke_config`. For Rancher v2.5.8 or above (list maxitems:1)
* **New Argument:** `rancher2_cloud_credential.google_credential_config` - (Optional) Google config for the Cloud Credential (list maxitems:1)

ENHANCEMENTS:

* Updated `rancher2_catalog_v2` schema resource, defining conflict between `git_repo` and `url` arguments
* Improved `rancher2_cluster_sync` with new cluster state check method and new option to wait until alerting is enabled
* Updated go mod to support Rancher `v2.5.8`
* Updated acceptance tests to use Rancher `v2.5.8`

BUG FIXES:

* Fix `rancher2_node_pool` resource, adding `forcenew` property to not updatable arguments
* Fix `rancher2_cluster` resource, fixing provider crash if `cluster_monitoring_input` argument is deleted
* Fix `rancher2_project` resource, fixing provider crash if `project_monitoring_input` argument is deleted
* Fix `rancher2_catalog_v2` resource, just setting default `git_branch` value if `git_repo` is specified
* Fix `rancher2_cluster.eks_config_v2` argument, setting `private_access`, `public_access` and `secrets_encryption` as computed argument, removing default value

## 1.13.0 (March 31, 2021)

FEATURES:

* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.image_id` - (Optional) The EKS node group image ID (string)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.launch_template` - (Optional) The EKS node groups launch template (list Maxitem: 1)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.launch_template.id` - (Required) The EKS node group launch template ID (string)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.launch_template.name` - (Optional/Computed) The EKS node group launch template name (string)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.launch_template.version` - (Optional) The EKS node group launch template version. Default: `1` (int)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.request_spot_instances` - (Optional) Enable EKS node group request spot instances (bool)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.resource_tags` - (Optional) The EKS node group resource tags (map)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.spot_instance_types` - (Optional) The EKS node group sport instace types (list string)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.subnets` - (Optional) The EKS node group subnets (list string)
* **New Argument:** `rancher2_cluster.eks_config_v2.node_groups.user_data` - (Optional) The EKS node group user data (string)
* **New Argument:** `rancher2_cluster_sync.wait_catalogs` - (Optional) Wait until all catalogs are downloaded and active. Default: `false` (bool)
* **New Attribute:** `rancher2_cluster.eks_config_v2.node_groups.version` - (Computed) The EKS node group version (string)
* **New Attribute:** `rancher2_app_v2.system_default_registry` - (Computed) The system default registry of the app (string)
* **New Data Source:** `rancher2_secret_v2` - Provides a Rancher V2 Secret V2 data source
* **New Resource:** `rancher2_secret_v2` - Provides a Rancher V2 Secret V2 resource

ENHANCEMENTS:

* Updated go mod to support Rancher `v2.5.7`
* Updated acceptance tests to use Rancher `v2.5.7`
* Updated `rancher2_cluster_sync` to allow wait until all catalogs are downloaded and active

BUG FIXES:

* Fix `rancher2_app_v2` to respect Rancher system default registry
* Fix `rancher2_cluster.eks_config_v2` to deploy properly EKS clusters
* Fix `rancher2_catalog_v2` to wait until `downloaded` status

## 1.12.0 (March 05, 2021)

FEATURES:

* **New Argument:** `rancher2_node_template.node_taints` - (Optional) Node taints. For Rancher v2.3.3 or above (List)
* **New Argument:** `rancher2_cluster.aks_config.load_balancer_sku` - (Optional/Computed) Load balancer type (basic | standard). Must be standard for auto-scaling
* **New Argument:** `rancher2_cluster.rke_config.services.etc.backup_config.timeout` - (Optional/Computed) Set timeout in seconds for etcd backup. Just for Rancher v2.5.6 and above
* **New Data Source:** `rancher2_global_role` - Provides a Rancher V2 Global Role data source
* **New Resource:** `rancher2_global_role` - Provides a Rancher V2 Global Role resource
* **New Resource:** `rancher2_feature` - Provides a Rancher V2 Feature resource. Just for Rancher v2.5.0 and above

ENHANCEMENTS:

* Updated `rancher2_node_template.openstack_config` to support `boot_from_volume` and related arguments
* Added `password` as valid `cluster_template_questions` type to `rancher2_cluster` resource
* Preserve `cluster_template_answers` for `cluster_template_questions` of type `password` in `rancher2_cluster` resource to avoid misleading diffs
* Added `nodes` attribute reference to `rancher2_cluster_sync` resource
* Updated go mod to support Rancher `v2.5.6`
* Updated acceptance tests to use Rancher `v2.5.6`
* Added retry to get k8s default version, if getting forbidden or server error
* Added retry to get V2 catalogs and apps, if getting server error

BUG FIXES:

* Fixed cluster and project resource for update monitoring version properly
* Fixed `rancher2_app_v2` resource, added retry to GetAppV2OperationByID if got apierr 500
* Fixed `rancher2_cluster` docs, annotations and labels argument description

## 1.11.0 (January 08, 2021)

FEATURES:

* **New Argument:** `rancher2_node_template.hetzner_config` - (Optional) Hetzner config for the Node Template (list maxitems:1)
* **New Argument:** `rancher2_cluster.rke_config.dns.linear_autoscaler_params` - (Optional) LinearAutoScalerParams dns config (list Maxitem: 1)
* **New Argument:** `rancher2_cluster.rke_config.dns.update_strategy` - (Optional) DNS update strategy (list Maxitems: 1)
* **New Argument:** `rancher2_notifier.dingtalk_config` - (Optional) Dingtalk config for notifier (list maxitems:1)
* **New Argument:** `rancher2_notifier.msteams_config` - (Optional) MSTeams config for notifier (list maxitems:1)
* **New Data Source:** `rancher2_global_dns_provider` - Provides a Rancher V2 Global DNS Provider data source
* **New Resource:** `rancher2_global_dns` - Provides a Rancher V2 Global DNS resource
* **New Resource:** `rancher2_global_dns_provider` - Provides a Rancher V2 Global DNS Provider resource

ENHANCEMENTS:

* Updated `rancher2_app_v2.chart_version` as optional/computed argument. Deploying latest app v2 version if `chart_version` is not provided
* Updated `rancher2_app_v2.wait` default value to `true`
* Updated go mod to support Rancher `v2.5.4`
* Updated acceptance tests to use Rancher `v2.5.4`

BUG FIXES:

* Fixed `rancher2_cluster` resource, added retry when enabling cluster monitoring and got apierr 500. https://github.com/rancher/rancher/issues/30188
* Fixed `rancher2_cluster` datasource error, when `rke_config.services.kube_api.secrets_encryption_config.custom_config` or `rke_config.services.kube_api.event_rate_limit.configuration` are set. https://github.com/rancher/terraform-provider-rancher2/issues/546
* Fixed `rancher2_cluster_template` required argument definition on docs
* Fixed `Apps & marketplace` guide for Rancher v2.5.0 format
* Fixed doc examples for activedirectory, freeipa and openldap auth providers
* Fixed `rancher2_app_v2` resource to properly pass global values to sub charts. https://github.com/rancher/terraform-provider-rancher2/issues/545
* Fixed `rancher2_app_v2` resource to don't override name nor namespace on App v2 not certified by rancher
* Fixed `rancher2_cluster` docs, adding missed `gke_config.enable_master_authorized_network` argument

## 1.10.6 (November 11, 2020)

FEATURES:



ENHANCEMENTS:


BUG FIXES:

* Fixed `flattenClusterTemplateRevisions` func to avoid crash on `rancher2_cluster_template` resource at some circumstances

## 1.10.5 (November 11, 2020)

FEATURES:

* **Deprecated Argument:** `rancher2_cluster.eks_import` - (Optional) Use `rancher2_cluster.eks_config_v2` instead. Just for Rancher v2.5.0 and above
* **New Argument:** `rancher2_cluster.eks_config_v2` - (Optional) EKS cluster import and new management support. Just for Rancher v2.5.0 and above

ENHANCEMENTS:

* Updated go mod to support Rancher `v2.5.2`
* Updated acceptance tests to use Rancher `v2.5.2`
* Improved `rancher2_bootstrap` on resource creation. Number of retires on `bootstrapDoLogin` function can be configured with `retries` provider argument
* Updated `rancher2_catalog_v2` contextualized resource id with `cluster_id` prefix
* Updated `rancher2_app_v2` contextualized resource id with `cluster_id` prefix
* Updated `rancher2_app_v2` to show helm operation log if fail
* Updated `rancher2_app_v2.values` argument as sensitive

BUG FIXES:

* Fixed `rancher2_cluster.rke_config.upgrade_strategy.drain` argument to set false value properly
* Fixed `Apps & marketplace` guide for Rancher v2.5.0 format
* Fixed `rancher2_app_v2.values` argument to avoid false diff
* Fixed `rancher2_cluster_role_template_binding` and  `rancher2_cluster_role_template_binding` arguments to forceNew on update

## 1.10.4 (October 29, 2020)

FEATURES:

* **New Argument:** `rancher2_cluster.oke_config` - (Optional) Oracle OKE configuration
* **New Argument:** `rancher2_node_template.openstack_config.application_credential_id` - (Optional) OpenStack application credential id
* **New Argument:** `rancher2_node_template.openstack_config.application_credential_name` - (Optional) OpenStack application credential name
* **New Argument:** `rancher2_node_template.openstack_config.application_credential_secret` - (Optional) OpenStack application credential secret
* **New Argument:** `rancher2_notifier.dingtal_config` - (Optional) Dingtalk config for notifier. For Rancher v2.4.0 and above (list maxitems:1)
* **New Argument:** `rancher2_notifier.msteams_config` - (Optional) MSTeams config for notifier. For Rancher v2.4.0 and above (list maxitems:1)
* **New Argument:** `rancher2_cluster.eks_import` - (Optional) EKS cluster import and new management support. Just for Rancher v2.5.0 and above
* **New Argument:** `rancher2_bootstrap.ui_default_landing` - (Optional) Set default ui landing on Rancher bootstrap. Just for Rancher v2.5.0 and above
* **New Data Source:** `rancher2_catalog_v2` - Support new Rancher catalog V2 datasource. Just for Rancher v2.5.0 and above
* **New Resource:** `rancher2_catalog_v2` - Support new Rancher catalog V2 resource. Just for Rancher v2.5.0 and above
* **New Resource:** `rancher2_app_v2` - Support new Rancher app V2 resource. Just for Rancher v2.5.0 and above

ENHANCEMENTS:

* Added new computed `ca_cert` argument at `rancher2_cluster` resource and datasource
* Delete `rancher2_app` if created and got timeout to be active
* Updated golang to v1.14.9 and removing vendor folder
* Updated go mod to support Rancher `v2.5.1`
* Added dingtal_config and msteams_config arguments at rancher2_notifier resource. go code and docs
* Improved `rancher2_cluster_sync` wait for cluster monitoring
* Improved `rancher2_bootstrap` on resource creation. `bootstrapDoLogin` function will retry 3 times user/pass login before fail
* Updated acceptance tests to use Rancher `v2.5.1`, k3s `v1.18.9-k3s1` and cert-manager `v1.0.1`
* Added new `Apps & marketplace` guide for Rancher v2.5.0

BUG FIXES:

* Fix `rke_config.monitoring.replicas` argument to set default value to 1 if monitoring enabled
* Fix Rancher auth config apply on activedirectory, freeipa and openldap providers
* Fix `rancher2_cluster.rke_config.upgrade_strategy.drain` argument to set false value properly


## 1.10.3 (September 14, 2020)

FEATURES:



ENHANCEMENTS:



BUG FIXES:

* Fix `Error: string is required` upgrading rancher2 provider from v1.10.0 or lower

## 1.10.2 (September 10, 2020)

FEATURES:



ENHANCEMENTS:

* Updated go mod, vendor files and provider tests to support rancher 2.4.8 and k3s v1.18.8-k3s1
* Added `rancher2_cluster_sync.state_confirm` argument to wait until active status is confirmed a number of times
* Added `syslog_config.enable_tls` argument to cluster and project logging

BUG FIXES:

* Fix `rke_config.cloud_provider.name` argument to not be validated
* Fix `rancher2_certificate` resource update
* Fix false diff if `rancher2_project.project_monitoring_input` not specified
* Fix `rancher2_token.ttl` argument to work properly on Rancher up to v2.4.7
* Fix `rancher2_namespace.resource_quota` argument to computed
* Fix `rancher2_app` resource to wait until created/updated

## 1.10.1 (August 27, 2020)

FEATURES:



ENHANCEMENTS:

* Added `nsg` support on `azure_config` argument on `rancher2_node_template` resource
* Updated go mod, vendor files and provider tests to support rancher 2.4.6
* Added aws kms key id support to `rancher2_node_template`

BUG FIXES:

* Fix `rke_config.event_rate_limit.configuration` argument to work properly
* Fix cluster and project role template binding doc files name
* Fix `rancher2_cluster_sync` resource error if referred cluster deleted out of band
* Fix `rancher2_namespace` and `rancher2_project` resources error if destroyed by not global admin user
* Fix `rancher2_app` resource error if referred project deleted out of band
* Fix `rancher2_app` doc typo on `target_namespace` argument description
* Fix `rancher2_cluster` and `rancher2_project` resources error if created with monitoring enabled by not global admin user
* Fix `rancher2_token` to set annotations and labels as computed attibutes
* Fix `rke_config.secrets_encryption_config.custom_config` argument to work properly
* Fix `rancher2_token.ttl` argument to work properly on Rancher v2.4.6
* Fix `rancher2_project` resource applying `pod_security_policy_template_id` argument on creation

## 1.10.0 (July 29, 2020)

FEATURES:

* **Deprecated Argument:** `rancher2_cluster.enable_cluster_istio` - Deploy istio using `rancher2_app` resource instead
* **New Argument:** `rancher2_cluster.istio_enabled` - (Computed) Is istio enabled at cluster?

ENHANCEMENTS:

* Added `wait` argument to rancher2_app
* Added configurable retry logic when Rancher responds with "405 method not allowed" for `rancher2_node_template` resource
* Added drone pipeline definition to publish provider at terraform registry
* Updated docs to terraform registry format

BUG FIXES:

* Fixes on `rancher2_cluster_template` resource:
  * Update default revision. Related to https://github.com/rancher/terraform-provider-rancher2/issues/393
  * Import. Related to https://github.com/rancher/terraform-provider-rancher2/issues/386
  * Delete old template revisions. Related to https://github.com/rancher/terraform-provider-rancher2/issues/397
* Fixed import resource description on doc files
* Fixed bootstrap link on doc website

## 1.9.0 (June 29, 2020)

FEATURES:



ENHANCEMENTS:

* Updated acceptance tests:
  * run rancher HA environment on k3s v1.18.2-k3s1
  * integrated rancher update scenario from v2.3.6 to v2.4.5
* Updated local cluster on `rancher2_bootstrap` resource, due to issue https://github.com/rancher/rancher/issues/16213
* Added `load_balancer_sku` argument to `azure_cloud_provider` configuration
* Added `nodelocal` argument to `rke_config.dns` argument on `rancher2_cluster` resource
* Added `view` verb to `rules` argument for `rancher2_node_template` resource
* Updated golang to v1.13, modules and vendor files
* Updated Rancher support to v2.4.5
* Added full feature to `rke_config.monitoring` argument
* Added `external` as allowed value on `rke_config.cloud_provider` argument on `rancher2_cluster` resource
* Added `region` argument on `gke_config` for `rancher2_cluster` resource
* Updated `annotations` and `labels` arguments to supress diff when name contains `cattle.io/` or `rancher.io/`

BUG FIXES:

* Fixed `nodeTemplateStateRefreshFunc` function on `rancher2_node_template` resource to check if returned error is forbidden
* Updated `rancher2_app` resource to fix local cluster scoped catalogs
* Updated api bool fields with default=true to `*bool`. Related to https://github.com/rancher/types/pull/1083
* Fixed update on `rancher2_cluster_template` resource. Related to https://github.com/terraform-providers/terraform-provider-rancher2/issues/365

## 1.8.3 (April 09, 2020)

FEATURES:



ENHANCEMENTS:



BUG FIXES:

* Fix project alert group and alert rule datasource and resoruce documentation
* Added `version` argument to `cluster_monitoring_input` argument on `rancher2_cluster` and `rancher2_project` resources
* Fixed rancher timeout on bootstrapping

## 1.8.2 (April 02, 2020)

FEATURES:



ENHANCEMENTS:

* Added `fixNodeTemplateID` to fix `rancher2_node_template` ID upgrading up to v2.3.3. Issue [#195](https://github.com/terraform-providers/terraform-provider-rancher2/issues/195)
* Updated rnacher to v2.4.2 on acceptance test

BUG FIXES:

* Fix `upgrading` state on resourceRancher2ClusterUpdate() function
* Updated process for getting rke supported kubernetes version
* Set `version` argument on `rancher2_catalog` as computed

## 1.8.1 (March 31, 2020)

FEATURES:



ENHANCEMENTS:



BUG FIXES:

* Fix init provider if api_url is dependent of infra that is not yet deployed

## 1.8.0 (March 31, 2020)

FEATURES:

* **New Data Source:** `rancher2_cluster_scan`

ENHANCEMENTS:

* Added `wait_monitoring` argument to `rancher2_cluster_sync` resource
* Added `retries` config argument and `isRancherActive()` function
* Updated go modules and vendor files to rancher v2.4.0
* Updated rancher to v2.4.0 and k3s to v1.17.4-k3s1 on acceptance tests
* New rancher v2.4.0 features:
  * Added `group_principal_id` argument to `rancher2_global_role_binding` resource
  * Added `k3s_config` argument to `rancher2_cluster` datasource and resource
  * Added `version` argument to `rancher2_catalog` datasource and resource
  * Added `upgrade_strategy` argument to `rke_config` on `rancher2_cluster` resource
  * Added `scheduled_cluster_scan` argument on `rancher2_cluster` and `rancher2_cluster_template` resources
  * Added `rancher2_cluster_scan` datasource
* Added `fixNodeTemplateID` to fix `rancher2_node_template` ID upgrading up to v2.3.3. Issue [#195](https://github.com/terraform-providers/terraform-provider-rancher2/issues/195)

BUG FIXES:

* Added `enable_json_parsing` argument to cluster and project logging
* Sync resource delete with rancher API
* Fix recipient update on cluster and project alert groups

## 1.7.3 (March 24, 2020)

FEATURES:

* **New Data Source:** `rancher2_pod_security_policy_template`
* **New Resource:** `rancher2_pod_security_policy_template`

ENHANCEMENTS:

* Updated `rancher/norman` go modules and vendor files
* Added `plugin` optional value `none` to `rke_config` argument on `rancher2_cluster` resource
* Updated multiline arguments to trim spaces by default and avoid false diff
* Updated `rancher/types` go modules and vendor files
* Added `mtu` argument to `rke_config.network` argument on `rancher2_cluster` resource
* Added `custom_target_config` argument to `rancher2_cluster_logging` and `rancher2_project_logging` resources
* Updated `aks_config`, `eks_config` and `gke_config` arguments ti proper updte `rancher2_cluster` resource

BUG FIXES:

* Fix `audit_log.configuration.policy` argument to `rke_config.services.kube_api` argument on `rancher2_cluster` resource
* Added `plugin` optional value `none` to `rke_config` argument on `rancher2_cluster` resource
* Updated multiline arguments to trim spaces by default and avoid false diff
* Updated `private_key_file` definition for openstack driver on `rancher2_node_template` docs
* Updated `private_key_file` definition for openstack driver on `rancher2_node_template` docs
* Fixed `rke_config.cloud_provider.aws_cloud_provider.global` argument as computed to avoid false diff

## 1.7.2 (January 28, 2020)

FEATURES:



ENHANCEMENTS:

* Added `refresh` argument to `rancher2_catalog` resource
* Added `name` and `is_external` arguments to `rancher2_user` datasource
* Added `delete_not_ready_after_secs` and `node_taints` arguments to `node_pool` resource
* Added `delete_not_ready_after_secs` and `node_taints` arguments to `rancher2_node_pool` resource
* Updated `github.com/rancher/types` and `github.com/rancher/norman` go modules and vendor files to support rancher v2.3.3
* Splitted schema, structure and test `cluster_rke_config_services` files for every rke service
* Added `ssh_cert_path` argument to `rke_config` argument on `rancher2_cluster` resource
* Added `audit_log`, `event_rate_limit` and `secrets_encryption_config` arguments to `rke_config.services.kube_api` argument on `rancher2_cluster` resource
* Added `generate_serving_certificate` argument to `rke_config.services.kubelet` argument on `rancher2_cluster` resource
* Added `driver_id` argument on `rancher2_node_template` resource to reference user created `rancher2_node_driver`

BUG FIXES:

* Fix `template_revisions` update on `rancher2_cluster_template` resource
* Fix `rke_config.services.kube_api.policy` argument on `rancher2_cluster` resource
* Fix `data` argument set as sensitive on `rancher2_secret` resource

## 1.7.1 (December 04, 2019)

FEATURES:



ENHANCEMENTS:

* Added GetRancherVersion function to provider config
* Updated `vsphere_config` argument schema on `rancher2_node_template` resource to support Rancher v2.3.3 features
* Updated rancher to v2.3.3 and k3s to v0.10.2 on acceptance tests

BUG FIXES:

* Set `annotations` argument as computed on `rancher2_node_template` resource
* Added `rancher2_node_template` resource workaround on docs when upgrade Rancher to v2.3.3

## 1.7.0 (November 20, 2019)

FEATURES:

* **New Resource:** `rancher2_token`

ENHANCEMENTS:

* Added `always_pull_images` argument on `kube_api` argument on `rke_config` argument for `rancher2_clusters` resource
* Added resource deletion if not getting active state on creation for `rancher2_catalog` resource
* Updated rancher to v2.3.2 and k3s to v0.10.1 on acceptance tests
* Added `desired nodes` support on `eks_config` argument on `rancher2_cluster` resource
* Added `managed disk` support on `azure_config` argument on `rancher2_node_template` resource
* Migrated provider to use `terraform-plugin-sdk`
* Updated `rancher2_etcd_backup` documentation

BUG FIXES:

* Fix `password` argument update for `rancher2_catalog` resource
* Fix `rancher2_app` update issue on Rancher v2.3.2
* Fix: set `key` argument as sensitive on `rancher2_certificate` resource.
* Fix continuous diff issues on `rancher2_project` resource
* Fix `pod_security_policy_template_id` update on `rancher2_project` resource
* Fix continuous diff issues on `rancher2_namespace` resource

## 1.6.0 (October 08, 2019)

FEATURES:

* **New Data Source:** `rancher2_cluster_alert_group`
* **New Data Source:** `rancher2_cluster_alert_rule`
* **New Data Source:** `rancher2_cluster_template`
* **New Data Source:** `rancher2_notifier`
* **New Data Source:** `rancher2_project_alert_group`
* **New Data Source:** `rancher2_project_alert_rule`
* **New Data Source:** `rancher2_role_template`
* **New Resource:** `rancher2_auth_config_keycloak`
* **New Resource:** `rancher2_auth_config_okta`
* **New Resource:** `rancher2_cluster_alert_group`
* **New Resource:** `rancher2_cluster_alert_rule`
* **New Resource:** `rancher2_cluster_sync`
* **New Resource:** `rancher2_cluster_template`
* **New Resource:** `rancher2_notifier`
* **New Resource:** `rancher2_project_alert_group`
* **New Resource:** `rancher2_project_alert_rule`
* **New Resource:** `rancher2_role_template`

ENHANCEMENTS:

* Added `monitoring_input` argument to define monitoring config for `rancher2_cluster` and `rancher2_project`
* Improved capitalization/spelling/grammar/etc in docs

BUG FIXES:

* Fix `expandAppExternalID` function on `rancher2_app` resource. Function was generating a wrong `ExternalID` catalog URL, on `cluster` and `project` scope
* Fix `flattenMultiClusterApp` function on `rancher2_multi-cluster_app` resource. Function wasn't updating fine `catalog_name`, `template_name` and/or `template_version` arguments, when contains char `-`
* Fix: set `value_yaml` multiline argument as base64 encoded
* Fix: removed `restricted` and `unrestricted` values checking for `default_pod_security_policy_template_id` argument on `rancher2_cluster` resource

## 1.5.0 (September 06, 2019)

FEATURES:

* **New Data Source:** `rancher2_app`
* **New Data Source:** `rancher2_certificate`
* **New Data Source:** `rancher2_multi_cluster_app`
* **New Data Source:** `rancher2_node_template`
* **New Data Source:** `rancher2_secret`
* **New Resource:** `rancher2_certificate`
* **New Resource:** `rancher2_app`
* **New Resource:** `rancher2_multi_cluster_app`
* **New Resource:** `rancher2_secret`

ENHANCEMENTS:

* Updated default image to `canonical:UbuntuServer:18.04-LTS:latest` on Azure node template
* Added `folder` argument on `s3_backup_config`
* Updated `github.com/rancher/types` and `github.com/rancher/norman` go modules and vendor files to support rancher v2.2.8
* Updated rancher to v2.2.8 and k3s to v0.8.0 on acceptance tests
* Added `key_pair_name` argument on `eks_config` argument on `rancher2_cluster` resource
* Set `kubernetes_version` argument as required on `eks_config` argument on `rancher2_cluster` resource
* Set `quantity` argument as optional with default value `1` on `rancher2_node_pool` resource. Added validation that value >= 1

BUG FIXES:

* Fix: `container_resource_limit` argument update issue on `rancher2_namespace` and `rancher2_project` resources update
* Fix: `sidebar_current` definition on datasources docs
* Fix: set `access_key` and `secret_key` arguments as optional on `s3_backup_config`
* Fix: crash `rancher2_cluster`  datasource and resource if `enableNetworkPolicy` doesn't exist
* Fix: don't delete builtin cluster nor node drivers from rancher on tf destroy
* Fix: wrong updates on not changed sensitive arguments on `rancher2_cluster_logging` and `rancher2_project_logging` resources

## 1.4.1 (August 16, 2019)

FEATURES:

ENHANCEMENTS:

BUG FIXES:

* Fix: auth issue when using `access_key` and `secret_key`

## 1.4.0 (August 15, 2019)

FEATURES:

* **New Data Source:** `rancher2_catalog`
* **New Data Source:** `rancher2_cloud_credential`
* **New Data Source:** `rancher2_cluster`
* **New Data Source:** `rancher2_cluster_driver`
* **New Data Source:** `rancher2_cluster_logging`
* **New Data Source:** `rancher2_cluster_role_template_binding`
* **New Data Source:** `rancher2_etcd_backup`
* **New Data Source:** `rancher2_global_role_binding`
* **New Data Source:** `rancher2_namespace`
* **New Data Source:** `rancher2_node_driver`
* **New Data Source:** `rancher2_node_pool`
* **New Data Source:** `rancher2_project_logging`
* **New Data Source:** `rancher2_project_role_template_binding`
* **New Data Source:** `rancher2_registry`
* **New Data Source:** `rancher2_user`
* **New Resource:** `rancher2_global_role_binding`
* **New Resource:** `rancher2_registry`
* **New Resource:** `rancher2_user`

ENHANCEMENTS:

* Set `session_token` argument as sensitive on `eks_config` argument on `rancher2_cluster` resource
* Added `wait_for_cluster` argument on `rancher2_namespace` and `rancher2_project` resources
* Set default value to `engine_install_url` argument on `rancher2_node_template` resource
* Added `enable_cluster_monitoring` argument to `rancher2_cluster` resource and datasource
* Added `enable_project_monitoring` argument to `rancher2_project` resource and datasource
* Added `token` argument on `cluster_registration_token` argument to rancher2_cluster resource and datasource
* Set default value to `engine_install_url` argument on `rancher2_node_template` resource
* Added `custom_ca` argument on etcd `s3_backup_config` on `rancher2_cluster` and `rancher2_etcd_backup` resources
* Updated `github.com/rancher/types` and `github.com/rancher/norman` go modules and vendor files to support rancher v2.2.6
* Updated rancher to v2.2.6 and k3s to v0.7.0 on acceptance tests
* Added cluster and project scope support on `rancher2_catalog` resource and datasource
* Updated `provider` config validation to enable bootstrap and resource creation at same run
* Added `container_resource_limit` argument on `rancher2_namespace` and `rancher2_project` resources and datasources
* Added `pod_security_policy_template_id` on `rancher2_project` resource

BUG FIXES:

* Fix: `toArrayString` and `toMapString` functions to check `nil` values
* Fix: Set `kubernetes_version` argument as required on `aks_config` argument on `rancher2_cluster` resource
* Fix: Set `security_groups`, `service_role`, `subnets` and `virtual_network` arguments as optional to `eks_config` argument on `rancher2_cluster` resource
* Fix: Removed `docker_version` argument from `rancher2_node_template` resource

## 1.3.0 (June 26, 2019)

FEATURES:

ENHANCEMENTS:

* Added `scheduler` argument to `services`-`rke_config` argument on `rancher2_cluster` resource

BUG FIXES:

* Fix: index out of range issue on `vsphere_cloud_provider`-`cloud_provider`-`rke_config` argument on `rancher2_cluster` resource

## 1.2.0 (June 12, 2019)

FEATURES:

* **New Data Source:** `rancher2_project`

ENHANCEMENTS:

* Added `cluster_auth_endpoint` argument to `rancher2_cluster` resource
* Added `default_pod_security_policy_template_id` argument to `rancher2_cluster` resource
* Added `enable_network_policy` argument to `rancher2_cluster` resource
* Updated acceptance tests
  * k3s version updated to v0.5.0
  * Rancher version updated to v2.2.4

BUG FIXES:

* Fix: set default value to `true` on `ignore_docker_version`-`rke_config` argument on `rancher2_cluster` resource
* Fix: set default value to `false` on `pod_security_policy`-`services`-`rke_config` argument on `rancher2_cluster` resource
* Fix: typo on `boot2docker_url`-`vsphere_config` argument name on `rancher2_node_template` resource docs
* Fix: set `monitor_delay` and `monitor_timeout` fields as string type for openstack load_balancer config on `cloud_provider`-`rke_config` argument on `rancher2_cluster` resource
* Fix: Updated `rancher2_etcd_backup` resource to work on rancher v2.2.4

## 1.1.0 (May 29, 2019)

FEATURES:

ENHANCEMENTS:

* Added `default_project_id` & `system_project_id` attributes to `rancher2_cluster` resource
* Added support to move `rancher2_namespace` resource to a rancher project when import
* Added support to terraform 0.12

BUG FIXES:

* Fix: Updated `flattenNamespace` function on `rancher2_namespace` resource to avoid no empty plan if `resource_quota` is not specified
* Fix: Updated `rke_config` argument for openstack cloud_provider on `rancher2_cluster` resource:
  * Removed `used_id` field on global argument in favour of `username` following [k8s openstack cloud provider docs](https://github.com/kubernetes/cloud-provider-openstack/blob/master/docs/provider-configuration.md#global-required-parameters)
  * Set computed=true on optional field to avoid no empty plan if not specified

## 1.0.0 (May 14, 2019)

* Initial Terraform Ecosystem Release


## v0.2.0-rc5 (Unreleased)

FEATURES:

ENHANCEMENTS:

* Updated `rancher2_cluster` `rke_config` argument to support `aws_cloud_provider` config
* Updated k3s version to v0.4.0 to run acceptance tests
* Added support to openstack and vsphere drivers on `rancher2_cloud_credential` resource
* Added support to openstack and vsphere drivers on `rancher2_node_template` resource

BUG FIXES:

* Fix: Updated `rancher2_cluster` resource to save correctly S3 and cloud providers passwords on `rke_config`
* Updated `rancher2_cloud_credential` resource to save correctly S3 password
* Updated `rancher2_etcd_backup` resource to save correctly S3 password

## v0.2.0-rc4 (Unreleased)

FEATURES:

* **New Resource:** `rancher2_bootstrap`
* **New Resource:** `rancher2_cloud_credential`
* **New Resource:** `rancher2_cluster_driver`
* **New Resource:** `rancher2_etcd_backup`

ENHANCEMENTS:

* Added `.drone.yml` file to also support run rancher pipeline
* Added `rancher2_node_pool` resource tests
* Added `rancher2_auth_config_*` resource tests
* Updated and reviewed docs format
* Added support to rancher v2.2.x
* Updated `rancher2_cluster` `rke_config` argument to support:
  * etcd service `backup_config` with local and S3 storage backends
  * `dns` config
  * `weave` network provider
* Splitted resources into own schema, structure and import files.
* Added support to amazonec2, azure and digitalocean drivers on `rancher2_cloud_credential` resource
* Added support to local and S3 storage backends on `rancher2_etcd_backup` resource

BUG FIXES:

* Fix: drone build image to golang:1.12.3 to fix go fmt issues
* Fix: removed test on apply for `rancher2_auth_config_*` resources
* Fix: updated `api_url` field as required on provider.go
* Fix: updated `rancher2_namespace` move to a project after import it from k8s cluster

## v0.2.0-rc3 (Unreleased)

FEATURES:

ENHANCEMENTS:

* Added `Sensitive: true` option to fields with sensible data

BUG FIXES:

* Fix: set rke cluster `cloud_provider_vsphere` disk and network as optional and computed fields

## v0.2.0-rc2 (Unreleased)

FEATURES:

ENHANCEMENTS:

* Added `Sensitive: true` option to fields with sensible data
* Added `kube_config` computed field on cluster resources
* Added `ami` and `associate_worker_node_public_ip` fields for `eks_config` on cluster resources
* Added all available fields for rke_config on cluster resources
* Added `manifest_url` and `windows_node_command` fields for `cluster_registration_token` on cluster resources
* Added `creation` argument on `etcd` service for rke_config on cluster resources

BUG FIXES:

* Fix: added updating pending state on cluster resource update
* Fix: checking if `cluster_registration_token` exists on cluster resource creation
* Fix: typo on `gke_config` credential field on cluster resource
* Fix: Updated auth resources to avoid permission denied error

## 0.1.0-rc1 (Unreleased)

FEATURES:

* **New Resource:** `rancher2_auth_config_activedirectory`
* **New Resource:** `rancher2_auth_config_adfs`
* **New Resource:** `rancher2_auth_config_azuread`
* **New Resource:** `rancher2_auth_config_freeipa`
* **New Resource:** `rancher2_auth_config_github`
* **New Resource:** `rancher2_auth_config_openldap`
* **New Resource:** `rancher2_auth_config_ping`
* **New Resource:** `rancher2_catalog`
* **New Resource:** `rancher2_cluster`
* **New Resource:** `rancher2_cluster_logging`
* **New Resource:** `rancher2_cluster_role_template_binding`
* **New Resource:** `rancher2_namespace`
* **New Resource:** `rancher2_node_driver`
* **New Resource:** `rancher2_node_pool`
* **New Resource:** `rancher2_node_template`
* **New Resource:** `rancher2_project`
* **New Resource:** `rancher2_project_logging`
* **New Resource:** `rancher2_project_role_template_binding`
* **New Resource:** `rancher2_setting`

ENHANCEMENTS:

* First release candidate of the rancher2 provider.
* resource/rancher2_cluster: support for providers:
  * Amazon EKS
  * Azure AKS
  * Google GKE
  * Imported
  * RKE
    * Cloud providers adding node pools
    * Custom
* resource/rancher2_cluster_logging: support for providers:
  * Elasticsearch
  * Fluentd
  * Kafka
  * Splunk
  * Syslog
* resource/rancher2_namespace: quota limits support on Rancher v2.1.x or higher
  * Amazon EC2
  * Azure
  * Digitalocean
* resource/rancher2_project: quota limits support on Rancher v2.1.x or higher
* resource/rancher2_project_logging: support for providers:
  * Elasticsearch
  * Fluentd
  * Kafka
  * Splunk
  * Syslog
* resource/rancher2_node_template: support for providers:

BUG FIXES:
