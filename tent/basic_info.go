// Steve Phillips / elimisteve
// 2012.10.02

package tent

import "encoding/json"

type BasicInfo struct {
	Name      string `json:"name,omitempty"`
	AvatarURL URL    `json:"avatar_url,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Location  string `json:"location,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Bio       string `json:"bio,omitempty"`
}

// func (info *BasicInfo) Marshal() ([]byte, error) {
// 	jsonStr, err := json.Marshal(info)
// 	return jsonStr, err
// }
