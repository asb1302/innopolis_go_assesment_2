package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalSequence(t *testing.T) {
	type args struct {
		mtx [][]int
		ua  []int
	}

	mtx1 := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	mtx2 := [][]int{
		{0, 1, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 1},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0},
	}

	tests := []struct {
		name      string
		args      args
		want      int
		expectErr string
	}{
		{
			name: "mtx 5 vertices 100%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 0, 2},
			},
			want:      100,
			expectErr: "",
		},
		{
			name: "mtx 5 vertices 0%",
			args: args{
				mtx: mtx1,
				ua:  []int{},
			},
			want:      0,
			expectErr: "",
		},
		{
			name: "mtx 5 vertices 50%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 0},
			},
			want:      50,
			expectErr: "",
		},
		{
			name: "mtx 7 vertices 100%",
			args: args{
				mtx: mtx2,
				ua:  []int{0, 1, 3, 5, 2, 6, 4},
			},
			want:      100,
			expectErr: "",
		},
		{
			name: "mtx 7 vertices 83%",
			args: args{
				mtx: mtx2,
				ua:  []int{0, 1, 3, 5, 2, 6},
			},
			want:      83,
			expectErr: "",
		},
		{
			name: "invalid matrix",
			args: args{
				mtx: [][]int{
					{0, 1},
					{1, 0, 0},
				},
				ua: []int{0, 1},
			},
			want:      0,
			expectErr: "invalid matrix",
		},
		{
			name: "invalid user answer",
			args: args{
				mtx: mtx1,
				ua:  []int{0, 5},
			},
			want:      0,
			expectErr: "invalid user answer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalSequence(tt.args.mtx, tt.args.ua)
			if tt.expectErr != "" {
				assert.EqualError(t, err, tt.expectErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
