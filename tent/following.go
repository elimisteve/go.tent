// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Following struct {
	RemoteId    string      `json:"remote_id,omitempty"`
	Entity      Entity      `json:"entity,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
	Id          Id          `json:"id,omitempty"`
	CreatedAt   int64       `json:"created_at,omitempty"`
	UpdatedAt   int64       `json:"updated_at,omitempty"`
	Groups      []Group     `json:"groups,omitempty"`
	Profile     Profile     `json:"profile,omitempty"`
	Licenses    []URL       `json:"licenses,omitempty"`
}
