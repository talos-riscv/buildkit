package testutil

import (
	"testing"

	"github.com/talos-riscv/buildkit/solver"
)

func TestMemoryCacheStorage(t *testing.T) {
	RunCacheStorageTests(t, solver.NewInMemoryCacheStorage)
}
