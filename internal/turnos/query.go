package turnos

var (
	QueryInsertTurno           = `INSERT INTO appointments (id_patient, id_dentist, description, appointment_date, appointment_time) VALUES (?,?,?,?,?)`
	QueryGetAllTurno           = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments`
	QueryGetAllTurnoByPaciente = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments WHERE id_patient = ?`
	QueryGetByIdTurno          = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments WHERE id = ?`
	QueryUpdateTurno           = `UPDATE appointments SET id_patient = ?, id_dentist = ?, description = ?, appointment_date = ?, appointment_time = ? WHERE id = ?`
	QueryDeleteTurno           = `DELETE FROM appointments WHERE id = ?`
)
