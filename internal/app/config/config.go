package config

// Config представляет структуру для хранения конфигурационных параметров приложения.
type Config struct {
	lineDB string
}

// New создает и возвращает новый объект Config с заданными параметрами.
func New(lineDB string) *Config {
	return &Config{
		lineDB: lineDB,
	}
}

func (c *Config) GetLineDB() string {
	return c.lineDB
}
