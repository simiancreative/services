package server

import (
	"github.com/simiancreative/services/service"
)

var services service.Collection

func AddService(addedservice service.Config) {
	services = append(services, addedservice)
}

func AddServices(addedservices []service.Config) {
	services = append(services, addedservices...)
}
