// Steve Phillips / elimisteve
// 2012.10.02

package tent

type Post struct {
	Id          Id           `json:"id"`
	Entity      Entity       `json:"entity"`
	PublishedAt int64        `json:"published_at"`
	ReceivedAt  int64        `json:"received_at"`
	Mentions    []Mention    `json:"mentions"`
	Licenses    []URL        `json:"licenses"`
	Type        URL          `json:"type"`
	Content     Content      `json:"content"`
	Attachments []Attachment `json:"attachments"`
	App         App          `json:"app"`
	Views       []View       `json:"views"`
	Permissions Permissions  `json:"permissions"`
}
