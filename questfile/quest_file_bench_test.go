package questfile

import (
	"bytes"
	"testing"
)

func BenchmarkRead_Minimal(b *testing.B) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	if err := Write(&buf, q); err != nil {
		b.Fatal(err)
	}
	data := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Read(bytes.NewReader(data))
	}
}

func BenchmarkWrite_Minimal(b *testing.B) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = Write(&buf, q)
	}
}

func BenchmarkRead_Maximal(b *testing.B) {
	q := minimalValidQuestFile()
	for i := range q.Objectives {
		q.Objectives[i].Block[0] = TypeDROP
		q.Objectives[i].Block[92] = 255
		q.Objectives[i].Name = make([]byte, 255)
	}
	var buf bytes.Buffer
	if err := Write(&buf, q); err != nil {
		b.Fatal(err)
	}
	data := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Read(bytes.NewReader(data))
	}
}

func BenchmarkWrite_Maximal(b *testing.B) {
	q := minimalValidQuestFile()
	for i := range q.Objectives {
		q.Objectives[i].Block[0] = TypeDROP
		q.Objectives[i].Block[92] = 255
		q.Objectives[i].Name = make([]byte, 255)
	}
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = Write(&buf, q)
	}
}
