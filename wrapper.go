package wendyevents

import "github.com/Meduzz/wendy"

type (
	ModuleWithEvents struct {
		*wendy.Module
		*EventListener
	}
)

func NewModuleWithEvents(name string) *ModuleWithEvents {
	mod := wendy.NewModule(name)
	event := NewEventListener(name)

	return &ModuleWithEvents{mod, event}
}
