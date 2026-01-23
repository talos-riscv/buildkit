// Package npipe provides connhelper for npipe://<address>
package npipe

import "github.com/talos-riscv/buildkit/client/connhelper"

func init() {
	connhelper.Register("npipe", Helper)
}
