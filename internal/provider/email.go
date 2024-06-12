package provider

import (
	"love_knot/internal/config"
	"love_knot/pkg/email"
)

func NewEmailClient(conf *config.Config) *email.Client {
	return email.NewEmailClient(&email.ClientConf{
		Host:     conf.Email.Host,
		Smtp:     conf.Email.Smtp,
		Addr:     conf.Email.Addr,
		Name:     conf.Email.Name,
		Password: conf.Email.Password,
	})
}
