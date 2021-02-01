package split

import (
	"reflect"
	"testing"
)

//测试 go test
//覆盖率 go test -cover
//覆盖率 go test -cover -coverprofile=c.out
//覆盖率 go tool cover -html=c.out   用HTML 数据

//func TestSplit(t *testing.T) {
//	got := Split("我爱你", "爱")
//	want := []string{"我", "你"}
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("want: %v got: %v", want, got)
//	}
//}

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"simple1": {input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"simple2": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"simple3": {input: "我爱你", sep: "爱", want: []string{"我", "你"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name： %s,want: %v got: %v", name, tc.want, got)
			}
		})
	}
}


//基准测试
//go test -bench=Split
//BenchmarkSplit-4(进程数)         3114676(执行次数)               391 ns/op(每一次操作耗费的时间)
//PASS
//ok      example.com/m/v2        2.802s

//go test -bench=Split -benchmem   内存相关数据
//BenchmarkSplit-4         1341330               894 ns/op             112 B/op(每次操作字节数)          3 allocs/op(三次申请内存操作)
//PASS
//ok      example.com/m/v2        4.151s


//优化之后
//BenchmarkSplit-4         5124068               226 ns/op              48 B/op          1 allocs/op
//PASS
//ok      example.com/m/v2        2.426s


func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙也有河", "沙")
	}
}
