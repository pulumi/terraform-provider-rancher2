package rancher2

import (
	"reflect"
	"testing"

	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

var (
	testClusterEKSConfigV2NodeGroupLaunchTemplateConf      *managementClient.LaunchTemplate
	testClusterEKSConfigV2NodeGroupLaunchTemplateInterface []interface{}
	testClusterEKSConfigV2NodeGroupConf                    []managementClient.NodeGroup
	testClusterEKSConfigV2NodeGroupInterface               []interface{}
	testClusterEKSConfigV2Conf                             *managementClient.EKSClusterConfigSpec
	testClusterEKSConfigV2Interface                        []interface{}
)

func init() {
	LtVersion := int64(1)
	testClusterEKSConfigV2NodeGroupLaunchTemplateConf = &managementClient.LaunchTemplate{
		Name:    "launch_template",
		Version: &LtVersion,
	}
	testClusterEKSConfigV2NodeGroupLaunchTemplateInterface = []interface{}{
		map[string]interface{}{
			"name":    "launch_template",
			"version": 1,
		},
	}
	size := int64(3)
	testClusterEKSConfigV2NodeGroupConf = []managementClient.NodeGroup{
		{
			NodegroupName: "name",
			InstanceType:  "instance_type",
			DesiredSize:   &size,
			DiskSize:      &size,
			Ec2SshKey:     "ec2_ssh_key",
			Gpu:           newTrue(),
			ImageID:       "image_id",
			Labels: map[string]string{
				"label1": "one",
				"label2": "two",
			},
			LaunchTemplate:       testClusterEKSConfigV2NodeGroupLaunchTemplateConf,
			MaxSize:              &size,
			MinSize:              &size,
			RequestSpotInstances: newTrue(),
			ResourceTags: map[string]string{
				"rstag1": "one",
				"rstag2": "two",
			},
			SpotInstanceTypes: []string{"spot1", "spot2"},
			Subnets:           []string{"net1", "net2"},
			Tags: map[string]string{
				"tag1": "one",
				"tag2": "two",
			},
			UserData: "user_data",
			Version:  "kubernetes_version",
		},
	}
	testClusterEKSConfigV2NodeGroupInterface = []interface{}{
		map[string]interface{}{
			"name":          "name",
			"instance_type": "instance_type",
			"desired_size":  3,
			"disk_size":     3,
			"ec2_ssh_key":   "ec2_ssh_key",
			"gpu":           true,
			"image_id":      "image_id",
			"labels": map[string]interface{}{
				"label1": "one",
				"label2": "two",
			},
			"launch_template":        testClusterEKSConfigV2NodeGroupLaunchTemplateInterface,
			"max_size":               3,
			"min_size":               3,
			"request_spot_instances": true,
			"resource_tags": map[string]interface{}{
				"rstag1": "one",
				"rstag2": "two",
			},
			"spot_instance_types": []interface{}{"spot1", "spot2"},
			"subnets":             []interface{}{"net1", "net2"},
			"tags": map[string]interface{}{
				"tag1": "one",
				"tag2": "two",
			},
			"user_data": "user_data",
			"version":   "kubernetes_version",
		},
	}
	testClusterEKSConfigV2Conf = &managementClient.EKSClusterConfigSpec{
		AmazonCredentialSecret: "test",
		DisplayName:            "eksimport",
		LoggingTypes:           []string{"type1", "type2"},
		NodeGroups:             testClusterEKSConfigV2NodeGroupConf,
		Region:                 "test",
		KmsKey:                 "kms_key",
		Imported:               false,
		KubernetesVersion:      "kubernetes_version",
		PrivateAccess:          newTrue(),
		PublicAccess:           newTrue(),
		PublicAccessSources:    []string{"access1", "access2"},
		SecretsEncryption:      newTrue(),
		SecurityGroups:         []string{"sec1", "sec2"},
		ServiceRole:            "service_role",
		Subnets:                []string{"net1", "net2"},
		Tags: map[string]string{
			"value1": "one",
			"value2": "two",
		},
	}
	testClusterEKSConfigV2Interface = []interface{}{
		map[string]interface{}{
			"cloud_credential_id":   "test",
			"kubernetes_version":    "kubernetes_version",
			"imported":              false,
			"kms_key":               "kms_key",
			"logging_types":         []interface{}{"type1", "type2"},
			"node_groups":           testClusterEKSConfigV2NodeGroupInterface,
			"name":                  "eksimport",
			"private_access":        true,
			"public_access":         true,
			"public_access_sources": []interface{}{"access1", "access2"},
			"region":                "test",
			"secrets_encryption":    true,
			"security_groups":       []interface{}{"sec1", "sec2"},
			"service_role":          "service_role",
			"subnets":               []interface{}{"net1", "net2"},
			"tags": map[string]interface{}{
				"value1": "one",
				"value2": "two",
			},
		},
	}
}

func TestFlattenClusterEKSConfigV2NodeGroups(t *testing.T) {

	cases := []struct {
		Input          []managementClient.NodeGroup
		ExpectedOutput []interface{}
	}{
		{
			testClusterEKSConfigV2NodeGroupConf,
			testClusterEKSConfigV2NodeGroupInterface,
		},
	}

	for _, tc := range cases {
		output := flattenClusterEKSConfigV2NodeGroups(tc.Input, tc.ExpectedOutput)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterEKSConfigV2(t *testing.T) {

	cases := []struct {
		Input          *managementClient.EKSClusterConfigSpec
		ExpectedOutput []interface{}
	}{
		{
			testClusterEKSConfigV2Conf,
			testClusterEKSConfigV2Interface,
		},
	}

	for _, tc := range cases {
		output := flattenClusterEKSConfigV2(tc.Input, tc.ExpectedOutput)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterEKSConfigV2NodeGroups(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput []managementClient.NodeGroup
	}{
		{
			testClusterEKSConfigV2NodeGroupInterface,
			testClusterEKSConfigV2NodeGroupConf,
		},
	}

	for _, tc := range cases {
		output := expandClusterEKSConfigV2NodeGroups(tc.Input, []string{"net1", "net2"}, "kubernetes_version")
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterEKSConfigV2(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.EKSClusterConfigSpec
	}{
		{
			testClusterEKSConfigV2Interface,
			testClusterEKSConfigV2Conf,
		},
	}

	for _, tc := range cases {
		output := expandClusterEKSConfigV2(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
