package dao

import (
	"net"
	"time"
	"config"
	"log"
	"utils"
	"bufio"
	"fmt"
	"bytes"
	"strconv"
)

/**
	处理读出消息
 */
func HandleRead(conn net.Conn)  {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for{
		peekbytes,err := reader.Peek(3)
		log.Printf("recv:% X\n",peekbytes)
		if err!=nil{
			log.Println("peek err:",err)
			break
		}

		if bytes.Contains(peekbytes,[]byte{0x03,0x68}){
			//读取 字节
			dianbiaoNum,_ := reader.ReadByte()

			if dianbiaoNum==1{
				a,_:= strconv.Atoi(config.ReadConfig().MeterIds["1"]);
				dianbiaoNum=byte(a)
			}
			//丢弃5个字节
			log.Println(dianbiaoNum)
			reader.Discard(4)
			recvbytes := make([]byte,0x68)
			reader.Read(recvbytes)
			log.Printf("% X\n",recvbytes)
			config.SMap1[dianbiaoNum] = recvbytes
			continue
		}
		if bytes.Contains(peekbytes,[]byte{0x03,240}){

			//读取 字节
			dianbiaoNum,_ := reader.ReadByte()
			if dianbiaoNum==1{
				a,_:= strconv.Atoi(config.ReadConfig().MeterIds["1"]);
				dianbiaoNum=byte(a)
			}
			//丢弃5个字节
			log.Println(dianbiaoNum)
			reader.Discard(4)
			recvbytes := make([]byte,0,240)
			tmp := make([]byte,240)
			total :=0
			for{
				n,err := reader.Read(tmp)
				if err!=nil{
					break
				}
				recvbytes = append(recvbytes,tmp[:n]...)
				total += n;
				if total>=240{
					break
				}
			}
			fmt.Println("totalsize",len(recvbytes))
			fmt.Printf("% X\n",recvbytes)
			config.SMap2[dianbiaoNum] = recvbytes
			continue
		}
		if bytes.Contains(peekbytes,[]byte{0x03,16}){
			//读取 字节
			dianbiaoNum,_ := reader.ReadByte()
			if dianbiaoNum==1{
				a,_:= strconv.Atoi(config.ReadConfig().MeterIds["1"]);
				dianbiaoNum=byte(a)
			}
			//丢弃5个字节
			fmt.Print(dianbiaoNum)
			reader.Discard(4)
			recvbytes := make([]byte,16)
			reader.Read(recvbytes)
			fmt.Printf("% X\n",recvbytes)
			config.SMap3[dianbiaoNum] = recvbytes
			//处理，将数据送入数据库
			go utils.HandData(dianbiaoNum)
			continue
		}

		//将字节取出不做处理
		recvbytes :=make([]byte,reader.Buffered())
		reader.Read(recvbytes)
		fmt.Printf("heart:% X\n",recvbytes)
		time.Sleep(1*time.Second)
	}




}

/**
	处理写入消息
 */
func HandleWrite(conn net.Conn)  {
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