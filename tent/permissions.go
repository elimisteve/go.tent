// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Permissions struct {
	Groups []Group `json:"groups"`
	// Entities []Entity `json:"entities,omitempty"` // TODO: FIXME: See primitives.go
	Entities map[Entity]bool `json:"entities"`
	Public   bool            `json:"public"`
}
