package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)


type Handler =  func(gopacket.Packet) error

type capture struct {
	options Options
	snaplen int32
	filter  string
	handler Handler
}

func New(options Options, f Handler) *capture{
	 capture := new(capture)
	 capture.options = options
	 capture.snaplen = 1600
	 capture.filter = "tcp and port 80"
	 capture.handler = f
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

	if err := handle.SetBPFFilter(c.options.Filter); err != nil {
		return err
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		if err := c.handler(packet); err != nil {
			return err
		}
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