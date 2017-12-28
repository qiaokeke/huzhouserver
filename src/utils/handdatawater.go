package utils

import (
	"strconv"
	"fmt"
	"time"
	"daosql"
	"config"
)

/**
解析水表数据
 */
func ParseWaterData(bytes []byte)  float64{
	//－33H
	r :=0.0
	for j:=0;j<4;j++ {
		hexs := strconv.FormatInt(int64((bytes[j]-0x33)&0xff),16)
		//fmt.Println(hexs)
		i, err := strconv.Atoi(hexs)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		k := 0.0
		if j==0{
			k=0.01
		}
		if j==1{
			k=1
		}
		if j==2{
			k=100
		}
		if j==3{
			k=10000
		}
		r += k * float64(i)
	}
	fmt.Println(r)
	return r
}

/**
解析水表编号
 */
func ParseWaterId(bytes []byte)  string{
	str := ""
	for i:=5;i>=0;i--{
		hexs := strconv.FormatInt(int64((bytes[i])&0xff),16)
		if len(hexs)==1{
			str+="0"
		}
		str += hexs
	}
	fmt.Println(str)
	return str
}


//整理数据，将数据放入数据库
func HandDataWater(bytes []byte) {
	recvBytes:=bytes[0:18]
	data := ParseWaterData(recvBytes[12:16])
	id := ParseWaterId(recvBytes[1:7])
	daosql.InsertWater(id,fmt.Sprintf("%.2f",data),time.Now().Format("20060102150405"),config.ReadConfig().DataSource)
	daosql.InsertWater(id,fmt.Sprintf("%.2f",data),time.Now().Format("20060102150405"),config.ReadConfig().DataSource2)

}
