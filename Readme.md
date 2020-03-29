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