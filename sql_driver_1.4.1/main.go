package main

import (
	"git.qutoutiao.net/vlean/sd14/model"
	"log"
	"sync"
	"time"
)

func main() {
	go model.DbStat()
	start := time.Now()
	for {
		group := sync.WaitGroup{}
		group.Add(5)
		for i:=0; i<5;i++ {
			go func(v int) {
				query, err := model.SlowQuery()
				log.Printf("query done: err:%+v, result:%d, gort:%d", err, query, v)
				group.Done()
			}(i)
		}
		group.Wait()
		log.Printf("run %.2f second", time.Since(start).Seconds())
		time.Sleep(time.Second*6)
	}
}