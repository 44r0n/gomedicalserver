DROP TABLE IF EXISTS Historiales;
DROP TABLE IF EXISTS Tratamientos;
DROP TABLE IF EXISTS Doctores;
DROP TABLE IF EXISTS Pacientes;
DROP TABLE IF EXISTS Admin;

CREATE TABLE IF NOT EXISTS Pacientes
(
	Id int Primary Key Auto_Increment,
	DNI blob,
	Nombre blob,
	Apellidos blob
);

CREATE TABLE IF NOT EXISTS Doctores
(
	Id int Primary Key Auto_Increment,
	DNI blob,
	Nombre blob,
	Apellidos blob,
	Password blob,
	Salt blob

);

CREATE TABLE IF NOT EXISTS Tratamientos
(
	Id int Primary Key Auto_Increment,
	NombreEnfermedad blob,
	Observaciones blob
);

CREATE TABLE IF NOT EXISTS Historiales
(
	Id int Primary Key Auto_Increment,
	Paciente int,
	Doctor int,
	Observaciones blob,
	Fecha blob,
	Tratamiento int,
	CONSTRAINT fk_Historiales_Pacientes FOREIGN KEY (Paciente) REFERENCES Pacientes (Id) ON UPDATE SET NULL ON DELETE SET NULL,
	CONSTRAINT fk_Historiales_Doctores FOREIGN KEY (Doctor) REFERENCES Doctores (Id) ON UPDATE SET NULL ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS Admin
(
	Id int Primary Key Auto_Increment,
	Nombre blob,
	Password blob,
	Salt blob
);