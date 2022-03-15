package tenants

import (
	"github.com/m-to-n/common/channels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTenantConfigToJsonString1(t *testing.T) {
	t.Parallel()

	tenantConfig := TenantConfig{
		TenantId: "tenant-123",
		Name:     "dummyTenant",
		Desc:     "dummy tenant for development & testing",
		Channels: make([]TenantChannelConfig, 0),
	}

	tenantChannelConfig := TenantChannelConfig{
		Channel: channels.CHANNELS_WHATSAPP,
	}

	tenantConfig.Channels = append(tenantConfig.Channels, tenantChannelConfig)

	tcStr, _ := TenantConfigToJsonString(tenantConfig)
	assert.Equal(t, *tcStr, `{"tenantId":"tenant-123","name":"dummyTenant","desc":"dummy tenant for development \u0026 testing","channels":[{"channel":"whatsapp","data":{"whatsapp":{"accountSid":"","authToken":"","numbers":null}}}]}`)
}

func TestTenantConfigToJsonString2(t *testing.T) {
	t.Parallel()

	tenantConfig := TenantConfig{
		TenantId: "tenant-123",
		Name:     "dummyTenant",
		Desc:     "dummy tenant for development & testing",
		Channels: make([]TenantChannelConfig, 0),
	}

	tenantChannelConfigWhatsApp := TenantChannelConfigWhatsApp{
		AccountSid: "ACC123",
		AuthToken:  "b32456",
		Numbers: []ChannelConfigWhatsAppNumbers{
			ChannelConfigWhatsAppNumbers{
				PhoneNumber: "+420123456789",
				Language:    "en",
			},
		},
	}

	tenantChannelConfig := TenantChannelConfig{
		Channel: channels.CHANNELS_WHATSAPP,
		Data: TenantChannelConfigData{
			WhatsApp: tenantChannelConfigWhatsApp,
		},
	}

	tenantConfig.Channels = append(tenantConfig.Channels, tenantChannelConfig)

	tcStr, _ := TenantConfigToJsonString(tenantConfig)
	assert.Equal(t, *tcStr, `{"tenantId":"tenant-123","name":"dummyTenant","desc":"dummy tenant for development \u0026 testing","channels":[{"channel":"whatsapp","data":{"whatsapp":{"accountSid":"ACC123","authToken":"b32456","numbers":[{"phoneNumber":"+420123456789","language":"en"}]}}}]}`)

}
