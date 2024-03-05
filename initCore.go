package core

import (
	"connection_test/core/config"
)

//pushing config to struct

func init() {

	config.LoadConf("conf.txt")

}
