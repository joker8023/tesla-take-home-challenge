package main

import (
	"log"
	"net/http"
	"tesla-take-home-challenge/internal/entity"
	"tesla-take-home-challenge/internal/handler"
	"time"
)

func main() {
	i := entity.NewInventory()

	// 每秒钟 检查库存够不够
	go func() {
		for {
			_ = i.AddCar()
			time.Sleep(time.Second)
		}
	}()

	// 每分钟更新销售率
	go func() {
		for {
			i.CalR()
			time.Sleep(time.Minute)
		}
	}()

	http.HandleFunc("/age", handler.HandlerAge(i))
	http.HandleFunc("/car", handler.HandlerCar(i))
	http.HandleFunc("/rate", handler.HandlerRate(i))
	http.HandleFunc("/buffer", handler.HandlerBuffer(i))

	log.Printf(" [*] http serve Waiting for connection:%d", 8081)

	http.ListenAndServe(":8081", nil)
}
