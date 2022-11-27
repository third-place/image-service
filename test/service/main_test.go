package service_test

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("CI") == "" {
		_ = godotenv.Load()
	}
	os.Exit(m.Run())
}
