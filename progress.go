package jprogress

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

// Out is the default writer to render progress bars to
var Out = os.Stdout

// RefreshInterval in the default time duration to wait for refreshing the output
var RefreshInterval = time.Millisecond * 10

// defaultProgress is the default progress
var defaultProgress = New()

// Progress represents the container that renders progress bars
type Progress struct {
	// Out is the writer to render progress bars to
	Out io.Writer

	// Width is the width of the progress bars
	Width int

	// Bars is the collection of progress bars
	Bars []*Bar

	// RefreshInterval in the time duration to wait for refreshing the output
	RefreshInterval time.Duration

	lw     *uilive.Writer
	ticker *time.Ticker
	tdone  chan bool
	mtx    *sync.RWMutex
}

// New returns a new progress bar with defaults
func New() *Progress {
	lw := uilive.New()
	lw.Out = Out

	return &Progress{
		Width:           Width,
		Out:             Out,
		Bars:            make([]*Bar, 0),
		RefreshInterval: RefreshInterval,

		tdone: make(chan bool),
		lw:    lw,
		mtx:   &sync.RWMutex{},
	}
}

// AddBar creates a new progress bar and adds it to the default progress container
func AddBar(total int) *Bar {
	return defaultProgress.AddBar(total)
}
func Default(total int, description ...string) *Bar {
	return defaultProgress.Default(total, description...)
}

// RemoveBar remove bar from progress
func RemoveBar(bar *Bar) error {
	return defaultProgress.RemoveBar(bar)
}

// RemoveBarOnComplete remove bar from progress when complete
func RemoveBarOnComplete(bar *Bar) error {
	return defaultProgress.RemoveBarOnComplete(bar)
	//if bar.IsComplete() {
	//	return defaultProgress.RemoveBar(bar)
	//} else {
	//	return nil
	//}
}

// Start starts the rendering the progress of progress bars using the DefaultProgress. It listens for updates using `bar.Set(n)` and new bars when added using `AddBar`
func Start() {
	defaultProgress.Start()
}

// Stop stops listening
func Stop() {
	defaultProgress.Stop()
}

// Listen listens for updates and renders the progress bars
func Listen() {
	defaultProgress.Listen()
}

func (p *Progress) SetOut(o io.Writer) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.Out = o
	p.lw.Out = o
}

func (p *Progress) SetRefreshInterval(interval time.Duration) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.RefreshInterval = interval
}

// AddBar creates a new progress bar and adds to the container
func (p *Progress) AddBar(total int) *Bar {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	bar := NewBar(total)
	bar.Width = p.Width
	p.Bars = append(p.Bars, bar)
	return bar
}

// AddBar creates a new progress bar and adds to the container
func (p *Progress) Default(total int, description ...string) *Bar {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	bar := NewBar(total)
	bar.Width = p.Width
	if len(description) > 0 {
		bar.PrependCompleted().PrependDesc(description[0]).AppendStr("(").AppendSlashNum().AppendElapsed().AppendStr(":").AppendETA().AppendStr(")")
	} else {
		bar.PrependCompleted().PrependDesc(description[0]).AppendStr("(").AppendSlashNum().AppendElapsed().AppendStr(":").AppendETA().AppendStr(")")
	}
	p.Bars = append(p.Bars, bar)
	return bar
}

// Listen listens for updates and renders the progress bars
func (p *Progress) Listen() {
	for {

		p.mtx.Lock()
		interval := p.RefreshInterval
		p.mtx.Unlock()

		select {
		case <-time.After(interval):
			p.print()
		case <-p.tdone:
			p.print()
			close(p.tdone)
			return
		}
	}
}

func (p *Progress) print() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	//dDel := -1
	for _, bar := range p.Bars {
		fmt.Fprintln(p.lw, bar.String())
		//if bar.current == bar.Total {
		//	dDel = k
		//}
	}
	//if dDel > -1 {
	//	if len(p.Bars) > dDel+1 {
	//		p.Bars = append(p.Bars[:dDel], p.Bars[dDel+1:]...)
	//	} else {
	//		p.Bars = p.Bars[:dDel]
	//	}
	//}
	p.lw.Flush()
}

// Start starts the rendering the progress of progress bars. It listens for updates using `bar.Set(n)` and new bars when added using `AddBar`
func (p *Progress) Start() {
	go p.Listen()
}

// Stop stops listening
func (p *Progress) Stop() {
	p.tdone <- true
	<-p.tdone
}

// Bypass returns a writer which allows non-buffered data to be written to the underlying output
func (p *Progress) Bypass() io.Writer {
	return p.lw.Bypass()
}

func (p *Progress) RemoveBar(bar *Bar) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	dDel := -1
	for k, b := range p.Bars {
		if b == bar {
			dDel = k
			break
		}
	}
	if dDel > -1 {
		if len(p.Bars) > dDel+1 {
			p.Bars = append(p.Bars[:dDel], p.Bars[dDel+1:]...)
		} else {
			p.Bars = p.Bars[:dDel]
		}
	}
	return nil
}

// RemoveBarOnComplete remove bar from progress when complete
func (p *Progress) RemoveBarOnComplete(bar *Bar) error {
	if bar == nil {
		return fmt.Errorf("bar not exist")
	}
	if bar.IsComplete() {
		return defaultProgress.RemoveBar(bar)
	} else {
		return nil
	}
}
