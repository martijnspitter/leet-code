package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueryResults(t *testing.T) {
	tests := []struct {
		limit   int
		queries [][]int
		want    []int
	}{
		{
			limit:   4,
			queries: [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}},
			want:    []int{1, 2, 2, 3},
		},
		{
			limit:   4,
			queries: [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}},
			want:    []int{1, 2, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("limit=%d,queries=%v", tt.limit, tt.queries), func(t *testing.T) {
			got := queryResults(tt.limit, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResults(%d, %v) = %v; want %v", tt.limit, tt.queries, got, tt.want)
			}
		})
	}
}
