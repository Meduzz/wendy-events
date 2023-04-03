package wendyevents

import "testing"

func TestModuleWithEvent(t *testing.T) {
	subject := NewModuleWithEvents("test")

	// RPC way
	subject.WithHandler("rpc", nil)
	// Event way :yay:
	subject.WithListener("event", nil)
}
