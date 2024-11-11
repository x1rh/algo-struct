package sstable

import (
	"os"
	"testing"
)

func TestMetadata(t *testing.T) {
    path := "./test.metadata"
    fp, err := os.Create(path)
    if err != nil {
        t.Fatal(err)
    }

    meta1 := metablock{
        dataOffset: 0,
        dataLen: 1,
        indexOffset: 2,
        indexLen: 3,
    }

    if err := meta1.store(fp); err != nil {
        t.Fatal(err)
    }

    meta2 := metablock{}
    if err := meta2.load(fp); err != nil {
        t.Fatal(err)
    }

    if meta1 != meta2 {
        t.Fatal("not equal")
    } 
}
