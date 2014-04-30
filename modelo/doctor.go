package modelo

import "hospital/datos"
import "hospital/util"

type Doctor struct {
	id int
	DNI string
	Nombre string
	Apellidos string
}

func NuevoDoctor() *Doctor{
	doc := new(Doctor)
	return doc
}

func (doc *Doctor) encrypt() ([]byte, []byte, []byte) {
	cifra := util.NuevoCifrador()
	return cifra.Encrypt([]byte(doc.DNI)), cifra.Encrypt([]byte(doc.Nombre)), cifra.Encrypt([]byte(doc.Apellidos))
}

func (doc *Doctor) decrypt(dni []byte, nombre []byte, apellidos []byte) {
	cifra := util.NuevoCifrador()
	doc.DNI = string(cifra.Decrypt(dni))
	doc.Nombre = string(cifra.Decrypt(nombre))
	doc.Apellidos = string(cifra.Decrypt(apellidos))
}

func (doc *Doctor) GetById(id int) *Doctor {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM doctores WHERE id = ?",id)
	rows.Next()
	var dni, nombre, apellidos []byte
	rows.Scan(&doc.id,&dni,&nombre,&apellidos)
	doc.decrypt(dni,nombre,apellidos)
	return doc
}

func (doc *Doctor) GetId() int {
	return doc.id
}

func (doc *Doctor) insert() {
	if doc.id != 0 {
		return
	}
	database.Connect()
	defer database.Close()
	dni, nombre , apellidos := doc.encrypt()
	database.ExecuteNonQuery("INSERT INTO doctores (DNI, Nombre, Apellidos) VALUES (?,?,?)", dni, nombre, apellidos);
	rows := database.ExecuteQuery("SELECT MAX(id) FROM doctores")
	rows.Next()
	var last int
	rows.Scan(&last)
	doc.id = last
}

func (doc *Doctor) update() bool {
	if doc.id  == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	dni, nombre , apellidos := doc.encrypt()
	database.ExecuteNonQuery("UPDATE pacientes SET DNI = ?, Nombre = ?, Apellidos = ? WHERE id = ?",dni,nombre,apellidos,doc.id)
	return true
}

func (doc *Doctor) Delete() bool {
	if doc.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("DELETE FROM pacientes WHERE id = ?",doc.id)
	return true
}

func (doc *Doctor) Save() bool {
	if(doc.id != 0) {
		return doc.update()
	} else {
		doc.insert()
		if doc.id != 0 {
			return true
		}
	}
	return false
}