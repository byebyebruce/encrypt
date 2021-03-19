package encrypt

import "testing"

func makeTestData() []byte {
	const n = 2048
	testBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		testBytes[i] = n % 256
	}
	return testBytes
}

func TestDecrypt(t *testing.T) {
	var b = makeTestData()
	var b1 = makeTestData()
	var key uint16 = 23154
	Encrypt(b, key)
	Decrypt(b, key)
	for k, c := range b {
		if b1[k] != c {
			t.Fail()
		}
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bb := makeTestData()
		b.StartTimer()
		Encrypt(bb, 48649)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bb := makeTestData()
		b.StartTimer()
		Decrypt(bb, 51927)
	}
}
