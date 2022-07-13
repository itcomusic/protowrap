// Package protowrap provides functions of converting and helpers for wrapperspb types.
package protowrap

import (
	"time"

	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Int32Value returns *wrapperspb.Int32Value from *int.
func Int32Value[T constraints.Integer](v *T) *wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	return &wrapperspb.Int32Value{Value: int32(*v)}
}

// Int64Value returns *wrapperspb.Int64Value from *int.
func Int64Value[T constraints.Integer](v *T) *wrapperspb.Int64Value {
	if v == nil {
		return nil
	}
	return &wrapperspb.Int64Value{Value: int64(*v)}
}

// IntF32V returns *int from *wrapperspb.Int32Value.
func IntF32V[T constraints.Integer](v *wrapperspb.Int32Value) *T {
	if v == nil {
		return nil
	}
	res := T(v.Value)
	return &res
}

// IntF64V returns *int from *wrapperspb.Int64Value.
func IntF64V[T constraints.Integer](v *wrapperspb.Int64Value) *T {
	if v == nil {
		return nil
	}
	res := T(v.Value)
	return &res
}

// FloatValue returns *wrapperspb.FloatValue from *float.
func FloatValue[T constraints.Float](v *T) *wrapperspb.FloatValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.FloatValue{Value: float32(*v)}
}

// DoubleValue returns *wrapperspb.DoubleValue from *float.
func DoubleValue[T constraints.Float](v *T) *wrapperspb.DoubleValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.DoubleValue{Value: float64(*v)}
}

// FloatF32V returns *float from *wrapperspb.FloatValue.
func FloatF32V[T constraints.Float](v *wrapperspb.FloatValue) *T {
	if v == nil {
		return nil
	}
	res := T(v.Value)
	return &res
}

// FloatF64V returns *float from *wrapperspb.DoubleValue.
func FloatF64V[T constraints.Float](v *wrapperspb.DoubleValue) *T {
	if v == nil {
		return nil
	}
	res := T(v.Value)
	return &res
}

// StringValue returns *wrapperspb.StringValue from *string.
func StringValue(v *string) *wrapperspb.StringValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.StringValue{Value: *v}
}

// String returns *string from *wrapperspb.StringValue.
func String(v *wrapperspb.StringValue) *string {
	if v == nil {
		return nil
	}
	return &v.Value
}

// BoolValue returns *wrapperspb.BoolValue from *bool.
func BoolValue(v *bool) *wrapperspb.BoolValue {
	if v == nil {
		return nil
	}
	return &wrapperspb.BoolValue{Value: *v}
}

// Bool returns *bool from *wrapperspb.BoolValue.
func Bool(v *wrapperspb.BoolValue) *bool {
	if v == nil {
		return nil
	}
	return &v.Value
}

// Timestamp returns *timestamppb.Timestamp from *time.Time.
func Timestamp(v *time.Time) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}

	if (*v).IsZero() {
		return nil
	}
	return timestamppb.New(*v)
}

// Time returns *time.Time from *timestamppb.Timestamp.
func Time(v *timestamppb.Timestamp) *time.Time {
	if v == nil {
		return nil
	}

	if !v.IsValid() {
		return nil
	}

	res := v.AsTime()
	if res.IsZero() {
		return nil
	}
	return &res
}

// Ints returns slice of a new type.
func Ints[OUT constraints.Integer, IN constraints.Integer](v []IN) []OUT {
	if v == nil {
		return nil
	}

	res := make([]OUT, len(v))
	for i, v := range v {
		res[i] = OUT(v)
	}
	return res
}
