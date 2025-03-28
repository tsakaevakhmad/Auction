package config

type JWTConfig struct {
	Secret                          string `yaml:"secret"`
	ExpirationDateInMinutes         int    `yaml:"expirationDateInMinutes"`
	RefreshTokenExpirationInMinutes int    `yaml:"refreshTokenExpirationInMinutes"`
}
