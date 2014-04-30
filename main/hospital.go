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

	p := modelo.NuevoPaciente()
	p.DNI = "12345678S"
	p.Nombre = "Manolo"
	p.Apellidos = "Solo"
	p.Save()
	fmt.Println("Paciente Guardado")

	p2 := modelo.NuevoPaciente()

	//p.Search("12345678S")

	p2.GetById(1)	
	fmt.Println(p2.DNI)
	fmt.Println(p2.Nombre)
	fmt.Println(p2.Apellidos)
	/*p.Nombre = "Angel"
	p.Save()
	fmt.Println(p.DNI)
	fmt.Println(p.Nombre)
	fmt.Println(p.Apellidos)
	fmt.Println("Modificado")
	if(p.Delete()) {
		fmt.Println("Eliminado")
	}*/

	

	d := modelo.NuevoDoctor()
	d.DNI = "23456789B"
	d.Nombre = "Pepe"
	d.Apellidos = "Ese"
	d.Save()
	fmt.Println("Doctor Guardado")

	d2 := modelo.NuevoDoctor()
	d2.GetById(1)
	fmt.Println(d2.DNI)
	fmt.Println(d2.Nombre)
	fmt.Println(d2.Apellidos)

	
	t := new(modelo.Tratamiento)
	t.NombreEnfermedad = "Una"
	t.Observaciones = "Te veo"
	t.Save()
	fmt.Println("Tratamiento Guardado")

	t2 := new(modelo.Tratamiento)
	t2.GetById(1)
	fmt.Println(t2.NombreEnfermedad)
	fmt.Println(t2.Observaciones)

	/*
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