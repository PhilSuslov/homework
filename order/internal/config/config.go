package config

import (
	"os"

	"github.com/PhilSuslov/homework/order/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger      LoggerConfig
	PaymentGRPC paymentGRPCConfig
	InventoryGRPC inventoryGRPCConfig
	OrderHTTP orderHTTPConfig
	PostgresCfg PostgresConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	paymentGRPCCfg, err := env.NewPaymentGRPCConfig()
	if err != nil {
		return err
	}

	inventoryGRPCcfg, err := env.NewInventoryGRPCConfig()
	if err != nil{
		return err
	}

	orderHTTPCfg, err := env.NewOrderHTTPConfig()
	if err != nil{
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil{
		return err
	}

	appConfig = &config{
		Logger:      loggerCfg,
		PaymentGRPC: paymentGRPCCfg,
		InventoryGRPC: inventoryGRPCcfg,
		OrderHTTP: orderHTTPCfg,
		PostgresCfg: postgresCfg,
	}

	return nil
}

func AppConfig() *config { return appConfig }
