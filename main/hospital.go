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
	//"code.google.com/p/go.crypto/scrypt"
	//"math/rand"
    //"time"
    //"bytes"
    //"hospital/util"
)

/*func randomString () string {
	//min 33 - 126
	min := 33
	max := 126
    var result   bytes.Buffer
    var temp string
    l := randInt(32,128)
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
}*/

func main() {	

	
	/*texto := randomString()
	fmt.Println(texto)
	dk, _ := scrypt.Key([]byte("password"), []byte(texto),65536, 16, 3, 32)

	fmt.Println(dk)*/

	/*p := modelo.NuevoPaciente()
	p.DNI = "12345678S"
	p.Nombre = "Manolo"
	p.Apellidos = "Solo"
	p.Save()
	fmt.Println("Paciente Guardado")
	p.Nombre = "Antonio"
	p.Save()*/

	//p2 := modelo.NuevoPaciente()

	//p.Search("12345678S")

	/*p2.GetById(1)	
	fmt.Println(p2.DNI)
	fmt.Println(p2.Nombre)
	fmt.Println(p2.Apellidos)*/
	/*p.Nombre = "Angel"
	p.Save()
	fmt.Println(p.DNI)
	fmt.Println(p.Nombre)
	fmt.Println(p.Apellidos)
	fmt.Println("Modificado")
	if(p.Delete()) {
		fmt.Println("Eliminado")
	}*/

	

	/*d := modelo.NuevoDoctor()
	d.DNI = "23456789B"
	d.Nombre = "Pepe"
	d.Apellidos = "Ese"
	if d.Save() {
		fmt.Println("Doctor Guardado con el id: ",d.GetId())
	} else {
		fmt.Println("No se ha guardado el doctor")
	}*/
	/*d.Nombre = "Juan"
	if !d.Save() {
		fmt.Println("No se ha modificado")
	}
	fmt.Println(d.DNI)
	fmt.Println(d.Nombre)
	fmt.Println(d.Apellidos)*/
	/*if d.SetPassword("taca") {
		fmt.Println("Contraseña establecida")
	} else {
		fmt.Println("La contraseña no se ha guardado")
	}

	if modelo.Authenticate("23456789B","taca") {
		fmt.Println("Autenticado")
	} else {
		fmt.Println("No autenticado")
	}*/

	/*d2 := modelo.NuevoDoctor()
	d2.GetById(1)
	fmt.Println(d2.DNI)
	fmt.Println(d2.Nombre)
	fmt.Println(d2.Apellidos)*/

	
	/*t := new(modelo.Tratamiento)
	t.NombreEnfermedad = "Una"
	t.Observaciones = "Te veo"
	t.Save()
	fmt.Println("Tratamiento Guardado")*/

	/*t2 := new(modelo.Tratamiento)
	t2.GetById(1)
	fmt.Println(t2.NombreEnfermedad)
	fmt.Println(t2.Observaciones)*/

	
	/*h := new(modelo.Historial)
	h.Paciente = p.GetId()
	h.Doctor = d.GetId()
	h.Tratamiento = t.GetId()
	h.Observaciones = "Observación de tratamiento"
	h.Fecha = time.Now()
	if h.Save() {
		fmt.Println("Historial Guardado")
	}
	h1 := new(modelo.Historial)
	h1.GetById(1)
	fmt.Println(h1.Fecha.Format("2006-01-02"))*/

	fmt.Println("Historiales totales: ")
	fmt.Println(modelo.Totales())

}