package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/stretchr/testify/assert"
)

// The tests below are example tests, you can find more information at
// https://cdk.tf/testing

var (
	stack = AppStack(cdktf.Testing_App(nil), "stack")
	synth = cdktf.Testing_Synth(stack, jsii.Bool(true))
)

func TestShouldContainK8sCluster(t *testing.T) {
	assertion := cdktf.Testing_ToHaveResource(synth, kubernetes.KubernetesProvider_TfResourceType())
	assert.True(t, *assertion)
}
