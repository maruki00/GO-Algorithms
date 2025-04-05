package main

import "testing"

func CanPlaceFlowersTest(t *testing.T) {
	tests := []struct {
		n       int
		flowers []int
		want    bool
	}{
		{
			n:       1,
			flowers: []int{1, 2, 3, 44, 5},
			want:    false,
		},
		{
			n:       1,
			flowers: []int{1, 2, 3, 44, 5},
			want:    false,
		},
		{
			n:       1,
			flowers: []int{1, 2, 3, 44, 5},
			want:    false,
		},
		{
			n:       1,
			flowers: []int{1, 2, 3, 44, 5},
			want:    false,
		},
	}
	for _, ts := range tests {
		w := CanPlaceFlowers(ts.flowers, ts.n)
		if w == ts.want {
			t.Errorf("Failed")
		}
	}
}
