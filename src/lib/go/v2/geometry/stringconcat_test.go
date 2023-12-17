package geometry

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func sprintf() string {
	return fmt.Sprintf("%d,%d", 5, 6)
}
func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintf()
	}
}

func join() string {
	return strings.Join([]string{strconv.Itoa(5), strconv.Itoa(6)}, ",")
}
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join()
	}
}

func concat() string {
	return strconv.Itoa(5) + "," + strconv.Itoa(6)
}
func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat()
	}
}

func stringBuilder() string {
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(5))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(6))
	return builder.String()
}
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilder()
	}
}
