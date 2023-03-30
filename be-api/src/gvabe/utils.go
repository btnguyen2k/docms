package gvabe

import (
	"bytes"
	"compress/zlib"
	"crypto/rsa"

	"github.com/btnguyen2k/goyai"
)

// global variables
var (
	DEBUG_MODE = false
	DEMO_MODE  = false

	rsaPrivKey *rsa.PrivateKey
	rsaPubKey  *rsa.PublicKey

	i18n goyai.I18n

	exterAppId   string
	exterBaseUrl string
)

// global constants
const (
	apiResultExtraAccessToken = "_access_token_"

	loginSessionTtl        = 3600 * 8
	loginSessionNearExpiry = 3600 * 3
)

// zlibCompress compresses data using zlib.
func zlibCompress(data []byte) []byte {
	var b bytes.Buffer
	w, _ := zlib.NewWriterLevel(&b, zlib.BestCompression)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

// available since template-v0.2.0
func zipAndEncrypt(data []byte) ([]byte, error) {
	zip := zlibCompress(data)
	return rsaEncrypt(RsaModeAuto, zip, rsaPubKey)
}
