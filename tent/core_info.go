// Steve Phillips / elimisteve
// 2012.10.02

package tent

// import "encoding/json"

// All fields optional
type CoreInfo struct {
	Entity      Entity      `json:"entity,omitempty"`
	Licenses    []URL       `json:"licenses,omitempty"`
	Servers     []URL       `json:"servers,omitempty"` // API roots
	Permissions Permissions `json:"permissions,omitempty"`
}

// func (info *CoreInfo) Marshal() ([]byte, error) {
// 	jsonStr, err := json.Marshal(info)
// 	return jsonStr, err
// }
