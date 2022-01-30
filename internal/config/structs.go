package config

import (
	"fmt"
	"net"
	"time"

	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dialect  string `mapstructure:"dialect"`
	Name     string `mapstructure:"name"`

	ConnectionValidation                    bool          `mapstructure:"connection_validation"`
	ConnectionValidationSQL                 string        `mapstructure:"connection_validation_sql"`
	ConnectionValidationRetryTimes          int64         `mapstructure:"connection_validation_retry_times"`
	ConnectionValidationRetryDuringEachTime time.Duration `mapstructure:"connection_validation_retry_during_each_time"`
}

func (o *Database) DSN() string {
	op := er.GetOperator()

	if o.Dialect == "mysql" {
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true",
			o.User, o.Password, o.Host, o.Port, o.Name,
		)
	} else if o.Dialect == "postges" {
		return fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			o.User, o.Password, o.Host, o.Port, o.Name,
		)
	} else {
		panic(er.WrapOp(fmt.Errorf("database's dialect is wrong, [%v]", o), op))
	}
}

type EndPoint struct {
	Host    string        `mapstructure:"host"`
	Port    string        `mapstructure:"port"`
	TLS     bool          `mapstructure:"tls"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func (o *EndPoint) HTTPAddress() string {
	basicAddress := net.JoinHostPort(o.Host, o.Port)
	if o.TLS {
		return "https://" + basicAddress
	}
	return "http://" + basicAddress
}

type Google struct {
	OAuthCallbackURL string `mapstructure:"oauthCallbackUrl"`
	ClientID         string `mapstructure:"clientId"`
	ClientSecret     string `mapstrcuture:"clientSecret"`
}
