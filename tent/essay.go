// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Essay struct {
	Title   string   `json:"title"`
	Excerpt string   `json:"excerpt"`
	Body    string   `json:"body"`    // Required
	Tags    []string `json:"tags"`
}