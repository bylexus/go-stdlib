package ecli

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/term"
)

const (
	ansiiGreen       = "\033[32m"
	ansiiOrange      = "\033[33m"
	ansiiRed         = "\033[31m"
	ansiiOk          = ansiiGreen
	ansiiWarn        = ansiiOrange
	ansiiFail        = ansiiRed
	ansiiReset       = "\033[0m"
	ansiiResetColor  = "\033[39m"
	ansiiBold        = "\033[1m"
	ansiiUnbold      = "\033[22m"
	ansiiUnderline   = "\033[4m"
	ansiiNoUnderline = "\033[24m"

	ansiiCursorLeft = "\033[1D"
	ansiiDeleteChar = "\033[K"
)

type Cli struct {
	spinnerIsRunning bool
	spinnerChannel   chan bool
	spinnerWg        sync.WaitGroup
}

func ansiStr(str ...string) string {
	if IsTerm() {
		return strings.Join(str, "")
	} else {
		return ""
	}
}

func Green() string {
	return ansiStr(ansiiGreen)
}
func Orange() string {
	return ansiStr(ansiiOrange)
}

func Red() string {
	return ansiStr(ansiiRed)
}

func Ok() string {
	return Green()
}

func Warn() string {
	return Orange()
}

func Fail() string {
	return Red()
}

func Reset() string {
	return ansiStr(ansiiReset)
}

func ResetColor() string {
	return ansiStr(ansiiResetColor)
}

func Bold() string {
	return ansiStr(ansiiBold)
}
func BoldOff() string {
	return ansiStr(ansiiUnbold)
}

func Underline() string {
	return ansiStr(ansiiUnderline)
}

func UnderlineOff() string {
	return ansiStr(ansiiNoUnderline)
}

func CursorLeft() string {
	return ansiStr(ansiiCursorLeft)
}

func DelChar() string {
	return ansiStr(ansiiDeleteChar)
}

func DelLastChar() string {
	return ansiStr(ansiiCursorLeft, ansiiDeleteChar)
}

func New() *Cli {
	return &Cli{
		spinnerIsRunning: false,
		spinnerChannel:   make(chan bool),
		spinnerWg:        sync.WaitGroup{},
	}
}

func IsTerm() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
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
	c.Printf(Ok()+format+ResetColor(), args...)
}

func (c *Cli) OkPrintln(str ...string) {
	for _, s := range str {
		c.Printf("%s%s%s\n", Ok(), s, ResetColor())
	}
}

func (c *Cli) OkPrint(str ...string) {
	for _, s := range str {
		c.Printf(Ok() + s + ResetColor())
	}
}

func (c *Cli) WarnPrintf(format string, args ...interface{}) {
	c.Printf(Warn()+format+ResetColor(), args...)
}

func (c *Cli) WarnPrintln(str ...string) {
	for _, s := range str {
		c.Printf("%s%s%s\n", Warn(), s, ResetColor())
	}
}

func (c *Cli) WarnPrint(str ...string) {
	for _, s := range str {
		c.Printf(Warn() + s + ResetColor())
	}
}

func (c *Cli) ErrorPrintf(format string, args ...interface{}) {
	c.Printf(Fail()+format+ResetColor(), args...)
}

func (c *Cli) ErrorPrintln(str ...string) {
	for _, s := range str {
		c.Printf("%s%s%s\n", Fail(), s, ResetColor())
	}
}

func (c *Cli) ErrorPrint(str ...string) {
	for _, s := range str {
		c.Printf(Fail() + s + ResetColor())
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
				c.Print(DelLastChar())
			}
		}
	}()
}

func (c *Cli) SpinnerOff(print ...string) {
	close(c.spinnerChannel)
	c.spinnerWg.Wait()
	c.Print(print...)
}
