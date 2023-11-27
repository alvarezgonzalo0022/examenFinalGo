package pacientes

var (
	QueryInsertPaciente = `INSERT INTO pacientes(name, last_name, address, dni, ingress_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPacientes = `SELECT id, name, last_name, address, dni, ingress_date 
	FROM pacientes`
	QueryDeletePaciente  = `DELETE FROM pacientes WHERE id = ?`
	QueryGetPacienteById = `SELECT id, name, last_name, address, dni, ingress_date
	FROM pacientes WHERE id = ?`
	QueryUpdatePaciente = `UPDATE pacientes SET name = ?, last_name = ?, address = ?, dni = ?, ingress_date = ?
	WHERE id = ?`
)
