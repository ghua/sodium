package sodium

import (
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
)

type testInput struct {
	Clear   string
	Base64  string
	Variant int
}

var licenceContent, _ = ioutil.ReadFile("LICENSE")
var testInputs = []testInput{
	{
		Clear:   "test",
		Base64:  "dGVzdA==",
		Variant: SODIUM_BASE64_VARIANT_URLSAFE,
	},
	{
		Clear: string(licenceContent),
		Base64: "VGhlIE1JVCBMaWNlbnNlIChNSVQpCgpDb3B5cmlnaHQgKGMpIDIwMTYgSmFtZXMgUnVhbiA8cnVhbmJlaWhvbmdAZ21haWwuY29tP" +
			"goKUGVybWlzc2lvbiBpcyBoZXJlYnkgZ3JhbnRlZCwgZnJlZSBvZiBjaGFyZ2UsIHRvIGFueSBwZXJzb24gb2J0YWluaW5nIGEgY2" +
			"9weQpvZiB0aGlzIHNvZnR3YXJlIGFuZCBhc3NvY2lhdGVkIGRvY3VtZW50YXRpb24gZmlsZXMgKHRoZSAiU29mdHdhcmUiKSwgdG8" +
			"gZGVhbAppbiB0aGUgU29mdHdhcmUgd2l0aG91dCByZXN0cmljdGlvbiwgaW5jbHVkaW5nIHdpdGhvdXQgbGltaXRhdGlvbiB0aGUg" +
			"cmlnaHRzCnRvIHVzZSwgY29weSwgbW9kaWZ5LCBtZXJnZSwgcHVibGlzaCwgZGlzdHJpYnV0ZSwgc3VibGljZW5zZSwgYW5kL29yI" +
			"HNlbGwKY29waWVzIG9mIHRoZSBTb2Z0d2FyZSwgYW5kIHRvIHBlcm1pdCBwZXJzb25zIHRvIHdob20gdGhlIFNvZnR3YXJlIGlzCm" +
			"Z1cm5pc2hlZCB0byBkbyBzbywgc3ViamVjdCB0byB0aGUgZm9sbG93aW5nIGNvbmRpdGlvbnM6CgpUaGUgYWJvdmUgY29weXJpZ2h" +
			"0IG5vdGljZSBhbmQgdGhpcyBwZXJtaXNzaW9uIG5vdGljZSBzaGFsbCBiZSBpbmNsdWRlZCBpbiBhbGwKY29waWVzIG9yIHN1YnN0" +
			"YW50aWFsIHBvcnRpb25zIG9mIHRoZSBTb2Z0d2FyZS4KClRIRSBTT0ZUV0FSRSBJUyBQUk9WSURFRCAiQVMgSVMiLCBXSVRIT1VUI" +
			"FdBUlJBTlRZIE9GIEFOWSBLSU5ELCBFWFBSRVNTIE9SCklNUExJRUQsIElOQ0xVRElORyBCVVQgTk9UIExJTUlURUQgVE8gVEhFIF" +
			"dBUlJBTlRJRVMgT0YgTUVSQ0hBTlRBQklMSVRZLApGSVRORVNTIEZPUiBBIFBBUlRJQ1VMQVIgUFVSUE9TRSBBTkQgTk9OSU5GUkl" +
			"OR0VNRU5ULiBJTiBOTyBFVkVOVCBTSEFMTCBUSEUKQVVUSE9SUyBPUiBDT1BZUklHSFQgSE9MREVSUyBCRSBMSUFCTEUgRk9SIEFO" +
			"WSBDTEFJTSwgREFNQUdFUyBPUiBPVEhFUgpMSUFCSUxJVFksIFdIRVRIRVIgSU4gQU4gQUNUSU9OIE9GIENPTlRSQUNULCBUT1JUI" +
			"E9SIE9USEVSV0lTRSwgQVJJU0lORyBGUk9NLApPVVQgT0YgT1IgSU4gQ09OTkVDVElPTiBXSVRIIFRIRSBTT0ZUV0FSRSBPUiBUSE" +
			"UgVVNFIE9SIE9USEVSIERFQUxJTkdTIElOIFRIRQpTT0ZUV0FSRS4K",
		Variant: SODIUM_BASE64_VARIANT_URLSAFE,
	},
}

func TestBase642Bin(t *testing.T) {
	var binActual string

	for _, v := range testInputs {
		Base642bin(v.Base64, &binActual, v.Variant)
		if binActual != v.Clear {
			t.Fatalf("expected string %s, given: %s", v.Clear, binActual)
		}
	}
}

func TestBin2base64(t *testing.T) {
	var b64Actual string

	for _, v := range testInputs {
		Bin2base64(v.Clear, &b64Actual, v.Variant)
		if b64Actual != v.Base64 {
			t.Fatalf("expected base64 %s, given: %s", v.Base64, b64Actual)
		}
	}
}

func TestBase64EncodedLen(t *testing.T) {
	var binLen uint32
	var b64LenExpected uint32
	var b64LenActual uint32

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 100; i++ {
		binLen = uint32(r.Intn(100))
		b64LenExpected = (binLen+2)/3*4 + 1

		b64LenActual = Base64EncodedLen(binLen, SODIUM_BASE64_VARIANT_URLSAFE)
		if b64LenExpected != b64LenActual {
			t.Fatalf("expected len is %d, given: %d", b64LenExpected, b64LenActual)
		}
	}
}
