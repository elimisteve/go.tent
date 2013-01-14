// Steve Phillips / elimisteve
// 2012.10.02

package tent

import (
	"time"
)

type StatusPost struct {
	Id          Id           `json:"id,omitempty"`
	Entity      string       `json:"entity,omitempty"`
	PublishedAt int64        `json:"published_at,omitempty"`
	ReceivedAt  int64        `json:"received_at,omitempty"`
	UpdatedAt   int64        `json:"updated_at,omitempty"`
	Mentions    []Mention    `json:"mentions,omitempty"`
	Licenses    []URL        `json:"licenses,omitempty"`
	Type        URL          `json:"type,omitempty"`
	Version     int64        `json:"version,omitempty"`
	Content     Content      `json:"content,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	App         App          `json:"app,omitempty"`
	Views       []View       `json:"views,omitempty"`
	Permissions Permissions  `json:"permissions,omitempty"`
}

// NewStatus generates a new StatusPost using the given message
// intended for the given recipients
//
// TODO: Rename to `NewPrivateStatus` or `NewLimitedStatus`, whichever
// is more precise
func NewStatus(message string, recipients ...string) (status StatusPost) {
	permissions := Permissions{
		Entities: make(map[string]bool, len(recipients)),
		Public:   false,
	}
	mentions := []Mention{}
	// Mention each recipient and give them view permissions
	for _, entity := range recipients {
		permissions.Entities[entity] = true
		mentions = append(mentions, Mention{Entity: entity})
	}

	status = StatusPost{
		Type:        TYPE_STATUSv020,
		PublishedAt: time.Now().Unix(),
		Permissions: permissions,
		Mentions:    mentions,
		Licenses:    []URL{LICENSE_LIMITED},
		// Licenses: []URL{LICENSE_PRIVATE},
		Content: Content{Text: message},
	}
	return
}
