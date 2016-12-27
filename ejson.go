package ejson

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"time"
)

type ejsonBytes struct{ v *[]byte }

var _ Marshaler = (*ejsonBytes)(nil)
var _ Unmarshaler = (*ejsonBytes)(nil)

// Currently I'm patching a type check into encode.go at line 281
// but it might be better to patch encodeByteSlice at line 673

func (eb ejsonBytes) MarshalJSON() ([]byte, error) {
	s := *(eb.v)
	if s == nil {
		return []byte(`{"$binary":""}`), nil
	}
	e := &bytes.Buffer{}
	e.WriteString(`{"$binary":"`)
	// Code from encodeByteSlice in encode.go.
	if len(s) < 1024 {
		// for small buffers, using Encode directly is much faster.
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
		base64.StdEncoding.Encode(dst, s)
		e.Write(dst)
	} else {
		// for large buffers, avoid unnecessary extra temporary
		// buffer space.
		enc := base64.NewEncoder(base64.StdEncoding, e)
		enc.Write(s)
		enc.Close()
	}
	e.WriteString(`"}`)
	return e.Bytes(), nil
}

func (eb ejsonBytes) UnmarshalJSON(b []byte) error {
	aux := struct {
		Value string `json:"$binary"`
	}{}
	err := Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	// Code from literalStore in decode.go.
	by := make([]byte, base64.StdEncoding.DecodedLen(len(aux.Value)))
	n, err := base64.StdEncoding.Decode(by, []byte(aux.Value))
	if err != nil {
		return err
	}
	*(eb.v) = by[:n]
	return nil
}

type ejsonDate struct{ v *time.Time }

var _ Marshaler = (*ejsonDate)(nil)
var _ Unmarshaler = (*ejsonDate)(nil)

func (ed ejsonDate) MarshalJSON() ([]byte, error) {
	if ed.v == nil {
		return []byte("null"), nil
	}
	millis := (*ed.v).UnixNano() / 1000000
	e := &bytes.Buffer{}
	e.WriteString(`{"$date":`)
	e.WriteString(fmt.Sprint(millis))
	e.WriteString(`}`)
	return e.Bytes(), nil
}

func (ed *ejsonDate) UnmarshalJSON(b []byte) error {
	aux := struct {
		Value int64 `json:"$date"`
	}{}
	err := Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	d := time.Unix(0, aux.Value*1000000)
	*(ed.v) = d
	return nil
}
