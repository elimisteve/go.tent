package tent

const (
	MEDIA_TYPE  = "application/vnd.tent.v0+json"
	PROFILE_REL = "https://tent.io/rels/profile"
	VERSION     = "0.0.1"

	NONCE_LENGTH = 6

	STATUS_BODY = `{"type": "https://tent.io/types/post/status/v0.1.0","published_at": %d,"permissions": {"public": true},"licenses": ["http://creativecommons.org/licenses/by/3.0/"],"content": {"text": "%s"}}`

)

var (
	STATUS_HEADERS = map[string]string{
		"Content-Type": MEDIA_TYPE,
		"Accept": MEDIA_TYPE,
	}
)