package main

import (
	"hospital/modelo"
	"fmt"
	//"crypto/rc4"
	//"code.google.com/p/go.crypto/salsa20"
	//"crypto/aes"
	//"encoding/base64"
	//"crypto/cipher"
	//"time"
)

func main() {
	/*key := []byte("taca")
	texto := []byte("a crifrar")
	cifrador, _ := rc4.NewCipher(key)
	cifrador.XORKeyStream(texto,texto)
	salsa.XORKeyStream
	fmt.Println(texto)
	cifrador, _ = rc4.NewCipher(key)
	cifrador.XORKeyStream(texto,texto)
	fmt.Printf("after decoding, src=% s\n", texto)*/

	/*var key [32]byte
	var rckey []byte
	contra := "qwertyuiopasdfghjklñzxcvbnm12345"
	copy(key[:],contra)
	rckey = key[:]
	plaintext := []byte("To encrypt")

	salsa20.XORKeyStream(plaintext,plaintext,[]byte("12345678"),&key)
	fmt.Println(plaintext)
	cifrador, _ := rc4.NewCipher(rckey)
	cifrador.XORKeyStream(plaintext,plaintext)
	fmt.Println(plaintext)
	cifrador, _ = rc4.NewCipher(rckey)
	cifrador.XORKeyStream(plaintext,plaintext)
	salsa20.XORKeyStream(plaintext,plaintext,[]byte("12345678"),&key)
	texto := string(plaintext)
	fmt.Println(texto)*/
	

	p := modelo.NuevoPaciente()
	//p.Search("12345678S")

	p.GetById(1)	
	fmt.Println(p.DNI)
	fmt.Println(p.Nombre)
	fmt.Println(p.Apellidos)
	/*p.Nombre = "Angel"
	p.Save()
	fmt.Println(p.DNI)
	fmt.Println(p.Nombre)
	fmt.Println(p.Apellidos)
	fmt.Println("Modificado")
	if(p.Delete()) {
		fmt.Println("Eliminado")
	}*/

	/*p.DNI = "12345678S"
	p.Nombre = "Manolo"
	p.Apellidos = "Solo"
	p.Save()
	fmt.Println("Paciente Guardado")*/

	/*d := new(modelo.Doctor)
	d.DNI = "23456789B"
	d.Nombre = "Pepe"
	d.Apellidos = "Ese"
	d.Save()
	fmt.Println("Doctor Guardado")

	t := new(modelo.Tratamiento)
	t.NombreEnfermedad = "Una"
	t.Observaciones = "Te veo"
	t.Save()
	fmt.Println("Tratamiento Guardado")

	h := new(modelo.Historial)
	h.Paciente = p.GetId()
	h.Doctor = d.GetId()
	h.Tratamiento = t.GetId()
	h.Observaciones = "Observación de tratamiento"
	h.Fecha = time.Now()
	h.Save()
	fmt.Println("Historial Guardado")*/
	/*h := new(modelo.Historial)
	h.GetById(1)
	fmt.Println(h.Fecha.Format("2006-01-02"))*/


}