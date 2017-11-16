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
	fmt.Println("-------------------------")
	fmt.Println(uri.Host)
	fmt.Println(uri.Path)
	fmt.Println("-------------------------")

	return &EurekaAdapter{host: uri.Host, path: uri.Path}
}

type EurekaAdapter struct {
	host string
	path string
}

func (r *EurekaAdapter) Ping() error {

	return nil
}

func (r *EurekaAdapter) Register(service *bridge.Service) error {

	return nil
}

func (r *EurekaAdapter) Deregister(service *bridge.Service) error {

	return nil
}

func (r *EurekaAdapter) Refresh(service *bridge.Service) error {

	return nil
}

func (r *EurekaAdapter) Services() ([]*bridge.Service, error) {
	return nil, nil
}
