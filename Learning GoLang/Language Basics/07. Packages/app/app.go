package app

import (
	"fmt"

	"github.com/ShahJugalR/Packages/network"
	"github.com/ShahJugalR/Packages/storage"
)

func RunApp() {
	fmt.Println("Running application...")

	// Capitals allows to be exported to other modules too.
	network.GetInternet()
	storage.AttachDisk()
}
