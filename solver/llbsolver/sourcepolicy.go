package llbsolver

import (
	"context"

	"github.com/talos-riscv/buildkit/solver/pb"
)

type SourcePolicyEvaluator interface {
	Evaluate(ctx context.Context, op *pb.Op) (bool, error)
}
