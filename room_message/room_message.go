package room_message

type RoomMessage struct {
	Event string `json:"event"`
	Item  Item   `json:"item"`
}

type Item struct {
	Message       Message `json:"message"`
	Room          Room    `json:"room"`
	OauthClientId string  `json:"oauth_client_id"`
	WebhookId     string  `json:"webhook_id"`
}

type Message struct {
	Date string `json:"date"`
	File string `json:"file"`
	From struct {
		Id    int32 `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		MentionName string `json:"mention_name"`
		Name        string `json:"name"`
	} `json:"from"`
	Id       int32     `json:"id"`
	Mentions []Mention `json:"mentions"`
	Message  string    `json:"message"`
}

type Mention struct {
	Id    int32 `json:"id"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	MentionName string `json:"mention_name"`
	Name        string `json:"name"`
}

type Room struct {
	Id    int32 `json:"id"`
	Links struct {
		Members  string `json:"members"`
		Self     string `json:"self"`
		Webhooks string `json:"webhooks"`
	} `json:"links"`
	Name string `json:"name"`
}

type RoomNotification struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	MessageFormat string `json:"message_format"`
}
