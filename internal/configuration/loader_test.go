package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//TestLoadToStructs tests if data loads to structs correctly
func TestLoadToStructs(t *testing.T) {
	tests := []struct {
		name     string
		expected Configuration
	}{
		{
			name: "test",
			expected: Configuration{
				Api: Api{
					Version: "v2",
				},
				DataBase: DataBase{
					Type:     "postgres1",
					Host:     "localhost1",
					Port:     5431,
					User:     "postgres1",
					Password: "secret1",
					Name:     "postgres1",
				},
				Server: Server{
					Port: "8090",
				},
				Redis: Redis{
					Host: "127.0.0.1",
					Port: 6378,
				},
				JWT: JWT{
					Secret: "sedfqeSDFasfADSFASdfqe44FASFw4RRRW",
				},
			},
		},
	}

	actual, err := Load("test_configurations/configuration_test_pg.yml")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, err)
			require.NotNil(t, actual)
			assert.Equal(t, &tt.expected, actual)
		})
	}
}

//TestLoadMissingFile tests if Load function returns error if wrong file path set
func TestLoadConfigFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		err      string
	}{
		{
			name:     "expired",
			filePath: "configuration_test_missing.yml",
			err:      "Loaded non-existent file",
		},
		{
			name:     "invalid signature",
			filePath: "",
			err:      "Loaded empty file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := Load("test_configurations/" + tt.filePath)
			if err == nil {
				t.Fatalf("%v: %v", tt.err, err)
			}
		})
	}
}

//TestLoadWrongFile tests if full structured data loaded
func TestLoadWrongFile(t *testing.T) {
	_, err := Load("test_configurations/configuration_test_wrong.yml")
	if err != nil {
		t.Fatal("Loaded file with wrong configuration structure")
	}
}
