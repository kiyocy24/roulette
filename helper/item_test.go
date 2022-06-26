package helper

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLot(t *testing.T) {
	const Num = 100000

	type args struct {
		items map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int // histogram
	}{
		{
			name: "same weight, num 10",
			args: args{
				items: map[string]int{
					"1":  10,
					"2":  10,
					"3":  10,
					"4":  10,
					"5":  10,
					"6":  10,
					"7":  10,
					"8":  10,
					"9":  10,
					"10": 10,
				},
			},
			want: map[string]int{
				"1":  Num / 10,
				"2":  Num / 10,
				"3":  Num / 10,
				"4":  Num / 10,
				"5":  Num / 10,
				"6":  Num / 10,
				"7":  Num / 10,
				"8":  Num / 10,
				"9":  Num / 10,
				"10": Num / 10,
			},
		},
		{
			name: "weight 10, 20, 30",
			args: args{
				items: map[string]int{
					"1": 10,
					"2": 20,
					"3": 30,
				},
			},
			want: map[string]int{
				"1": Num * 10 / 60,
				"2": Num * 20 / 60,
				"3": Num * 30 / 60,
			},
		},
		{
			name: "weight 1, 99",
			args: args{
				items: map[string]int{
					"1": 1,
					"2": 99,
				},
			},
			want: map[string]int{
				"1": Num * 1 / 100,
				"2": Num * 99 / 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]int{}
			for i := 0; i < Num; i++ {
				got[Lot(tt.args.items)] += 1
			}
			for k, v := range got {
				diff := math.Abs(float64(v - tt.want[k]))
				if diff < float64(tt.want[k])*0.05 {
					got[k] = tt.want[k]
				}
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Log() is mismatch (-got +want):%s\n", diff)
			}
		})
	}
}
