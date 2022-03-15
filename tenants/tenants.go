package tenants

/**
extension interface pattern to replace inheritance with composition:
see https://medium.com/swlh/what-is-the-extension-interface-pattern-in-golang-ce852dcecaec
we are not using in the end.
*/

import (
	"encoding/json"
)

type TenantChannelConfig struct {
	// serializing this kind of enum into json does not work OOTB
	// Channel channels.ChannelType `json:"channel"`
	Channel string                  `json:"channel"`
	Data    TenantChannelConfigData `json:"data"`
}

type TenantChannelConfigData struct {
	WhatsApp TenantChannelConfigWhatsApp `json:"whatsapp"`
}

type TenantChannelConfigWhatsApp struct {
	AccountSid string                         `json:"accountSid"`
	AuthToken  string                         `json:"authToken"`
	Numbers    []ChannelConfigWhatsAppNumbers `json:"numbers"`
}

type ChannelConfigWhatsAppNumbers struct {
	PhoneNumber string `json:"phoneNumber"`
	Language    string `json:"language"`
}

// represents configuration of platform tenant
// for now this means tenant channels, later NLPs, human providers, etc.
type TenantConfig struct {
	TenantId string                `json:"tenantId"`
	Name     string                `json:"name"`
	Desc     string                `json:"desc"`
	Channels []TenantChannelConfig `json:"channels"`
}

type TenantConfigByTenantIdReq struct {
	TenantId string `json:"tenantId"`
}

type TenantConfigByTwilioAccIDAndReceiverNumReq struct {
	AccountSid          string `json:"accountSid"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber"`
}

func TenantConfigToJson(tc TenantConfig) ([]byte, error) {
	tcBytes, err := json.Marshal(tc)
	if err != nil {
		return nil, err
	}
	return tcBytes, nil
}

func TenantConfigToJsonString(tc TenantConfig) (*string, error) {
	tcBytes, err := TenantConfigToJson(tc)
	if err != nil {
		return nil, err
	}

	tcStr := string(tcBytes)

	return &tcStr, nil
}
