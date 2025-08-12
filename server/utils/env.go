package utils

import "os"

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvOrDefault(key, d string) string {
	v, ok := os.LookupEnv(key)

	if !ok {
		return d
	}

	return v
}
