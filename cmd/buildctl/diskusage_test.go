package main

import (
	"testing"

	"github.com/talos-riscv/buildkit/util/testutil/integration"
	"github.com/stretchr/testify/require"
)

func testDiskUsage(t *testing.T, sb integration.Sandbox) {
	cmd := sb.Cmd("du")
	err := cmd.Run()
	require.NoError(t, err)
}
