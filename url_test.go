package kgo

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestUrl_ParseStr(t *testing.T) {
	var res map[string]interface{}
	var err error

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri1, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri2, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri3, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri4, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri5, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri6, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri7, res)
	assert.Nil(t, err)

	//将不合法的参数名转换
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri8, res)
	assert.Nil(t, err)

	//错误的
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri9, res)
	assert.NotNil(t, err)

	//错误的
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri10, res)
	assert.NotNil(t, err)

	//错误的
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri11, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri12, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri13, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri14, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri15, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri16, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri17, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri18, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri19, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri20, res)
	assert.Nil(t, err)

	//key nvalid URL escape "%"
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri21, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri22, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri23, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri24, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri25, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri26, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri27, res)
	assert.Nil(t, err)

	//key nvalid URL escape "%"
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri28, res)
	assert.NotNil(t, err)
}

func BenchmarkUrl_ParseStr(b *testing.B) {
	b.ResetTimer()
	res := KArr.NewStrMapItf()
	for i := 0; i < b.N; i++ {
		_ = KStr.ParseStr(tesUri1, res)
	}
}

func TestUrl_ParseUrl(t *testing.T) {
	var res map[string]string
	var err error
	var chk bool

	res, err = KStr.ParseUrl(tesUrl01, -1)
	assert.Nil(t, err)

	res, err = KStr.ParseUrl(strHello, -1)
	assert.Nil(t, err)

	//错误的URL
	res, err = KStr.ParseUrl(tesUrl02, -1)
	assert.NotNil(t, err)
	assert.Empty(t, res)

	res, err = KStr.ParseUrl(tesUrl01, 1)
	_, chk = res["scheme"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 2)
	_, chk = res["host"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 4)
	_, chk = res["port"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 8)
	_, chk = res["user"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 16)
	_, chk = res["pass"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 32)
	_, chk = res["path"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 64)
	_, chk = res["query"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 128)
	_, chk = res["fragment"]
	assert.True(t, chk)
}

func BenchmarkUrl_ParseUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.ParseUrl(tesUrl01, -1)
	}
}

func TestUrl_UrlEncodeUrlDecode(t *testing.T) {
	var res1, res2 string
	var err error

	res1 = KStr.UrlEncode(tesStr1)
	res2, err = KStr.UrlDecode(res1)
	assert.Equal(t, res2, tesStr1)
	assert.Nil(t, err)
}

func BenchmarkUrl_UrlEncode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.UrlEncode(tesStr1)
	}
}

func BenchmarkUrl_UrlDecode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.UrlDecode(tesStr2)
	}
}

func TestUrl_RawUrlEncodeRawUrlDecode(t *testing.T) {
	var res1, res2 string
	var err error

	res1 = KStr.RawUrlEncode(tesStr3)
	res2, err = KStr.RawUrlDecode(res1)
	assert.Equal(t, res2, tesStr3)
	assert.Nil(t, err)
}

func BenchmarkUrl_RawUrlEncode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.RawUrlEncode(tesStr3)
	}
}

func BenchmarkUrl_RawUrlDecode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.RawUrlDecode(tesStr4)
	}
}

func TestUrl_HttpBuildQuery(t *testing.T) {
	var res string
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")

	res = KStr.HttpBuildQuery(params)
	assert.Contains(t, res, "&")
}

func BenchmarkUrl_HttpBuildQuery(b *testing.B) {
	b.ResetTimer()
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")
	for i := 0; i < b.N; i++ {
		KStr.HttpBuildQuery(params)
	}
}
