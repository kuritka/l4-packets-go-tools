package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (

	filter = "tcp and port 80"
	devfound = false
)


type capture struct {
	options CaptureOptions
	snaplen int32
}


func New(options CaptureOptions) *capture{
	 capture := new(capture)
	 capture.options = options
	 capture.snaplen = 1600
	 return capture
}

func (c *capture) Run() error {

	ex := deviceExists(c.options.NetworkInterface)
	if !ex {
		return fmt.Errorf("device %s not found", c.options.NetworkInterface)
	}

	handle, err := pcap.OpenLive(c.options.NetworkInterface, c.snaplen, false, pcap.BlockForever )
	if err != nil {
		return err
	}
	defer handle.Close()

	if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		return err
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		fmt.Println(packet.String())
	}
	return nil
}

func deviceExists(iface string) bool {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return false
	}

	for _, device := range devices {
		if device.Name == iface {
			return true
		}
	}

	return false
}