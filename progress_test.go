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
}
