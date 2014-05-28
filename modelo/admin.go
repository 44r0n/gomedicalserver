package modelo

import (
	"hospital/datos"
	"hospital/util"
)

type Admin struct {
	id int
	Nombre string
}

func NuevoAdmin () *Admin{
	admin := new(Admin)
	return admin
}

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
//                            * Métodos privados *                          //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

func (admin *Admin) encrypt() ([]byte) {
	cifra := util.NuevoCifrador()
	return cifra.Encrypt([]byte(admin.Nombre))
}

func (admin *Admin) decrypt(nombre []byte) {
	cifra := util.NuevoCifrador()
	admin.Nombre = string(cifra.Decrypt(nombre))
}

func (admin *Admin) insert() {
	if admin.id != 0 {
		return
	}

	database.Connect()
	defer database.Close()
	nombre := admin.encrypt()
	database.ExecuteNonQuery("INSERT INTO Admin (Nombre) VALUES (?)",nombre)
	rows := database.ExecuteQuery("SELECT MAX(id) FROM Admin")
	rows.Next()
	var last int
	rows.Scan(&last)
	admin.id = last
}

func (admin *Admin) update() bool {
	if admin.id == 0 {
		return false
	}
	database.Connect()
	defer database.Close()
	nombre := admin.encrypt()
	database.ExecuteNonQuery("UPDATE Admin SET Nombre = ? WHERE id = ?",nombre, admin.id)
	return true
}

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
//                            * Métodos Públicos *                          //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

func (admin *Admin) GetById(id int) *Admin {
	database.Connect()
	defer database.Close()
	rows := database.ExecuteQuery("SELECT Id, Nombre FROM Admin WHERE id = ?",id)
	rows.Next()
	var nombre []byte
	rows.Scan(&admin.id,&nombre)
	admin.decrypt(nombre)
	return admin
} 

func (admin *Admin) GetId() int {
	return admin.id
}

func (admin *Admin) Delete() bool {
	if admin.id == 0 {
		return false
	}

	database.Connect()
	defer database.Close()
	database.ExecuteNonQuery("DELETE FROM Admin WHERE id = ?",admin.id)
	return true
}

func (admin *Admin) Save() bool {
	if admin.id !=0 {
		return admin.update()
	} else {
		admin.insert()
		if admin.id != 0 {
			return true
		}
	}
	return false
}

func (admin *Admin) SetPassword(password string) bool {
	if admin.id == 0 {
		return false
	} else {
		database.Connect()
		defer database.Close()
		dk, salt := util.NuevaContraseña(password)
		database.ExecuteNonQuery("UPDATE Admin SET Password = ?, Salt = ? WHERE id = ?",dk, salt, admin.id)
		return true
	}
}

func AuthenticateAdmin(nombre string, password string) bool {
	database.Connect()
	defer database.Close()
	cifra := util.NuevoCifrador()
	nombree := cifra.Encrypt([]byte(nombre))
	rows := database.ExecuteQuery("SELECT Password, Salt FROM Admin WHERE Nombre = ?",nombree)
	rows.Next()
	var passwordbd, salt []byte
	rows.Scan(&passwordbd,&salt)
	return util.CheckPassword(password,passwordbd,salt)
}