package container

import (
	"net"

	"github.com/talos-riscv/buildkit/executor"
	"github.com/talos-riscv/buildkit/solver/pb"
	"github.com/pkg/errors"
)

func ParseExtraHosts(ips []*pb.HostIP) ([]executor.HostIP, error) {
	out := make([]executor.HostIP, len(ips))
	for i, hip := range ips {
		ip := net.ParseIP(hip.IP)
		if ip == nil {
			return nil, errors.Errorf("failed to parse IP %s", hip.IP)
		}
		out[i] = executor.HostIP{
			IP:   ip,
			Host: hip.Host,
		}
	}
	return out, nil
}
