package tarantool

import (
	"github.com/vmihailenco/msgpack/v5"
)

// IntKey is utility type for passing integer key to Select*, Update* and Delete*
// It serializes to array with single integer element.
type IntKey struct {
	I int
}

func (k IntKey) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error
	if err = enc.EncodeArrayLen(1); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(k.I)); err != nil {
		return err
	}
	return nil
}

// UintKey is utility type for passing unsigned integer key to Select*, Update* and Delete*
// It serializes to array with single integer element.
type UintKey struct {
	I uint
}

func (k UintKey) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error
	if err = enc.EncodeArrayLen(1); err != nil {
		return err
	}
	if err = enc.EncodeUint(uint64(k.I)); err != nil {
		return err
	}
	return nil
}

// UintKey is utility type for passing string key to Select*, Update* and Delete*
// It serializes to array with single string element.
type StringKey struct {
	S string
}

func (k StringKey) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error
	if err = enc.EncodeArrayLen(1); err != nil {
		return err
	}
	if err = enc.EncodeString(k.S); err != nil {
		return err
	}
	return nil
}

// IntIntKey is utility type for passing two integer keys to Select*, Update* and Delete*
// It serializes to array with two integer elements
type IntIntKey struct {
	I1, I2 int
}

func (k IntIntKey) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error
	if err = enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(k.I1)); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(k.I2)); err != nil {
		return err
	}
	return nil
}

// Op - is update operation
type Op struct {
	Op    string
	Field int
	Arg   interface{}
}

func (o Op) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error
	if err = enc.EncodeArrayLen(3); err != nil {
		return err
	}
	if err = enc.EncodeString(o.Op); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(o.Field)); err != nil {
		return err
	}
	return enc.Encode(o.Arg)
}

type OpSplice struct {
	Op      string
	Field   int
	Pos     int
	Len     int
	Replace string
}

func (o OpSplice) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error
	if err = enc.EncodeArrayLen(5); err != nil {
		return err
	}
	if err = enc.EncodeString(o.Op); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(o.Field)); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(o.Pos)); err != nil {
		return err
	}
	if err = enc.EncodeInt(int64(o.Len)); err != nil {
		return err
	}
	if err = enc.EncodeString(o.Replace); err != nil {
		return err
	}
	return nil
}
