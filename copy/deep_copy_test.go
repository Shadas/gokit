package copy

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestDeepCopyPointerToStruct(t *testing.T) {
	type Foo struct {
		Bar int
	}
	f := &Foo{Bar: 42}
	cpy := DeepCopy(f)
	if f == cpy {
		t.Error("f, cpy should not be the same pointer")
	}
	if !reflect.DeepEqual(f, cpy) {
		t.Error("the value of f and cpy should be the same")
	}
}

func TestUnexportedFields(t *testing.T) {
	u := &Unexported{
		A:  "A",
		B:  123,
		aa: "aa",
		bb: 456,
		cc: []int{1, 2, 3, 4},
		dd: map[string]string{"xx": "yy"},
	}
	cpy := DeepCopy(u).(*Unexported)
	if cpy == u {
		t.Error("u, cpy should not be the same pointer")
	}
	if u.A != cpy.A {
		t.Error("Unexported.A copy failed")
	}
	if u.B != cpy.B {
		t.Error("Unexported.B copy failed")
	}
	if cpy.aa != "" {
		t.Error("Unexported.aa shouldn't be set")
	}
	if cpy.bb != 0 {
		t.Error("Unexported.bb shouldn't be set")
	}
	if cpy.cc != nil {
		t.Error("Unexported.cc shouldn't be set")
	}
	if cpy.dd != nil {
		t.Error("Unexported.dd shouldn't be set")
	}
}

type Unexported struct {
	A  string
	B  int
	aa string
	bb int
	cc []int
	dd map[string]string
}

func TestDeepCopySlice(t *testing.T) {
	// []string
	strings := []string{"A", "B", "C"}
	copyS := DeepCopy(strings).([]string)
	if (*reflect.SliceHeader)(unsafe.Pointer(&strings)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&copyS)).Data {
		t.Error("[]string: data pointers should be different but not")
	}
	if len(copyS) != len(strings) {
		t.Error("[]string: lengths are different")
	}
	for i, v := range strings {
		if v != copyS[i] {
			t.Errorf("[]string: value not same, idx=%d", i)
		}
	}
	for i := range copyS {
		copyS[i] = string(append([]byte(copyS[i]), copyS[i]...))
	}
	for i, v := range strings {
		if v == copyS[i] {
			t.Errorf("[]string: value not different, idx=%d", i)
		}
	}

	// []bool
	bools := []bool{true, false, true, false}
	copyB := DeepCopy(bools).([]bool)
	if (*reflect.SliceHeader)(unsafe.Pointer(&bools)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&copyB)).Data {
		t.Error("[]bool: data pointers should be different but not")
	}
	if len(copyB) != len(bools) {
		t.Error("[]bool: lengths are different")
	}
	for i, v := range bools {
		if v != copyB[i] {
			t.Errorf("[]bool: value not same, idx=%d", i)
		}
	}
	for i := range copyB {
		copyB[i] = !copyB[i]
	}
	for i, v := range bools {
		if v == copyB[i] {
			t.Errorf("[]bool: value not different, idx=%d", i)
		}
	}

	// []byte
	bytes := []byte("hello")
	copyBt := DeepCopy(bytes).([]byte)
	if (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&copyBt)).Data {
		t.Error("[]byte: data pointers should be different but not")
	}
	if len(bytes) != len(copyBt) {
		t.Error("[]byte: lengths are different")
	}
	for i, v := range bytes {
		if v != copyBt[i] {
			t.Errorf("[]byte: value not same, idx=%d", i)
		}
	}
	for i := range copyBt {
		copyBt[i] = copyBt[i] + 1
	}
	for i, v := range bytes {
		if v == copyBt[i] {
			t.Errorf("[]byte: value not different, idx=%d", i)
		}
	}

	// []int
	ints := []int{1, 2, 3, 54}
	copyI := DeepCopy(ints).([]int)
	if (*reflect.SliceHeader)(unsafe.Pointer(&ints)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&copyI)).Data {
		t.Error("[]int: data pointers should be different but not")
	}
	if len(ints) != len(copyI) {
		t.Error("[]int: lengths are different")
	}
	for i, v := range ints {
		if v != copyI[i] {
			t.Errorf("[]int: value not same, idx=%d", i)
		}
	}
	for i := range copyI {
		copyI[i] = copyI[i] + 1
	}
	for i, v := range ints {
		if v == copyI[i] {
			t.Errorf("[]int: value not different, idx=%d", i)
		}
	}

	// []interface{}
	interfaces := []interface{}{10, "x", true, 3.14}
	copyIf := DeepCopy(interfaces).([]interface{})
	if (*reflect.SliceHeader)(unsafe.Pointer(&interfaces)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&copyIf)).Data {
		t.Error("[]interface: data pointers should be different but not")
	}
	if len(interfaces) != len(copyIf) {
		t.Error("[]interface: lengths are different")
	}
	for i, v := range interfaces {
		if v != copyIf[i] {
			t.Errorf("[]interface: value not same, idx=%d", i)
		}
	}
	for i := range copyI {
		copyIf[i] = i
	}
	for i, v := range interfaces {
		if v == copyIf[i] {
			t.Errorf("[]interface: value not different, idx=%d", i)
		}
	}
}

func TestDeepCopyInterfaceUse(t *testing.T) {
	i := &I{A: "self value"}
	cpyI := DeepCopy(i).(*I)
	if cpyI.A != customSpecificValue {
		t.Error("DeepCopy interface failed to use in DeepCopy function")
	}
	ni := &NestI{I: &I{A: "self value"}}
	cpyNI := DeepCopy(ni).(*NestI)
	if cpyNI.I.A != customSpecificValue {
		t.Error("nested: DeepCopy interface failed to use in DeepCopy function")
	}
}

var customSpecificValue = "custom copy value"

type I struct {
	A string
}

func (i *I) DeepCopy() interface{} {
	return &I{A: customSpecificValue}
}

type NestI struct {
	I *I
}
