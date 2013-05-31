package gpio

import "testing"

func TestOpeningValves(t *testing.T) {
	for _, i := range Valves() {
		Open(i)
	}
}

func TestValvesAreOpened(t *testing.T) {
	for _, i := range Valves() {
		opened := IsOpened(i)
		t.Logf("Relay #%d is %t\n", i, opened)
	}
}

func TestClosingValves(t *testing.T) {
	for _, i := range Valves() {
		Close(i)
	}
}

func TestValvesAreClosed(t *testing.T) {
	for _, i := range Valves() {
		opened := IsOpened(i)
		t.Logf("Relay #%d is %t\n", i, opened)
	}
}
