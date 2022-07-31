package main

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_setupLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.SugaredLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
