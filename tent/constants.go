// Steve Phillips / elimisteve
// 2012.10.02

package tent

type ProfileInfo string

// Info types
const (
	TYPE_CORE  ProfileInfo = "https://tent.io/types/info/core/v0.1.0"
	TYPE_BASIC ProfileInfo = "https://tent.io/types/info/basic/v0.1.0"
)

// Post types
const (
	TYPE_ALBUM    ProfileInfo = "https://tent.io/types/post/album/v0.1.0"
	TYPE_DELETE   ProfileInfo = "https://tent.io/types/post/delete/v0.1.0"
	TYPE_ESSAY    ProfileInfo = "https://tent.io/types/post/essay/v0.1.0"
	TYPE_MENTIONS ProfileInfo = "https://tent.io/types/post/mentions/v0.1.0"
	TYPE_PHOTO    ProfileInfo = "https://tent.io/types/post/photo/v0.1.0"
	TYPE_PROFILE  ProfileInfo = "https://tent.io/types/post/profile/v0.1.0"
	TYPE_REPOST   ProfileInfo = "https://tent.io/types/post/repost/v0.1.0"
	TYPE_STATUS   ProfileInfo = "https://tent.io/types/post/status/v0.1.0"
)

// Miscellaneous
const (
	MEDIA_TYPE  = "application/vnd.tent.v0+json"
	PROFILE_REL = "https://tent.io/rels/profile"
	VERSION     = "0.0.1"

	NONCE_LENGTH = 6

	STATUS_BODY = `{"type": "https://tent.io/types/post/status/v0.1.0","published_at": %d,"permissions": {"public": true},"licenses": ["http://creativecommons.org/licenses/by/3.0/"],"content": {"text": "%s"}}`
)

// These act like constants but technically aren't
var (
	STATUS_HEADERS = map[string]string{
		"Content-Type": MEDIA_TYPE,
		"Accept":       MEDIA_TYPE,
	}
)
