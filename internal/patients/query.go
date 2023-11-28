package patients

var (
	QueryInsertPatient = `INSERT INTO patients (first_name, last_name, address, document, admission_date) VALUES(?,?,?,?,?)`
	QueryGetAllPatients = `SELECT id, first_name, last_name, address, document, admission_date FROM patients`
	QueryDeletePatient  = `DELETE FROM _time WHERE id = ?`
	QueryGetPatientById = `SELECT id, first_name, last_name, address, document, admission_date FROM _time WHERE id = ?`
	QueryUpdatePatient = `UPDATE patients SET first_name = ?, last_name = ?, address = ?, document = ?, admission_date = ? WHERE id = ?`
)
