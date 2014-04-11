package util

import (
	"crypto/rc4"
	"code.google.com/p/go.crypto/salsa20"
)

type Cifrador struct {
	key [32]byte
	nonce[] byte
}

func NuevoCifrador(key [32]byte, nonce []byte) *Cifrador {
	cf := new(Cifrador)
	cf.key = key
	cf.nonce = nonce
	return cf
}

func (cf *Cifrador) Encrypt(aencriptar []byte) []byte {
	res := aencriptar
	salsa20.XORKeyStream(res, res, cf.nonce, &cf.key)
	rckey := cf.key[:]
	crc4, err := rc4.NewCipher(rckey)
	if err != nil {
		panic(err)
	}
	crc4.XORKeyStream(res, res)
	return res
}

func (cf *Cifrador) Decrypt(desencriptar []byte) []byte {
	res := desencriptar
	rckey := cf.key[:]
	crc4, err := rc4.NewCipher(rckey)
	if err != nil {
		panic (err)
	}
	crc4.XORKeyStream(res, res)
	salsa20.XORKeyStream(res,res,cf.nonce, &cf.key)
	return res
}