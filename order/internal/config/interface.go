package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type inventoryGRPCConfig interface {
	Address() string
}

type paymentGRPCConfig interface {
	Address() string
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
}

type orderHTTPConfig interface {
	Address() string
	Timeout() string
}
