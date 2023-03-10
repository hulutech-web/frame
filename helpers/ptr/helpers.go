package ptr

import (
	"math/big"

	"github.com/hulutech-web/frame/helpers/log"
	"github.com/hulutech-web/frame/helpers/zone"
	"github.com/hulutech-web/frame/model/types/bigfloat"
	"github.com/hulutech-web/frame/model/types/bigint"
)

func String(value string) *string {
	return &value
}

func MapStringInterface(value map[string]interface{}) *map[string]interface{} {
	return &value
}

func Bool(value bool) *bool {
	return &value
}

func Byte(value byte) *byte {
	return &value
}

func Float64(value float64) *float64 {
	return &value
}
func Float32(value float32) *float32 {
	return &value
}

func Int(value int) *int {
	return &value
}
func Int64(value int64) *int64 {
	return &value
}
func Int32(value int32) *int32 {
	return &value
}
func Int16(value int16) *int16 {
	return &value
}
func Int8(value int8) *int8 {
	return &value
}

func Uint(value uint) *uint {
	return &value
}
func Uint64(value uint64) *uint64 {
	return &value
}
func Uint32(value uint32) *uint32 {
	return &value
}
func Uint16(value uint16) *uint16 {
	return &value
}
func Uint8(value uint8) *uint8 {
	return &value
}

func BigInt(value *big.Int) *bigint.BigInt {
	v := bigint.BigInt{}
	_ = log.Error(v.Convert(value))
	return &v
}
func BigFloat(value string) (*bigfloat.BigFloat, error) {
	v := bigfloat.BigFloat{}
	err := v.CreateFromString(value, bigfloat.ToNearestEven)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
func Time(value zone.Time) *zone.Time {
	return &value
}
