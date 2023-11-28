package appointments

var (
	QueryInsertAppointment           = `INSERT INTO appointments (id_patient, id_dentist, description, appointment_date, appointment_time) VALUES (?,?,?,?,?)`
	QueryGetAllAppointment           = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments`
	QueryGetAllAppointmentByPatient = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments WHERE id_patient = ?`
	QueryGetByIdAppointment          = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments WHERE id = ?`
	QueryUpdateAppointment           = `UPDATE appointments SET id_patient = ?, id_dentist = ?, description = ?, appointment_date = ?, appointment_time = ? WHERE id = ?`
	QueryDeleteAppointment           = `DELETE FROM appointments WHERE id = ?`
)
