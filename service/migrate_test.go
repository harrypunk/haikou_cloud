package service

import (
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {
	os.Setenv("DB_NAME", "haikou_test")
	os.Setenv("DB_USER", "onetwo")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_URL", "localhost")
	err := Migrate()
	if err != nil {
		t.Fail()
	}
}
