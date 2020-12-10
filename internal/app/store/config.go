package store

//Config ...
type Config struct {
	DB_url string `toml:db_url`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		DB_url: "asterisk:asterisk@(192.168.209.149:3306)/queue_nao",
	}
}
