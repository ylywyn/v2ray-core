package assert

import (
	"bytes"

	"v2ray.com/core/common/serial"
)

func (v *Assert) Bytes(value []byte) *BytesSubject {
	return &BytesSubject{
		Subject: Subject{
			disp: serial.BytesToHexString(value),
			a:    v,
		},
		value: value,
	}
}

type BytesSubject struct {
	Subject
	value []byte
}

func (subject *BytesSubject) Equals(expectation []byte) {
	if !bytes.Equal(subject.value, expectation) {
		subject.Fail("is equal to", serial.BytesToHexString(expectation))
	}
}

func (subject *BytesSubject) NotEquals(expectation []byte) {
	if bytes.Equal(subject.value, expectation) {
		subject.Fail("is not equal to", serial.BytesToHexString(expectation))
	}
}
