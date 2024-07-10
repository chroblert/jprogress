package main

import (
	"github.com/chroblert/jasync"
	"github.com/chroblert/jlog"
	"github.com/chroblert/jprogress"
	"time"
)

func main() {
	waitTime := time.Millisecond * 100
	jprogress.Start()

	//var wg sync.WaitGroup
	jlog.Info("kkkkktest")
	bar1 := jprogress.Default(20, "1")
	bar2 := jprogress.Default(40, "2")
	bar3 := jprogress.Default(3000, "3")

	//wg.Add(1)
	a := jasync.NewAR(3)
	a.Init("i").CAdd(func(bar *jprogress.Bar) {
		for i := int64(0); i < bar.Total; i++ {
			time.Sleep(waitTime)
			bar.Incr()
		}
	}, bar1).CDO()
	a.Init("2").CAdd(func(bar *jprogress.Bar) {
		for i := int64(0); i < bar.Total; i++ {
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
		for i := int64(0); i < bar.Total; i++ {
			time.Sleep(waitTime)
			if !bar.Incr() {
				//jlog.Info("删除bar3")
				//jprogress.RemoveBar(bar)
			}
		}
	}, bar3).CDO()
	a.Wait()

}
