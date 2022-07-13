package protowrap

import (
	"math"
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestInt32Value(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*wrapperspb.Int32Value)(nil)
		if got := Int32Value[int32](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := 1
		want := &wrapperspb.Int32Value{Value: int32(v)}
		if got := Int32Value(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestInt64Value(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*wrapperspb.Int64Value)(nil)
		if got := Int64Value[int64](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := 1
		want := &wrapperspb.Int64Value{Value: int64(v)}
		if got := Int64Value(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestIntF32V(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*int32)(nil)
		if got := IntF32V[int32](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := int32(1)
		if got := IntF32V[int32](&wrapperspb.Int32Value{Value: v}); !reflect.DeepEqual(&v, got) {
			t.Errorf("want: %v, got: %v", &v, got)
		}
	})
}

func TestIntF64V(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*int64)(nil)
		if got := IntF64V[int64](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := int64(1)
		if got := IntF64V[int64](&wrapperspb.Int64Value{Value: v}); !reflect.DeepEqual(&v, got) {
			t.Errorf("want: %v, got: %v", &v, got)
		}
	})
}

func TestFloatValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*wrapperspb.FloatValue)(nil)
		if got := FloatValue[float32](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := float32(1.0)
		want := &wrapperspb.FloatValue{Value: 1}
		if got := FloatValue(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestDoubleValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*wrapperspb.DoubleValue)(nil)
		if got := DoubleValue[float64](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := 1.0
		want := &wrapperspb.DoubleValue{Value: v}
		if got := DoubleValue(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestFloatFromFloatValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*float32)(nil)
		if got := FloatF32V[float32](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := float32(1.0)
		if got := FloatF32V[float32](&wrapperspb.FloatValue{Value: 1}); !reflect.DeepEqual(&v, got) {
			t.Errorf("want: %v, got: %v", &v, got)
		}
	})
}

func TestFloatFromDoubleValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*float64)(nil)
		if got := FloatF64V[float64](nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := 1.0
		if got := FloatF64V[float64](&wrapperspb.DoubleValue{Value: v}); !reflect.DeepEqual(&v, got) {
			t.Errorf("want: %v, got: %v", &v, got)
		}
	})
}

func TestStringValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*wrapperspb.StringValue)(nil)
		if got := StringValue(nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := "v"
		want := &wrapperspb.StringValue{Value: v}
		if got := StringValue(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*string)(nil)
		if got := String(nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := "v"
		if got := String(&wrapperspb.StringValue{Value: v}); !reflect.DeepEqual(&v, got) {
			t.Errorf("want: %v, got: %v", &v, got)
		}
	})
}

func TestBoolValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*wrapperspb.BoolValue)(nil)
		if got := BoolValue(nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := true
		want := &wrapperspb.BoolValue{Value: v}
		if got := BoolValue(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestBool(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*bool)(nil)
		if got := Bool(nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := true
		if got := Bool(&wrapperspb.BoolValue{Value: v}); !reflect.DeepEqual(&v, got) {
			t.Errorf("want: %v, got: %v", &v, got)
		}
	})
}

func TestTimestamp(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*timestamppb.Timestamp)(nil)
		if got := Timestamp(nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		v := time.Now()
		want := &timestamppb.Timestamp{Seconds: v.Unix(), Nanos: int32(v.Nanosecond())}
		if got := Timestamp(&v); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("nil", func(t *testing.T) {
		want := (*timestamppb.Timestamp)(nil)
		if got := Timestamp(&time.Time{}); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestTime(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		want := (*time.Time)(nil)
		if got := Time(nil); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		want := time.Now().UTC()
		if got := Time(&timestamppb.Timestamp{
			Seconds: want.Unix(),
			Nanos:   int32(want.Nanosecond()),
		}); !reflect.DeepEqual(&want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		want := (*time.Time)(nil)
		if got := Time(&timestamppb.Timestamp{Seconds: math.MaxInt}); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("zero", func(t *testing.T) {
		want := (*time.Time)(nil)
		if got := Time(timestamppb.New(time.Time{})); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestInts(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		want := []int(nil)
		if got := Ints[int]([]int32(nil)); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("value", func(t *testing.T) {
		want := []int{1, 2}
		if got := Ints[int]([]int32{1, 2}); !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}
