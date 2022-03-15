package dapr

import (
	daprc "github.com/dapr/go-sdk/client"
	daprcommon "github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"log"
	"os"
	"sync"
)

var (
	once       sync.Once
	instance_c daprc.Client
	instance_d daprcommon.Service
)

func DaprClientNotWorking(daprGrpcPort string) daprc.Client {
	once.Do(func() {
		client, err := daprc.NewClientWithPort(daprGrpcPort)

		if err != nil {
			// fmt.Printf("DaprClient error %s", err.Error())
			// return
			log.Fatalf("failed to start dapr client: %v", err)
		}
		instance_c = client
	})
	return instance_c
}

// daprGrpcPort - grpc port where local sidecar is listening
func DaprClient(daprGrpcPort string) daprc.Client {

	// set DAPR_GRPC_PORT otherwise NewClient will try to init on default grpc port (50001)
	// while our sidecar runs on DAPR_GRPC_PORT!
	os.Setenv("DAPR_GRPC_PORT", daprGrpcPort)
	client, err := daprc.NewClient() // NewClient uses internally doOnce.Do to prevent double init of grpc connection!

	if err != nil {
		log.Fatalf("failed to start dapr client: %v", err)
	}

	return client
}

func DaprClientClose() {
	if instance_c != nil {
		instance_c.Close()
	}
}

// daprGrpcPort - grpc port where local sidecar is listening prefixed with colon
// e.g :50001
func DaprService(daprGrpcAddr string) daprcommon.Service {
	once.Do(func() {

		s, err := daprd.NewService(daprGrpcAddr)
		if err != nil {
			log.Fatalf("failed to start dapr server: %v", err)
		}

		instance_d = s
	})
	return instance_d
}
