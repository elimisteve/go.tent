// Steve Phillips / elimisteve
// 2012.10.02

package tent

type StatusPost struct {
	Id          Id           `json:"id,omitempty"`
	Entity      Entity       `json:"entity,omitempty"`
	PublishedAt int64        `json:"published_at,omitempty"`
	ReceivedAt  int64        `json:"received_at,omitempty"`
	Mentions    []Mention    `json:"mentions,omitempty"`
	Licenses    []URL        `json:"licenses,omitempty"`
	Type        URL          `json:"type,omitempty"`
	Content     Content      `json:"content,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	App         App          `json:"app,omitempty"`
	Views       []View       `json:"views,omitempty"`
	Permissions Permissions  `json:"permissions,omitempty"`
}
