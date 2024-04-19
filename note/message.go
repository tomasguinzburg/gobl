package note

import (
	"encoding/json"

	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/pkg/here"
	"github.com/invopop/gobl/uuid"
	"github.com/invopop/jsonschema"
	"github.com/invopop/validation"
)

// Message represents a simple message object with a title and some
// content meant.
type Message struct {
	// UUID of the message.
	UUID uuid.UUID `json:"uuid,omitempty" jsonschema:"title=UUID"`
	// Summary of the message content
	Title string `json:"title,omitempty" jsonschema:"title=Title"`
	// Details of what exactly this message wants to communicate.
	Content string `json:"content" jsonschema:"title=Content"`
	// Any additional semi-structured data that might be useful.
	Meta cbc.Meta `json:"meta,omitempty" jsonschema:"title=Meta Data"`
}

// Validate ensures the message contains everything it should.
func (m *Message) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.UUID),
		validation.Field(&m.Content, validation.Required),
		validation.Field(&m.Meta),
	)
}

// JSONSchemaExtend adds examples to the JSON Schema.
func (Message) JSONSchemaExtend(s *jsonschema.Schema) {
	exs := here.Bytes(`
		[
			{
				"title": "Example Title",
				"content": "This is an example message."
			}
		]`)
	if err := json.Unmarshal(exs, &s.Examples); err != nil {
		panic(err)
	}
}
