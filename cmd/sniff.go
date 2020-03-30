package cmd

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/kuritka/l4packet/capture"
	"github.com/spf13/cobra"
)

var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "capture any text in unencrypted stream",
	Long: `capture any text in unencrypted stream`,

	Run: func(cmd *cobra.Command, args []string) {

		captureOptions := capture.Options{
			NetworkInterface: options.NetworkInterface,
			Filter:           fmt.Sprintf("tcp and dst port %v", options.Port),
		}
		err := capture.New(captureOptions, sniffing).Run()
		if err != nil {
			logger.Fatal().Err(err).Msg("error")
		}
	},
}


func sniffing(p gopacket.Packet) error {
	fmt.Println(p.String())
	return nil
}


func init(){
	sniffCmd.Flags().StringVarP(&options.NetworkInterface, "interface", "i", "", "network interface")
	sniffCmd.Flags().IntVarP(&options.Port, "port", "p", 80,  "port")
	sniffCmd.MarkFlagRequired("interface")
	sniffCmd.MarkFlagRequired("port")
	rootCmd.AddCommand(sniffCmd)
}