package main

import (
	"github.com/mitchellh/packer/packer/plugin"
  "github.com/mcansky/packer-builder-release/builder/amazon/release"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(release.Builder))
	server.Serve()
}
