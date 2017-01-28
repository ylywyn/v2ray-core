package serial

import "strconv"

func Uint16ToBytes(value uint16, b []byte) []byte {
	return append(b, byte(value>>8), byte(value))
}

func Uint16ToString(value uint16) string {
	return strconv.Itoa(int(value))
}

func WriteUint16(value uint16) func([]byte) (int, error) {
	return func(b []byte) (int, error) {
		b = Uint16ToBytes(value, b[:0])
		return 2, nil
	}
}

func Uint32ToBytes(value uint32, b []byte) []byte {
	return append(b, byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func Uint32ToString(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

func WriteUint32(value uint32) func([]byte) (int, error) {
	return func(b []byte) (int, error) {
		b = Uint32ToBytes(value, b[:0])
		return 4, nil
	}
}

func IntToBytes(value int, b []byte) []byte {
	return append(b, byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func IntToString(value int) string {
	return Int64ToString(int64(value))
}

func Int64ToBytes(value int64, b []byte) []byte {
	return append(b,
		byte(value>>56),
		byte(value>>48),
		byte(value>>40),
		byte(value>>32),
		byte(value>>24),
		byte(value>>16),
		byte(value>>8),
		byte(value))
}

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}
