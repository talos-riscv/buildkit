package local

import (
	"github.com/talos-riscv/buildkit/solver/llbsolver/provenance"
	provenancetypes "github.com/talos-riscv/buildkit/solver/llbsolver/provenance/types"
	"github.com/talos-riscv/buildkit/source"
	srctypes "github.com/talos-riscv/buildkit/source/types"
	"github.com/tonistiigi/fsutil"
)

type LocalIdentifier struct {
	Name               string
	SessionID          string
	IncludePatterns    []string
	ExcludePatterns    []string
	FollowPaths        []string
	SharedKeyHint      string
	Differ             fsutil.DiffType
	MetadataOnly       bool
	MetadataExceptions []string
}

func NewLocalIdentifier(str string) (*LocalIdentifier, error) {
	return &LocalIdentifier{Name: str}, nil
}

func (*LocalIdentifier) Scheme() string {
	return srctypes.LocalScheme
}

var _ source.Identifier = (*LocalIdentifier)(nil)

func (id *LocalIdentifier) Capture(c *provenance.Capture, pin string) error {
	c.AddLocal(provenancetypes.LocalSource{
		Name: id.Name,
	})
	return nil
}
