package daosql

import (
	"fmt"
	"database/sql"
)

/**
插入到数据库
 */
func InsertWater(W_ADDRESS string,W_READINGS string,W_TIME string,dataSorce string)  {
	defer func() {
		if x := recover();x!=nil{
			fmt.Println("insert,water,err,flag")
			return
		}
	}()
	sqlstring := "insert into tbl_water_info(W_ADDRESS,W_READINGS,W_TIME)"
	sqlstring += "values("+W_ADDRESS+","+W_READINGS+","+W_TIME+")"
	fmt.Println(sqlstring)

	db, e:= sql.Open("mysql", dataSorce)
	db.SetMaxIdleConns(1000)
	defer db.Close()
	if e!=nil{
		fmt.Print(e)
		return
	}

	row,err2:=db.Query(sqlstring)
	defer row.Close()
	if err2!=nil{
		fmt.Println(err2)
	}
	fmt.Print("写入成功")
}