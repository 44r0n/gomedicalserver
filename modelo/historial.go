package modelo

import "hospital/datos"
import "hospital/util"
import "time"

type Historial struct {
	id int
	Paciente int
	Doctor int
	Tratamiento int
	Observaciones string 
	Fecha time.Time
}

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
//                            * Métodos privados *                          //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

func (hs *Historial) encrypt() ([]byte, []byte) {
	cifra := util.NuevoCifrador()
	return cifra.Encrypt([]byte(hs.Observaciones)), cifra.Encrypt([]byte(hs.Fecha.Format("2006-01-02")))
}

func (hs *Historial) decrypt(observaciones, fecha []byte) {
	cifra := util.NuevoCifrador()
	hs.Observaciones = string(cifra.Decrypt(observaciones))
	hs.Fecha, _ = time.Parse("2006-01-02",string(cifra.Decrypt(fecha)))
}

func (hs *Historial) insert() {
	if hs.id != 0 {
		return
	}
	database.Connect()
	defer database.Close()
	observaciones, fecha := hs.encrypt()
	database.ExecuteNonQuery("INSERT INTO historiales (Paciente, Doctor, Observaciones, Tratamiento, Fecha) VALUES (?,?,?,?,?)",hs.Paciente,hs.Doctor,observaciones,hs.Tratamiento,fecha)
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
	observaciones, fecha := hs.encrypt()
	database.ExecuteNonQuery("UPDATE historiales SET Paciente = ?, Doctor = ?, Observaciones = ?, Tratamiento = ?, Fecha = ? WHERE id = ?",hs.Paciente,hs.Doctor,observaciones,hs.Tratamiento,fecha,hs.id)
	return true
}

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
//                            * Métodos Públicos *                          //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

func(hs *Historial) GetById(id int) *Historial {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM historiales WHERE id = ?",id)
	rows.Next()
	var observaciones, fecha []byte
	rows.Scan(&hs.id,&hs.Paciente,&hs.Doctor,&observaciones,&fecha,&hs.Tratamiento)
	hs.decrypt(observaciones,fecha)
	return hs
}

func(hs *Historial) GetId() int {
	return hs.id
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