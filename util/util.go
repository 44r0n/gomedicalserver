package util

import (
	"crypto/rc4"
	"code.google.com/p/go.crypto/salsa20"
	"math/rand"
    "time"
    "code.google.com/p/go.crypto/scrypt"
	"bytes"
)

type Cifrador struct {
	key [32]byte
	nonce[] byte
	rckey [32]byte
}

var n, r, p, keylen int

func init() {
	n = 65536
	r = 16
	p = 3
	keylen = 32
}

func NuevaContraseña(password string) ([]byte, []byte) {
	salt := RandomString()
	dk, _ := scrypt.Key([]byte(password), []byte(salt),n, r, p, keylen)
	cifra := NuevoCifrador()
	return dk, cifra.Encrypt([]byte(salt))
}

func CheckPassword(plainpassword string, password, salte []byte) bool {
	cifra := NuevoCifrador()
	dk, _ := scrypt.Key([]byte(plainpassword), cifra.Decrypt(salte), n, r, p, keylen)
	return bytes.Equal(dk,password)
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

func RandomString() string {
	min := 33
	max := 126
    var result   bytes.Buffer
    var temp string
    l := 64
    for i:=0 ; i<l ;  {
        if string(randInt(min,max))!=temp {
        temp = string(randInt(min,max))
        result.WriteString(temp)
        i++
      }
    }
	return result.String()
}

func randInt(min int , max int) int {
        rand.Seed( time.Now().UTC().UnixNano())
        return min + rand.Intn(max-min)
}