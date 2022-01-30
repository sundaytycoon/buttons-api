package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type Config struct {
	Env   string
	Debug bool

	ButtonsAdminServer *EndPoint `mapstructure:"buttonsAdminServer"`
	ButtonsAdminWeb    *EndPoint `mapstructure:"buttonsAdminWeb"`
	BatchDB            *Database `mapstructure:"batchDatabase"`
	ServiceDB          *Database `mapstructure:"serviceDatabase"`
	HTTPEndPoint       *EndPoint `mapstructure:"httpServer"`
	GRPCEndPoint       *EndPoint `mapstructure:"grpcServer"`
	Google             *Google   `mapstructure:"google"`
}

const (
	SSMPathFormat = "/lambda/%s/%s/%s"
	FileName      = "application.yml"
)

func New() (*Config, error) {
	var (
		cfg = &Config{}
		env = os.Getenv("ENV")
		op  = er.GetOperator()
	)

	viper.SetConfigType("yaml")

	if env == "alpha" || env == "prod" {
		// yaml, err := getYamlFromSSM(env, buttonsapi.App)
		// if err != nil {
		// 	return nil, er.WrapOp(err, op)
		// }
		// if err = viper.ReadConfig(strings.NewReader(yaml)); err != nil {
		// 	return nil, er.WrapOp(err, op)
		// }
	} else if env == "local" || env == "" {
		viper.AddConfigPath(".")
		viper.SetConfigFile(FileName)
		if err := viper.ReadInConfig(); err != nil {
			return nil, er.WrapOp(err, op)
		}
	} else {
		return nil, er.WrapOp(fmt.Errorf("ENV only accept [alpha, prod, local(or empty)]"), op)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, er.WrapOp(err, op)
	}
	cfg.Env = env
	return cfg, nil
}
