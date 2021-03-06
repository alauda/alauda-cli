package lb

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster string
	service string
}

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Short: "List load balancers",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	lsCmd.Flags().StringVarP(&opts.service, "service", "s", "", "Service")

	return lsCmd
}

// NewLbsCmd creates a new alauda lbs command, which is a shortcut to the lb ls command.
func NewLbsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "lbs"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing load balancers")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	params := client.ListLoadBalancersParams{
		Cluster: cluster,
		Service: opts.service,
	}

	result, err := alauda.ListLoadBalancers(&params)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListLoadBalancersResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "TYPE", "ADDRESS", "ADDRESS TYPE", "CREATED WITH", "CREATED AT"}
}

func buildLsTableContent(result *client.ListLoadBalancersResult) [][]string {
	var content [][]string

	for _, lb := range result.LoadBalancers {
		content = append(content, []string{lb.Name, lb.Type, lb.Address, lb.AddressType, lb.CreatedWith, lb.CreatedAt})
	}

	return content
}
