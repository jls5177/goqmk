package flags

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via"
)

const (
	VendorId  = 0x3297
	ProductId = 0x1969
)

type UsbDevice struct {
	VendorId uint16
	ProductId uint16

	device *via.Device
}

func (u *UsbDevice) AddFlags(cmd *cobra.Command) {
	cmd.Flags().Uint16Var(&u.VendorId, "vendor-id", VendorId, "USB VendorId for Via device")
	cmd.Flags().Uint16Var(&u.ProductId, "product-id", ProductId, "USB ProductId for Via device")
}

func (u *UsbDevice) PreRun() error {
	device, err := via.New(u.VendorId, u.ProductId)
	if err != nil {
		return err
	}
	u.device = device
	return nil
}

func (u UsbDevice) Device() *via.Device {
	return u.device
}
