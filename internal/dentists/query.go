package dentists

var (
	QueryInsertDentist = `INSERT INTO dentists (first_name, last_name, registration_id)
	VALUES(?,?,?)`
	QueryGetAllDentists = `SELECT id, first_name, last_name, registration_id 
	FROM dentists`
	QueryDeleteDentist  = `DELETE FROM dentists WHERE id = ?`
	QueryGetDentistById = `SELECT id, first_name, last_name, registration_id
	FROM dentists WHERE id = ?`
	QueryUpdateDentist = `UPDATE dentists SET first_name = ?, last_name = ?, registration_id = ?
	WHERE id = ?`
)
