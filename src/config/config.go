package config

import (
	"io/ioutil"
	"encoding/json"
	"runtime"
	"strings"
	"log"
)

/** 配置文件路径
 */
var ConfigPath string

//电表的3个区段
var SMap1 map[byte] []byte
var SMap2 map[byte] []byte
var SMap3 map[byte] []byte


type Config struct {
	Port string
	Cmds [][][]byte
	MeterIds map[string] string
	DataSource string
	DataSource2 string
}

func InitData(path string)  {
	ConfigPath = path
	SMap1 = make(map[byte] []byte)
	SMap2 = make(map[byte] []byte)
	SMap3 = make(map[byte] []byte)
}

/**
	读取配置文件
 */
func ReadConfig()  Config{
	/**
	判断操作系统类型
	 */
	os := runtime.GOOS
	path :=""
	if strings.Compare(os,"windows")==0{
		path = "C:\\work\\go\\hzserver\\huzhouserver\\src\\config\\"+ConfigPath
	}else {
		path = "/root/work/go/huzhouserver/src/config/"+ConfigPath
	}

	file,err:=ioutil.ReadFile(path)
	if err!=nil{
		log.Println("config file err:",err)
	}
	config := Config{}
	err2 := json.Unmarshal(file,&config)
	if err2!=nil{
		log.Println("json err:",err2)
	}
	return config
}


