package tenants

import (
	"github.com/m-to-n/common/channels"
)

// use extension interface pattern to replace inheritance with composition
// see https://medium.com/swlh/what-is-the-extension-interface-pattern-in-golang-ce852dcecaec
type TenantChannelConfig struct {
	Channel channels.ChannelType `json:"channel"`
}

type ChannelConfigWhatsAppNumbers struct {
	PhoneNumber string `json:"phoneNumber"`
	Language    string `json:"language"`
}

type TenantChannelConfigWhatsApp struct {
	TenantChannelConfig
	AccountSid string                         `json:"accountSid"`
	Numbers    []ChannelConfigWhatsAppNumbers `json:"numbers"`
}

// represents configuration of platform tenant
// for now this means tenant channels, later NLPs, human providers, etc.
type TenantConfig struct {
	TenantId string                `json:"tenantId"`
	Name     string                `json:"name"`
	Desc     string                `json:"desc"`
	Channels []TenantChannelConfig `json:"channels"`
}
