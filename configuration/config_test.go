package configuration

import (
	"testing"
)

func TestConfig(t *testing.T) {
	testcase := []struct {
		key   string
		value string
	}{
		{
			"DB_HOST",
			"localhost",
		},
		{
			"DB_PORT",
			"5432",
		},
		{
			"SECRET_KEY",
			"secret",
		},
		{
			"DB_NAME",
			"test",
		},
	}

	// we use `../` before file name when we calling NewConfig from non-root path
	con := NewConfig("../.env")
	for _, i := range testcase {
		t.Run(i.key, func(t *testing.T) {
			value, _ := con.GetConfig(i.key)
			if value != "" {
				if i.value != value {
					t.Error("error : expected", i.value, "received", value)
				}
			}
		})
	}

}
