package samples

import (
	"fmt"
	"os"
	"testing"

	"github.com/subratohld/config"
)

func TestConfigReader(t *testing.T) {
	// It will read from environment variable because no config file is passed
	// It will look for FILE_PATh variable
	configFile := os.Getenv("CONFIG_FILE")

	config, err := config.New(configFile)
	fmt.Println(err)
	fmt.Println(config.GetString("server.port"))
}
