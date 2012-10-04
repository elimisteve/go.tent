// Steve Phillips / elimisteve
// 2012.10.02

package tent

type BasicInfo struct {
	Name      string `json:"name"`
	AvatarURL URL    `json:"avatar_url"`
	Birthdate string `json:"birthdate"`
	Location  string `json:"location"`
	Gender    string `json:"gender"`
	Bio       string `json:"bio"`
}
