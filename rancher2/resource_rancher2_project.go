package rancher2

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

func resourceRancher2Project() *schema.Resource {
	return &schema.Resource{
		Create: resourceRancher2ProjectCreate,
		Read:   resourceRancher2ProjectRead,
		Update: resourceRancher2ProjectUpdate,
		Delete: resourceRancher2ProjectDelete,
		Importer: &schema.ResourceImporter{
			State: resourceRancher2ProjectImport,
		},

		Schema: projectFields(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceRancher2ProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	project := expandProject(d)

	active, _, err := meta.(*Config).isClusterActive(project.ClusterID)
	if err != nil {
		if IsNotFound(err) || IsForbidden(err) {
			return fmt.Errorf("[ERROR] Creating Project: Cluster ID %s not found or is forbidden", project.ClusterID)
		}
		return err
	}
	if !active {
		if v, ok := d.Get("wait_for_cluster").(bool); ok && !v {
			return fmt.Errorf("[ERROR] Creating Project: Cluster ID %s is not active", project.ClusterID)
		}

		stateCluster := &resource.StateChangeConf{
			Pending:    []string{},
			Target:     []string{"active"},
			Refresh:    clusterStateRefreshFunc(client, project.ClusterID),
			Timeout:    d.Timeout(schema.TimeoutCreate),
			Delay:      1 * time.Second,
			MinTimeout: 3 * time.Second,
		}
		_, waitClusterErr := stateCluster.WaitForState()
		if waitClusterErr != nil {
			return fmt.Errorf("[ERROR] waiting for cluster ID (%s) to be active: %s", project.ClusterID, waitClusterErr)
		}
	}

	log.Printf("[INFO] Creating Project %s on Cluster ID %s", project.Name, project.ClusterID)

	// Creating cluster with monitoring disabled
	project.EnableProjectMonitoring = false
	newProject, err := client.Project.Create(project)
	if err != nil {
		return err
	}
	newProject.EnableProjectMonitoring = d.Get("enable_project_monitoring").(bool)
	d.SetId(newProject.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"initializing", "configuring", "active"},
		Target:     []string{"active"},
		Refresh:    projectStateRefreshFunc(client, newProject.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}
	_, waitErr := stateConf.WaitForState()
	if waitErr != nil {
		return fmt.Errorf(
			"[ERROR] waiting for project (%s) to be created: %s", newProject.ID, waitErr)
	}

	monitoringInput := expandMonitoringInput(d.Get("project_monitoring_input").([]interface{}))
	if newProject.EnableProjectMonitoring {
		if len(newProject.Actions[monitoringActionEnable]) == 0 {
			newProject, err = client.Project.ByID(newProject.ID)
			if err != nil {
				return err
			}
		}
		err = client.Project.ActionEnableMonitoring(newProject, monitoringInput)
		if err != nil {
			return err
		}
	}

	if pspID, ok := d.Get("pod_security_policy_template_id").(string); ok && len(pspID) > 0 {
		pspInput := &managementClient.SetPodSecurityPolicyTemplateInput{
			PodSecurityPolicyTemplateName: pspID,
		}
		_, err = client.Project.ActionSetpodsecuritypolicytemplate(newProject, pspInput)
		if err != nil {
			// Checking error due to ActionSetpodsecuritypolicytemplate() issue
			if error.Error(err) != "unexpected end of JSON input" {
				return err
			}
		}
	}

	return resourceRancher2ProjectRead(d, meta)
}

func resourceRancher2ProjectRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Refreshing Project ID %s", d.Id())
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	project, err := client.Project.ByID(d.Id())
	if err != nil {
		if IsNotFound(err) || IsForbidden(err) {
			log.Printf("[INFO] Project ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	var monitoringInput *managementClient.MonitoringInput
	if len(project.Annotations[monitoringInputAnnotation]) > 0 {
		monitoringInput = &managementClient.MonitoringInput{}
		err = jsonToInterface(project.Annotations[monitoringInputAnnotation], monitoringInput)
		if err != nil {
			return err
		}

	}

	return flattenProject(d, project, monitoringInput)
}

func resourceRancher2ProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating Project ID %s", d.Id())
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	project, err := client.Project.ByID(d.Id())
	if err != nil {
		return err
	}

	resourceQuota, nsResourceQuota := expandProjectResourceQuota(d.Get("resource_quota").([]interface{}))
	update := map[string]interface{}{
		"name":                          d.Get("name").(string),
		"description":                   d.Get("description").(string),
		"containerDefaultResourceLimit": expandProjectContainerResourceLimit(d.Get("container_resource_limit").([]interface{})),
		"namespaceDefaultResourceQuota": nsResourceQuota,
		"resourceQuota":                 resourceQuota,
		"annotations":                   toMapString(d.Get("annotations").(map[string]interface{})),
		"labels":                        toMapString(d.Get("labels").(map[string]interface{})),
	}

	newProject, err := client.Project.Update(project, update)
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"active"},
		Target:     []string{"active"},
		Refresh:    projectStateRefreshFunc(client, newProject.ID),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}
	_, waitErr := stateConf.WaitForState()
	if waitErr != nil {
		return fmt.Errorf(
			"[ERROR] waiting for project (%s) to be updated: %s", newProject.ID, waitErr)
	}

	if d.HasChange("pod_security_policy_template_id") {
		pspInput := &managementClient.SetPodSecurityPolicyTemplateInput{
			PodSecurityPolicyTemplateName: d.Get("pod_security_policy_template_id").(string),
		}
		_, err = client.Project.ActionSetpodsecuritypolicytemplate(newProject, pspInput)
		if err != nil {
			// Checking error due to ActionSetpodsecuritypolicytemplate() issue
			if error.Error(err) != "unexpected end of JSON input" {
				return err
			}
		}
	}

	if d.HasChange("enable_project_monitoring") || d.HasChange("project_monitoring_input") {
		enableMonitoring := d.Get("enable_project_monitoring").(bool)
		if !enableMonitoring && len(newProject.Actions[monitoringActionDisable]) > 0 {
			err = client.Project.ActionDisableMonitoring(newProject)
			if err != nil {
				return err
			}
		}
		if enableMonitoring {
			monitoringInput := expandMonitoringInput(d.Get("project_monitoring_input").([]interface{}))
			if len(newProject.Actions[monitoringActionEnable]) > 0 {
				err = client.Project.ActionEnableMonitoring(newProject, monitoringInput)
				if err != nil {
					return err
				}
			} else {
				monitorVersionChanged := false
				if d.HasChange("project_monitoring_input") {
					old, new := d.GetChange("project_monitoring_input")
					oldInput := old.([]interface{})
					oldInputLen := len(oldInput)
					newInput := new.([]interface{})
					newInputLen := len(newInput)
					monitorVersionChanged = (oldInputLen != newInputLen)
					if newInputLen > 0 && !monitorVersionChanged {
						oldRow, oldOK := oldInput[0].(map[string]interface{})
						newRow, newOK := newInput[0].(map[string]interface{})
						if oldOK && newOK {
							if oldRow["version"] != newRow["version"] {
								monitorVersionChanged = true
							}
						}
					}
				}
				if monitorVersionChanged {
					err = client.Project.ActionDisableMonitoring(newProject)
					if err != nil {
						return err
					}
					time.Sleep(5 * time.Second)
					newProject, err = client.Project.ByID(newProject.ID)
					if err != nil {
						return err
					}
					err = client.Project.ActionEnableMonitoring(newProject, monitoringInput)
					if err != nil {
						return err
					}
				} else {
					err = client.Project.ActionEditMonitoring(newProject, monitoringInput)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return resourceRancher2ProjectRead(d, meta)
}

func resourceRancher2ProjectDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting Project ID %s", d.Id())
	id := d.Id()
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	project, err := client.Project.ByID(id)
	if err != nil {
		if IsNotFound(err) || IsForbidden(err) {
			log.Printf("[INFO] Project ID %s not found.", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = client.Project.Delete(project)
	if err != nil {
		return fmt.Errorf("Error removing Project: %s", err)
	}

	log.Printf("[DEBUG] Waiting for project (%s) to be removed", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"removing"},
		Target:     []string{"removed"},
		Refresh:    projectStateRefreshFunc(client, id),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      1 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, waitErr := stateConf.WaitForState()
	if waitErr != nil {
		return fmt.Errorf(
			"[ERROR] waiting for project (%s) to be removed: %s", id, waitErr)
	}

	d.SetId("")
	return nil
}

// projectStateRefreshFunc returns a resource.StateRefreshFunc, used to watch a Rancher Project.
func projectStateRefreshFunc(client *managementClient.Client, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		obj, err := client.Project.ByID(projectID)
		if err != nil {
			if IsNotFound(err) || IsForbidden(err) {
				return obj, "removed", nil
			}
			return nil, "", err
		}

		return obj, obj.State, nil
	}
}
