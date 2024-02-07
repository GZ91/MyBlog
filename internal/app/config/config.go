package config

// Config представляет структуру для хранения конфигурационных параметров приложения.
type Config struct {
	lineDB   string
	mainPath string
}

// New создает и возвращает новый объект Config с заданными параметрами.
func New(lineDB string, mainPath string) *Config {
	return &Config{
		lineDB:   lineDB,
		mainPath: mainPath,
	}
}

func (c *Config) GetLineDB() string {
	return c.lineDB
}

func (c *Config) GetMainPath() string {
	return c.mainPath
}
