package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
    for i := 0; i < b.N; i++ {
        HelloWorld("Niki")
    }
}


func BenchmarkHellWordHamdani(b *testing.B) {
	for i := 0; i < b.N; i++{
		HelloWorld("Hamdani")
	}
}






func TestMain(m *testing.M) {
	// Seperti Layout Pada Component
	// Test ini akan di jalankan pertama kali
	// Dan akan di akhiri setelah semua test sudah di jalankan
	fmt.Println("Sebelum unit test")
	m.Run()
	fmt.Println("Setelah unit test")
}

func TestHelloWordTable(t *testing.T) {
	test := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Hello(Niki)",
			request: "Niki",
			expected: "Hello Niki",
		},
		{
			name: "Hello(Ahmad)",
			request: "Ahmad",
			expected: "Hello Ahmad",
		},
		{
			name: "Hello(Hamdani)",
			request: "Hamdani",
			expected: "Hello Hamdani",
		},
	}

	for _, val := range test {
		t.Run(val.name, func(t *testing.T) {
			result := HelloWorld(val.request)
			assert.Equal(t, val.expected, result)
		})
	}
}


func TestSubTest(t *testing.T) {
	t.Run("Niki", func(t *testing.T) {
		result := HelloWorld("Niki")
		require.Equal(t, "Hello Niki", result, "Harusnya menampilkan Niki")
	})
	t.Run("Hamdani", func(t *testing.T) {
		result := HelloWorld("Hamdani")
		require.Equal(t, "Hello Hamdani", result, "Harusnya menampilkan Niki")
	})

	//jika hanya ingin menjalankan subtest nyasaja
	//go test -v -run=TestSubTest/NamaSubtest
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows"{
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Niki")
	assert.Equal(t, "Hello Niki", result, "Harusnya menampilkan Hello Niki")
	fmt.Println("TestHelloWorld with Assert Done")
}


func TestHelloWordAssert(t *testing.T) {
	result := HelloWorld("Niki")
	//assert akan menjalankan code di bawahnya
	assert.Equal(t, "Hello Niki", result, "Harusnya menampilkan Hello Niki")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWordRequire(t *testing.T) {
	result := HelloWorld("Niki")
	//require akan langsung berhenti dan perintah di bawahnya tidak akan di jalankan
	require.Equal(t, "Hello Niki", result, "Harusnya menampilkan Hello Niki")
	fmt.Println("TestHelloWorld with Require Done")
}


//Tanpa package stretchr/testify/assert
// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("Niki")
// 	if result != "Hello Niki" {
// 		//t.Fail() //Kalo fail akan tetap menjalangkan kode di bawahnya
// 		t.Fatal("Harusnya menampilkan Hi Niki")
// 	}
// 	fmt.Println("Test HelloWorld")
// }

// func TestHelloWorldHamdani(t *testing.T) {
// 	result := HelloWorld("Hamdani")
// 	if result != "Hello Hamdani" {
// 		//t.FailNow() //failnow akan langsung memberhentikan code, dan kode di bawahnya tidak akan di jalankan
// 		t.Fatal("Harusnya menampiplkan Hi Hamdani")
// 	}

// 	fmt.Println("Test HelloWordHamdani")
// }
//End tanpa package stretchr/testify/assert