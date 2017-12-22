package test

import "testing"
import (
	"config"
	"strconv"
)

func Test_ReadConfig(t *testing.T)  {
	config.ConfigPath="huzhou6002.json"

	config :=config.ReadConfig()
	m1:=config.MeterIds

	t.Log(strconv.Atoi(m1["1"]))

	t.Log(config.Port)
	t.Log(config.MeterIds)
	t.Log(config.Cmds)
	t.Log(config.DataSource)

}
