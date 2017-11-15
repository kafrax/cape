package wechat

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/kafrax/chaos"
)

func checkSign(msg map[string]interface{}, key, sign string) bool {
	signCalc := genSign(msg, key)
	if sign == signCalc {
		return true
	}
	return false
}

func genSign(mReq map[string]string, key string) (sign string) {
	var sortedKeys []string
	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	signStrings = signStrings + "key=" + key
	md5Ctx := md5.New()
	md5Ctx.Write(chaos.String2Byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
