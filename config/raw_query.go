package config

const (
	SelectSchedulePagination = `
	SELECT id, user_id, date, start_time, end_time, documentation
	FROM schedules
  WHERE deleted_at IS NULL
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2`

	// User
	InsertUser           = `INSERT INTO users (name,email,password,role) VALUES ($1,$2,$3,$4) RETURNING id`
	SelectUserPagination = `SELECT id, name, email, role FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectUserByID       = `SELECT id, name, email, role FROM users WHERE id = $1 AND deleted_at IS NULL`
	SelectUserByEmail    = `SELECT id, name, email, password, role FROM users WHERE email = $1 AND deleted_at IS NULL`
	SelectUserByRole     = `SELECT id, name, email, role FROM users WHERE role = $3 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2 `
	UpdateUser           = `UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5`
	DeleteUser           = `UPDATE users SET deleted_at = $1 WHERE id = $2`

	// Questions
	InsertQuestions             = `INSERT INTO questions (user_id,schedule_id,description) VALUES ($1,$2,$3) RETURNING id`
	SelectQuestionsByID         = `SELECT id, schedule_id, description, status FROM questions WHERE id = $1 AND deleted_at IS NULL`
	SelectQuestionsByScheduleID = `SELECT id, schedule_id, description, status FROM questions WHERE schedule_id = $1 AND deleted_at IS NULL`
	DeleteQuestions             = `UPDATE questions SET deleted_at = $1 WHERE id = $2`
	UpdateQuestions             = `UPDATE questions SET status = $1, updated_at = $2 WHERE id = $3`

	// Schedule
	InsertSchedule       = `INSERT INTO schedules (user_id, date, start_time, end_time, documentation) VALUES ($1, $2, $3, $4, $5) RETURNING id,user_id, date, start_time, end_time, documentation`
	SelectScheduleByID   = `SELECT id, user_id, date, start_time, end_time, documentation FROM schedules WHERE id = $1`
	SelectScheduleByRole = `SELECT s.id, s.user_id, s.date, s.start_time, s.end_time, s.documentation FROM schedules s JOIN users u ON s.user_id = u.id WHERE u.role = $3 AND s.deleted_at IS NULL ORDER BY s.created_at DESC LIMIT $1 OFFSET $2`
	DeleteSchedule       = `UPDATE schedules SET deleted_at = $1 WHERE id = $2`
	UpdateScheduleDoc    = `UPDATE schedules SET documentation = $1, updated_at = $2 WHERE id = $3`

	// Attendance
	DeleteAttendance = `UPDATE attendances SET deleted_at = $1 WHERE id = $2`
)
