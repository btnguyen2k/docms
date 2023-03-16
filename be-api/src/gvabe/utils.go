package gvabe

import (
	"bytes"
	"compress/zlib"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/hex"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/consu/semita"
	"github.com/btnguyen2k/goyai"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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

func encryptPassword(salt, rawPassword string) string {
	out := sha1.Sum([]byte(salt + "." + rawPassword))
	return strings.ToLower(hex.EncodeToString(out[:]))
}

// zlibCompress compresses data using zlib.
func zlibCompress(data []byte) []byte {
	var b bytes.Buffer
	w, _ := zlib.NewWriterLevel(&b, zlib.BestCompression)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

// // zlibDecompress decompressed compressed-data using zlib.
// func zlibDecompress(compressedData []byte) ([]byte, error) {
// 	r, err := zlib.NewReader(bytes.NewReader(compressedData))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var b bytes.Buffer
// 	_, err = io.Copy(&b, r)
// 	r.Close()
// 	return b.Bytes(), err
// }

// available since template-v0.2.0
func zipAndEncrypt(data []byte) ([]byte, error) {
	zip := zlibCompress(data)
	return rsaEncrypt(RsaModeAuto, zip, rsaPubKey)
}

// // available since template-v0.2.0
// func decryptAndUnzip(encdata []byte) ([]byte, error) {
// 	if zip, err := rsaDecrypt(RsaModeAuto, encdata, rsaPrivKey); err != nil {
// 		return nil, err
// 	} else {
// 		return zlibDecompress(zip)
// 	}
// }

// func genLoginToken(u *user.User) (string, error) {
// 	t := time.Now()
// 	data := map[string]interface{}{
// 		loginAttrUsername:  u.GetUsername(),
// 		loginAttrGroupId:   u.GetGroupId(),
// 		loginAttrTimestamp: t.Unix(),
// 		loginAttrExpiry:    t.Unix() + loginSessionTtl,
// 	}
// 	if js, err := json.Marshal(data); err != nil {
// 		return "", err
// 	} else {
// 		zip := zlibCompress(js)
// 		if enc, err := aesEncrypt([]byte(u.GetAesKey()), zip); err != nil {
// 			return "", err
// 		} else {
// 			return base64.StdEncoding.EncodeToString(enc), nil
// 		}
// 	}
// }

// func decodeLoginToken(username string, loginToken string) (map[string]interface{}, error) {
// 	if user, err := userDao.Get(username); user == nil || err != nil {
// 		return nil, err
// 	} else if enc, err := base64.StdEncoding.DecodeString(loginToken); err != nil {
// 		return nil, err
// 	} else if zip, err := aesDecrypt([]byte(user.GetAesKey()), enc); err != nil {
// 		return nil, err
// 	} else if js, err := zlibDecompress(zip); err != nil {
// 		return nil, err
// 	} else {
// 		var data map[string]interface{}
// 		if err := json.Unmarshal(js, &data); err != nil {
// 			return nil, nil
// 		}
// 		return data, nil
// 	}
// }

// func verifyLoginToken(username string, loginToken string) (int, error) {
// 	if data, err := decodeLoginToken(username, loginToken); err != nil {
// 		return sessionStatusError, err
// 	} else if data == nil {
// 		return sessionStatusUserNotFound, nil
// 	} else {
// 		if u, err := reddo.ToString(data[loginAttrUsername]); err != nil {
// 			return sessionStatusError, err
// 		} else if u != username {
// 			return sessionStatusInvalid, nil
// 		} else if expiry, err := reddo.ToInt(data[loginAttrExpiry]); err != nil {
// 			return sessionStatusError, err
// 		} else if expiry < time.Now().Unix() {
// 			return sessionStatusExpired, nil
// 		}
// 	}
// 	return sessionStatusOk, nil
// }

/*----------------------------------------------------------------------*/

var muxSytemInfo sync.Mutex
var systemInfoArr = make([]map[string]interface{}, 0)

func routineUpdateSystemInfo() {
	for {
		go doUpdateSystemInfo()
		<-time.After(10 * time.Second)
	}
}

func lastSystemInfo() map[string]interface{} {
	muxSytemInfo.Lock()
	defer muxSytemInfo.Unlock()
	return systemInfoArr[len(systemInfoArr)-1]
}

func doUpdateSystemInfo() {
	muxSytemInfo.Lock()
	defer muxSytemInfo.Unlock()

	data := make(map[string]interface{})
	{
		load, err := load.Avg()
		cpuLoad := -1.0
		if err == nil && load != nil {
			cpuLoad = math.Floor(load.Load1*100) / 100
		}
		historyLoad := make([]float64, 0)
		for _, data := range systemInfoArr {
			s := semita.NewSemita(data)
			load, _ := s.GetValueOfType("cpu.load", reddo.TypeFloat)
			historyLoad = append(historyLoad, load.(float64))
		}
		data["cpu"] = map[string]interface{}{
			"cores":        runtime.NumCPU(),
			"load":         cpuLoad,
			"history_load": historyLoad,
		}
	}

	{
		mem, err := mem.VirtualMemory()
		memFree := uint64(0)
		if err == nil && mem != nil {
			memFree = mem.Free
		}
		historyFree := make([]uint64, 0)
		historyFreeKb := make([]float64, 0)
		historyFreeMb := make([]float64, 0)
		historyFreeGb := make([]float64, 0)
		for _, data := range systemInfoArr {
			s := semita.NewSemita(data)
			free, _ := s.GetValueOfType("memory.free", reddo.TypeUint)
			historyFree = append(historyFree, free.(uint64))
			freeKb, _ := s.GetValueOfType("memory.freeKb", reddo.TypeFloat)
			historyFreeKb = append(historyFreeKb, freeKb.(float64))
			freeMb, _ := s.GetValueOfType("memory.freeMb", reddo.TypeFloat)
			historyFreeMb = append(historyFreeMb, freeMb.(float64))
			freeGb, _ := s.GetValueOfType("memory.freeGb", reddo.TypeFloat)
			historyFreeGb = append(historyFreeGb, freeGb.(float64))
		}
		data["memory"] = map[string]interface{}{
			"free":           memFree,
			"freeKb":         math.Floor(100.0*(float64(memFree)/1024)) / 100.0,
			"freeMb":         math.Floor(100.0*(float64(memFree)/1024/1024)) / 100.0,
			"freeGb":         math.Floor(100.0*(float64(memFree)/1024/1024/1024)) / 100.0,
			"history_free":   historyFree,
			"history_freeKb": historyFreeKb,
			"history_freeMb": historyFreeMb,
			"history_freeGb": historyFreeGb,
		}
	}

	{
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		historyUsed := make([]uint64, 0)
		historyUsedKb := make([]float64, 0)
		historyUsedMb := make([]float64, 0)
		historyUsedGb := make([]float64, 0)
		for _, data := range systemInfoArr {
			s := semita.NewSemita(data)
			used, _ := s.GetValueOfType("app_memory.used", reddo.TypeUint)
			historyUsed = append(historyUsed, used.(uint64))
			usedKb, _ := s.GetValueOfType("app_memory.usedKb", reddo.TypeFloat)
			historyUsedKb = append(historyUsedKb, usedKb.(float64))
			usedMb, _ := s.GetValueOfType("app_memory.usedMb", reddo.TypeFloat)
			historyUsedMb = append(historyUsedMb, usedMb.(float64))
			usedGb, _ := s.GetValueOfType("app_memory.usedGb", reddo.TypeFloat)
			historyUsedGb = append(historyUsedGb, usedGb.(float64))
		}
		data["app_memory"] = map[string]interface{}{
			"used":           m.Alloc,
			"usedKb":         math.Floor(100.0*(float64(m.Alloc)/1024)) / 100.0,
			"usedMb":         math.Floor(100.0*(float64(m.Alloc)/1024/1024)) / 100.0,
			"usedGb":         math.Floor(100.0*(float64(m.Alloc)/1024/1024/1024)) / 100.0,
			"history_used":   historyUsed,
			"history_usedKb": historyUsedKb,
			"history_usedMb": historyUsedMb,
			"history_usedGb": historyUsedGb,
		}
	}

	{
		history := make([]int, 0)
		for _, data := range systemInfoArr {
			s := semita.NewSemita(data)
			n, _ := s.GetValueOfType("go_routines.num", reddo.TypeInt)
			history = append(history, int(n.(int64)))
		}
		data["go_routines"] = map[string]interface{}{
			"num":     runtime.NumGoroutine(),
			"history": history,
		}
	}

	systemInfoArr = append(systemInfoArr, data)
	if len(systemInfoArr) > 10 {
		systemInfoArr[0] = nil
		systemInfoArr = systemInfoArr[1:]
	}
}
