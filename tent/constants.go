// Steve Phillips / elimisteve
// 2012.10.02

package tent

// Info types
const (
	TYPE_CORE  URL = "https://tent.io/types/info/core/v0.1.0"
	TYPE_BASIC URL = "https://tent.io/types/info/basic/v0.1.0"
)

// Post types
const (
	TYPE_ALBUMv010    URL = "https://tent.io/types/post/album/v0.1.0"
	TYPE_ALBUMv020    URL = "https://tent.io/types/post/album/v0.2.0"
	TYPE_DELETEv010   URL = "https://tent.io/types/post/delete/v0.1.0"
	TYPE_DELETEv020   URL = "https://tent.io/types/post/delete/v0.2.0"
	TYPE_ESSAYv010    URL = "https://tent.io/types/post/essay/v0.1.0"
	TYPE_ESSAYv020    URL = "https://tent.io/types/post/essay/v0.2.0"
	TYPE_MENTIONSv010 URL = "https://tent.io/types/post/mentions/v0.1.0"
	TYPE_MENTIONSv020 URL = "https://tent.io/types/post/mentions/v0.2.0"
	TYPE_PHOTOv010    URL = "https://tent.io/types/post/photo/v0.1.0"
	TYPE_PHOTOv020    URL = "https://tent.io/types/post/photo/v0.2.0"
	TYPE_PROFILEv010  URL = "https://tent.io/types/post/profile/v0.1.0"
	TYPE_PROFILEv020  URL = "https://tent.io/types/post/profile/v0.2.0"
	TYPE_REPOSTv010   URL = "https://tent.io/types/post/repost/v0.1.0"
	TYPE_REPOSTv020   URL = "https://tent.io/types/post/repost/v0.2.0"
	TYPE_STATUSv010   URL = "https://tent.io/types/post/status/v0.1.0"
	TYPE_STATUSv020   URL = "https://tent.io/types/post/status/v0.2.0"
)

// Licenses
const (
	LICENSE_CREATIVE_COMMONS URL = "http://creativecommons.org/licenses/by/3.0/"
	LICENSE_GPL_2            URL = "http://www.gnu.org/licenses/gpl-2.0.html"
	LICENSE_GPL_3            URL = "http://www.gnu.org/licenses/gpl-3.0.html"
	LICENSE_GPL_NEWEST       URL = "http://www.gnu.org/licenses/gpl.html"
	LICENSE_AGPL_3           URL = "http://www.gnu.org/licenses/agpl-3.0.html"
	LICENSE_AGPL_NEWEST      URL = "http://www.gnu.org/licenses/agpl.html"
	LICENSE_PRIVATE          URL = "http://fakedomain.org/licenses/private.html"
	LICENSE_LIMITED          URL = "http://fakedomain.org/intended-parties-only.html"
	// TODO: Create private license
)

// Miscellaneous
const (
	MEDIA_TYPE      = "application/vnd.tent.v0+json"
	PROFILE_REL URL = "https://tent.io/rels/profile"
	VERSION         = "0.0.1"

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
