package main

import (
	// _ "github.com/hashicorp/terraform"
	// _ "github.com/terraform-providers/terraform-provider-aws"
	"fmt"
	"os"

	tfPlugin "github.com/hashicorp/terraform/plugin"
	tfPluginDiscovery "github.com/hashicorp/terraform/plugin/discovery"
	goplugin "github.com/hashicorp/go-plugin"
)

func main() {
	var (
		cli *goplugin.Client
		metadata tfPluginDiscovery.PluginMeta
	)

	metadata = tfPluginDiscovery.PluginMeta{
		Name: "terraform-provider-aws",
		Version: "2.29.0",
		Path: "./binaries",
	}

	cli = tfPlugin.Client(metadata)
	fmt.Printf("%+v\n", cli)

	defer cli.Kill()

	// Connect via RPC
	rpcClient, err := cli.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%+v\n", rpcClient)

	// ProviderPluginName    = "provider"
	raw, err := rpcClient.Dispense(tfPlugin.ProviderPluginName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// store the client so that the plugin can kill the child process
	p := raw.(*tfPlugin.GRPCProvider)
	p.PluginClient = cli

	fmt.Printf("%+v\n", p)
}


