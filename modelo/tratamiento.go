package modelo

import "hospital/datos"

type Tratamiento struct {
	id int
	NombreEnfermedad string
	Observaciones string
}

func (tr *Tratamiento) GetById(id int) *Tratamiento {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM tratamientos WHERE id = ?",id)
	rows.Next()
	rows.Scan(&tr.id,&tr.NombreEnfermedad,&tr.Observaciones)
	return tr
}

func (tr *Tratamiento) GetId() int {
	return tr.id
}

func (tr *Tratamiento) insert() {
	if tr.id != 0 {
		return
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("INSERT INTO tratamientos (NombreEnfermedad, Observaciones) VALUES (?,?)",tr.NombreEnfermedad,tr.Observaciones)
	rows:= database.ExecuteQuery("SELECT MAX(id) FROM tratamientos")
	rows.Next()
	var last int
	rows.Scan(&last)
	tr.id = last
}

func (tr * Tratamiento) update() bool {
	if tr.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("UPDATE tratamientos SET NombreEnfermedad = ?, Observaciones = ? WHERE id = ?",tr.NombreEnfermedad,tr.Observaciones,tr.id)
	return true
}

func (tr *Tratamiento) Delete() bool {
	if tr.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("DELETE FROM tratamientos WHERE id = ?",tr.id)
	return true
}

func (tr *Tratamiento) Save() bool {
	if tr.id != 0 {
		return tr.update()
	} else {
		tr.insert()
		if tr.id != 0 {
			return true
		}
	}
	return false
}