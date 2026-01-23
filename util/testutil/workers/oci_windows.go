package workers

import "github.com/talos-riscv/buildkit/util/bklog"

func initOCIWorker() {
	bklog.L.Info("OCI Worker not supported on Windows.")
}
