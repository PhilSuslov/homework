package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type paymentGRPCConfig interface {
	Address() string
}
