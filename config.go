package cloudlinux

type Config struct {
	Username  string
	SecretKey string
}

func NewConfig(username, secretKey string) *Config {
	return &Config{
		Username:  username,
		SecretKey: secretKey,
	}
}
