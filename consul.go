package svc

import (
	"fmt"
	"os"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var (
	consul *consulapi.Client
	err    error
)

func ConsulRegister(name string, port int, tags []string) (string, error) {
	consul, err = consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		return "", err
	}

	hn, err := os.Hostname()
	if err != nil {
		return "", err
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = name + "-" + time.Now().Format("20060102150405")
	registration.Name = name
	registration.Address = hn
	registration.Port = port
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/health", hn, port)
	registration.Check.Interval = "10s"
	registration.Check.Timeout = "5s"
	registration.Tags = tags
	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		return "", err
	}
	return registration.ID, nil
}

func ConsulDeregister(id string) error {
	return consul.Agent().ServiceDeregister(id)
}
