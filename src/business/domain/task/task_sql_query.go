package task

const (
	createTask = `
	INSERT INTO task (title, fk_user_id, title, priority, task_status, periodic, due_time, status, created_by)
		VALUES(:title, :fk_user_id, :title, :priority, :task_status, :periodic, :due_time, :status, :created_by)`

	getTask = `
	SELECT
		id,
		fk_user_id,
		title,
		priority,
		task_status,
		periodic,
		due_time,
		status,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM
		task`

	updateTask = `
	UPDATE
		task`

	readTaskCount = `
	SELECT
		COUNT(*)
	FROM
		task`
)
