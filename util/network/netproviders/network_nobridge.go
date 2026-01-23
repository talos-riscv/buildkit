//go:build !linux

package netproviders

import (
	"runtime"

	"github.com/talos-riscv/buildkit/util/network"
	"github.com/talos-riscv/buildkit/util/network/cniprovider"
	"github.com/pkg/errors"
)

func getBridgeProvider(_ cniprovider.Opt) (network.Provider, error) {
	return nil, errors.Errorf("bridge network is not supported on %s yet", runtime.GOOS)
}
