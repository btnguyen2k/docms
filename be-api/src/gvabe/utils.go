package gvabe

import (
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
)

// global constants
const (
	apiResultExtraAccessToken = "_access_token_"

	loginSessionTtl        = 3600 * 8
	loginSessionNearExpiry = 3600 * 3
)
