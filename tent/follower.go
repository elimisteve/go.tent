// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Follower struct {
	Entity      Entity      `json:"entity,omitempty"`
	Id          Id          `json:"id,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
}
