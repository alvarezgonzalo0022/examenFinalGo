package pacientes

var (
	QueryInsertPaciente = `INSERT INTO patients (first_name, last_name, address, document, admission_date) VALUES(?,?,?,?,?)`
	QueryGetAllPacientes = `SELECT id, first_name, last_name, address, document, admission_date FROM patients`
	QueryDeletePaciente  = `DELETE FROM _time WHERE id = ?`
	QueryGetPacienteById = `SELECT id, first_name, last_name, address, document, admission_date FROM _time WHERE id = ?`
	QueryUpdatePaciente = `UPDATE patients SET first_name = ?, last_name = ?, address = ?, document = ?, admission_date = ? WHERE id = ?`
)
