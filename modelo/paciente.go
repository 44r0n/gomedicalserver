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

func (pac *Paciente) encrypt() ([]byte, []byte, []byte) {
	cifra := util.NuevoCifrador()
	return cifra.Encrypt([]byte(pac.DNI)), cifra.Encrypt([]byte(pac.Nombre)), cifra.Encrypt([]byte(pac.Apellidos))
}

func (pac *Paciente) decrypt(dni []byte, nombre []byte, apellidos []byte) {
	cifra := util.NuevoCifrador()
	pac.DNI = string(cifra.Decrypt(dni))
	pac.Nombre = string(cifra.Decrypt(nombre))
	pac.Apellidos = string(cifra.Decrypt(apellidos))
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
	pac.decrypt(dni,nombre,apellidos)
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
	pac.decrypt(dni,nombre,apellidos)
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
	dni, nombre, apellidos := pac.encrypt()
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
	database.ExecuteNonQuery("DELETE FROM pacientes WHERE Id = ?",pac.id)
	return true
}

func (pac *Paciente) update() bool {
	if pac.id  == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	dni, nombre, apellidos := pac.encrypt()
	database.ExecuteNonQuery("UPDATE pacientes SET DNI = ?, Nombre = ?, Apellidos = ? WHERE Id = ?",dni,nombre,apellidos,pac.id)
	return true
}