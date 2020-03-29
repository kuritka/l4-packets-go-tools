package devices

import (
	"fmt"
	"github.com/google/gopacket/pcap"
)


func Run() error{
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	for _, device := range devices {
		fmt.Println("DEVICE : "+device.Name + " " + device.Description)
		for _,addr := range device.Addresses {
			fmt.Printf("   IP:   %s\n", addr.IP)
			fmt.Printf("   Mask: %s\n", addr.Netmask)
		}
		fmt.Println()
	}
	return nil
}
