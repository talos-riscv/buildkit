//go:build !linux

package cniprovider

import (
	"context"

	resourcestypes "github.com/talos-riscv/buildkit/executor/resources/types"
)

func (ns *cniNS) sample() (*resourcestypes.NetworkSample, error) {
	return nil, nil
}

func withDetachedNetNSIfAny(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}
