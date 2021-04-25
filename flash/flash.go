package flash

import (
	"fmt"
	"github.com/jls5177/goqmk/via"
	wallycli "github.com/jls5177/wally-cli"
	"time"
)

func FlashBoard(filePath string, viaDevice *via.Device) error {
	b, err := wallycli.New(wallycli.DfuBoard)
	if err != nil {
		return err
	}

	// reset the device
	fmt.Printf("reseting keyboard....")
	if err := viaDevice.GotoBootloader(); err != nil {
		return err
	}
	fmt.Printf("success!\r\n")

	// start flashing
	fmt.Printf("starting flashing, please wait....\r\n")
	if err := b.FlashAsync(filePath); err != nil {
		return err
	}

	for !b.Finished() {
		time.Sleep(500 * time.Millisecond)
		if b.Running() {
			fmt.Printf("\r%v/%v", b.CompletedSteps(), b.TotalSteps())
		}
	}
	fmt.Printf("\rflashing completed!\r\n")
	return b.FlashError()
}
