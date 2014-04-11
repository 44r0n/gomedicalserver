package modelo

import "hospital/datos"
import "time"

type Historial struct {
	id int
	Paciente int
	Doctor int
	Tratamiento int
	Observaciones string 
	Fecha time.Time
}

func(hs *Historial) GetById(id int) *Historial {
	var fecha string
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM historiales WHERE id = ?",id)
	rows.Next()
	rows.Scan(&hs.id,&hs.Paciente,&hs.Doctor,&hs.Observaciones,&fecha,&hs.Tratamiento)
	hs.Fecha,_  = time.Parse("2006-01-02",fecha)
	return hs
}

func(hs *Historial) GetId() int {
	return hs.id
}

func (hs *Historial) insert() {
	if hs.id != 0 {
		return
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("INSERT INTO historiales (Paciente, Doctor, Observaciones, Tratamiento, Fecha) VALUES (?,?,?,?,?)",hs.Paciente,hs.Doctor,hs.Observaciones,hs.Tratamiento,hs.Fecha)
	rows := database.ExecuteQuery("SELECT MAX(id) FROM historiales")
	rows.Next()
	var last int
	rows.Scan(&last)
	hs.id = last
}

func (hs *Historial) update() bool {
	if hs.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("UPDATE historiales SET Paciente = ?, Doctor = ?, Observaciones = ?, Tratamiento = ?, Fecha = ? WHERE id = ?",hs.Paciente,hs.Doctor,hs.Observaciones,hs.Tratamiento,hs.Fecha,hs.id)
	return true
}

func (hs *Historial) Delete() bool{
	if  hs.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("DELETE FROM historiales WHERE id = ?",hs.id)
	return true
}

func (hs *Historial) Save() bool {
	if hs.id != 0 {
		return hs.update()
	} else {
		hs.insert()
		if hs.id != 0 {
			return true
		}
	}
	return false
}