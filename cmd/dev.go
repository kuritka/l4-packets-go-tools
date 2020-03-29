package cmd

import (
	"github.com/kuritka/gext/guard"
	"github.com/kuritka/l4packet/devices"
	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "network interfaces",
	Long: `network interfaces`,

	Run: func(cmd *cobra.Command, args []string) {

		err := devices.Run()

		guard.FailOnError(err,"list of network interfaces")
	},
}

func init(){
	rootCmd.AddCommand(devicesCmd)
}