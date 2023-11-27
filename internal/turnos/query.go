package turnos

var (
	QueryInsertTurno           = `INSERT INTO turno (id_paciente, id_odontologo, description, fecha, hora) VALUES (?,?,?,?,?)`
	QueryGetAllTurno           = `SELECT id, id_paciente, id_odontologo, description, fecha, hora FROM turno`
	QueryGetAllTurnoByPaciente = `SELECT id, id_paciente, id_odontologo, description, fecha, hora FROM turno WHERE id_paciente = ?`
	QueryGetByIdTurno          = `SELECT id, id_paciente, id_odontologo, description, fecha, hora FROM turno WHERE id = ?`
	QueryUpdateTurno           = `UPDATE turno SET id_paciente = ?, id_odontologo = ?, description = ?, fecha = ?, hora = ? WHERE id = ?`
	QueryDeleteTurno           = `DELETE FROM turno WHERE id = ?`
)