package main

import (
	"github.com/bazzaware/packer-post-processor-tool-ovf"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(ovftool.OVFPostProcessor))
	server.Serve()
}
