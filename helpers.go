package sodium

// #cgo pkg-config: libsodium
// #include <sodium.h>
// #include <sodium/utils.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

const (
	SODIUM_BASE64_VARIANT_ORIGINAL            = C.sodium_base64_VARIANT_ORIGINAL
	SODIUM_BASE64_VARIANT_ORIGINAL_NO_PADDING = C.sodium_base64_VARIANT_ORIGINAL_NO_PADDING
	SODIUM_BASE64_VARIANT_URLSAFE             = C.sodium_base64_VARIANT_URLSAFE
	SODIUM_BASE64_VARIANT_URLSAFE_NO_PADDING  = C.sodium_base64_VARIANT_URLSAFE_NO_PADDING
)

var (
	ErrParsingCanNotBeFinished = errors.New("more than bin_maxlen bytes would be required to store the parsed string, " +
		"or if the string couldn't be fully parsed, but a valid pointer for b64_end was not provided")
)

func Base64EncodedLen(binLen uint32, variant int) uint32 {
	return uint32(C.sodium_base64_encoded_len(C.ulong(binLen), C.int(variant)))
}

func Bin2base64(binString []byte, b64String *string, variant int) int {
	bin := C.CBytes(binString)
	binLength := len(binString)
	b64MaxLength := C.size_t(Base64EncodedLen(uint32(binLength), variant))
	var b64 C.char
	defer C.free(unsafe.Pointer(bin))

	//char * sodium_bin2base64(char * const b64, const size_t b64_maxlen,
	//	const unsigned char * const bin, const size_t bin_len,
	//	const int variant
	//)
	r := C.sodium_bin2base64(&b64, b64MaxLength, (*C.uchar)(unsafe.Pointer(bin)), C.size_t(binLength), C.int(variant))

	*b64String = C.GoString(&b64)

	return int(*r)
}

func Base642bin(b64 string, binString *[]byte, variant int) error {
	var b64Length = len(b64)
	var binMaxLength = b64Length/4*3 + 10
	var binRealLength C.ulong

	b64C := C.CString(b64)
	var bin C.uchar
	defer C.free(unsafe.Pointer(b64C))

	//int sodium_base642bin(unsigned char * const bin, const size_t bin_maxlen,
	//	const char * const b64, const size_t b64_len,
	//	const char * const ignore, size_t * const bin_len,
	//	const char ** const b64_end, const int variant
	//)
	r := int(C.sodium_base642bin((*C.uchar)(unsafe.Pointer(&bin)), C.ulong(binMaxLength),
		b64C, C.ulong(b64Length),
		nil, &binRealLength, nil, C.int(variant)))
	if -1 == r {
		return ErrParsingCanNotBeFinished
	}

	*binString = C.GoBytes(unsafe.Pointer(&bin), C.int(binRealLength))

	return nil
}
