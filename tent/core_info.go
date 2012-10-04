// Steve Phillips / elimisteve
// 2012.10.02

package tent

// All fields optional
type CoreInfo struct {
	Entity      Entity      `json:"entity"`
	Licenses    []URL       `json:"licenses"`
	Servers     []URL       `json:"servers"` // API roots
	Permissions Permissions `json:"permissions"`
}
