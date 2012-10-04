// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Permissions struct {
	Groups   []Group  `json:"groups"`
	Entities Entities `json:"entities"` // TODO: FIXME: See primitives.go
	Public   bool     `json:"public"`
}
