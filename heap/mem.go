package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func bigBytes() *[]byte {
	s := make([]byte, 100000000)
	return &s
}

func main() {
	var wg sync.WaitGroup

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}
	runtime.
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)
/*
   RES – 将会显示在当前时刻进程的内存用量，但可能不包含任何尚未换入或已经换出的页面。
   mem.Alloc – 已经被配并仍在使用的字节数
   mem.TotalAlloc – 从开始运行到现在分配的内存总数
   mem.HeapAlloc – 堆当前的用量
   mem.HeapSys – 包含堆当前和已经被释放但尚未归还操作系统的用量
 */
	wg.Add(1)
	wg.Wait()

}
