package restricter

import "encoding/json"

// Datastore gives the restricter the full.
type Datastore interface {
	GetMany([]string) map[string]json.RawMessage
	GetAll() map[string]json.RawMessage
}

// Element knows how to restrict one element.
// TODO: Find a better name.
type Element interface {
	Restrict(int, json.RawMessage) (json.RawMessage, error)
}

// HasPermer tells if a user has a specivic perm.
type HasPermer interface {
	HasPerm(uid int, perm string) bool
	IsSuperadmin(uid int) bool
	InGroups(uid int, groups []int) bool
	UserRequired(uid int) []string
	Get(collection string, id int, v interface{}) error
}
