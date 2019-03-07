package helpers

import (
	"io/ioutil"

	"github.com/enfipy/cronpub/src/config"

	yaml "gopkg.in/yaml.v2"
)

func GetSettings(path string) *config.CronpubSettings {
	settings := new(config.CronpubSettings)

	yamlFile, err := ioutil.ReadFile(path)
	PanicOnError(err)
	err = yaml.Unmarshal(yamlFile, settings)
	PanicOnError(err)

	return settings
}
