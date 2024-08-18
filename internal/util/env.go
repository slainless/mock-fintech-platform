package util

import "os"

func IsPositiveEnv(v string) bool {
	return v == "true" || v == "1" || v == "TRUE" || v == "y" || v == "Y" || v == "yes" || v == "YES"
}

func UseMockServices() bool {
	v, ok := os.LookupEnv("USE_MOCK_SERVICES")
	return ok || IsPositiveEnv(v)
}
