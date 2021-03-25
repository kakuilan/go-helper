package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile_GetExt(t *testing.T) {
	var ext string

	ext = KFile.GetExt(fileGo)
	assert.Equal(t, "go", ext)

	ext = KFile.GetExt(fileGitkee)
	assert.Equal(t, "gitkeep", ext)

	ext = KFile.GetExt(fileSongs)
	assert.Equal(t, "txt", ext)

	ext = KFile.GetExt(fileNone)
	assert.Empty(t, ext)
}

func BenchmarkFile_GetExt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetExt(fileMd)
	}
}

func TestFile_ReadFile(t *testing.T) {
	var bs []byte
	var err error

	bs, err = KFile.ReadFile(fileMd)
	assert.NotEmpty(t, bs)
	assert.Nil(t, err)

	//不存在的文件
	bs, err = KFile.ReadFile(fileNone)
	assert.NotNil(t, err)
}

func BenchmarkFile_ReadFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadFile(fileMd)
	}
}
