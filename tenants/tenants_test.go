package tenants

import (
	"github.com/m-to-n/common/channels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientValidateIncomingRequest(t *testing.T) {
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
	assert.Equal(t, *tcStr, `{"tenantId":"tenant-123","name":"dummyTenant","desc":"dummy tenant for development \u0026 testing","channels":[{"channel":"whatsapp"}]}`)
}
