package interview

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	题目：有多个 site 类型，每个都有 query 方法，实现两个方法 queryAll 和 queryOne
	     遍历一个 site 类型的切片，用协程执行里边每个元素的 query 方法
		 区别是 queryAll 要保证每个 site 都执行了，而 queryOne 只要有一个执行就结束
	考察点：多态，Interface，WaitGroup，Context
*/

func Solution1() string {
	sites := make([]Site, 0)
	var baidu Site = &Baidu{}
	sites = append(sites, baidu)
	var weibo Site = &Weibo{}
	sites = append(sites, weibo)
	queryAll(sites)
	queryOne(sites)
	return "done!"
}

// 准备原材料
type Site interface {
	Query()
}

type Baidu struct {
	Site string
}

type Weibo struct {
	Site string
}

func (b *Baidu) Query() {
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("query baidu site")
}

func (w *Weibo) Query() {
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("query weibo site")
}

var wg sync.WaitGroup

func queryAll(sites []Site) {
	for _, v := range sites {
		wg.Add(1)
		go func(v Site) {
			v.Query()
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func queryOne(sites []Site) {
	done := make(chan string, len(sites))
	root, cancel := context.WithCancel(context.Background())
	for _, v := range sites {
		go func(v Site) {
			sub, cancel := context.WithCancel(root)
			defer cancel()
			select {
			case <-sub.Done():
			default:
				v.Query()
				done <- "done"
			}
		}(v)
	}
	<-done
	cancel()
}
