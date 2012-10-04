// Steve Phillips / elimisteve
// 2012.10.02

package tent

type ContentStatus struct {
	Text     string `json:"text,omitempty"`
	Location Point  `json:"location,omitempty"`
}

type ContentPost struct {
	Types  []string `json:"types,omitempty"`  // Required
	Action string   `json:"action,omitempty"` // Required
}
