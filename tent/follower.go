// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Follower struct {
	Entity      string      `json:"entity,omitempty"`
	Id          string      `json:"id,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
}
