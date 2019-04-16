package sparkline

import (
	"testing"
)

func TestSparkLine_cleanup(t *testing.T) {
	type Sparkline struct {
		maxSize      int
		expectedSize int
		values       []float64
	}
	tests := []struct {
		name      string
		sparkline Sparkline
	}{
		{"should cleanup values to the max size", Sparkline{maxSize: 5, expectedSize: 5, values: []float64{1, 2, 3, 4, 5, 6, 7, 8}}},
		{"should not cleanup values if max size is bigger than values len", Sparkline{maxSize: 5, expectedSize: 3, values: []float64{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SparkLine{
				values: tt.sparkline.values,
			}
			s.cleanup(tt.sparkline.maxSize)
			if len(s.values) != tt.sparkline.expectedSize {
				t.Errorf("Values size after cleanup is %v, but needed to be %v", len(s.values), tt.sparkline.expectedSize)
			}
		})
	}
}
