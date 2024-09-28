package ecli

import (
	"fmt"
	"sync"
	"time"
)

const (
	AnsiiGreen       = "\033[32m"
	AnsiiOrange      = "\033[33m"
	AnsiiRed         = "\033[31m"
	AnsiiOk          = AnsiiGreen
	AnsiiWarn        = AnsiiOrange
	AnsiiError       = AnsiiRed
	AnsiiReset       = "\033[0m"
	AnsiiBold        = "\033[1m"
	AnsiiUnbold      = "\033[22m"
	AnsiiUnderline   = "\033[4m"
	AnsiiNoUnderline = "\033[24m"

	AnsiiCursorLeft = "\033[1D"
	AnsiiDeleteChar = "\033[K"
)

type Cli struct {
	spinnerIsRunning bool
	spinnerChannel   chan bool
	spinnerWg        sync.WaitGroup
}

func New() *Cli {
	return &Cli{
		spinnerIsRunning: false,
		spinnerChannel:   make(chan bool),
		spinnerWg:        sync.WaitGroup{},
	}
}

func delLastChar() {
	fmt.Print(AnsiiCursorLeft + AnsiiDeleteChar)
}

func (c *Cli) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (c *Cli) Println(str ...string) {
	for _, s := range str {
		c.Printf("%s\n", s)
	}
}

func (c *Cli) Print(str ...string) {
	for _, s := range str {
		c.Printf("%s", s)
	}
}

func (c *Cli) OkPrintf(format string, args ...interface{}) {
	c.Printf(AnsiiOk+format+AnsiiReset, args...)
}

func (c *Cli) OkPrintln(str ...string) {
	for _, s := range str {
		c.Printf("%s%s%s\n", AnsiiOk, s, AnsiiReset)
	}
}

func (c *Cli) OkPrint(str ...string) {
	for _, s := range str {
		c.Printf(AnsiiOk + s + AnsiiReset)
	}
}

func (c *Cli) WarnPrintf(format string, args ...interface{}) {
	c.Printf(AnsiiWarn+format+AnsiiReset, args...)
}

func (c *Cli) WarnPrintln(str ...string) {
	for _, s := range str {
		c.Printf("%s%s%s\n", AnsiiWarn, s, AnsiiReset)
	}
}

func (c *Cli) WarnPrint(str ...string) {
	for _, s := range str {
		c.Printf(AnsiiWarn + s + AnsiiReset)
	}
}

func (c *Cli) ErrorPrintf(format string, args ...interface{}) {
	c.Printf(AnsiiError+format+AnsiiReset, args...)
}

func (c *Cli) ErrorPrintln(str ...string) {
	for _, s := range str {
		c.Printf("%s%s%s\n", AnsiiError, s, AnsiiReset)
	}
}

func (c *Cli) ErrorPrint(str ...string) {
	for _, s := range str {
		c.Printf(AnsiiError + s + AnsiiReset)
	}
}

func (c *Cli) SpinnerOn() {
	spinners := []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
	if c.spinnerIsRunning {
		close(c.spinnerChannel)
		c.spinnerWg.Wait()
	}
	c.spinnerChannel = make(chan bool)
	c.spinnerWg.Add(1)
	go func() {
		idx := 0
		for {
			select {
			case <-c.spinnerChannel:
				c.spinnerWg.Done()
				return
			default:
				c.Printf("%s", string(spinners[idx]))
				idx = (idx + 1) % len(spinners)
				time.Sleep(100 * time.Millisecond)
				delLastChar()
			}
		}
	}()
}

func (c *Cli) SpinnerOff(print ...string) {
	close(c.spinnerChannel)
	c.spinnerWg.Wait()
	c.Print(print...)
}
