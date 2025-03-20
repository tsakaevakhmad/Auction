package configurations

type MainConfig struct {
	Database struct {
		ConnectionString string `yaml:"connectionString"`
	} `yaml:"database"`

	Server struct {
		Domain       string   `yaml:"domain"`
		AllowOrigins []string `yaml:"allowOrigins"`
		Port         int      `yaml:"port"`
	} `yaml:"server"`

	JWTConfig JWTConfig `yaml:"jwtConfig"`
}
