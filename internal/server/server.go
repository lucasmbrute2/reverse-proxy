package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/lucasmbrute2/reverse-proxy-from-scratch/internal/configs"
)

func Run() error {
	//load configurations from config file
	config, err := configs.NewConfiguration()
	if err != nil {
		fmt.Errorf("could not load configuration: %v", err)
	}
	// Creates a new router
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)

	// Iterating through the configuration resource and registering
	// them into the router
	for _, resource := range config.Resources {
		url, _ := url.Parse(resource.Destination_url)
		proxy := NewProxy(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
	}
	if err := http.ListenAndServe(config.Server.Host+":"+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
