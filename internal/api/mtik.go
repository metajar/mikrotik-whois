package api

import (
	"fmt"
	"github.com/go-routeros/routeros"
	"github.com/metajar/mikrotik-whois/internal/config"
)

type MTClient struct {
	Config *config.MikrotikConfig
	Client *routeros.Client
}

func New(c *config.MikrotikConfig) MTClient {
	return MTClient{
		Config: c,
		Client: nil,
	}
}

func (m *MTClient) Connect() error {
	c, err := routeros.Dial(fmt.Sprintf("%v:%v", m.Config.Address, m.Config.Port), m.Config.Username, m.Config.Password)
	if err != nil {
		return err
	}
	m.Client = c
	return nil
}

func (m *MTClient) Close() error {
	m.Client.Close()
	return nil
}

func (m *MTClient) RunCommand(c ...string) (*routeros.Reply, error) {
	out, err := m.Client.Run(c...)
	if err != nil {
		return &routeros.Reply{}, err
	}
	return out, nil

}

func (m *MTClient) GetDHCPHost(ip string) (string, error) {
	a, err := m.RunCommand("/ip/dhcp-server/lease/print", fmt.Sprintf("?address=%s", ip))
	if err != nil {
		return "", err
	}
	return a.Re[0].Map["host-name"], nil
}
