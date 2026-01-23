package llbsolver

import (
	"context"

	cacheconfig "github.com/talos-riscv/buildkit/cache/config"
	"github.com/talos-riscv/buildkit/frontend"
	"github.com/talos-riscv/buildkit/session"
	"github.com/talos-riscv/buildkit/solver"
	"github.com/talos-riscv/buildkit/solver/llbsolver/provenance"
	"github.com/talos-riscv/buildkit/worker"
	"github.com/pkg/errors"
)

type Result struct {
	*frontend.Result
	Provenance *provenance.Result
}

type Attestation = frontend.Attestation

func workerRefResolver(refCfg cacheconfig.RefConfig, all bool, g session.Group) func(ctx context.Context, res solver.Result) ([]*solver.Remote, error) {
	return func(ctx context.Context, res solver.Result) ([]*solver.Remote, error) {
		ref, ok := res.Sys().(*worker.WorkerRef)
		if !ok {
			return nil, errors.Errorf("invalid result: %T", res.Sys())
		}

		return ref.GetRemotes(ctx, true, refCfg, all, g)
	}
}
