package sodium

// #cgo pkg-config: libsodium
// #include <sodium.h>
// #include <sodium/utils.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

const (
	SODIUM_BASE64_VARIANT_ORIGINAL            = C.sodium_base64_VARIANT_ORIGINAL
	SODIUM_BASE64_VARIANT_ORIGINAL_NO_PADDING = C.sodium_base64_VARIANT_ORIGINAL_NO_PADDING
	SODIUM_BASE64_VARIANT_URLSAFE             = C.sodium_base64_VARIANT_URLSAFE
	SODIUN_BASE64_VARIANT_URLSAFE_NO_PADDING  = C.sodium_base64_VARIANT_URLSAFE
)

func Base64EncodedLen(binLen uint32, variant int) uint32 {
	return uint32(C.sodium_base64_encoded_len(C.ulong(binLen), C.int(variant)))
}

func Bin2base64(binString string, b64String *string, variant int) int {
	bin := C.CString(binString)
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

func Base642bin(b64 string, binString *string, variant int) int {
	var b64Length = len(b64)
	var binLength = b64Length/4*3 + 2

	b64C := C.CString(b64)
	var bin C.char
	defer C.free(unsafe.Pointer(b64C))

	//int sodium_base642bin(unsigned char * const bin, const size_t bin_maxlen,
	//	const char * const b64, const size_t b64_len,
	//	const char * const ignore, size_t * const bin_len,
	//	const char ** const b64_end, const int variant
	//)
	r := C.sodium_base642bin((*C.uchar)(unsafe.Pointer(&bin)), C.ulong(binLength), b64C, C.ulong(b64Length), nil, nil, nil, C.int(variant))

	*binString = C.GoString(&bin)

	return int(r)
}
