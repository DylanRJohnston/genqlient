package optional

import (
	"encoding/json"
)

type Value[V any] struct {
	value V
	ok    bool
}

func Some[V any](value V) Value[V] {
	return Value[V]{value: value, ok: true}
}

func None[V any]() Value[V] {
	return Value[V]{ok: false}
}

func (v Value[V]) Unpack() (V, bool) {
	return v.value, v.ok
}

func (v Value[V]) Get(fallback V) V {
	if v.ok {
		return v.value
	}

	return fallback
}

func FromPtr[V any](ptr *V) Value[V] {
	if ptr == nil {
		return None[V]()
	}

	return Some(*ptr)
}

func (value Value[V]) MarshalJSON() ([]byte, error) {
	if value.ok {
		return json.Marshal(value.value)
	} else {
		return json.Marshal((*V)(nil))
	}
}

func (value *Value[V]) UnmarshalJSON(data []byte) error {
	v := (*V)(nil)

	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	if v != nil {
		value.value = *v
		value.ok = true
	}

	return nil
}
