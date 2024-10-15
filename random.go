package random

import (
	"math/rand"
	"reflect"
	"slices"
	"unsafe"
)

// IntBetween returns random int in [min,max]
func IntBetween(min, max int) int {
	return IntN(max-min+1) + min
}

func IntDice(faces int) int {
	return IntN(faces) + 1
}

func IntDiceN(faces, n int) int {
	tmp := 0
	for i := 0; i < n; i = i + 1 {
		tmp = tmp + IntDice(faces)
	}
	return tmp
}

func IntDiceNBetween(n, min, max int) int {
	tmp := 0
	for i := 0; i < n; i = i + 1 {
		tmp = tmp + IntBetween(min, max)
	}
	return tmp
}

func IntDiceNBetweenTimes(n, argMin, argMax, times int) int {
	r := 0
	for i := 0; i < times; i = i + 1 {
		r = max(r, IntDiceNBetween(n, argMin, argMax))
	}
	return r
}

// IntN returns random int in [0,n)
func IntN(n int) int {
	return rand.Intn(n)
}

func IntPositive() int {
	return rand.Int()
}

func Bool() bool {
	return IntN(2) == 0
}

func Int() int {
	return int(rand.Uint64())
}

func Int8() int8 {
	return int8(Uint32())
}

func Int16() int16 {
	return int16(Uint32())
}

func Int32() int32 {
	return int32(Uint32())
}

func Int64() int64 {
	return int64(Uint64())
}

func Uint() uint {
	return uint(Uint64())
}

func Uint8() uint8 {
	return uint8(Uint32())
}

func Uint16() uint16 {
	return uint16(Uint32())
}

func Uint32() uint32 {
	return rand.Uint32()
}

func Uint64() uint64 {
	return rand.Uint64()
}

func Float3201() float32 {
	return rand.Float32()
}

func Float6401() float64 {
	return rand.Float64()
}

func Float32() float32 {
	return float32(rand.NormFloat64())
}

func Float64() float64 {
	return rand.NormFloat64()
}

func Complex64() complex64 {
	return complex(Float32(), Float32())
}

func Complex128() complex128 {
	return complex(Float64(), Float64())
}

func Uintptr() uintptr {
	return uintptr(Uint64())
}

func String(length int) string {
	str := make([]rune, length)

	for i := 0; i < len(str); i = i + 1 {
		b := Int32()
		str[i] = b
	}

	return string(str)
}

const (
	StringFilteredNumbers byte = 1 << iota //47-57
	StringFilteredLetters                  //65-90, 97-122
	StringFilteredSigns                    //33-47, 58-64, 91-96, 123-126
	StringFilteredSpacers                  //9, 10, 13, 32
)

var (
	numbers        = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	letters        = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	capitalLetters = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	signs          = []byte{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
	spacer         = []byte{' ', '\t', '\n', '\r'}
)

func StringFiltered(flags byte, length int) string {
	str := ""

	for {
		if len([]rune(str)) == length {
			break
		} else if len([]rune(str)) > length {
			str = ""
		}

		nextByte := Uint8()

		for {
			if (flags&StringFilteredNumbers != 0 && slices.Contains(numbers, nextByte)) ||
				(flags&StringFilteredLetters != 0 && slices.Contains(letters, nextByte) && slices.Contains(capitalLetters, nextByte)) ||
				(flags&StringFilteredSigns != 0 && slices.Contains(signs, nextByte)) ||
				(flags&StringFilteredSpacers != 0 && slices.Contains(spacer, nextByte)) {
				break
			} else {
				nextByte = Uint8()
			}
		}

		str = str + string([]byte{nextByte})

	}

	return str
}

func Pick[T any](elems ...T) T {
	return elems[IntN(len(elems))]
}

func Kind() reflect.Kind {
	return reflect.Kind(IntN(27))
}

var minArrayLen = 0
var maxArrayLen = 20

var minSliceLen = 0
var maxSliceLen = 20

var minFuncArg = 0
var maxFuncArg = 20

var minFuncOut = 0
var maxFuncOut = 20

var minMapEntries = 0
var maxMapEntries = 20

var minStringLen = 0
var maxStringLen = 20

var minStructFields = 0
var maxStructFields = 20

var minBufferLen = 1
var maxBufferLen = 20

var minStructNameLen = 1
var maxStructNameLen = 20

var minPkgPathLen = 1
var maxPkgPathLen = 20

var maxDepthType = reflect.TypeOf(true)

func Any(maxDepth int) any {
	return Get(Type(maxDepth), maxDepth)
}

func Get(t reflect.Type, maxDepth int) any {
	switch t.Kind() {
	case reflect.Invalid:
		return nil //todo (sense?)

	case reflect.Bool:
		return Bool()

	case reflect.Int:
		return Int()

	case reflect.Int8:
		return Int8()

	case reflect.Int16:
		return Int16()

	case reflect.Int32:
		return Int32()

	case reflect.Int64:
		return Int64()

	case reflect.Uint:
		return Uint()

	case reflect.Uint8:
		return Uint8()

	case reflect.Uint16:
		return Uint16()

	case reflect.Uint32:
		return Uint32()

	case reflect.Uint64:
		return Uint64()

	case reflect.Uintptr:
		return Uintptr()

	case reflect.Float32:
		return Float32()

	case reflect.Float64:
		return Float64()

	case reflect.Complex64:
		return Complex64()

	case reflect.Complex128:
		return Complex128()

	case reflect.Array:
		arr := reflect.New(t).Elem()
		for i := 0; i < arr.Len(); i = i + 1 {
			arr.Index(i).Set(reflect.ValueOf(Get(t.Elem(), maxDepth-1)))
		}
		return arr.Interface()

	case reflect.Chan:
		buffered := Bool()
		bufferLen := 0
		if buffered {
			bufferLen = IntBetween(minBufferLen, maxBufferLen)
		}

		//reflect.MakeChan cant make unidirectional channels
		if t.ChanDir() != reflect.BothDir {
			tBothChan := reflect.ChanOf(reflect.BothDir, t.Elem())
			vBothChan := reflect.MakeChan(tBothChan, bufferLen)
			vChan := vBothChan.Convert(t)
			return vChan.Interface()
		}

		vChan := reflect.MakeChan(t, bufferLen)
		return vChan.Interface()

	case reflect.Func:
		vFunc := reflect.New(t).Elem()
		return vFunc.Interface()

	case reflect.Interface:
		vInter := reflect.New(t).Elem()
		vInter.Set(reflect.ValueOf(Any(maxDepth - 1))) //todo -1 correct?
		return vInter.Interface()

	case reflect.Map:
		vMap := reflect.MakeMap(t)
		entries := IntBetween(minMapEntries, maxMapEntries)
		for i := 0; i < entries; i++ {
			key := reflect.ValueOf(Get(t.Key(), maxDepth-1))
			for !key.Type().Comparable() { //todo interface types can lead to not comparable after Type call
				key = reflect.ValueOf(Get(t.Key(), maxDepth-1))
			}
			vMap.SetMapIndex(key, reflect.ValueOf(Get(t.Elem(), maxDepth-1)))
		}
		return vMap.Interface()

	case reflect.Pointer:
		ptr := reflect.New(t.Elem())
		ptr.Elem().Set(reflect.ValueOf(Get(t.Elem(), maxDepth-1)))
		return ptr.Interface()

	case reflect.Slice:
		length := IntBetween(minSliceLen, maxSliceLen)
		slice := reflect.MakeSlice(t, length, length)
		for i := 0; i < length; i = i + 1 {
			slice.Index(i).Set(reflect.ValueOf(Get(t.Elem(), maxDepth-1)))
		}
		return slice.Interface()

	case reflect.String:
		return String(IntBetween(minStringLen, maxStringLen))

	case reflect.Struct:
		vStruct := reflect.New(t).Elem()
		for i := 0; i < vStruct.NumField(); i++ {
			if !vStruct.Type().Field(i).IsExported() { //todo to evil?
				vField := reflect.NewAt(vStruct.Type().Field(i).Type, vStruct.Field(i).Addr().UnsafePointer())
				vField.Elem().Set(reflect.ValueOf(Get(t.Field(i).Type, maxDepth-1)))
			} else {
				vStruct.Field(i).Set(reflect.ValueOf(Get(t.Field(i).Type, maxDepth-1)))
			}
		}

		return vStruct.Interface()

	case reflect.UnsafePointer:
		return unsafe.Pointer(Uintptr())
	}

	return nil
}

func Type(maxDepth int) reflect.Type {
	if maxDepth <= 0 {
		return maxDepthType
	}
	typ := Kind()
	switch typ {
	case reflect.Invalid:
		return Type(maxDepth) //todo
	case reflect.Bool:
		return reflect.TypeFor[bool]()
	case reflect.Int:
		return reflect.TypeFor[int]()
	case reflect.Int8:
		return reflect.TypeFor[int8]()
	case reflect.Int16:
		return reflect.TypeFor[int16]()
	case reflect.Int32:
		return reflect.TypeFor[int32]()
	case reflect.Int64:
		return reflect.TypeFor[int64]()
	case reflect.Uint:
		return reflect.TypeFor[uint]()
	case reflect.Uint8:
		return reflect.TypeFor[uint8]()
	case reflect.Uint16:
		return reflect.TypeFor[uint16]()
	case reflect.Uint32:
		return reflect.TypeFor[uint32]()
	case reflect.Uint64:
		return reflect.TypeFor[uint64]()
	case reflect.Uintptr:
		return reflect.TypeFor[uintptr]()
	case reflect.Float32:
		return reflect.TypeFor[float32]()
	case reflect.Float64:
		return reflect.TypeFor[float64]()
	case reflect.Complex64:
		return reflect.TypeFor[complex64]()
	case reflect.Complex128:
		return reflect.TypeFor[complex128]()
	case reflect.Array:
		return reflect.ArrayOf(IntBetween(minArrayLen, maxArrayLen), Type(maxDepth-1))
	case reflect.Chan:
		dir := Pick(reflect.RecvDir, reflect.SendDir, reflect.BothDir)
		tElem := Type(maxDepth - 1)
		for tElem.Size() > 1<<16 {
			tElem = Type(maxDepth - 1)
		}
		return reflect.ChanOf(dir, tElem)
	case reflect.Func:
		variadic := Bool()
		tArgs := make([]reflect.Type, IntBetween(minFuncArg, maxFuncArg))
		for i := range tArgs {
			tArgs[i] = Type(maxDepth - 1)
		}
		if variadic {
			if len(tArgs) < 1 {
				tArgs = append(tArgs, reflect.SliceOf(Type(maxDepth-1)))
			} else {
				tArgs[len(tArgs)-1] = reflect.SliceOf(Type(maxDepth - 1))
			}
		}
		tOut := make([]reflect.Type, IntBetween(minFuncOut, maxFuncOut))
		for i := range tOut {
			tOut[i] = Type(maxDepth - 1)
		}
		return reflect.FuncOf(tArgs, tOut, variadic)
	case reflect.Interface:
		return reflect.TypeFor[any]()

	case reflect.Map:
		tKey := Type(maxDepth - 1)
		for !tKey.Comparable() {
			tKey = Type(maxDepth - 1)
		}
		tValue := Type(maxDepth - 1)
		return reflect.MapOf(tKey, tValue)

	case reflect.Pointer:
		return reflect.PointerTo(Type(maxDepth - 1))

	case reflect.Slice:
		return reflect.SliceOf(Type(maxDepth - 1))

	case reflect.String:
		return reflect.TypeFor[string]()

	case reflect.Struct:
		fields := make([]reflect.StructField, IntBetween(minStructFields, maxStructFields))
		offset := uintptr(0)
		pkgPath := StringFiltered(
			StringFilteredLetters,
			IntBetween(minPkgPathLen, maxPkgPathLen))
		names := make(map[string]bool)
		for i := range fields {
			fields[i].Type = Type(maxDepth - 1)

			// anonymous fields are not supported
			fields[i].Anonymous = false

			// generate initial field name
			fields[i].Name = StringFiltered(
				StringFilteredLetters,
				IntBetween(minStructNameLen, maxStructNameLen))

			// no duplicates
			for names[fields[i].Name] {
				fields[i].Name = StringFiltered(
					StringFilteredLetters,
					IntBetween(minStructNameLen, maxStructNameLen))
			}

			// register field name to avoid duplicates
			names[fields[i].Name] = true

			// exported field has no PkgPath
			fields[i].PkgPath = pkgPath
			if len([]rune(fields[i].Name)) > 0 {
				first := []rune(fields[i].Name)[0]
				if first >= 'A' && first <= 'Z' {
					fields[i].PkgPath = ""
				}
			}

			fields[i].Tag = reflect.StructTag(String(IntBetween(minStringLen, maxStringLen)))
			fields[i].Offset = offset
			fields[i].Index = []int{i}
			offset += fields[i].Type.Size()
		}
		return reflect.StructOf(fields)

	case reflect.UnsafePointer:
		return reflect.TypeFor[unsafe.Pointer]()
	}

	return nil
}
