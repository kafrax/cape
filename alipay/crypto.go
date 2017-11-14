package alipay

import (
	"net/url"
	"strings"
	"crypto"
	"encoding/pem"
	"encoding/base64"
	"errors"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
	"sort"
	"net/http"

	"github.com/kafrax/chaos"
)

func SignRsa2(keys []string, param *url.Values, privateKey []byte) (string, error) {
	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var src = strings.Join(pList, "&")
	var sig, err = signPKCS1v15(chaos.String2Byte(src), privateKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

func signPKCS1v15(src, key []byte, hash crypto.Hash) ([]byte, error) {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, pri, hash, hashed)
}

func verifyPKCS1v15(src, sig, key []byte, hash crypto.Hash) (err error) {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return errors.New("publick key error")
	}
	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, sig)
}

func ValidRSA2(req *http.Request, key []byte) (ok bool, err error) {
	sign, err := base64.StdEncoding.DecodeString(req.PostForm.Get("sign"))
	if err != nil {
		return false, err
	}
	var keys []string
	for key, value := range req.PostForm {
		if key == "sign" || key == "sign_type" {
			continue
		}
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	var pList []string
	for _, key := range keys {
		var value = strings.TrimSpace(req.PostForm.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var s = strings.Join(pList, "&")

	err = verifyPKCS1v15(chaos.String2Byte(s), sign, key, crypto.SHA256)
	if err != nil {
		return false, err
	}
	return true, nil
}
