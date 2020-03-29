package cmd

import (
	"github.com/kuritka/l4packet/capture"
	"github.com/spf13/cobra"
)

var (
	captureOptions capture.CaptureOptions
)


var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "live capturing",
	Long: `live capturing`,

	Run: func(cmd *cobra.Command, args []string) {

		err := capture.New(captureOptions).Run()
		if err != nil {
			logger.Fatal().Err(err).Msg("error")
		}
	},
}

func init(){
	captureCmd.Flags().StringVarP(&captureOptions.NetworkInterface, "interface", "i", "", "network interface")
	rootCmd.AddCommand(captureCmd)
}
