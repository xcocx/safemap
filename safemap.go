package safemap

import (
	"sync"
	"time"
)

type safeMap struct {
	*sync.Map
}

func New() Interface {
	return &safeMap{Map: &sync.Map{}}
}

func (s *safeMap) Delete(key any) {
	s.Map.Delete(key)
}

func (s *safeMap) Load(key any) (value any, ok bool) {
	return s.Map.Load(key)
}

func (s *safeMap) LoadAndDelete(key any) (value any, loaded bool) {
	return s.Map.LoadAndDelete(key)
}

func (s *safeMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	return s.Map.LoadOrStore(key, value)
}

func (s *safeMap) Range(f func(key, value any) bool) {
	s.Map.Range(f)
}

func (s *safeMap) Store(key, value any) {
	s.Map.Store(key, value)
}

func (s *safeMap) Set(key, value any) {
	s.Store(key, value)
}

func (s *safeMap) Get(key any) (value any, ok bool) {
	return s.Load(key)
}

func (s *safeMap) GetString(key any) (value string) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(string)
	}
	return
}

func (s *safeMap) GetBool(key string) (value bool) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(bool)
	}
	return
}

func (s *safeMap) GetInt8(key string) (value int8) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(int8)
	}
	return
}

func (s *safeMap) GetInt16(key string) (value int16) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(int16)
	}
	return
}

func (s *safeMap) GetInt32(key string) (value int32) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(int32)
	}
	return
}

func (s *safeMap) GetInt(key string) (value int) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(int)
	}
	return
}

func (s *safeMap) GetInt64(key string) (value int64) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(int64)
	}
	return
}

func (s *safeMap) GetUint8(key string) (value uint8) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(uint8)
	}
	return
}

func (s *safeMap) GetUint16(key string) (value uint16) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(uint16)
	}
	return
}

func (s *safeMap) GetUint32(key string) (value uint32) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(uint32)
	}
	return
}

func (s *safeMap) GetUint(key string) (value uint) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(uint)
	}
	return
}

func (s *safeMap) GetUint64(key string) (value uint64) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(uint64)
	}
	return
}

func (s *safeMap) GetFloat32(key string) (value float32) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(float32)
	}
	return
}

func (s *safeMap) GetFloat64(key string) (value float64) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(float64)
	}
	return
}

func (s *safeMap) GetTime(key string) (value time.Time) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(time.Time)
	}
	return
}

func (s *safeMap) GetDuration(key string) (value time.Duration) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(time.Duration)
	}
	return
}

func (s *safeMap) GetByteSlice(key any) (value []byte) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]byte)
	}
	return
}

func (s *safeMap) GetRuneSlice(key any) (value []rune) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]rune)
	}
	return
}

func (s *safeMap) GetIntSlice(key string) (value []int) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]int)
	}
	return
}

func (s *safeMap) GetInt64Slice(key string) (value []int64) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]int64)
	}
	return
}

func (s *safeMap) GetUintSlice(key string) (value []uint) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]uint)
	}
	return
}

func (s *safeMap) GetUint64Slice(key string) (value []uint64) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]uint64)
	}
	return
}

func (s *safeMap) GetFloat32Slice(key string) (value []float32) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]float32)
	}
	return
}

func (s *safeMap) GetFloat64Slice(key string) (value []float64) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]float64)
	}
	return
}

func (s *safeMap) GetStringSlice(key string) (value []string) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.([]string)
	}
	return
}

func (s *safeMap) GetStringMap(key string) (value map[string]any) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(map[string]any)
	}
	return
}

func (s *safeMap) GetStringMapString(key string) (value map[string]string) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(map[string]string)
	}
	return
}

func (s *safeMap) GetStringMapStringSlice(key string) (value map[string][]string) {
	if val, ok := s.Load(key); ok && val != nil {
		value, _ = val.(map[string][]string)
	}
	return
}

func (s *safeMap) Debug() (value map[any]any) {
	value = make(map[any]any)
	s.Range(func(k, v any) bool {
		value[k] = v
		return true
	})
	return
}
