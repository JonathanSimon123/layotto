package persistence

import (
	"github.com/layotto/layotto/pkg/filter/network/tcpcopy/strategy"
	"mosn.io/pkg/log"
	"testing"
)

func TestIsPersistence_switch_off(t *testing.T) {
	log.DefaultLogger.SetLogLevel(log.DEBUG)

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "test switch off",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPersistence(); got != tt.want {
				t.Errorf("IsPersistence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPersistence_not_in_sample_duration(t *testing.T) {
	log.DefaultLogger.SetLogLevel(log.DEBUG)

	strategy.DumpSwitch = true
	strategy.DumpSampleFlag = 0

	got := IsPersistence()

	want := false
	if got != want {
		t.Errorf("IsPersistence() = %v, want %v", got, want)
	}
}

func TestIsPersistence_cpu_exceed_max_rate(t *testing.T) {
	log.DefaultLogger.SetLogLevel(log.DEBUG)

	strategy.DumpSwitch = true
	strategy.DumpSampleFlag = 1
	strategy.DumpCpuMaxRate = 0

	got := IsPersistence()

	want := false
	if got != want {
		t.Errorf("IsPersistence() = %v, want %v", got, want)
	}
}

func TestIsPersistence_mem_exceed_max_rate(t *testing.T) {

	log.DefaultLogger.SetLogLevel(log.DEBUG)

	strategy.DumpSwitch = true
	strategy.DumpSampleFlag = 1
	strategy.DumpMemMaxRate = 0

	got := IsPersistence()

	want := false
	if got != want {
		t.Errorf("IsPersistence() = %v, want %v", got, want)
	}
}

func TestIsPersistence_success(t *testing.T) {

	log.DefaultLogger.SetLogLevel(log.DEBUG)

	strategy.DumpSwitch = true
	strategy.DumpSampleFlag = 1
	strategy.DumpCpuMaxRate = 100
	strategy.DumpMemMaxRate = 100

	got := IsPersistence()

	want := true
	if got != want {
		t.Errorf("IsPersistence() = %v, want %v", got, want)
	}
}