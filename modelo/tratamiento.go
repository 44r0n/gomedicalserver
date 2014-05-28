package modelo

import "hospital/datos"
import "hospital/util"

type Tratamiento struct {
	id int
	NombreEnfermedad string
	Observaciones string
}

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
//                            * Métodos privados *                          //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

func (tr *Tratamiento) encrypt() ([]byte, []byte) {
	cifra := util.NuevoCifrador()
	return cifra.Encrypt([]byte(tr.NombreEnfermedad)), cifra.Encrypt([]byte(tr.Observaciones))
}

func (tr *Tratamiento) decrypt(nombreEnfermedad []byte, observaciones []byte) {
	cifra := util.NuevoCifrador()
	tr.NombreEnfermedad = string(cifra.Decrypt(nombreEnfermedad))
	tr.Observaciones = string(cifra.Decrypt(observaciones))
}

func (tr *Tratamiento) insert() {
	if tr.id != 0 {
		return
	}
	database.Connect()
	defer database.Close()
	nombreEnfermedad, observaciones := tr.encrypt()
	database.ExecuteNonQuery("INSERT INTO tratamientos (NombreEnfermedad, Observaciones) VALUES (?,?)",nombreEnfermedad,observaciones)
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
	nombreEnfermedad, observaciones := tr.encrypt()
	database.ExecuteNonQuery("UPDATE tratamientos SET NombreEnfermedad = ?, Observaciones = ? WHERE id = ?",nombreEnfermedad,observaciones,tr.id)
	return true
}

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
//                            * Métodos Públicos *                          //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

func NuevoTratamiento() *Tratamiento {
	tr := new(Tratamiento)
	return tr
}

func (tr *Tratamiento) GetById(id int) *Tratamiento {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT * FROM tratamientos WHERE id = ?",id)
	rows.Next()
	var nombreEnfermedad, observaciones []byte
	rows.Scan(&tr.id,&nombreEnfermedad,&observaciones)
	tr.decrypt(nombreEnfermedad,observaciones)
	return tr
}

func (tr *Tratamiento) GetId() int {
	return tr.id
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