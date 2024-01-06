package task

const (
	createTask = `INSERT INTO task (fk_user_id, fk_category_id, title, priority, task_status, periodic, due_time, created_by)
	VALUES (:fk_user_id, :fk_category_id, :title, :priority, :task_status, :periodic, :due_time, :created_by)`

	getTask = `
		SELECT
			id,
			fk_user_id,
			fk_category_id,
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
