package ejson

import (
	"reflect"
	"testing"
	"time"
)

func TestEJSON_marshalByteSlice(t *testing.T) {
	b := []byte{0, 1}
	ejson := `{"$binary":"AAE="}`

	res, err := Marshal(b)
	if err != nil {
		t.Fatal("[]byte Marshal error", err)
	}

	if string(res) != ejson {
		t.Error("[]byte Marshal fail, expected", ejson, "got", string(res))
	}
}

func TestEJSON_marshalByteSlicePointer(t *testing.T) {
	b := []byte{0, 1}
	ejson := `{"$binary":"AAE="}`

	res, err := Marshal(&b)
	if err != nil {
		t.Fatal("[]byte Marshal error", err)
	}

	if string(res) != ejson {
		t.Error("[]byte Marshal fail, expected", ejson, "got", string(res))
	}
}

func TestEJSON_marshalNilByteSlice(t *testing.T) {
	var b []byte
	ejson := `{"$binary":""}`

	res, err := Marshal(b)
	if err != nil {
		t.Fatal("[]byte Marshal error", err)
	}

	if string(res) != ejson {
		t.Error("[]byte Marshal fail, expected", ejson, "got", string(res))
	}
}

func TestEJSON_unmarshalByteSlice(t *testing.T) {
	ejson := `{"$binary":"AAE="}`
	b := []byte{0, 1}

	var res []byte
	err := Unmarshal([]byte(ejson), &res)
	if err != nil {
		t.Fatal("[]byte Marshal error", err)
	}

	if !reflect.DeepEqual(b, res) {
		t.Error("[]byte Marshal fail, expected", b, "got", res)
	}
}

func TestEJSON_marshalTime(t *testing.T) {
	d := time.Unix(0, 1358205756553*1000000)
	ejson := `{"$date":1358205756553}`

	res, err := Marshal(d)
	if err != nil {
		t.Fatal("time.Time Marshal error", err)
	}

	if string(res) != ejson {
		t.Error("time.Time Marshal fail, expected", ejson, "got", string(res))
	}
}

func TestEJSON_marshalTimePointer(t *testing.T) {
	d := time.Unix(0, 123456789*1000000)
	ejson := `{"$date":123456789}`

	res, err := Marshal(&d)
	if err != nil {
		t.Fatal("time.Time Marshal error", err)
	}

	if string(res) != ejson {
		t.Error("time.Time Marshal fail, expected", ejson, "got", string(res))
	}
}

func TestEJSON_marshalNilTimePointer(t *testing.T) {
	var d *time.Time
	ejson := "null"

	res, err := Marshal(&d)
	if err != nil {
		t.Fatal("time.Time Marshal error", err)
	}

	if string(res) != ejson {
		t.Error("time.Time Marshal fail, expected", ejson, "got", string(res))
	}
}

func TestEJSON_unmarshalTime(t *testing.T) {
	ejson := `{"$date":123456789}`
	d := time.Unix(0, 123456789*1000000)

	var res time.Time
	err := Unmarshal([]byte(ejson), &res)
	if err != nil {
		t.Fatal("time.Time Marshal error", err)
	}

	if !reflect.DeepEqual(d, res) {
		t.Error("time.Time Marshal fail, expected", d, "got", res)
	}
}
