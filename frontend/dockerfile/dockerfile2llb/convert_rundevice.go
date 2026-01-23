//go:build dfrundevice

package dockerfile2llb

import (
	"github.com/talos-riscv/buildkit/client/llb"
	"github.com/talos-riscv/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunDevices(c *instructions.RunCommand) ([]llb.RunOption, error) {
	var out []llb.RunOption
	for _, device := range instructions.GetDevices(c) {
		deviceOpts := []llb.CDIDeviceOption{
			llb.CDIDeviceName(device.Name),
		}
		if !device.Required {
			deviceOpts = append(deviceOpts, llb.CDIDeviceOptional)
		}
		out = append(out, llb.AddCDIDevice(deviceOpts...))
	}
	return out, nil
}
