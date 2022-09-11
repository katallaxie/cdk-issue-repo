package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func AppStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	kubernetes.NewKubernetesProvider(
		stack,
		jsii.String("demo"),
		&kubernetes.KubernetesProviderConfig{},
	)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	AppStack(app, "cuddly-garbanzo")

	app.Synth()
}
