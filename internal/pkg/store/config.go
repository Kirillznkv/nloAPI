package store

type Config struct {
	DatabaseURL    string `toml:"database_url"`
	MigrateDirPath string `toml:"migrate_dirpath""`
}

func NewConfig() *Config {
	return &Config{
		MigrateDirPath: "db/migrations",
	}
}
