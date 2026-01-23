package main

import (
	"flag"
	"fmt"
	"os"

	dockerfile "github.com/talos-riscv/buildkit/frontend/dockerfile/builder"
	"github.com/talos-riscv/buildkit/frontend/gateway/grpcclient"
	"github.com/talos-riscv/buildkit/util/appcontext"
	"github.com/talos-riscv/buildkit/util/bklog"
	_ "github.com/talos-riscv/buildkit/util/grpcutil/encoding/proto"
	"github.com/talos-riscv/buildkit/util/stack"
)

func init() {
	stack.SetVersionInfo(Version, Revision)
}

func main() {
	var version bool
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()

	if version {
		fmt.Printf("%s %s %s %s\n", os.Args[0], Package, Version, Revision)
		os.Exit(0)
	}

	if err := grpcclient.RunFromEnvironment(appcontext.Context(), dockerfile.Build); err != nil {
		bklog.L.Errorf("fatal error: %+v", err)
		panic(err)
	}
}
