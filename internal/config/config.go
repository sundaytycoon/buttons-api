package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"

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
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Config struct {
	Env   string
	Debug bool

	BatchDB      *Database `mapstructure:"batchDatabase"`
	ServiceDB    *Database `mapstructure:"serviceDatabase"`
	HTTPEndPoint *EndPoint `mapstructure:"httpServer"`
	GRPCEndPoint *EndPoint `mapstructure:"grpcServer"`
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
