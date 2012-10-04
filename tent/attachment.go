// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Attachment struct {
	Id          Id     `json:"id"`
	ContentType string `json:"type"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Data        string `json:"data"`
	Size        int64  `json:"size"`
	Filename    string `json:"filename"`
}
