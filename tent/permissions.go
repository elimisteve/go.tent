// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Permissions struct {
	Groups   []Group  `json:"groups,omitempty"`
	Entities []Entity `json:"entities,omitempty"` // TODO: FIXME: See primitives.go
	Public   bool     `json:"public,omitempty"`
}
