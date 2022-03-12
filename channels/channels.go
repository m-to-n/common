package channels

import (
	"errors"
)

const CHANNELS_WHATSAPP = "whatsapp"
const CHANNELS_UNKNOWN = "unknown"

// https://threedots.tech/post/safer-enums-in-go/
type ChannelType struct {
	channelType string `json:"channelType"`
}

var (
	Unknown  = ChannelType{channelType: ""}
	WhatsApp = ChannelType{channelType: "whatsapp"}
)

func (cht ChannelType) String() string {
	return cht.channelType
}

func ChannelTypeFromString(s string) (ChannelType, error) {
	switch s {
	case WhatsApp.channelType:
		return WhatsApp, nil
	}

	return Unknown, errors.New("unknown channel type: " + s)
}
