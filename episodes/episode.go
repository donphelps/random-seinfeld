package episodes

import (
	"encoding/json"
)

// Episode represents a single Seinfeld episode.
type Episode struct {
	Season  int `json:"season"`
	Episode int `json:"episode"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Director string `json:"director"`
	Writers  string `json:"writers"`

	HuluLink string `json:"link"`

	asJSON []byte
}

// AsJSONbytes returns the episode represented in JSON.
func (e *Episode) AsJSONbytes() []byte {
	if e.asJSON == nil {
		e.asJSON, _ = json.Marshal(e)
	}
	return e.asJSON
}
