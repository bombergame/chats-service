package rest

const (
	ChatIDPathVar = "chat_id"
	ChatIDFormat  = "[0-9-a-z]+"
	ChatIDPath    = "{" + ChatIDPathVar + ":" + ChatIDFormat + "}"
)
