package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTables(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "EKo",
			request: "Eko",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello World'")
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello World'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before test")

	m.Run()

	fmt.Println("After test")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello World'")
	fmt.Println("TestHelloWorld With Assert Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS")
	}

	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello World'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello World'")
	fmt.Println("TestHelloWorld With Assert Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		t.Error("Result must be 'Hello Khannedy'")
	}

	fmt.Println("Test HelloWorldKhannedy Done")
}
