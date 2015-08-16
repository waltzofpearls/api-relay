package rapi_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/relay-api/rapi"
)

func TestEmptyConfig(t *testing.T) {
	rapi.NewConfig()
}

func TestValidConfigFile(t *testing.T) {
	fixtures := []string{
		`{}`,
		`{"ListenAddr":""}`,
		`{"downstream":""}`,
		`{"extPathPrefix":""}`,
		`{"intPathPrefix":""}`,
	}

	for testn, fix := range fixtures {
		func() {

			// Setup
			config, err := ioutil.TempFile("", "rapi-config")
			require.Nil(t, err)
			if os.Getenv("TEST_PRESERVE") == "" {
				defer os.Remove(config.Name())
			}

			_, err = config.WriteString(fix)
			require.Nil(t, err)
			err = config.Close()
			require.Nil(t, err)

			// Verification
			var conf *rapi.Config

			assert.NotPanics(t, func() { conf = rapi.NewConfigFile(config.Name()) }, "[%d:%s] parse errors", testn, fix)
			assert.NotNil(t, conf, "[%d:%s] invalid config", testn, fix)
		}()
	}
}
