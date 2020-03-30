#  Packets
Captuing packets off the wire 


## intstall

```bash
sudo yum install libpcap-dev

# resp

sudo yum install libpcap-devel
```

## Testing

listener requires `sudo`

following command will install packets into `$GOBIN`
```bash
go install packets.go
```

to get list of available network interfaces run 
```bash
packets devices
```
and you will get something like this:

```text
Logger configured
DEVICE : wlp61s0 
   IP:   192.168.0.46
   Mask: ffffff00
   IP:   fe80::f13b:e8b6:591d:c94e
   Mask: ffffffffffffffff0000000000000000

DEVICE : lo 
   IP:   127.0.0.1
   Mask: ff000000
   IP:   ::1
   Mask: ffffffffffffffffffffffffffffffff

DEVICE : any Pseudo-device that captures on all interfaces

DEVICE : virbr0 
   IP:   192.168.122.1
   Mask: ffffff00

```

one of interface is your primary and listening on localhost. In my example it is interface called `lo`. Open three terminals and run commands

**Terminal1**
```bash
sudo $GOBIN/packets capture -i lo
```

**Terminal2**
```bash
nc -l -p 80
```

**Terminal3**
```bash
telnet 127.0.0.1 80
write something
```


you get something like this:
```text
Logger configured
PACKET: 74 bytes, wire length 74 cap length 74 @ 2020-03-30 20:02:35.533002 +0200 CEST
- Layer 1 (14 bytes) = Ethernet {Contents=[..14..] Payload=[..60..] SrcMAC=00:00:00:00:00:00 DstMAC=00:00:00:00:00:00 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..40..] Version=4 IHL=5 TOS=16 Length=60 Id=58955 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=22110 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (40 bytes) = TCP      {Contents=[..40..] Payload=[] SrcPort=56126 DstPort=80(http) Seq=553418654 Ack=0 DataOffset=10 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65495 Checksum=65072 Urgent=0 Options=[..5..] Padding=[]}

PACKET: 74 bytes, wire length 74 cap length 74 @ 2020-03-30 20:02:35.533064 +0200 CEST
- Layer 1 (14 bytes) = Ethernet {Contents=[..14..] Payload=[..60..] SrcMAC=00:00:00:00:00:00 DstMAC=00:00:00:00:00:00 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..40..] Version=4 IHL=5 TOS=0 Length=60 Id=0 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=15546 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (40 bytes) = TCP      {Contents=[..40..] Payload=[] SrcPort=80(http) DstPort=56126 Seq=1516930799 Ack=553418655 DataOffset=10 FIN=false SYN=true RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=65483 Checksum=65072 Urgent=0 Options=[..5..] Padding=[]}

PACKET: 66 bytes, wire length 66 cap length 66 @ 2020-03-30 20:02:35.533108 +0200 CEST
- Layer 1 (14 bytes) = Ethernet {Contents=[..14..] Payload=[..52..] SrcMAC=00:00:00:00:00:00 DstMAC=00:00:00:00:00:00 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4     {Contents=[..20..] Payload=[..32..] Version=4 IHL=5 TOS=16 Length=52 Id=58956 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=22117 SrcIP=127.0.0.1 DstIP=127.0.0.1 Options=[] Padding=[]}
- Layer 3 (32 bytes) = TCP      {Contents=[..32..] Payload=[] SrcPort=56126 DstPort=80(http) Seq=553418655 Ack=1516930800 DataOffset=8 FIN=false SYN=false RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=512 Checksum=65064 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:1745747525/1745747525 0x680dfe45680dfe45)] Padding=[]}

```