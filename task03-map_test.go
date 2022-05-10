package homework

import (
	"reflect"
	"testing"
)

func Test_sortMapValues(t *testing.T) {
	type args struct {
		input map[int]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{name: "map1", args: struct{ input map[int]string }{input: map[int]string{20: "twenty", 10: "ten"}}, wantResult: []string{"ten", "twenty"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := sortMapValues(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("sortMapValues() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
