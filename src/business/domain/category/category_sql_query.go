package category

const (
	createCategory = `
	INSERT INTO category (name, created_by)
	    VALUES (:name, :created_by)`

	getCategory = `
		SELECT
			id,
			name,
			status,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			category`

	updateCategory = `
	UPDATE
		category`

	readCategoryCount = `
		SELECT
			COUNT(*)
		FROM
			category`
)
