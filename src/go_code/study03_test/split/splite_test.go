package split

import (
	"reflect"
	"testing"
)

// 测试
// TestSplit 单个测试用例
// func TestSplit(t *testing.T) {
// 	got := Split("a:b:c", "b")
// 	want := []string{"a", "c"}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("want:%v, got:%v", want, got)
// 	}
// }

// TestSplit 多个测试用例
func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	// b.N不是固定数
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
