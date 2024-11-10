package jprogress

import (
	"bytes"
	"fmt"
	"github.com/chroblert/jlog"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestStoppingPrintout(t *testing.T) {
	progress := New()
	progress.SetRefreshInterval(time.Millisecond * 10)

	var buffer = &bytes.Buffer{}
	progress.SetOut(buffer)

	bar := progress.AddBar(100)
	progress.Start()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for i := 0; i <= 80; i = i + 10 {
			bar.Set(i)
			time.Sleep(time.Millisecond * 5)
		}

		wg.Done()
	}()

	wg.Wait()

	progress.Stop()
	fmt.Fprintf(buffer, "foo")

	var wantSuffix = "[======================================================>-------------]\nfoo"

	if !strings.HasSuffix(buffer.String(), wantSuffix) {
		t.Errorf("Content that should be printed after stop not appearing on buffer.")
	}
}

func TestSize0Bar(t *testing.T) {
	t.Run("size 0 bar", func(t *testing.T) {
		jp := New()
		jp.Start()
		defer jp.Stop()
		bar := jp.Default64(0, "size 0")
		for _, v := range []string{"1", "2", "3"} {
			bar.Add(1)
			time.Sleep(5 * time.Second)
			jlog.Debug(v)
		}

	})

	t.Run("stop后start", func(t *testing.T) {
		Start()
		bar := Default64(10, "bar1")
		for k := 0; k < 10; k++ {
			time.Sleep(10 * time.Millisecond)
			bar.Add(1)
		}
		Stop()
		bar.Set(9)
		bar.Incr()
		Start()
		bar2 := Default64(12, "bar2")
		for k := 0; k < 10; k++ {
			time.Sleep(10 * time.Millisecond)
			bar2.Add(1)
		}
		jlog.Debug("before stop in bar1 x")
		Stop()
	})
	t.Run("remove bar", func(t *testing.T) {
		Start()
		bar := Default64(10, "bar1")
		for k := 0; k < 10; k++ {
			time.Sleep(10 * time.Millisecond)
			bar.Add(1)
		}
		Stop()
		Stop()
		Start()
		bar2 := Default64(12, "bar2")
		for k := 0; k < 12; k++ {
			time.Sleep(10 * time.Millisecond)
			if k > 5 {
				RemoveBar(bar)
			}
			bar2.Add(1)
		}
	})
	t.Run("sleep", func(t *testing.T) {
		Start()
		bar := Default64(10, "bar1")
		for k := 0; k < 10; k++ {
			//time.Sleep(10 * time.Millisecond)
			bar.Add(1)
			jlog.Debug(k)
			time.Sleep(20 * time.Second)
		}
		bar.Finish()
		Stop()
	})
}
