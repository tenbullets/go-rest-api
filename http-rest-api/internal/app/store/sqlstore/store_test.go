package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=postgres password=qwikWell12 dbname=go_site_test sslmode=disable"
	}

	os.Exit(m.Run())
}
