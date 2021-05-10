package rancher2

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRancher2ClusterSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceRancher2ClusterSyncCreate,
		Read:   resourceRancher2ClusterSyncRead,
		Update: resourceRancher2ClusterSyncUpdate,
		Delete: resourceRancher2ClusterSyncDelete,

		Schema: clusterSyncFields(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
	}
}

func resourceRancher2ClusterSyncCreate(d *schema.ResourceData, meta interface{}) error {
	clusterID := d.Get("cluster_id").(string)

	cluster, err := meta.(*Config).WaitForClusterState(clusterID, clusterActiveCondition, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	if cluster.EnableClusterMonitoring && d.Get("wait_monitoring").(bool) {
		_, err := meta.(*Config).WaitForClusterState(clusterID, clusterMonitoringEnabledCondition, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("[ERROR] waiting for cluster ID (%s) monitoring to be running: %v", clusterID, err)
		}
	}

	if cluster.EnableClusterAlerting && d.Get("wait_alerting").(bool) {
		_, err := meta.(*Config).WaitForClusterState(clusterID, clusterAlertingEnabledCondition, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("[ERROR] waiting for cluster ID (%s) alerting to be running: %v", clusterID, err)
		}
	}

	if d.Get("wait_catalogs").(bool) {
		_, err := meta.(*Config).WaitAllCatalogV2Downloaded(clusterID)
		if err != nil {
			return fmt.Errorf("[ERROR] waiting for cluster ID (%s) downloading catalogs: %v", clusterID, err)
		}
	}

	d.SetId(clusterID)

	return resourceRancher2ClusterSyncRead(d, meta)
}

func resourceRancher2ClusterSyncRead(d *schema.ResourceData, meta interface{}) error {
	clusterID := d.Get("cluster_id").(string)

	active, clus, err := meta.(*Config).isClusterActive(clusterID)
	if err != nil {
		if IsNotFound(err) || IsForbidden(err) {
			log.Printf("[INFO] Cluster ID %s not found.", clusterID)
			d.SetId("")
			return nil
		}
		return err
	}

	if active {
		defaultProjectID, systemProjectID, err := meta.(*Config).GetClusterSpecialProjectsID(clusterID)
		if err != nil {
			return err
		}
		d.Set("default_project_id", defaultProjectID)
		d.Set("system_project_id", systemProjectID)

		client, err := meta.(*Config).ManagementClient()
		if err != nil {
			return err
		}
		kubeConfig, err := client.Cluster.ActionGenerateKubeconfig(clus)
		if err != nil {
			return err
		}
		d.Set("kube_config", kubeConfig.Config)
		nodes, err := meta.(*Config).GetClusterNodes(clusterID)
		if err != nil {
			return err
		}
		d.Set("nodes", flattenClusterNodes(nodes))

		if clus.EnableClusterMonitoring && d.Get("wait_monitoring").(bool) {
			monitor, _, err := meta.(*Config).isClusterMonitoringEnabledCondition(clusterID)
			if err != nil {
				return err
			}
			if !monitor {
				d.Set("synced", false)
				return nil
			}
		}

		if clus.EnableClusterAlerting && d.Get("wait_alerting").(bool) {
			alert, _, err := meta.(*Config).isClusterAlertingEnabledCondition(clusterID)
			if err != nil {
				return err
			}
			if !alert {
				d.Set("synced", false)
				return nil
			}
		}

		if d.Get("wait_catalogs").(bool) {
			_, err := meta.(*Config).WaitAllCatalogV2Downloaded(clusterID)
			if err != nil {
				return err
			}
		}
	}

	d.Set("synced", active)
	return nil
}

func resourceRancher2ClusterSyncUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceRancher2ClusterSyncCreate(d, meta)
}

func resourceRancher2ClusterSyncDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
