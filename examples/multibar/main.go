package main

import (
	"github.com/chroblert/jasync"
	"github.com/chroblert/jprogress"
	"time"
)

func main() {
	waitTime := time.Millisecond * 100
	jprogress.Start()

	//var wg sync.WaitGroup

	bar1 := jprogress.AddBar(20).AppendCompleted().PrependElapsed()
	bar2 := jprogress.AddBar(40).AppendCompleted().PrependElapsed()
	bar3 := jprogress.AddBar(3000).AppendCompleted().PrependSlashNum().PrependDesc("Test").AppendETA()

	//wg.Add(1)
	a := jasync.NewAR(3)
	a.Init("i").CAdd(func(bar *jprogress.Bar) {
		for i := 0; i < bar.Total; i++ {
			time.Sleep(waitTime)
			bar.Incr()
		}
	}, bar1).CDO()
	a.Init("2").CAdd(func(bar *jprogress.Bar) {
		for i := 0; i < bar.Total; i++ {
			time.Sleep(waitTime)
			bar.Incr()
		}
	}, bar2).CDO()
	a.Init("3").CAdd(func(bar *jprogress.Bar) {
		for i := 0; i < bar.Total; i++ {
			time.Sleep(waitTime)
			bar.Incr()
		}
	}, bar3).CDO()
	a.Wait()

}
