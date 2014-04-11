DROP TABLE IF EXISTS Historiales;
DROP TABLE IF EXISTS Tratamientos;
DROP TABLE IF EXISTS Doctores;
DROP TABLE IF EXISTS Pacientes;

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
	DNI char(9),
	Nombre varchar(20),
	Apellidos varchar(40),
	Password blob
);

CREATE TABLE IF NOT EXISTS Tratamientos
(
	Id int Primary Key Auto_Increment,
	NombreEnfermedad varchar(40),
	Observaciones varchar(255)
);

CREATE TABLE IF NOT EXISTS Historiales
(
	Id int Primary Key Auto_Increment,
	Paciente int,
	Doctor int,
	Observaciones varchar(255),
	Fecha date,
	Tratamiento int,
	CONSTRAINT fk_Historiales_Pacientes FOREIGN KEY (Paciente) REFERENCES Pacientes (Id) ON UPDATE SET NULL ON DELETE SET NULL,
	CONSTRAINT fk_Historiales_Doctores FOREIGN KEY (Doctor) REFERENCES Doctores (Id) ON UPDATE SET NULL ON DELETE SET NULL
);