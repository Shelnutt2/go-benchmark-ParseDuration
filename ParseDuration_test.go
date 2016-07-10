package durationtest

import (
	"testing"
	"time"
)

var parseDurationTests = []struct {
	in   string
	ok   bool
	want time.Duration
}{
	// simple
	{"0", true, 0},
	{"5s", true, 5 * time.Second},
	{"30s", true, 30 * time.Second},
	{"1478s", true, 1478 * time.Second},
	// sign
	{"-5s", true, -5 * time.Second},
	{"+5s", true, 5 * time.Second},
	{"-0", true, 0},
	{"+0", true, 0},
	// decimal
	{"5.0s", true, 5 * time.Second},
	{"5.6s", true, 5*time.Second + 600*time.Millisecond},
	{"5.s", true, 5 * time.Second},
	{".5s", true, 500 * time.Millisecond},
	{"1.0s", true, 1 * time.Second},
	{"1.00s", true, 1 * time.Second},
	{"1.004s", true, 1*time.Second + 4*time.Millisecond},
	{"1.0040s", true, 1*time.Second + 4*time.Millisecond},
	{"100.00100s", true, 100*time.Second + 1*time.Millisecond},
	// different units
	{"10ns", true, 10 * time.Nanosecond},
	{"11us", true, 11 * time.Microsecond},
	{"12µs", true, 12 * time.Microsecond}, // U+00B5
	{"12μs", true, 12 * time.Microsecond}, // U+03BC
	{"13ms", true, 13 * time.Millisecond},
	{"14s", true, 14 * time.Second},
	{"15m", true, 15 * time.Minute},
	{"16h", true, 16 * time.Hour},
	// composite durations
	{"3h30m", true, 3*time.Hour + 30*time.Minute},
	{"10.5s4m", true, 4*time.Minute + 10*time.Second + 500*time.Millisecond},
	{"-2m3.4s", true, -(2*time.Minute + 3*time.Second + 400*time.Millisecond)},
	{"1h2m3s4ms5us6ns", true, 1*time.Hour + 2*time.Minute + 3*time.Second + 4*time.Millisecond + 5*time.Microsecond + 6*time.Nanosecond},
	{"39h9m14.425s", true, 39*time.Hour + 9*time.Minute + 14*time.Second + 425*time.Millisecond},
	// large value
	{"52763797000ns", true, 52763797000 * time.Nanosecond},
	// more than 9 digits after decimal point, see https://golang.org/issue/6617
	{"0.3333333333333333333h", true, 20 * time.Minute},
	// 9007199254740993 = 1<<53+1 cannot be stored precisely in a float64
	{"9007199254740993ns", true, (1<<53 + 1) * time.Nanosecond},
	// largest duration that can be represented by int64 in nanoseconds
	{"9223372036854775807ns", true, (1<<63 - 1) * time.Nanosecond},
	{"9223372036854775.807us", true, (1<<63 - 1) * time.Nanosecond},
	{"9223372036s854ms775us807ns", true, (1<<63 - 1) * time.Nanosecond},
	// large negative value
	{"-9223372036854775807ns", true, -1<<63 + 1*time.Nanosecond},

	// errors
	{"", false, 0},
	{"3", false, 0},
	{"-", false, 0},
	{"s", false, 0},
	{".", false, 0},
	{"-.", false, 0},
	{".s", false, 0},
	{"+.s", false, 0},
	{"3000000h", false, 0},                  // overflow
	{"9223372036854775808ns", false, 0},     // overflow
	{"9223372036854775.808us", false, 0},    // overflow
	{"9223372036854ms775us808ns", false, 0}, // overflow
	// largest negative value of type int64 in nanoseconds should fail
	// see https://go-review.googlesource.com/#/c/2461/
	{"-9223372036854775808ns", false, 0},
}

var parseDurationBytesTests = []struct {
	in   []byte
	ok   bool
	want time.Duration
}{
	// simple
	{[]byte("0"), true, 0},
	{[]byte("5s"), true, 5 * time.Second},
	{[]byte("30s"), true, 30 * time.Second},
	{[]byte("1478s"), true, 1478 * time.Second},
	// sign
	{[]byte("-5s"), true, -5 * time.Second},
	{[]byte("+5s"), true, 5 * time.Second},
	{[]byte("-0"), true, 0},
	{[]byte("+0"), true, 0},
	// decimal
	{[]byte("5.0s"), true, 5 * time.Second},
	{[]byte("5.6s"), true, 5*time.Second + 600*time.Millisecond},
	{[]byte("5.s"), true, 5 * time.Second},
	{[]byte(".5s"), true, 500 * time.Millisecond},
	{[]byte("1.0s"), true, 1 * time.Second},
	{[]byte("1.00s"), true, 1 * time.Second},
	{[]byte("1.004s"), true, 1*time.Second + 4*time.Millisecond},
	{[]byte("1.0040s"), true, 1*time.Second + 4*time.Millisecond},
	{[]byte("100.00100s"), true, 100*time.Second + 1*time.Millisecond},
	// different units
	{[]byte("10ns"), true, 10 * time.Nanosecond},
	{[]byte("11us"), true, 11 * time.Microsecond},
	{[]byte("12µs"), true, 12 * time.Microsecond}, // U+00B5
	{[]byte("12μs"), true, 12 * time.Microsecond}, // U+03BC
	{[]byte("13ms"), true, 13 * time.Millisecond},
	{[]byte("14s"), true, 14 * time.Second},
	{[]byte("15m"), true, 15 * time.Minute},
	{[]byte("16h"), true, 16 * time.Hour},
	// composite durations
	{[]byte("3h30m"), true, 3*time.Hour + 30*time.Minute},
	{[]byte("10.5s4m"), true, 4*time.Minute + 10*time.Second + 500*time.Millisecond},
	{[]byte("-2m3.4s"), true, -(2*time.Minute + 3*time.Second + 400*time.Millisecond)},
	{[]byte("1h2m3s4ms5us6ns"), true, 1*time.Hour + 2*time.Minute + 3*time.Second + 4*time.Millisecond + 5*time.Microsecond + 6*time.Nanosecond},
	{[]byte("39h9m14.425s"), true, 39*time.Hour + 9*time.Minute + 14*time.Second + 425*time.Millisecond},
	// large value
	{[]byte("52763797000ns"), true, 52763797000 * time.Nanosecond},
	// more than 9 digits after decimal point, see https://golang.org/issue/6617
	{[]byte("0.3333333333333333333h"), true, 20 * time.Minute},
	// 9007199254740993 = 1<<53+1 cannot be stored precisely in a float64
	{[]byte("9007199254740993ns"), true, (1<<53 + 1) * time.Nanosecond},
	// largest duration that can be represented by int64 in nanoseconds
	{[]byte("9223372036854775807ns"), true, (1<<63 - 1) * time.Nanosecond},
	{[]byte("9223372036854775.807us"), true, (1<<63 - 1) * time.Nanosecond},
	{[]byte("9223372036s854ms775us807ns"), true, (1<<63 - 1) * time.Nanosecond},
	// large negative value
	{[]byte("-9223372036854775807ns"), true, -1<<63 + 1*time.Nanosecond},

	// errors
	{[]byte(""), false, 0},
	{[]byte("3"), false, 0},
	{[]byte("-"), false, 0},
	{[]byte("s"), false, 0},
	{[]byte("."), false, 0},
	{[]byte("-."), false, 0},
	{[]byte(".s"), false, 0},
	{[]byte("+.s"), false, 0},
	{[]byte("3000000h"), false, 0},                  // overflow
	{[]byte("9223372036854775808ns"), false, 0},     // overflow
	{[]byte("9223372036854775.808us"), false, 0},    // overflow
	{[]byte("9223372036854ms775us808ns"), false, 0}, // overflow
	// largest negative value of type int64 in nanoseconds should fail
	// see https://go-review.googlesource.com/#/c/2461/
	{[]byte("-9223372036854775808ns"), false, 0},
}

//BenchmarkParseDuration Benchmark Parse Duration
func BenchmarkParseDuration(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range parseDurationTests {
			d, err := time.ParseDuration(tc.in)
			if tc.ok && (err != nil || d != tc.want) {
				b.Errorf("ParseDuration(%q) = %v, %v, want %v, nil", tc.in, d, err, tc.want)
			} else if !tc.ok && err == nil {
				b.Errorf("ParseDuration(%q) = _, nil, want _, non-nil", tc.in)
			}
		}
	}
}

//BenchmarkParseDurationBytes Benchmark Parse Duration
func BenchmarkParseDurationBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range parseDurationBytesTests {
			d, err := time.ParseDurationBytes(tc.in)
			if tc.ok && (err != nil || d != tc.want) {
				b.Errorf("ParseDuration(%q) = %v, %v, want %v, nil", tc.in, d, err, tc.want)
			} else if !tc.ok && err == nil {
				b.Errorf("ParseDuration(%q) = _, nil, want _, non-nil", tc.in)
			}
		}
	}
}

//BenchmarkParseDurationBytesConversion Benchmark Parse Duration
func BenchmarkParseDurationBytesConversion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range parseDurationTests {
			d, err := time.ParseDurationBytes([]byte(tc.in))
			if tc.ok && (err != nil || d != tc.want) {
				b.Errorf("ParseDuration(%q) = %v, %v, want %v, nil", tc.in, d, err, tc.want)
			} else if !tc.ok && err == nil {
				b.Errorf("ParseDuration(%q) = _, nil, want _, non-nil", tc.in)
			}
		}
	}
}
