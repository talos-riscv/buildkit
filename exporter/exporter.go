package exporter

import (
	"context"

	"github.com/talos-riscv/buildkit/cache"
	"github.com/talos-riscv/buildkit/exporter/containerimage/exptypes"
	"github.com/talos-riscv/buildkit/solver/result"
	"github.com/talos-riscv/buildkit/util/compression"
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
)

type Source = result.Result[cache.ImmutableRef]

type Attestation = result.Attestation[cache.ImmutableRef]

type Exporter interface {
	Resolve(context.Context, int, map[string]string) (ExporterInstance, error)
}

type ExporterInstance interface {
	ID() int
	Name() string
	Config() *Config
	Type() string
	Attrs() map[string]string
	Export(ctx context.Context, src *Source, buildInfo ExportBuildInfo) (map[string]string, DescriptorReference, error)
}

type ExportBuildInfo struct {
	Ref         string
	InlineCache exptypes.InlineCache
	SessionID   string
}

type DescriptorReference interface {
	Release() error
	Descriptor() ocispecs.Descriptor
}

type Config struct {
	// Make the field private in case it is initialized with nil compression.Type
	compression compression.Config
}

func NewConfig() *Config {
	return &Config{
		compression: compression.Config{
			Type: compression.Default,
		},
	}
}

func NewConfigWithCompression(comp compression.Config) *Config {
	return &Config{
		compression: comp,
	}
}

func (c *Config) Compression() compression.Config {
	return c.compression
}
