package pointerinit

// returns str as *str
func StrPointer(s string) *string { return &s }

// returns float32 as *float32
func Float32Pointer(f float32) *float32 { return &f }

// returns float64 as *float64
func Float64Pointer(f float64) *float64 { return &f }

// return int as *int
func IntPointer(i int) *int { return &i }

// return int8 as *int8
func Int8pointer(i int8) *int8 { return &i }

// return int16 as *int16
func Int16pointer(i int16) *int16 { return &i }

// return int32 as *int32
func Int32pointer(i int32) *int32 { return &i }

// return int64 as *int64
func Int64pointer(i int64) *int64 { return &i }

// return uint16 as *uint16
func UInt16pointer(u uint16) *uint16 { return &u }
