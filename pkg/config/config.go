package config

import "os"

type Environment string

const (
	Dev     Environment = "DEV"
	Test    Environment = "TEST"
	Staging Environment = "STAGE"
	Prod    Environment = "PROD"
)

func (e Environment) String() string {
	return string(e)
}

func (e Environment) Valid() bool {
	return e == Dev || e == Test || e == Staging || e == Prod
}

type Config struct {
	Env Environment

	Port string
}

func New() *Config {
	env := Environment(getEnv("ENV", "DEV"))
	if !env.Valid() {
		panic("invalid environment, must be one of DEV, TEST, STAGE, PROD")
	}

	return &Config{
		Env: env,

		Port: getEnv("PORT", "8080"),
	}
}

// getEnv returns the value of the environment variable
// or the default value if the variable is not set
func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
