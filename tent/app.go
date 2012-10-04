// Steve Phillips / elimisteve
// 2012.09.26

package tent

type App struct {
	Name         *string           `json:"name,omitempty"`
	URL          *URL              `json:"url,omitempty"`
	Description  string            `json:"description,omitempty"`
	Icon         URL               `json:"icon,omitempty"`
	RedirectURIs []URL             `json:"redirect_uris,omitempty"`
	Scopes       map[string]string `json:"scopes,omitempty"`
}

// {
//   "name": "FooApp",
//   "description": "Does amazing foos with your data",
//   "url": "http://example.com",
//   "icon": "http://example.com/icon.png",
//   "redirect_uris": [
//     "https://app.example.com/tent/callback"
//   ],
//   "scopes": {
//     "write_profile": "Uses an app profile section to describe foos",
//     "read_followings": "Calculates foos based on your followings"
//   }
// }
