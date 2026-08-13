package config

type Config struct {
	AppName string
	// через сколько дней чистить логи
	// путь где хранить json с базой
	// путь где хранить логи

}

func New() *Config {
	return &Config{
		AppName: "todosher",
	}
}
