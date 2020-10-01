package  main
import (
	"runtime/debug"
	"time"
)
func main()  {
	num := 500000
	var bigmap = make(map[int]*[512]float32)
	for i := 0;i < num;i++ {
		bigmap[i] = &[512]float32{float32(i)}
	}

	println(len(bigmap))
	time.Sleep(15e9)
	for i := 0;i < num;i++ {
		delete(bigmap,i)
	}

	debug.FreeOSMemory()
	time.Sleep(1000e9)
}