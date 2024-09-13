package zlogtime

import (
	"reflect"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Test_getTimeDuration(t *testing.T) {
	type args struct {
		timeDuration time.Duration
		unit         string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "just 1 second with ms unit",
			args: args{
				timeDuration: time.Duration(1) * time.Second,
				unit:         "milli",
			},
			want: 1000,
		},
		{
			name: "just 1 second with us unit",
			args: args{
				timeDuration: time.Duration(1) * time.Second,
				unit:         "micro",
			},
			want: 1000000,
		},
		{
			name: "4.5 second with ns unit",
			args: args{
				timeDuration: time.Duration(4.5*1000) * time.Millisecond,
				unit:         "nano",
			},
			want: 4500000000,
		},
		{
			name: "with unsupported unit",
			args: args{
				timeDuration: time.Duration(1) * time.Second,
				unit:         "terasecond",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTimeDuration(tt.args.timeDuration, tt.args.unit); got != tt.want {
				t.Errorf("getTimeDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShortDurationUnit(t *testing.T) {
	type args struct {
		unit string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ms unit",
			args: args{
				unit: "milli",
			},
			want: "ms",
		},
		{
			name: "us unit",
			args: args{
				unit: "micro",
			},
			want: "us",
		},
		{
			name: "ns unit",
			args: args{
				unit: "nano",
			},
			want: "ns",
		},
		{
			name: "wrong unit",
			args: args{
				unit: "superposition",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShortDurationUnit(tt.args.unit); got != tt.want {
				t.Errorf("getShortDurationUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLogLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want *zerolog.Event
	}{
		{
			name: "debug level",
			args: args{
				level: "debug",
			},
			want: log.Logger.Debug(),
		},
		{
			name: "info level",
			args: args{
				level: "info",
			},
			want: log.Logger.Info(),
		},
		{
			name: "warn level",
			args: args{
				level: "warn",
			},
			want: log.Logger.Warn(),
		},
		{
			name: "error level",
			args: args{
				level: "error",
			},
			want: log.Logger.Error(),
		},
		{
			name: "wrong level",
			args: args{
				level: "",
			},
			want: log.Logger.Log(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLogLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_getLogLevel_Fatal(t *testing.T) {
// 	type args struct {
// 		level string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *zerolog.Event
// 	}{
// 		{
// 			name: "fatal level",
// 			args: args{
// 				level: "fatal",
// 			},
// 			want: log.Logger.Fatal(),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := getLogLevel(tt.args.level)
// 			got.Msg("Fatal")
// 		})
// 	}
// }

func Test_getLogLevel_Panic(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want *zerolog.Event
	}{
		{
			name: "panic level",
			args: args{
				level: "panic",
			},
			want: log.Logger.Panic(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code is not panic")
				}
			}()
			got := getLogLevel(tt.args.level)
			got.Msg("should panicking")
		})
	}
}
