// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Photo struct {
	Caption string   `json:"caption"`
	Albums  []Id     `json:"albums"`
	Tags    []string `json:"tags"`
	Exif    string   `json:"exif"`
}
