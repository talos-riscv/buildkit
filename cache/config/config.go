package config

import "github.com/talos-riscv/buildkit/util/compression"

type RefConfig struct {
	Compression            compression.Config
	PreferNonDistributable bool
}
