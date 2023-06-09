/*
Package utils provides a utility/helper functions that can be used by the application.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@since template-v0.4.r1
*/
package utils

import (
	"bytes"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	olaf2 "github.com/btnguyen2k/consu/olaf"
)

// global variables
var (
	// Location should be initialized during application bootstrap
	Location *time.Location
)

func getMacAddr() string {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				return i.HardwareAddr.String()
			}
		}
	}
	return ""
}

func getMacAddrAsLong() int64 {
	mac, _ := strconv.ParseInt(strings.Replace(getMacAddr(), ":", "", -1), 16, 64)
	return mac
}

var olaf = olaf2.NewOlaf(getMacAddrAsLong())

// UniqueId generates a unique id.
func UniqueId() string {
	return olaf.Id128Hex()
}

// UniqueIdSmall generates a unique id, shorter length than which is generated by UniqueId.
func UniqueIdSmall() string {
	return olaf.Id64Ascii()
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string with specified length.
func RandomString(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// VersionToNumber converts a version string to number representation
func VersionToNumber(version string) uint64 {
	var result uint64 = 0
	for tokens := strings.Split(version, "."); len(tokens) > 0; {
		if i, err := strconv.Atoi(tokens[0]); err == nil {
			result = result<<8 + (uint64(i) & 255)
		} else {
			result = result << 8
		}
		tokens = tokens[1:]
	}
	return result
}
