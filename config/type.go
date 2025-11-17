package config

type (
	Config struct {
		Server ServerConfig `yaml:"server"`
		Oracle OracleConfig `yaml:"oracle"`
	}

	ServerConfig struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	OracleConfig struct {
		Primary DatabaseConfig `yaml:"primary"`
	}

	DatabaseConfig struct {
		Dialect                 string `yaml:"dialect"`
		ConnString              string `yaml:"connString"`
		MigrationConnString     string `yaml:"migrateConnString"`
		MaxOpenConn             int    `yaml:"maxOpenConn"`
		MaxIdleConn             int    `yaml:"maxIdleConn"`
		MaxConnLifeTimeInSecond int    `yaml:"maxConnLifeTimeInSecond"`
	}
)
