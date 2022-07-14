package safemap

import "time"

type Interface interface {
	Delete(key any)
	Load(key any) (value any, ok bool)
	LoadAndDelete(key any) (value any, loaded bool)
	LoadOrStore(key, value any) (actual any, loaded bool)
	Range(f func(key, value any) bool)
	Store(key, value any)

	// Set alias for Store
	Set(key, value any)
	// Get alias for Load
	Get(key any) (value any, ok bool)

	GetString(key any) (value string)
	GetBool(key string) (value bool)

	GetInt8(key string) (value int8)
	GetInt16(key string) (value int16)
	GetInt32(key string) (value int32)
	GetInt(key string) (value int)
	GetInt64(key string) (value int64)

	GetUint8(key string) (value uint8)
	GetUint16(key string) (value uint16)
	GetUint32(key string) (value uint32)
	GetUint(key string) (value uint)
	GetUint64(key string) (value uint64)

	GetFloat32(key string) (value float32)
	GetFloat64(key string) (value float64)

	GetTime(key string) (value time.Time)
	GetDuration(key string) (value time.Duration)

	GetByteSlice(key any) (value []byte)
	GetRuneSlice(key any) (value []rune)
	GetIntSlice(key string) (value []int)
	GetInt64Slice(key string) (value []int64)
	GetUintSlice(key string) (value []uint)
	GetUint64Slice(key string) (value []uint64)
	GetFloat32Slice(key string) (value []float32)
	GetFloat64Slice(key string) (value []float64)
	GetStringSlice(key string) (value []string)

	GetStringMap(key string) (value map[string]any)
	GetStringMapString(key string) (value map[string]string)
	GetStringMapStringSlice(key string) (value map[string][]string)

	Debug() (value map[any]any)
}
