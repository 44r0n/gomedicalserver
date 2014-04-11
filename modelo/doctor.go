package modelo

import "hospital/datos"

type Doctor struct {
	id int
	DNI string
	Nombre string
	Apellidos string
}

func (doc *Doctor) GetById(id int) *Doctor {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM doctores WHERE id = ?",id)
	rows.Next()
	rows.Scan(&doc.id,&doc.DNI,&doc.Nombre,&doc.Apellidos)
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
	database.ExecuteNonQuery("INSERT INTO doctores (DNI, Nombre, Apellidos) VALUES (?,?,?)", doc.DNI, doc.Nombre, doc.Apellidos);
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
	database.ExecuteNonQuery("UPDATE pacientes SET DNI = ?, Nombre = ?, Apellidos = ? WHERE id = ?",doc.DNI,doc.Nombre,doc.Apellidos,doc.id)
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