package mdns

import (
	"context"
	"strings"
	"time"

	"github.com/grandcat/zeroconf"
)

type ZeroconfService struct{}

func NewZeroconfService() *ZeroconfService {
	return &ZeroconfService{}
}

func (zc *ZeroconfService)Scan(serviceType string, timeout time.Duration, callback func(*ServiceEntry)) error{
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil{
		return err
	}

	entries := make(chan *zeroconf.ServiceEntry)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})

	go func(){
		for entry := range entries{
			addrs := []string{}
			for _, ip := range entry.AddrIPv4{
				addrs = append(addrs, ip.String())
			}
			for _, ip := range entry.AddrIPv6 {
				addrs = append(addrs, ip.String())
			}
			
			txt := map[string]string{}

			for _, t := range entry.Text{
				parts := strings.Split(t, "=")
				if len(parts) == 2{
					txt[parts[0]] = parts[1]
				}
			}

			callback(&ServiceEntry{
				Instance: entry.Instance,
				Service: entry.Service,
				Domain: entry.Domain,
				Addresses: addrs,
				Text: txt,
			})
		}
		close(done)
	}()
	
	err = resolver.Browse(ctx, serviceType, "local.", entries)
	if err != nil {
		return err
	}

	<-done
	return nil
}