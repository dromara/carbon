package carbon

import "testing"

func BenchmarkSetTestNow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetTestNow(Yesterday())
	}
}

func BenchmarkCleanTestNow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CleanTestNow()
	}
}

func BenchmarkIsTestNow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsTestNow()
	}
}

func BenchmarkUnSetTestNow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnSetTestNow()
	}
}

func BenchmarkIsSetTestNow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsSetTestNow()
	}
}
