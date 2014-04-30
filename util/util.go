package util

import (
	"crypto/rc4"
	"code.google.com/p/go.crypto/salsa20"
)

type Cifrador struct {
	key [32]byte
	nonce[] byte
	rckey [32]byte
}

func NuevoCifrador() *Cifrador {
	cf := new(Cifrador)
	copy(cf.key[:], "O?SvUec4TXkV58CL-p?C*;K./Rt.v+BL")
	cf.nonce = []byte("-Ts#RVaZY?2Hy_geC36(KhH;")
	copy(cf.rckey[:], "Hh;F.#@>7>7~1#nu9K9gQ6_^M)&&X*8w")
	return cf
}

func (cf *Cifrador) Encrypt(aencriptar []byte) []byte {
	res := aencriptar
	salsa20.XORKeyStream(res, res, cf.nonce, &cf.key)
	rckey := cf.rckey[:]
	crc4, err := rc4.NewCipher(rckey)
	if err != nil {
		panic(err)
	}
	crc4.XORKeyStream(res, res)
	return res
}

func (cf *Cifrador) Decrypt(desencriptar []byte) []byte {
	res := desencriptar
	rckey := cf.rckey[:]
	crc4, err := rc4.NewCipher(rckey)
	if err != nil {
		panic (err)
	}
	crc4.XORKeyStream(res, res)
	salsa20.XORKeyStream(res,res,cf.nonce, &cf.key)
	return res
}