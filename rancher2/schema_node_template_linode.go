package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	linodeConfigDriver = "linode"
)

//Types

type linodeConfig struct {
	AuthorizedUsers string `json:"authorizedUsers,omitempty" yaml:"authorizedUsers,omitempty"`
	CreatePrivateIP bool   `json:"createPrivateIp,omitempty" yaml:"createPrivateIp,omitempty"`
	DockerPort      string `json:"dockerPort,omitempty" yaml:"dockerPort,omitempty"`
	Image           string `json:"image,omitempty" yaml:"image,omitempty"`
	InstanceType    string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	Label           string `json:"label,omitempty" yaml:"label,omitempty"`
	Region          string `json:"region,omitempty" yaml:"region,omitempty"`
	RootPass        string `json:"rootPass,omitempty" yaml:"rootPass,omitempty"`
	SSHPort         string `json:"sshPort,omitempty" yaml:"sshPort,omitempty"`
	SSHUser         string `json:"sshUser,omitempty" yaml:"sshUser,omitempty"`
	StackScript     string `json:"stackscript,omitempty" yaml:"stackscript,omitempty"`
	StackscriptData string `json:"stackscriptData,omitempty" yaml:"stackscriptData,omitempty"`
	SwapSize        string `json:"swapSize,omitempty" yaml:"swapSize,omitempty"`
	Tags            string `json:"tags,omitempty" yaml:"tags,omitempty"`
	Token           string `json:"token,omitempty" yaml:"token,omitempty"`
	UAPrefix        string `json:"uaPrefix,omitempty" yaml:"uaPrefix,omitempty"`
}

//Schemas

func linodeConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"authorized_users": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   false,
			Description: "Linode user accounts (seperated by commas) whose Linode SSH keys will be permitted root access to the created node",
		},
		"create_private_ip": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Create private IP for the instance",
		},
		"docker_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "2376",
			Description: "Docker Port",
		},
		"image": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "linode/ubuntu18.04",
			Description: "Specifies the Linode Instance image which determines the OS distribution and base files",
		},
		"instance_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "g6-standard-4",
			Description: "Specifies the Linode Instance type which determines CPU, memory, disk size, etc.",
		},
		"label": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Linode Instance Label",
		},
		"region": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "us-east",
			Description: "Specifies the region (location) of the Linode instance",
		},
		"root_pass": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Root Password",
		},
		"ssh_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "22",
			Description: "Linode Instance SSH Port",
		},
		"ssh_user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the user as which docker-machine should log in to the Linode instance to install Docker.",
		},
		"stackscript": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the Linode StackScript to use to create the instance",
		},
		"stackscript_data": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A JSON string specifying data for the selected StackScript",
		},
		"swap_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "512",
			Description: "Linode Instance Swap Size (MB)",
		},
		"tags": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A comma separated list of tags to apply to the the Linode resource",
		},
		"token": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Linode API Token",
		},
		"ua_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Prefix the User-Agent in Linode API calls with some 'product/version'",
		},
	}

	return s
}
