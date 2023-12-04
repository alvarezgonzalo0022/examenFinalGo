package patients

var (
	QueryInsertPatient = `INSERT INTO patients (first_name, last_name, address, document, admission_date) VALUES(?,?,?,?,?)`
	QueryGetAllPatients = `SELECT id, first_name, last_name, address, document, admission_date FROM patients`
	QueryDeletePatient  = `DELETE FROM patients WHERE id = ?`
	QueryGetPatientById = `SELECT id, first_name, last_name, address, document, admission_date FROM patients WHERE id = ?`
	QueryUpdatePatient = `UPDATE patients SET first_name = ?, last_name = ?, address = ?, document = ?, admission_date = ? WHERE id = ?`
)
