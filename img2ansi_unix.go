//build: unix

package main

import (
	"os"

	"github.com/snowsnail/gosshold/ssh/terminal"
)

func getTermDim() (w, h int, err error) {
	return terminal.GetSize(int(os.Stdout.Fd()))
}
