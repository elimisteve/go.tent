// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Mention struct {
	Entity Entity `json:"entity,omitempty"`
	Post   Id     `json:"post,omitempty"`
}
