package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"tesla-take-home-challenge/internal/constant"
	"tesla-take-home-challenge/pkg"
	"time"
)

type Inventory struct {
	N     []int64    // 库存
	lockN sync.Mutex //互斥锁
	R     int        // 销售速率
	X     int        // 小时
	S     sync.Map   //卖出的车
	file  string     // 缓存文件
}

// 入库一辆车
func (i *Inventory) AddCar() (err error) {
	i.lockN.Lock() //加锁

	// 判断库存的size 是否小于等于 销售速率*时间
	if int(len(i.N)) <= i.R*i.X {
		// 入库一辆车
		i.N = append(i.N, time.Now().UnixNano())
		// 写入文件
		err = i.Write(i.N)
	}
	i.lockN.Unlock() //解锁
	return
}

// 卖出一辆车
func (i *Inventory) SellCar() (car int64, err error) {
	i.lockN.Lock() //加锁

	// 判断当前库存是否还有车
	if len(i.N) > 0 {
		// 卖出一辆车
		car = i.N[0]
		i.N = i.N[1:]
		//卖出的车 存入s中
		i.S.Store(time.Now().UnixNano(), car)
		err = i.Write(i.N)
	} else {
		err = fmt.Errorf("No cars in inventory")
	}
	i.lockN.Unlock() //解锁
	return
}

// 获取库存
func (i *Inventory) GetN() (n int) {
	i.lockN.Lock() //加锁
	n = len(i.N)
	i.lockN.Unlock() //解锁
	return
}

// 库存写入文件
func (i *Inventory) Write(n []int64) (err error) {
	var fileContent []byte
	fileContent, err = json.Marshal(n)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(i.file, fileContent, 0666)
	if err != nil {
		return
	}
	return
}

// 从文件读取库存
func (i *Inventory) Read() (err error) {
	var fileContent []byte
	fileContent, err = ioutil.ReadFile(i.file)
	if err != nil {
		return
	}
	_ = json.Unmarshal(fileContent, &i.N)
	return
}

// 计算最近一小时的销售速率
func (i *Inventory) CalR() {
	now := time.Now().UnixNano()
	totalSell := 0
	i.S.Range(func(key, value interface{}) bool {
		car, ok := key.(int64)

		if ok {
			if now-car >= constant.ExpirationTime {
				i.S.Delete(car)
			} else {

				totalSell += 1
			}
		}
		return true
	})
	i.R = totalSell
	return
}

func NewInventory() (i *Inventory) {
	i = &Inventory{
		N:     []int64{},
		lockN: sync.Mutex{},
		R:     0,
		X:     2,
		S:     sync.Map{},
		file:  ".cache",
	}
	pkg.CreateFile(i.file)
	err := i.Read()
	if err != nil {
		panic(err)
	}

	return
}
