package yaml

import (
	// "fmt"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/aceberg/LinuxExecBot/internal/check"
	"github.com/aceberg/LinuxExecBot/internal/models"
)

// Read - read .yaml file to []struct
func Read(path string) models.Data {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var data models.Data

	err = yaml.Unmarshal(file, &data)
	check.IfError(err)

	return data
}
