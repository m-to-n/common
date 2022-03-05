package tenants

import (
	"github.com/m-to-n/common/channels"
)

// use extension interface pattern to replace inheritance with composition
// see https://medium.com/swlh/what-is-the-extension-interface-pattern-in-golang-ce852dcecaec
type TenantChannelConfig struct {
	Channel channels.ChannelType
}

type ChannelConfigWhatsAppNumbers struct {
	PhoneNumber string
	Language    string
}

type TenantChannelConfigWhatsApp struct {
	TenantChannelConfig
	AccountSid string
	Numbers    []ChannelConfigWhatsAppNumbers
}

// represents configuration of platform tenant
// for now this means tenant channels, later NLPs, human providers, etc.
type TenantConfig struct {
	TenantId string
	Name     string
	Desc     string
	Channels []TenantChannelConfig
}
