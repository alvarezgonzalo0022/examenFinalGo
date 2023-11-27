package odontologos

var (
	QueryInsertOdontologo = `INSERT INTO odontologos(name, last_name, code)
	VALUES(?,?,?)`
	QueryGetAllOdontologos = `SELECT id, name, last_name, code 
	FROM odontologos`
	QueryDeleteOdontologo  = `DELETE FROM odontologos WHERE id = ?`
	QueryGetOdontologoById = `SELECT id, name, last_name, code
	FROM odontologos WHERE id = ?`
	QueryUpdateOdontologo = `UPDATE odontologos SET name = ?, last_name = ?, code = ?
	WHERE id = ?`
)
