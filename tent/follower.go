// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Follower struct {
	Entity string `json:"entity"`
	Id string `json:"id"`
	Permissions Permissions `json:"permissions"`
}
