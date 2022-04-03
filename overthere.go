package Pointerinit

// returns str as *str
func StrPointer(s string) *string { return &s }

// returns []str as *[]str
func StrArrPointer(s []string) *[]string { return &s }

// returns float32 as *float32
func Float32Pointer(f float32) *float32 { return &f }

// returns float32 as *float32
func Float32ArrPointer(f []float32) *[]float32 { return &f }

// returns float64 as *float64
func Float64Pointer(f float64) *float64 { return &f }

// returns []float64 as *[]float64
func Float64ArrPointer(f []float64) *[]float64 { return &f }

// return int as *int
func IntPointer(i int) *int { return &i }

// return []int as *[]int
func IntArrPointer(i []int) *[]int { return &i }

// return int8 as *int8
func Int8Pointer(i int8) *int8 { return &i }

// return []int8 as *[]int8
func Int8ArrPointer(i []int8) *[]int8 { return &i }

// return int16 as *int16
func Int16Pointer(i int16) *int16 { return &i }

// return []int16 as *[]int16
func Int16ArrPointer(i []int16) *[]int16 { return &i }

// return int32 as *int32
func Int32Pointer(i int32) *int32 { return &i }

// return []int32 as *[]int32
func Int32ArrPointer(i []int32) *[]int32 { return &i }

// return int64 as *int64
func Int64Pointer(i int64) *int64 { return &i }

// return []int64 as *[]int64
func Int64ArrPointer(i []int64) *[]int64 { return &i }

// return int as *int
func UIntPointer(i uint) *uint { return &i }

// return []int as *[]int
func UIntArrPointer(i []uint) *[]uint { return &i }

// return int8 as *int8
func UInt8Pointer(i uint8) *uint8 { return &i }

// return []int8 as *[]int8
func UInt8ArrPointer(i []uint8) *[]uint8 { return &i }

// return int16 as *int16
func UInt16Pointer(i uint16) *uint16 { return &i }

// return []int16 as *[]int16
func UInt16ArrPointer(i []uint16) *[]uint16 { return &i }

// return int32 as *int32
func UInt32Pointer(i uint32) *uint32 { return &i }

// return []int32 as *[]int32
func UInt32ArrPointer(i []uint32) *[]uint32 { return &i }

// return int64 as *int64
func UInt64Pointer(i uint64) *uint64 { return &i }

// return []int64 as *[]int64
func UInt64ArrPointer(i []uint64) *[]uint64 { return &i }

// return bool as *bool
func BoolPointer(b bool) *bool { return &b }

// return []bool as *[]bool
func BoolArrPointer(b []bool) *[]bool { return &b }

// return struct as *struct
func StructPointer(s struct{}) *struct{} { return &s }

// return []struct as *[]struct
func StructArrPointer(s struct{}) *struct{} { return &s }
