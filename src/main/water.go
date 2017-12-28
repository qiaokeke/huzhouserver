package main

import (

	"dao"
	"config"
	"log"
	"net"
	"os"
)

func main()  {
	/**
	全局map,保存读出的数据
	 */
	log.Println("water")
	configPath:=os.Args[1]
	log.Println(configPath)

	config.InitData(configPath)
	/**
	获取参数文件dd
	 */

	port := config.ReadConfig().Port
	log.Println(port)
	listener,err := net.Listen("tcp",""+":"+port)
	defer listener.Close()
	if err!=nil{
		log.Println("listen err:",err)
		os.Exit(1)
	}
	log.Println("listening on:",port)
	for{
		conn,err := listener.Accept()
		if err !=nil{
			log.Println("accept err:",err)
			break
		}
		log.Println("connect from :",conn.RemoteAddr(),conn.LocalAddr())
		go dao.HandleReadWater(conn)
		go dao.HandleWriteWater(conn)
	}


}
