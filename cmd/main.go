package main

import (
	"flag"
	"fmt"
	"github.com/youthke/questack-api/conf"
	"github.com/youthke/questack-api/pkg/server"
)

var (
	state = flag.String("s", "local", "local/prd")
)

func main() {

	flag.Parse()
	if err := conf.Setup(fmt.Sprintf("conf/env/%s.toml", *state)); err != nil {
		fmt.Sprintf("%s", err)
		return
	}

	conf.Init()
	server.Init()
}