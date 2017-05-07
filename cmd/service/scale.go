package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewScaleCmd creates a new scale service command.
func NewScaleCmd(alauda client.APIClient) *cobra.Command {
	scaleCmd := &cobra.Command{
		Use:   "scale NAME=NUMBER",
		Short: "Scale a service to the specified number of instances",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service start expects NAME=NUMBER")
			}
			return doScale(alauda, args[0])
		},
	}

	return scaleCmd
}

func doScale(alauda client.APIClient, desc string) error {
	name, number, err := parseScale(desc)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] Starting", name)

	util.InitializeClient(alauda)

	data := client.ScaleServiceData{
		TargetInstances: number,
	}

	err = alauda.ScaleService(name, &data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
