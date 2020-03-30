package cmd

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/kuritka/l4packet/capture"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "live capturing",
	Long: `live capturing`,

	Run: func(cmd *cobra.Command, args []string) {

		captureOptions := capture.Options{
			NetworkInterface: options.NetworkInterface,
			Filter:           "tcp and port 80",
		}
		err := capture.New(captureOptions, tcpDump).Run()
		if err != nil {
			logger.Fatal().Err(err).Msg("error")
		}
	},
}

func tcpDump(p gopacket.Packet) error {
	fmt.Println(p.String())
	return nil
}

func init(){
	listenCmd.Flags().StringVarP(&options.NetworkInterface, "interface", "i", "", "network interface")
	listenCmd.MarkFlagRequired("interface")
	rootCmd.AddCommand(listenCmd)
}
