package dao

import (
"net"
"time"
"config"
"log"
"utils"
"fmt"
"bytes"
)

/**
	处理读出消息
 */
func HandleReadWater(conn net.Conn)  {
	defer conn.Close()
	for{
		recvBytes := make([]byte,0,100)
		tmp := make([]byte,100)
		n,err := conn.Read(tmp)
		if err!=nil{
			fmt.Println(err)
			break
		}
		if n<=0{
			continue
		}
		fmt.Printf("% X\n",tmp[0:n])

		if bytes.HasPrefix(tmp,[]byte{0x68}){
			recvBytes = append(recvBytes,tmp[:n]...)
			if n < 18{
				for{
					i,err := conn.Read(tmp)
					if err!=nil{
						fmt.Println(err)
						break
					}
					if i<=0{
						continue
					}
					fmt.Printf("% X\n",tmp[0:i])
					recvBytes = append(recvBytes,tmp[:i]...)
					n +=i
					if n>=18{
						break
					}
				}
			}
			fmt.Printf("% X\n",recvBytes)
			go utils.HandDataWater(recvBytes)
		}
	}

}

/**
	处理写入消息
 */
func HandleWriteWater(conn net.Conn)  {
	defer conn.Close()
	for{
		cmds := config.ReadConfig().Cmds
		for i:=0;i<len(cmds);i++{
			for j:=0;j<len(cmds[i]);j++ {
				log.Printf("write:% X\n",cmds[i][j])
				_, e := conn.Write(cmds[i][j])
				if e != nil {
					log.Println("write err:", e)
					return;
				}
				time.Sleep(40*time.Second)
			}
		}
		time.Sleep(5*time.Second)
	}
}
