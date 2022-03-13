# tesla-take-home-challenge

## Table of Contents

- [tesla-take-home-challenge](#tesla-take-home-challenge)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
  - [Usage](#usage)
  - [cache](#cache)
  - [build](#build)
  - [Topic 2](#topic-2)
  - [Topic 3](#topic-3)

## Requirements

requires the following to run:

- go ^1.17

## Usage

```
    go run main.go
```


## cache

可以删除.cache文件

## build

```
   go build -o main main.go
```


## Topic 2

>Please describe how would you monitor above inventory management program after it 
was deployed in production environment. (Please write your answer in README of your 
project)
 1. 健康检查/探活 如果部署在k8s上，可以配置健康检查，或者可以自行写一个探活程序。
 2. 指标监控：即各种指标监控，比如基础资源指标，服务性能指标，业务的调用指标。
 3. 日志：服务的运行日志监控。
 4. 调用链和链路追踪。
   
   以上还需要配合事件通知，故障响应

## Topic 3
>Please describe how would you optimize the above inventory management program. 
(Please write your answer in README of your project).

1. N(库存)数据缓存在本地，单体性能有上限，横向扩展困难，后续可以把数据缓存在redis上，提高服务的qps
2. S（卖出的车）当前是通过互斥锁的机制实现内存安全的，但会损耗服务性能。可以通过kafka 等mq中间件 解耦出来，提高性能上限
