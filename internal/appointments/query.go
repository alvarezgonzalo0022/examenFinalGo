package appointments

var (
	QueryInsertAppointment          = `INSERT INTO appointments (id_patient, id_dentist, description, appointment_date, appointment_time) VALUES (?,?,?,?,?)`
	QueryGetAllAppointment          = `SELECT appointments.*, dentists.last_name AS dentist_lastname, patients.last_name AS patient_lastname FROM appointments INNER JOIN dentists ON appointments.dentist_id = dentists.id INNER JOIN patients ON appointments.patient_id = patients.id `
	QueryGetAllAppointmentByPatient = `SELECT id, id_patient, id_dentist, description, appointment_date, appointment_time FROM appointments WHERE id_patient = ?`
	QueryGetByIdAppointment        = `SELECT appointments.*, dentists.last_name AS dentist_lastname, patients.last_name AS patient_lastname FROM appointments INNER JOIN dentists ON appointments.dentist_id = dentists.id INNER JOIN patients ON appointments.patient_id = patients.id WHERE appointments.id = ?`
	QueryUpdateAppointment          = `UPDATE appointments SET id_patient = ?, id_dentist = ?, description = ?, appointment_date = ?, appointment_time = ? WHERE id = ?`
	QueryDeleteAppointment          = `DELETE FROM appointments WHERE id = ?`
)
