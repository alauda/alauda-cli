package namespace

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newSetCmd(alauda client.APIClient) *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "set NAME",
		Short: "Set the current cluster namespace",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("namespace set expects NAME")
			}
			return doSet(alauda, args[0])
		},
	}
	return getCmd
}

func doSet(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Setting the current cluster namespace to", name)

	viper.Set(util.SettingNamespace, name)

	err := util.SaveConfig()

	return err
}
