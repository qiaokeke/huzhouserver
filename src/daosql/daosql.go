package daosql

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"container/list"
	"fmt"
	"config"
)

func Insert(li *list.List,dataSource string)  {
	defer func() {
		if x := recover();x!=nil{
			log.Println("inset,err,flag")
			return
		}
	}()
	//sq_command := "insert into tbl_power_info_v2(P_A_DIANYA,P_B_DIANYA,P_C_DIANYA,P_UAB_XIANDIANYA,P_UBC_XIANDIANYA,P_UCA_XIANDIANYA,P_A_DIANLIU,P_B_DIANLIU,P_C_DIANLIU,P_A_YGGL,P_B_YGGL,P_C_YGGL,P_HXYGGL,P_A_WGGL,P_B_WGGL,P_C_WGGL,P_HXWGGL,P_A_SZGL,P_B_SZGL,P_C_SZGL,P_HXSZGL,P_A_GLYS,P_B_GLYS,P_C_GLYS,P_HXGLYS,P_DWPL,P_BY_KwhZ,P_BY_KwhJ,P_BY_KwhF,P_BY_KwhP,P_BY_KwhG,P_BY_HKwhZ,P_BY_HKwhJ,P_BY_HKwhF,P_BY_HKwhP,P_BY_HKwhG, P_BY_KvarhZ,P_BY_KvarhJ,P_BY_KvarhF,P_BY_KvarhP,P_BY_KvarhG, P_BY_HKvarhZ,P_BY_HKvarhJ,P_BY_HKvarhF,P_BY_HKvarhP,P_BY_HKvarhG, P_SY_KwhZ,P_SY_KwhJ,P_SY_KwhF,P_SY_KwhP,P_SY_KwhG, P_SY_HKwhZ,P_SY_HKwhJ,P_SY_HKwhF,P_SY_HKwhP,P_SY_HKwhG, P_SY_KvarhZ,P_SY_KvarhJ,P_SY_KvarhF,P_SY_KvarhP,P_SY_KvarhG, P_SY_HKvarhZ,P_SY_HKvarhJ,P_SY_HKvarhF,P_SY_HKvarhP,P_SY_HKvarhG, P_SSY_KwhZ,P_SSY_KwhJ,P_SSY_KwhF,P_SSY_KwhP,P_SSY_KwhG, P_SSY_HKwhZ,P_SSY_HKwhJ,P_SSY_HKwhF,P_SSY_HKwhP,P_SSY_HKwhG, P_SSY_KvarhZ,P_SSY_KvarhJ,P_SSY_KvarhF,P_SSY_KvarhP,P_SSY_KvarhG, P_SSY_HKvarhZ,P_SSY_HKvarhJ,P_SSY_HKvarhF,P_SSY_HKvarhP,P_SSY_HKvarhG,P_TIME,P_CODE,P_ZXYGDN,P_FXYGDN,P_ZXWGDN,P_FXWGDN)values(?, ?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?, ?, ?,?, ?,?,?, ?,?, ?,?, ?, ?,?, ?,?,?, ?,?, ?,?, ?, ?,?, ?,?,?, ?,?, ?,?, ?, ?,?, ?,?,?, ?,?, ?,?, ?, ?,?, ?,?,?, ?,?,?,?, ?,?)"
	sqlstring := "insert into tbl_power_info_v2(P_A_DIANYA,P_B_DIANYA,P_C_DIANYA,P_UAB_XIANDIANYA,P_UBC_XIANDIANYA,P_UCA_XIANDIANYA,P_A_DIANLIU,P_B_DIANLIU,P_C_DIANLIU,P_A_YGGL,P_B_YGGL,P_C_YGGL,P_HXYGGL,P_A_WGGL,P_B_WGGL,P_C_WGGL,P_HXWGGL,P_A_SZGL,P_B_SZGL,P_C_SZGL,P_HXSZGL,P_A_GLYS,P_B_GLYS,P_C_GLYS,P_HXGLYS,P_DWPL,P_BY_KwhZ,P_BY_KwhJ,P_BY_KwhF,P_BY_KwhP,P_BY_KwhG,P_BY_HKwhZ,P_BY_HKwhJ,P_BY_HKwhF,P_BY_HKwhP,P_BY_HKwhG, P_BY_KvarhZ,P_BY_KvarhJ,P_BY_KvarhF,P_BY_KvarhP,P_BY_KvarhG, P_BY_HKvarhZ,P_BY_HKvarhJ,P_BY_HKvarhF,P_BY_HKvarhP,P_BY_HKvarhG, P_SY_KwhZ,P_SY_KwhJ,P_SY_KwhF,P_SY_KwhP,P_SY_KwhG, P_SY_HKwhZ,P_SY_HKwhJ,P_SY_HKwhF,P_SY_HKwhP,P_SY_HKwhG, P_SY_KvarhZ,P_SY_KvarhJ,P_SY_KvarhF,P_SY_KvarhP,P_SY_KvarhG, P_SY_HKvarhZ,P_SY_HKvarhJ,P_SY_HKvarhF,P_SY_HKvarhP,P_SY_HKvarhG, P_SSY_KwhZ,P_SSY_KwhJ,P_SSY_KwhF,P_SSY_KwhP,P_SSY_KwhG, P_SSY_HKwhZ,P_SSY_HKwhJ,P_SSY_HKwhF,P_SSY_HKwhP,P_SSY_HKwhG, P_SSY_KvarhZ,P_SSY_KvarhJ,P_SSY_KvarhF,P_SSY_KvarhP,P_SSY_KvarhG, P_SSY_HKvarhZ,P_SSY_HKvarhJ,P_SSY_HKvarhF,P_SSY_HKvarhP,P_SSY_HKvarhG,P_CODE,P_TIME,P_ZXYGDN,P_FXYGDN,P_ZXWGDN,P_FXWGDN)"
	//sqlstring += log.Sprintf("values(%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s, %s,%s,%s, %s,%s,%s,%s, %s,%s)",li)
	sqlstring += "values("
	for e :=li.Front();e!=nil;e=e.Next(){
		sqlstring += fmt.Sprint(e.Value) +","
	}
	sqlstring = sqlstring[:len(sqlstring)-1] + ")"
	log.Println(sqlstring)

	//datasource:=config.ReadConfig().DataSource;
	// sqlstring += log.Sprintf("values(%v)",li)
	db, e:= sql.Open("mysql", dataSource)
	db.SetMaxIdleConns(1000)
	defer db.Close()
	if e!=nil{
		log.Print(e)
	}

	row,err2:=db.Query(sqlstring)
	defer row.Close()
	if err2!=nil{
		log.Println(err2)
	}
	err3 := row.Err()
	if err3!=nil{
		log.Printf(err3.Error())
	}
	log.Print("写入成功")
}
