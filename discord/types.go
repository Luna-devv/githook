package discord

type WebhookPayload struct {
	Content    string               `json:"content"`
	Username   string               `json:"username,omitempty"`
	AvatarURL  string               `json:"avatar_url,omitempty"`
	Embeds     []Embed              `json:"embeds,omitempty"`
	Components []ActionRowComponent `json:"components,omitempty"`
}

type Embed struct {
	Author      EmbedAuthor  `json:"author,omitempty"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	URL         string       `json:"url,omitempty"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
	Footer      EmbedFooter  `json:"footer,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedAuthor struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type ComponentType int

const (
	ActionRow ComponentType = 1
	Button    ComponentType = 2
)

type ActionRowComponent struct {
	Type        ComponentType     `json:"type"`
	Componments []ButtonComponent `json:"components"`
}

type ButtonStyle int

const (
	Primary   ButtonStyle = 1
	Secondary ButtonStyle = 2
	Success   ButtonStyle = 3
	Danger    ButtonStyle = 4
	Link      ButtonStyle = 5
)

type ButtonComponent struct {
	Type     ComponentType `json:"type"`
	Style    ButtonStyle   `json:"style"`
	Label    string        `json:"label"`
	CustomID string        `json:"custom_id"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
}

type Emoji struct {
	Name     string `json:"name"`
	ID       string `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}
