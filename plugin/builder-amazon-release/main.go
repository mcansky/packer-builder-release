package main

import (
	"github.com/mcansky/packer-builder-release/builder/amazon/release"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(release.Builder))
	server.Serve()
}
