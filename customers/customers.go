package customers

import (
	"github.com/m-to-n/common/channels"
	"time"
)

type CustomerProfile struct {
	CustomerId  string
	PhoneNumber string
	Email       string
	Name        string
	Surname     string
}

type CustomerSession struct {
	SessionId  string
	CustomerId string
	Channel    channels.ChannelType
}

type SessionMessage struct {
	SessionId   string
	CreatedAt   time.Time
	MessageText string
}
