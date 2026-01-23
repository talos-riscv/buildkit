//go:build !linux && !windows

package ops

import (
	"github.com/talos-riscv/buildkit/snapshot"
	"github.com/talos-riscv/buildkit/solver/pb"
	"github.com/talos-riscv/buildkit/worker"
	"github.com/pkg/errors"
	copy "github.com/tonistiigi/fsutil/copy"
)

func getReadUserFn(_ worker.Worker) func(chopt *pb.ChownOpt, mu, mg snapshot.Mountable) (*copy.User, error) {
	return readUser
}

func readUser(chopt *pb.ChownOpt, mu, mg snapshot.Mountable) (*copy.User, error) {
	if chopt == nil {
		return nil, nil
	}
	return nil, errors.New("only implemented in linux and windows")
}
