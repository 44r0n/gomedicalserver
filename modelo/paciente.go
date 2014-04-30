package modelo

import "hospital/datos"
import "hospital/util"

type Paciente struct {
	id int
	DNI string
	Nombre string
	Apellidos string
}

func NuevoPaciente() *Paciente {
	pac := new(Paciente)
	return pac
}

func (pac *Paciente) Save() bool {
	if(pac.id != 0) {
		return pac.update()
	} else {
		pac.insert()
		if pac.id != 0 {
			return true
		}
	}
	return false
}

func (pac *Paciente) GetById(id int) *Paciente {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM pacientes WHERE id = ?",id)
	rows.Next()
	var dni, nombre, apellidos []byte
	rows.Scan(&pac.id,&dni,&nombre,&apellidos)
	cifra := util.NuevoCifrador()
	pac.DNI = string(cifra.Decrypt(dni))
	pac.Nombre = string(cifra.Decrypt(nombre))
	pac.Apellidos = string(cifra.Decrypt(apellidos))
	return pac
}

func (pac *Paciente) Search(dnis string) *Paciente {
	database.Connect()
	defer database.Close()
	cifra := util.NuevoCifrador()
	dnib := cifra.Encrypt([]byte(dnis))
	rows := database.ExecuteQuery("SELECT * FROM pacientes WHERE DNI = ?",dnib)
	rows.Next()
	var dni, nombre, apellidos []byte
	rows.Scan(&pac.id,&dni,&nombre,&apellidos)
	pac.DNI = string(cifra.Decrypt(dni))
	pac.Nombre = string(cifra.Decrypt(nombre))
	pac.Apellidos = string(cifra.Decrypt(apellidos))
	return pac
}

func (pac *Paciente) GetId() int {
	return pac.id
}

func (pac *Paciente) insert() {
	if pac.id != 0 {
		return
	}
	database.Connect()
	defer database.Close()
	cifra := util.NuevoCifrador()
	dni := cifra.Encrypt([]byte(pac.DNI))
	nombre := cifra.Encrypt([]byte(pac.Nombre))
	apellidos := cifra.Encrypt([]byte(pac.Apellidos))
	database.ExecuteNonQuery("INSERT INTO pacientes (DNI, Nombre, Apellidos) VALUES (?,?,?)", dni,nombre,apellidos);
	rows := database.ExecuteQuery("SELECT MAX(id) FROM pacientes")
	rows.Next()
	var last int
	rows.Scan(&last)
	pac.id = last
}

func (pac *Paciente) Delete() bool {
	if pac.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("DELETE FROM pacientes WHERE id = ?",pac.id)
	return true
}

func (pac *Paciente) update() bool {
	if pac.id  == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("UPDATE pacientes SET DNI = ?, Nombre = ?, Apellidos = ? WHERE id = ?",pac.DNI,pac.Nombre,pac.Apellidos,pac.id)
	return true
}