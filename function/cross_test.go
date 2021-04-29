package function

import (
	"reflect"
	"testing"
)

func TestCross_CrossOver(t *testing.T) {
	type fields struct {
		Calculator *crossCalculator
	}
	type args struct {
		first  float64
		second float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"TestCross_CrossOver", fields{Calculator: &crossCalculator{
		}}, args{
			first:  0,
			second: 0,
		}, false},
		{"TestCross_CrossOver", fields{Calculator: &crossCalculator{
			Count: 1,
		}}, args{
			first:  0,
			second: 0,
		}, false},
		{"TestCross_CrossOver", fields{Calculator: &crossCalculator{
			First:       -0.8,
			FirstValue:  0,
			Second:      -0.5,
			SecondValue: 0,
			Count:       3,
		}}, args{
			first:  1,
			second: 0,
		}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cross{
				Calculator: tt.fields.Calculator,
			}
			if got := c.CrossOver(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CrossOver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCross_CrossUnder(t *testing.T) {
	type fields struct {
		Calculator *crossCalculator
	}
	type args struct {
		first  float64
		second float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"TestCross_CrossUnder", fields{Calculator: &crossCalculator{
		}}, args{
			first:  0,
			second: 0,
		}, false},
		{"TestCross_CrossUnder", fields{Calculator: &crossCalculator{
			Count: 1,
		}}, args{
			first:  0,
			second: 0,
		}, false},
		{"TestCross_CrossUnder", fields{Calculator: &crossCalculator{
			First:       8.0,
			FirstValue:  0,
			Second:      6.0,
			SecondValue: 0,
			Count:       3,
		}}, args{
			first:  -0.8,
			second: 0,
		}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cross{
				Calculator: tt.fields.Calculator,
			}
			if got := c.CrossUnder(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CrossUnder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCross(t *testing.T) {
	tests := []struct {
		name string
		want *Cross
	}{
		{"TestNewCross", nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCross(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCross_CrossOver(b *testing.B) {
	cross := NewCross()
	for n := 0; n < b.N; n++ {
		cross.CrossOver(float64(n), 0)
	}
}

func BenchmarkCross_CrossUnder(b *testing.B) {
	cross := NewCross()
	for n := 0; n < b.N; n++ {
		cross.CrossUnder(float64(n), 0)
	}
}
