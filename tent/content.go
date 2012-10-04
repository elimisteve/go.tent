// Steve Phillips / elimisteve
// 2012.10.02

package tent

// Content is a superset of Status, Repost, Profile, and Delete. Visit
// http://tent.io/docs/post-types to see wtf I'm talking about.
type Content struct {
	Entity   Entity `json:"entity,omitempty"`
	Id       Id     `json:"id,omitempty"`
	Types    []URL  `json:"types,omitempty"`
	Action   string `json:"action,omitempty"`
	Text     string `json:"text,omitempty"`
	Location Point  `json:"location,omitempty"`
}
