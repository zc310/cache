package cache

import (
	"bytes"
	"testing"
	"time"
)

func BenchmarkBin(b *testing.B) {
	b.ReportAllocs()
	cv := &Value{}

	var b1 []byte
	var err error
	for i := 0; i < b.N; i++ {
		b1, err = Marshal(cv)
		if err != nil {
			b.Fatal(err)
		}
		err = Unmarshal(b1, cv)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshal(t *testing.T) {
	var cv Value
	cv.Size = 3
	cv.Timeout = time.Now().Unix()
	cv.Value = []byte("123")

	b, err := Marshal(&cv)
	if err != nil {
		t.Fatal(err)
	}
	var cv2 Value
	err = Unmarshal(b, &cv2)
	if err != nil {
		t.Fatal(err)
	}
	if cv.Timeout != cv2.Timeout {
		t.Fatal("err Timeout")
	}
	if cv.Size != cv2.Size {
		t.Fatal("err Size")
	}
	if bytes.Compare(cv.Value, cv2.Value) != 0 {
		t.Fatal("err Value")
	}
}
