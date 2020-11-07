package adapter

import "testing"

func TestMain(m *testing.M) {
	client := &client{}
	mac := &mac{}

	client.insertSquareUsbInComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{windowMachine:windowsMachine}

	client.insertSquareUsbInComputer(windowsMachineAdapter)
}