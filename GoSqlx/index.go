package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
	"time"
)

func main() {

	// MYSQL dsn格式：{username}:{password}@({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	var dsn = "root:root@(localhost:3520)/test?charset=utf8&parseTime=True&loc=Local"
	var db *sqlx.DB
	var err error

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 建表
	/*var ddl = "CREATE TABLE `place` (`id` BIGINT NOT NULL AUTO_INCREMENT, `country` VARCHAR(128), `city` VARCHAR(128), `tel_code` VARCHAR(128), create_time DATETIME DEFAULT CURRENT_TIMESTAMP, update_time DATETIME DEFAULT CURRENT_TIMESTAMP, remark TEXT, PRIMARY KEY (`id`) ) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;"
	_, err = db.Exec(ddl)
	if err != nil {
		panic(err)
	}*/

	// 插入
	/*var insertSql = `insert into place(country, city, tel_code) values (?, ?, ?)`
	res, err := db.Exec(insertSql, "中国", "香港", "852")
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	fmt.Println(reflect.TypeOf(id), id)*/ // int64 2

	// 更新
	/*var updateSql = "update `place` set `tel_code` = ?, `city` = ? where `id` = ?"
	res, err := db.Exec(updateSql, "123", "xyz", 1)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Println(rowsAffected)*/

	// 查询数据
	// 通过Get和Select查询数据，Get单记录查询，Select多记录查询
	//var p Place
	//var getSql = "select * from `place` limit ?"
	//err = db.Get(&p, getSql, 1)
	//fmt.Println(p)
	//fmt.Println(p.Id, p.Remark.String == "", p.Remark.Valid)
	//var total int
	//getSql = "select count(id) from `place`"
	//err = db.Get(&total, getSql)
	//fmt.Println(total)

	/*var placeList []Place
	var selectSql = "select * from `place` where `id` > 0"
	err = db.Select(&placeList, selectSql)
	fmt.Println(placeList)
	place := placeList[0]
	fmt.Println(place.CreateTime.Format("2006-01-02 15:04:05"))

	var cities []string
	selectSql = "select city from `place`"
	err = db.Select(&cities, selectSql)
	fmt.Println(cities)*/

	// 复杂查询
	var complexSql = "SELECT u.`name`, u.`age`, u.`gender`, u.create_time AS createTime, u.update_time AS updateTime,\n(SELECT p.country FROM place p WHERE p.id = u.`place_id`) AS country,\n(SELECT p.city FROM place p WHERE p.id = u.`place_id`) AS city,\n(SELECT p.tel_code FROM place p WHERE p.id = u.`place_id`) AS telCode,\n(SELECT p.remark FROM place p WHERE p.id = u.`place_id`) AS remark\nFROM `user` u"
	fmt.Println("\n------------------------------------------ 自定义struct接收-Begin ------------------------------------------")
	var res []ComplexRes
	err = db.Select(&res, complexSql)
	fmt.Println(res, "\n------------------------------------------ 自定义struct接收-End ------------------------------------------")

	fmt.Println("\n=========================== map切片接收-Begin ===========================")
	var mapRes []map[string]interface{}
	rows, err := db.Queryx(complexSql)
	for rows.Next() {
		var m = map[string]interface{}{}
		err = rows.MapScan(m)
		mapRes = append(mapRes, m)
	}
	fmt.Println(mapRes, "\n=========================== map切片接收-End ===========================")

	var raw interface{}
	var idxMap = mapRes[0]
	raw = idxMap["age"]
	fmt.Println("初始数据类型：", reflect.TypeOf(raw), "\t解析值：", string(raw.([]uint8)))
	city := string(idxMap["city"].([]uint8))
	createTime := idxMap["createTime"]
	fmt.Println(reflect.TypeOf(city), city)
	fmt.Println(reflect.TypeOf(createTime), createTime)

}

// Place 数据库结构体声明时候只允许“多”字段，不允许“少”字段，字段映射时候，默认会将首字母大写
type Place struct {
	Id int64 `db:"id"`
	Country string `db:"country"`
	City string `db:"city"`
	TelCode string `db:"tel_code"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	//Remark string
	Remark sql.NullString // 对于数据库表中允许为空的字段建议使用sql.NullString进行字段映射，否则可能由于映射失败而查询失败
}
// Go的字符串会给初始值""，这点和Java不一样，Java的字符串是一个对象，默认为null

type ComplexRes struct {
	Name sql.NullString `db:"name"`
	Age sql.NullInt16 `db:"age"`
	Gender sql.NullInt16 `db:"gender"`
	CreateTime sql.NullTime `db:"createTime"`
	UpdateTime sql.NullTime `db:"updateTime"`
	Country sql.NullString `db:"country"`
	City sql.NullString
	TelCode sql.NullString `db:"telCode"`
	Remark sql.NullString
}

