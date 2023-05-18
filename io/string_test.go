package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString1(t *testing.T) {
	s := "你好Go"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len("é"), len("。"), len("😀"), len("𫝈"))
}

func TestStringCut(t *testing.T) {
	s := "你好Go"
	c := s[0] // type uint8
	fmt.Println(c)
	fmt.Println(s[:2])
	fmt.Println(s[:3])

	// 截取字符时可以先将string转成rune切片,再进行切片操作:
	fmt.Println(string([]rune(s)[:2]))

}

func TestRange(t *testing.T) {
	s := "你好Go"
	//c := s[0] // uint8
	//fmt.Println(c)
	//fmt.Println(s[:2])
	//fmt.Println(s[:3])
	//fmt.Println()
	for i, v := range s { // []rune
		_ = i
		//fmt.Print(v, ",")
		fmt.Printf("%d: %q\t[% x]\n", i, v, []byte(string(v)))
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ",")
	}

	fmt.Println()

	fmt.Println(string([]rune(s)[:3]))
}

/*
strings.Builder 的特点
已存在的内容不可变 ， 但可以拼接更多的内容 ；
减少了内存分配和内容拷贝的次数 ；
可将内容重置，可重用值。
*/
func BenchmarkBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder := strings.Builder{}
		for j := 0; j < 100; j++ {
			builder.WriteString("Hello World\n")
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := ""
		for j := 0; j < 100; j++ {
			res += "Hello World\n"
		}
		_ = res
	}
	b.StopTimer()
}

func TestStringModify(t *testing.T) {
	s := "hello"
	s2 := s[:2]
	fmt.Printf("%p, %p", &s, &s2)
}

/*
Builder值不能再被复制是为了防止出现多个 Builder 值中的buf字节切片共用一个底层字节数组的情况。这样也就避免了多个同源的 Builder 值在拼接
内容时可能产生的冲突问题 。
*/
func TestStringBuilderCopy(t *testing.T) {
	builder := strings.Builder{}
	builder.WriteString("one!")

	func(b strings.Builder) {
		fmt.Println(b.String())
		b.WriteString("two!")
	}(builder)

	//func(b *strings.Builder) {
	//	b.WriteString("three!")
	//}(&builder)

	fmt.Println(builder.String())
}

// Reader的优势是维护一个已读计数器，知道下一次读的位置，读得更快。
func TestStringReader(t *testing.T) {
	s := "hello 世界"
	reader := strings.NewReader(s)

	fmt.Println(reader.ReadByte())
	fmt.Println(reader.ReadByte())

	_ = reader.Size() - int64(reader.Len()) // 已读计数

}
