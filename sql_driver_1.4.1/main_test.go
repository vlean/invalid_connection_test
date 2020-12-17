package main

import (
	"git.qutoutiao.net/vlean/sd14/model"
	"log"
	"testing"
)

func Benchmark_invalid_connection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := model.SlowQuery()
		if err != nil {
			log.Fatal(err)
		}
	}
}

