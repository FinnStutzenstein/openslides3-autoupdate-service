package projector

import "encoding/json"

// Datastore are the methods a usual projector slide needs.
type Datastore interface {
	Get(collection string, id int, v interface{}) error
	GetCollection(string) []json.RawMessage
	GetModels(string, []int) []json.RawMessage
	ConfigValue(string, interface{}) error
}

// Callable knows how to build the projector data for an element.
type Callable interface {
	Build(ds Datastore, element json.RawMessage, pid int) (json.RawMessage, error)
}

// CallableFunc is a functions that implements the Callable interface.
type CallableFunc func(ds Datastore, element json.RawMessage, pid int) (json.RawMessage, error)

// Build builds the data of a projector.
func (f CallableFunc) Build(ds Datastore, element json.RawMessage, pid int) (json.RawMessage, error) {
	return f(ds, element, pid)
}
