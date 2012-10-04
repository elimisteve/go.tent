// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Following struct {
	RemoteId    string      `json:"remoteid"`
	Entity      Entity      `json:"entity"`
	Permissions Permissions `json:"permissions"`
	Id          Id          `json:"id"`
	CreatedAt   int64       `json:"created_at"`
	UpdatedAt   int64       `json:"updated_at"`
	Groups      []Group     `json:"groups"`
	Profile     Profile     `json:"profile"`
	Licences    []URL       `json:"licences"`
}
