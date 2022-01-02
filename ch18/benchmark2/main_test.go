package main

import (
	"reflect"
	"testing"
)

func BenchmarkStreamDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreamDecode("art.json")
	}
}

func BenchmarkMemoryDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoryDecode("art.json")
	}
}

func TestStreamDecode(t *testing.T) {
	art := StreamDecode("art.json")
	art2 := Article{
		Id:      1,
		Content: "this is a art's content.",
		Comments: []Comment{
			{Id: 1, Content: "first comment content.", Uid: 1},
			{Id: 2, Content: "second comment content.", Uid: 1},
			{Id: 3, Content: "third comment content.", Uid: 2},
		},
		Uid: 1,
	}
	if !reflect.DeepEqual(art, art2) {
		t.Error("decode result is not ok")
		t.Log(art.String())
		t.Log(art2.String())
	}
}

func TestMemoryDecode(t *testing.T) {
	art := MemoryDecode("art.json")
	art2 := Article{
		Id:      1,
		Content: "this is a art's content.",
		Comments: []Comment{
			{Id: 1, Content: "first comment content.", Uid: 1},
			{Id: 2, Content: "second comment content.", Uid: 1},
			{Id: 3, Content: "third comment content.", Uid: 2},
		},
		Uid: 1,
	}
	if !reflect.DeepEqual(art, art2) {
		t.Error("decode result is not ok")
	}
}
