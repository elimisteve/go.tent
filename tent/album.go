// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Album struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Photos      []Id   `json:"photos"` // Required
	Cover       Id     `json:"cover"`
}
