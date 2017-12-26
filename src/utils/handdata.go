package utils

import (
	"time"
	"fmt"
	"config"
	"math"
	"container/list"
	"encoding/binary"
	"daosql"
)

//整理数据，将数据放入数据库
func HandData(num byte)  bool{
	value1,ok1 :=config.SMap1[num]
	value2,ok2 :=config.SMap2[num]
	value3,ok3 :=config.SMap3[num]
	if ok1&&ok2&&ok3{
		l1:=parseBytes(value1)
		l2:=parseBytes(value2)
		l1.PushBackList(l2)
		l1.PushBack(num)
		l1.PushBack(time.Now().Format("20060102150405"))
		l3:=parseBytes2int(value3)
		l1.PushBackList(l3)
		for e := l1.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value) //输出list的值,01234
		}
		//写入数据库
		daosql.Insert(l1,config.ReadConfig().DataSource)
		daosql.Insert(l1,config.ReadConfig().DataSource2)
		defer func() {
			if x := recover();x!=nil{
				fmt.Println("inset,err,flag")
				return
			}
		}()
		//fmt.Print(l1)
		//删除键值
		delete(config.SMap1,num)
		delete(config.SMap2,num)
		delete(config.SMap3,num)
		fmt.Println("读取完毕一个")
		return true
	}
	return false
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func parseBytes2int(bytes []byte) *list.List{
	l := list.New()
	n:= len(bytes)
	for i:=0;i<n;i+=4{
		curBytes := bytes[i:i+4]
		fmt.Printf("% X\n",curBytes)
		bits:=binary.BigEndian.Uint32(curBytes)
		fmt.Println(bits)
		l.PushBack(float32(bits)/1000.0)
	}
	return l
}


func parseBytes(bytes []byte)  *list.List{
	l := list.New()
	n := len(bytes)
	for i:=0;i<n;i+=4{
		curBytes := bytes[i:i+4]
		fmt.Printf("% X\n",curBytes)
		curFolat := ByteToFloat32(curBytes)
		fmt.Println(curFolat)
		l.PushBack(curFolat)
	}
	return l
}
