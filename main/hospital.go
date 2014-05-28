package main

import (
	"hospital/modelo"
	"fmt"
)

func main() {	
	p := modelo.NuevoPaciente()
	p.DNI = "12345678S"
	p.Nombre = "Manolo"
	p.Apellidos = "Solo"
	p.Save()
	fmt.Println("Paciente Guardado")
	p.Nombre = "Antonio"
	p.Save()
}