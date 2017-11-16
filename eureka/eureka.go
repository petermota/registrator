package eureka

import (
	"fmt"
	"net/url"

	"github.com/gliderlabs/registrator/bridge"
)

func init() {
	bridge.Register(new(Factory), "eureka")
}

type Factory struct{}

func (f *Factory) New(uri *url.URL) bridge.RegistryAdapter {
	return &EurekaAdapter{host: uri.Host, path: uri.Path}
}

type EurekaAdapter struct {
	host string
	path string
}

func (r *EurekaAdapter) Ping() error {
	fmt.Println("------------------------- Ping")
	fmt.Println(r.host)
	fmt.Println(r.path)
	fmt.Println("-------------------------")

	return r.getAllApps()
}

func (r *EurekaAdapter) Register(svc *bridge.Service) error {
	fmt.Println("------------------------- Register")
	fmt.Println("Name:", svc.Name)
	fmt.Println("ID:", svc.ID)
	fmt.Println("IP:", svc.IP)
	fmt.Println("Port:", svc.Port)
	fmt.Println("-------------------------")

	return r.registerApp(svc)
}

func (r *EurekaAdapter) Deregister(svc *bridge.Service) error {
	fmt.Println("------------------------- Unregister")
	fmt.Println("Name:", svc.Name)
	fmt.Println("ID:", svc.ID)
	fmt.Println("IP:", svc.IP)
	fmt.Println("Port:", svc.Port)
	fmt.Println("-------------------------")

	return r.deregisterApp(svc)
}

func (r *EurekaAdapter) Refresh(svc *bridge.Service) error {
	fmt.Println("------------------------- Refresh")
	fmt.Println("Name:", svc.Name)
	fmt.Println("ID:", svc.ID)
	fmt.Println("IP:", svc.IP)
	fmt.Println("Port:", svc.Port)
	fmt.Println("-------------------------")

	return nil
}

func (r *EurekaAdapter) Services() ([]*bridge.Service, error) {
	return nil, nil
}
