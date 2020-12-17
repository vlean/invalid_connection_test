package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB
func init()  {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3308)/test")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Second * 20)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	DB = db
}

func DbStat()  {
	for  {
		stats := DB.Stats()
		log.Printf("idle:%d inUse:%d maxIdleClosed:%d maxLifetimeClosed:%d openConnections:%d waitCount:%d waitSeconds:%2.f",
			stats.Idle, stats.InUse, stats.MaxIdleClosed, stats.MaxLifetimeClosed, stats.OpenConnections, stats.WaitCount, stats.WaitDuration.Seconds())
		time.Sleep(time.Second)
	}
}

func Query() (int, error)  {
	query, err := DB.Query("select count(*) as total from test")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer query.Close()

	var total int
	query.Scan(&total)
	return total, nil
}


func SlowQuery() (int, error)  {
	rows, err := DB.Query("select sleep(1),count(*) as total  from test")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var sleep int
	var total int
	rows.Next()
	rows.Scan(&sleep,  &total)
	return total, nil
}

func Insert(name string)  {
	DB.Exec("insert into test (name) value (?)", name)
}
