package durationtest

import (
	"encoding/json"
	"testing"
	"time"
)

var jsonDurationTests = []struct {
	json     string
	ok       bool
	duration time.Duration
}{
	// simple
	{`"0s"`, true, 0},
	{`"5s"`, true, 5 * time.Second},
	{`"30s"`, true, 30 * time.Second},
	{`"24m38s"`, true, 1478 * time.Second},
	// sign
	{`"-5s"`, true, -5 * time.Second},
	{`"5s"`, true, 5 * time.Second},
	// decimal
	{`"5s"`, true, 5 * time.Second},
	{`"5.6s"`, true, 5*time.Second + 600*time.Millisecond},
	{`"500ms"`, true, 500 * time.Millisecond},
	{`"1s"`, true, 1 * time.Second},
	{`"1s"`, true, 1 * time.Second},
	{`"1.004s"`, true, 1*time.Second + 4*time.Millisecond},
	{`"1.004s"`, true, 1*time.Second + 4*time.Millisecond},
	{`"1m40.001s"`, true, 100*time.Second + 1*time.Millisecond},
	// different units
	{`"10ns"`, true, 10 * time.Nanosecond},
	{`"11µs"`, true, 11 * time.Microsecond},
	{`"12µs"`, true, 12 * time.Microsecond}, // U+00B5
	{`"13ms"`, true, 13 * time.Millisecond},
	{`"14s"`, true, 14 * time.Second},
	{`"15m0s"`, true, 15 * time.Minute},
	{`"16h0m0s"`, true, 16 * time.Hour},
	// composite durations
	{`"3h30m0s"`, true, 3*time.Hour + 30*time.Minute},
	{`"4m10.5s"`, true, 4*time.Minute + 10*time.Second + 500*time.Millisecond},
	{`"-2m3.4s"`, true, -(2*time.Minute + 3*time.Second + 400*time.Millisecond)},
	{`"1h2m3.004005006s"`, true, 1*time.Hour + 2*time.Minute + 3*time.Second + 4*time.Millisecond + 5*time.Microsecond + 6*time.Nanosecond},
	{`"39h9m14.425s"`, true, 39*time.Hour + 9*time.Minute + 14*time.Second + 425*time.Millisecond},
	// large value
	{`"52.763797s"`, true, 52763797000 * time.Nanosecond},
	// more than 9 digits after decimal point, see https://golang.org/issue/6617
	{`"20m0s"`, true, 20 * time.Minute},
	// 9007199254740993 = 1<<53+1 cannot be stored precisely in a float64
	{`"2501h59m59.254740993s"`, true, (1<<53 + 1) * time.Nanosecond},
	// largest duration that can be represented by int64 in nanoseconds
	{`"2562047h47m16.854775807s"`, true, (1<<63 - 1) * time.Nanosecond},
	// large negative value
	{`"-2562047h47m16.854775807s"`, true, -1<<63 + 1*time.Nanosecond},
}

func BenchmarkDurationJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, dt := range jsonDurationTests {
			b.StopTimer()
			var jsonDuration time.Duration
			var jsonBytes []byte
			var err error
			if jsonBytes, err = json.Marshal(dt.duration); err != nil {
				b.Errorf("%v json.Marshal error = %v, want nil", dt.duration, err)
			}
			if string(jsonBytes) != dt.json {
				b.Errorf("%v JSON = %#q, want %#q", dt.duration, string(jsonBytes), dt.json)
			}
			b.StartTimer()
			if err = json.Unmarshal(jsonBytes, &jsonDuration); err != nil {
				b.Errorf("%v json.Unmarshal error = %v, want nil", dt.duration, err)
			}
			if jsonDuration != dt.duration {
				b.Errorf("Unmarshaled duration = %v, want %v", jsonDuration, dt.duration)
			}
		}
	}
}
