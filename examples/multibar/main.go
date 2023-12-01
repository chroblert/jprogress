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

	bar1 := jprogress.AddDefaultBar(20, "1")
	bar2 := jprogress.AddDefaultBar(40, "2")
	bar3 := jprogress.AddDefaultBar(3000, "3")

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
			if !bar.Incr() {
			}
			jprogress.RemoveBarOnComplete(bar)
		}
		if !bar.Incr() {
			//jlog.Info("删除bar2")
			//jprogress.RemoveBar(bar)
		}
	}, bar2).CDO()
	a.Init("3").CAdd(func(bar *jprogress.Bar) {
		for i := 0; i < bar.Total; i++ {
			time.Sleep(waitTime)
			if !bar.Incr() {
				//jlog.Info("删除bar3")
				//jprogress.RemoveBar(bar)
			}
		}
	}, bar3).CDO()
	a.Wait()

}
