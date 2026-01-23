package frontend

import (
	"context"

	"github.com/talos-riscv/buildkit/client/llb/sourceresolver"
	"github.com/talos-riscv/buildkit/executor"
	gw "github.com/talos-riscv/buildkit/frontend/gateway/client"
	"github.com/talos-riscv/buildkit/session"
	"github.com/talos-riscv/buildkit/solver"
	"github.com/talos-riscv/buildkit/solver/pb"
	"github.com/talos-riscv/buildkit/solver/result"
	digest "github.com/opencontainers/go-digest"
)

const (
	// KeySource is the option key used by the gateway frontend to represent
	// the source for the external frontend
	KeySource = "source"

	KeyDevelDeprecated = "gateway-devel"
)

type Result = result.Result[solver.ResultProxy]

type Attestation = result.Attestation[solver.ResultProxy]

type Frontend interface {
	Solve(ctx context.Context, llb FrontendLLBBridge, exec executor.Executor, opt map[string]string, inputs map[string]*pb.Definition, sid string, sm *session.Manager) (*Result, error)
}

type FrontendLLBBridge interface {
	sourceresolver.MetaResolver
	Solve(ctx context.Context, req SolveRequest, sid string) (*Result, error)
	Warn(ctx context.Context, dgst digest.Digest, msg string, opts WarnOpts) error
}

type SolveRequest = gw.SolveRequest

type CacheOptionsEntry = gw.CacheOptionsEntry

type WarnOpts = gw.WarnOpts
