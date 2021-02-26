package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_ArrayChunk(t *testing.T) {
	size := 3
	res := KArr.ArrayChunk(ssSingle, size)
	assert.Equal(t, 4, len(res))

	item := res[0]
	assert.Equal(t, size, len(item))

	KArr.ArrayChunk([]int{}, 1)
}

func TestArray_ArrayChunk_PanicSize(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayChunk]`size cannot be")
	}()
	KArr.ArrayChunk(ssSingle, 0)
}

func TestArray_ArrayChunk_PanicType(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayChunk]`arr type must be")
	}()
	KArr.ArrayChunk(strHello, 2)
}

func BenchmarkArray_ArrayChunk(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayChunk(ssSingle, 3)
	}
}

func TestArray_ArrayColumn_Struct(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayColumn]`arr type must be")
	}()

	var res []interface{}

	res = KArr.ArrayColumn(crowd, "Name")
	assert.NotEmpty(t, res)

	res = KArr.ArrayColumn(*orgS1, "Age")
	assert.NotEmpty(t, res)

	res = KArr.ArrayColumn(*orgS1, "age")
	assert.Empty(t, res)

	// type err
	KArr.ArrayColumn(orgS1, "Age")
}

func TestArray_ArrayColumn_Map(t *testing.T) {
	var arr map[string]interface{}
	var res []interface{}

	_ = KStr.JsonDecode([]byte(personsJson), &arr)

	res = KArr.ArrayColumn(arr, "Name")
	assert.Empty(t, res)

	res = KArr.ArrayColumn(arr, "name")
	assert.NotEmpty(t, res)

	//新元素类型错误
	arr["person5"] = strHello
	res2 := KArr.ArrayColumn(arr, "name")
	assert.Equal(t, len(res), len(res2))
}

func BenchmarkArray_ArrayColumn(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KArr.ArrayColumn(crowd, "Name")
	}
}

func TestArray_SlicePush_SlicePop(t *testing.T) {
	var ss []interface{}
	var item interface{}

	lenght := KArr.SlicePush(&ss, slItf...)
	assert.Greater(t, lenght, 1)

	for i := 0; i < lenght; i++ {
		item = KArr.SlicePop(&ss)
		assert.NotEmpty(t, item)
	}
}

func BenchmarkArray_SlicePush(b *testing.B) {
	var ss []interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss = nil
		KArr.SlicePush(&ss, slItf...)
	}
}

func BenchmarkArray_SlicePop(b *testing.B) {
	var ss [][]interface{}
	var sub []interface{}
	for j := 0; j < 10000000; j++ {
		sub = nil
		copy(sub, slItf)
		ss = append(ss, sub)
	}

	b.ResetTimer()
	for _, item := range ss {
		for i := 0; i < len(item); i++ {
			KArr.SlicePop(&item)
		}
	}
}

func TestArray_SliceUnshift_SliceShift(t *testing.T) {
	var ss []interface{}
	var item interface{}
	lenght := KArr.SliceUnshift(&ss, slItf...)
	assert.Greater(t, lenght, 1)

	for i := 0; i < lenght; i++ {
		item = KArr.SliceShift(&ss)
		assert.NotEmpty(t, item)
	}
}

func BenchmarkArray_SliceUnshift(b *testing.B) {
	var ss []interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss = nil
		KArr.SliceUnshift(&ss, slItf...)
	}
}

func BenchmarkArray_SliceShift(b *testing.B) {
	var ss [][]interface{}
	var sub []interface{}
	for j := 0; j < 10000000; j++ {
		sub = nil
		copy(sub, slItf)
		ss = append(ss, sub)
	}

	b.ResetTimer()
	for _, item := range ss {
		for i := 0; i < len(item); i++ {
			KArr.SliceShift(&item)
		}
	}
}

func TestArray_ArrayKeyExists(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayKeyExists]`arr type must be")
	}()

	chk1 := KArr.ArrayKeyExists(len(naturalArr)-1, naturalArr)
	assert.True(t, chk1)

	chk2 := KArr.ArrayKeyExists(len(slItf)-1, slItf)
	assert.True(t, chk2)

	chk3 := KArr.ArrayKeyExists("Name", personS1)
	chk4 := KArr.ArrayKeyExists("name", personS1)
	assert.True(t, chk3)
	assert.False(t, chk4)

	var persons map[string]interface{}
	_ = KStr.JsonDecode([]byte(personsJson), &persons)
	chk5 := KArr.ArrayKeyExists("person1", persons)
	chk6 := KArr.ArrayKeyExists("Age", persons)
	assert.True(t, chk5)
	assert.False(t, chk6)

	var key interface{}
	chk7 := KArr.ArrayKeyExists(key, persons)
	assert.False(t, chk7)

	KArr.ArrayKeyExists(1, nil)
}

func BenchmarkArray_ArrayKeyExists_Slice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists(1, naturalArr)
	}
}

func BenchmarkArray_ArrayKeyExists_Struct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists("Name", personS1)
	}
}

func BenchmarkArray_ArrayKeyExists_Map(b *testing.B) {
	var persons map[string]interface{}
	_ = KStr.JsonDecode([]byte(personsJson), &persons)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists("person1", persons)
	}
}

func TestArray_ArrayReverse(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayReverse]`arr type must be")
	}()

	res1 := KArr.ArrayReverse(naturalArr)
	itm1 := KArr.SlicePop(&res1)
	assert.Equal(t, 0, itm1)

	res2 := KArr.ArrayReverse(ssSingle)
	itm2 := KArr.SlicePop(&res2)
	assert.Equal(t, "a", itm2)

	KArr.ArrayReverse(strHello)
}

func BenchmarkArray_ArrayReverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayReverse(naturalArr)
	}
}

func TestArray_Implode(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[Implode]`arr type must be")
	}()

	//数组
	res1 := KArr.Implode(",", naturalArr)
	assert.Contains(t, res1, "0,1,2,3,4,5,6,7,8,9,10")

	//切片
	res2 := KArr.Implode(",", ssSingle)
	assert.Contains(t, res2, "a,b,c,d,e,f,g,h,i,j,k")

	//结构体
	res3 := KArr.Implode(",", personS1)
	assert.NotEmpty(t, res3)

	//map
	res4 := KArr.Implode(",", strMp1)
	assert.NotEmpty(t, res4)

	KArr.Implode(",", strHello)
}

func BenchmarkArray_Implode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.Implode(",", naturalArr)
	}
}

func TestArray_JoinStrings(t *testing.T) {
	res := KArr.JoinStrings(",", ssSingle)
	assert.Contains(t, res, "a,b,c,d,e,f,g,h,i,j,k")
}

func BenchmarkArray_JoinStrings(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.JoinStrings(",", ssSingle)
	}
}

func TestArray_JoinInts(t *testing.T) {
	ints := naturalArr[:]
	res := KArr.JoinInts(",", ints)
	assert.Contains(t, res, "0,1,2,3,4,5,6,7,8,9,10")
}

func BenchmarkArray_JoinInts(b *testing.B) {
	b.ResetTimer()
	ints := naturalArr[:]
	for i := 0; i < b.N; i++ {
		KArr.JoinInts(",", ints)
	}
}

func TestArray_UniqueInts(t *testing.T) {
	sl := naturalArr[:]
	sl = append(sl, 1, 2, 3, 4, 5, 6)
	res := KArr.UniqueInts(sl)
	assert.Equal(t, len(naturalArr), len(res))
}

func BenchmarkArray_UniqueInts(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.UniqueInts(intSlc)
	}
}

func TestArray_Unique64Ints(t *testing.T) {
	res := KArr.Unique64Ints(int64Slc)
	assert.Less(t, len(res), len(int64Slc))
}

func BenchmarkArray_Unique64Ints(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.Unique64Ints(int64Slc)
	}
}

func TestArray_UniqueStrings(t *testing.T) {
	sl := ssSingle[:]
	sl = append(sl, "a", "b", "c", "d", "e")
	res := KArr.UniqueStrings(sl)
	assert.Equal(t, len(ssSingle), len(res))
}

func BenchmarkArray_UniqueStrings(b *testing.B) {
	b.ResetTimer()
	sl := ssSingle[:]
	sl = append(sl, "a", "b", "c", "d", "e")
	for i := 0; i < b.N; i++ {
		KArr.UniqueStrings(sl)
	}
}

func TestArray_ArrayDiff(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayDiff]`arr1,arr2 type must be")
	}()

	var res, res2 map[interface{}]interface{}

	//数组-切片
	res = KArr.ArrayDiff(strSl1, strSl2, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayDiff(strSl1, strSl2, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayDiff(strSl1, strSl2, COMPARE_BOTH_KEYVALUE)
	assert.Greater(t, len(res2), len(res))

	res = KArr.ArrayDiff(strSlEmp, strSl1, COMPARE_ONLY_VALUE)
	assert.Empty(t, res)

	//数组-字典
	res = KArr.ArrayDiff(strSl1, strMp1, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayDiff(strSl1, strMp1, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayDiff(strSl1, strMp1, COMPARE_BOTH_KEYVALUE)
	assert.Greater(t, len(res2), len(res))

	res = KArr.ArrayDiff(strSlEmp, strMp1, COMPARE_ONLY_VALUE)
	assert.Empty(t, res)

	//字典-数组
	res = KArr.ArrayDiff(strMp1, strSl1, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayDiff(strMp1, strSl1, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayDiff(strMp1, strSl1, COMPARE_BOTH_KEYVALUE)
	assert.Greater(t, len(res2), len(res))

	res = KArr.ArrayDiff(strMpEmp, strSl1, COMPARE_ONLY_VALUE)
	assert.Empty(t, res)

	//字典-字典
	res = KArr.ArrayDiff(strMp1, strMp2, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayDiff(strMp1, strMp2, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayDiff(strMp1, strMp2, COMPARE_BOTH_KEYVALUE)
	assert.NotEmpty(t, res2)

	KArr.ArrayDiff(strHello, 1234, COMPARE_ONLY_VALUE)
}

func BenchmarkArray_ArrayDiff_A1A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayDiff(strSl1, strSl2, COMPARE_ONLY_VALUE)
	}
}

func BenchmarkArray_ArrayDiff_A1M(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayDiff(strSl1, strMp1, COMPARE_ONLY_VALUE)
	}
}

func BenchmarkArray_ArrayDiff_M1A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayDiff(strMp1, strSl1, COMPARE_ONLY_VALUE)
	}
}

func BenchmarkArray_ArrayDiff_M1M(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayDiff(strMp1, strMp2, COMPARE_ONLY_VALUE)
	}
}

func TestArray_ArrayIntersect(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayIntersect]`arr1,arr2 type must be")
	}()

	var res, res2 map[interface{}]interface{}

	//数组-切片
	res = KArr.ArrayIntersect(strSl1, strSl2, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayIntersect(strSl1, strSl2, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayIntersect(strSl1, strSl2, COMPARE_BOTH_KEYVALUE)
	assert.Less(t, len(res2), len(res))

	res = KArr.ArrayIntersect(strSlEmp, strSl1, COMPARE_ONLY_VALUE)
	assert.Empty(t, res)

	//数组-字典
	res = KArr.ArrayIntersect(strSl1, strMp1, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayIntersect(strSl1, strMp1, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayIntersect(strSl1, strMp1, COMPARE_BOTH_KEYVALUE)
	assert.Less(t, len(res2), len(res))

	res = KArr.ArrayIntersect(strSlEmp, strMp1, COMPARE_ONLY_VALUE)
	assert.Empty(t, res)

	//字典-数组
	res = KArr.ArrayIntersect(strMp1, strSl1, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayIntersect(strMp1, strSl1, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayIntersect(strMp1, strSl1, COMPARE_BOTH_KEYVALUE)
	assert.Less(t, len(res2), len(res))

	res = KArr.ArrayIntersect(strMpEmp, strSl1, COMPARE_ONLY_VALUE)
	assert.Empty(t, res)

	//字典-字典
	res = KArr.ArrayIntersect(strMp1, strMp2, COMPARE_ONLY_VALUE)
	assert.NotEmpty(t, res)

	res = KArr.ArrayIntersect(strMp1, strMp2, COMPARE_ONLY_KEY)
	assert.NotEmpty(t, res)

	res2 = KArr.ArrayIntersect(strMp1, strMp2, COMPARE_BOTH_KEYVALUE)
	assert.NotEmpty(t, res2)

	KArr.ArrayIntersect(strHello, 1234, COMPARE_ONLY_VALUE)
}

func BenchmarkArray_ArrayIntersect_A1A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayIntersect(strSl1, strSl2, COMPARE_ONLY_VALUE)
	}
}

func BenchmarkArray_ArrayIntersect_A1M(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayIntersect(strSl1, strMp1, COMPARE_ONLY_VALUE)
	}
}

func BenchmarkArray_ArrayIntersect_M1A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayIntersect(strMp1, strSl1, COMPARE_ONLY_VALUE)
	}
}

func BenchmarkArray_ArrayIntersect_M1M(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayIntersect(strMp1, strMp2, COMPARE_ONLY_VALUE)
	}
}

func TestArray_ArrayUnique(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayUnique]`arr type must be")
	}()

	var res map[interface{}]interface{}

	//数组切片
	res = KArr.ArrayUnique(intSlc)
	assert.Less(t, len(res), len(intSlc))

	//字典
	res = KArr.ArrayUnique(colorMp)
	assert.Less(t, len(res), len(colorMp))

	KArr.ArrayUnique(strHello)
}

func BenchmarkArray_ArrayUnique_Arr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayUnique(intSlc)
	}
}

func BenchmarkArray_ArrayUnique_Map(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayUnique(colorMp)
	}
}

func TestArray_ArraySearchItem(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArraySearchItem]`arr type must be")
	}()

	var res interface{}

	//子元素为字典
	cond1 := map[string]interface{}{"age": 21, "naction": "cn"}
	res = KArr.ArraySearchItem(personMps, cond1)
	assert.NotEmpty(t, res)

	//子元素为结构体
	cond2 := map[string]interface{}{"Gender": false}
	res = KArr.ArraySearchItem(perStuMps, cond2)
	assert.NotEmpty(t, res)

	KArr.ArraySearchItem(strHello, map[string]interface{}{"a": 1})
}

func BenchmarkArray_ArraySearchItem_Arr(b *testing.B) {
	b.ResetTimer()
	cond := map[string]interface{}{"age": 21, "naction": "cn"}
	for i := 0; i < b.N; i++ {
		KArr.ArraySearchItem(personMps, cond)
	}
}

func BenchmarkArray_ArraySearchItem_Map(b *testing.B) {
	b.ResetTimer()
	cond := map[string]interface{}{"Gender": false}
	for i := 0; i < b.N; i++ {
		KArr.ArraySearchItem(perStuMps, cond)
	}
}

func TestArray_ArraySearchMutil(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArraySearchMutil]`arr type must be")
	}()

	var res []interface{}

	//子元素为字典
	cond1 := map[string]interface{}{"age": 21, "naction": "cn"}
	res = KArr.ArraySearchMutil(personMps, cond1)
	assert.NotEmpty(t, res)

	//子元素为结构体
	cond2 := map[string]interface{}{"Gender": false}
	res = KArr.ArraySearchMutil(perStuMps, cond2)
	assert.NotEmpty(t, res)

	KArr.ArraySearchMutil(strHello, map[string]interface{}{"a": 1})
}

func BenchmarkArray_ArraySearchMutil_Arr(b *testing.B) {
	b.ResetTimer()
	cond := map[string]interface{}{"age": 21, "naction": "cn"}
	for i := 0; i < b.N; i++ {
		KArr.ArraySearchMutil(personMps, cond)
	}
}

func BenchmarkArray_ArraySearchMutil_Map(b *testing.B) {
	b.ResetTimer()
	cond := map[string]interface{}{"Gender": false}
	for i := 0; i < b.N; i++ {
		KArr.ArraySearchMutil(perStuMps, cond)
	}
}

func TestArray_ArrayShuffle(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayShuffle]`arr type must be")
	}()

	var res []interface{}
	res = KArr.ArrayShuffle(naturalArr)
	assert.NotEqual(t, toStr(res), toStr(naturalArr))

	res = KArr.ArrayShuffle(ssSingle)
	assert.NotEqual(t, toStr(res), toStr(ssSingle))

	KArr.ArrayShuffle(strHello)
}

func BenchmarkArray_ArrayShuffle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayShuffle(naturalArr)
	}
}

func TestArray_IsEqualArray(t *testing.T) {
	var res bool

	res = KArr.IsEqualArray(naturalArr, ssSingle)
	assert.False(t, res)

	arr := KArr.ArrayShuffle(ssSingle)
	res = KArr.IsEqualArray(arr, ssSingle)
	assert.True(t, res)
}

func TestArray_IsEqualArray_Panic_Expected(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[IsEqualArray]`expected type must be")
	}()

	KArr.IsEqualArray(strHello, ssSingle)
}

func TestArray_IsEqualArray_Panic_Actual(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[IsEqualArray]`actual type must be")
	}()

	KArr.IsEqualArray(ssSingle, strHello)
}

func BenchmarkArray_IsEqualArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsEqualArray(naturalArr, ssSingle)
	}
}

func TestArray_Length(t *testing.T) {
	var res int
	res = KArr.Length(naturalArr)
	assert.Equal(t, res, len(naturalArr))

	//非数组或切片
	res = KArr.Length(strHello)
	assert.Equal(t, -1, res)
}

func BenchmarkArray_Length(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.Length(naturalArr)
	}
}

func TestArray_IsArray(t *testing.T) {
	var res bool

	res = KArr.IsArray(naturalArr)
	assert.True(t, res)

	res = KArr.IsArray(intSlc)
	assert.False(t, res)

	res = KArr.IsArray(strHello)
	assert.False(t, res)
}

func BenchmarkArray_IsArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsArray(naturalArr)
	}
}

func TestArray_IsSlice(t *testing.T) {
	var res bool

	res = KArr.IsSlice(intSlc)
	assert.True(t, res)

	res = KArr.IsSlice(naturalArr)
	assert.False(t, res)

	res = KArr.IsSlice(strHello)
	assert.False(t, res)
}

func BenchmarkArray_IsSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsSlice(intSlc)
	}
}

func TestArray_IsArrayOrSlice(t *testing.T) {
	var res bool

	res = KArr.IsArrayOrSlice(intSlc)
	assert.True(t, res)

	res = KArr.IsArrayOrSlice(naturalArr)
	assert.True(t, res)

	res = KArr.IsArrayOrSlice(strHello)
	assert.False(t, res)
}

func BenchmarkArray_IsArrayOrSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsArrayOrSlice(intSlc)
	}
}

func TestArray_IsMap(t *testing.T) {
	var res bool

	res = KArr.IsMap(colorMp)
	assert.True(t, res)

	res = KArr.IsMap(strMpEmp)
	assert.True(t, res)

	res = KArr.IsMap(naturalArr)
	assert.False(t, res)
}

func BenchmarkArray_IsMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsMap(colorMp)
	}
}

func TestArray_IsStruct(t *testing.T) {
	var res bool

	res = KArr.IsStruct(personS1)
	assert.True(t, res)

	res = KArr.IsStruct(naturalArr)
	assert.False(t, res)
}

func BenchmarkArray_IsStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsStruct(personS1)
	}
}

func TestArray_DeleteSliceItems(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[DeleteSliceItems]`val type must be")
	}()

	var res []interface{}
	var del int

	res, del = KArr.DeleteSliceItems(naturalArr, 3, 5, 8)
	assert.Greater(t, len(naturalArr), len(res))

	res, del = KArr.DeleteSliceItems(int64Slc, 2, 4, 9)
	assert.Greater(t, del, 0)

	_, _ = KArr.DeleteSliceItems(strHello, 3, 5, 8)
}

func BenchmarkArray_DeleteSliceItems(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.DeleteSliceItems(naturalArr, 3, 5, 8)
	}
}

func TestArray_InArray(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[InArray]`haystack type must be")
	}()

	var res bool

	res = KArr.InArray(9, naturalArr)
	assert.True(t, res)

	res = KArr.InArray(personMp3, personMps)
	assert.True(t, res)

	res = KArr.InArray(personMp3, crowd)
	assert.False(t, res)

	KArr.InArray(9, strHello)
}

func BenchmarkArray_InArray_Arr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.InArray(9, naturalArr)
	}
}

func BenchmarkArray_InArray_Map(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.InArray(personMp3, personMps)
	}
}

func TestArray_InIntSlice(t *testing.T) {
	var res bool

	res = KArr.InIntSlice(9, intSlc)
	assert.True(t, res)

	res = KArr.InIntSlice(99, intSlc)
	assert.False(t, res)
}

func BenchmarkArray_InIntSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.InIntSlice(9, intSlc)
	}
}

func TestArray_InInt64Slice(t *testing.T) {
	var res bool

	res = KArr.InInt64Slice(9, int64Slc)
	assert.True(t, res)

	res = KArr.InInt64Slice(99, int64Slc)
	assert.False(t, res)
}

func BenchmarkArray_InInt64Slice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.InInt64Slice(9, int64Slc)
	}
}

func TestArray_InStringSlice(t *testing.T) {
	var res bool

	res = KArr.InStringSlice("c", ssSingle)
	assert.True(t, res)

	res = KArr.InStringSlice("w", ssSingle)
	assert.False(t, res)
}

func BenchmarkArray_InStringSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.InStringSlice("c", ssSingle)
	}
}

func TestArray_SliceFill(t *testing.T) {
	var res []interface{}

	res = KArr.SliceFill(strHello, 9)
	assert.Equal(t, 9, len(res))

	res = KArr.SliceFill(strHello, 0)
	assert.Empty(t, res)
}

func BenchmarkArray_SliceFill(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.SliceFill(strHello, 9)
	}
}
