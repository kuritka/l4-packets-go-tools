package cmd

import (
	"github.com/spf13/cobra"
	"os"

	"github.com/kuritka/gext/log"
)

var Verbose bool
var logger = log.Logger()

var rootCmd = &cobra.Command{
	Short: "l4packet",
	Long: `l4packet - Captuing packets off the wire`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Error().Msg("No parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msg("done..")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
