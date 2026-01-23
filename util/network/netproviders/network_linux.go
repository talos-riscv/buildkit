//go:build linux

package netproviders

import (
	"github.com/talos-riscv/buildkit/util/network"
	"github.com/talos-riscv/buildkit/util/network/cniprovider"
)

func getBridgeProvider(opt cniprovider.Opt) (network.Provider, error) {
	return cniprovider.NewBridge(opt)
}
