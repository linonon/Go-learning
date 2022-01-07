package countdown_test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown with second
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprint(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

}

func TestCountdown(t *testing.T) {
	t.Run("Countdown-ConfigurableSleeper", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		// spySleeper := &SpySleeper{}
		// Countdown(buffer, spySleeper)

		confSleeper := &ConfigurableSleeper{duration: 0}
		Countdown(buffer, confSleeper)

		got := buffer.String()
		want := `321Go!`

		if got != want {
			t.Errorf("got \n'%s' want \n'%s'", got, want)
		}

		// if spySleeper.Calls != 4 {
		// 	t.Errorf("not enough calls to sleeper, want 4 got %d\n", spySleeper.Calls)
		// }
	})

	t.Run("Sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}
