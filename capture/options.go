package capture

type Options struct {
	NetworkInterface string
	Filter string // "tcp and port 80" "tcp and dst port 21"
}
