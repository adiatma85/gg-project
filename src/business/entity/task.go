package entity

import (
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
)

const (
	// Task statuses
	TaskStatusTodo    = "todo"
	TaskStatusOnGoing = "ongoing"
	TaskStatusDone    = "done"

	// Task Periodic Num
	TaskPeriodicNone    = "none"
	TaskPeriodicDaily   = "daily"
	TaskPeriodicWeekly  = "weekly"
	TaskPeriodicMonthly = "monthly"
	TaskPeriodicYearly  = "yearly"
)

type Task struct {
	ID         int64       `db:"id" json:"id"`
	UserId     int64       `db:"fk_user_id" json:"userId"`
	CategoryID int64       `db:"fk_category_id" json:"categoryId"`
	Title      string      `db:"title" json:"title"`
	Priority   int64       `db:"priority" json:"priority"`
	TaskStatus string      `db:"task_status" json:"taskStatus"`
	Periodic   string      `db:"periodic" json:"periodic"`
	DueTime    null.Time   `db:"due_time" json:"dueTime"`
	Status     int64       `db:"status" json:"status" swaggertype:"integer"`
	CreatedAt  null.Time   `db:"created_at" json:"createdAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CreatedBy  null.String `db:"created_by" json:"createdBy" swaggertype:"string"`
	UpdatedAt  null.Time   `db:"updated_at" json:"updatedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy  null.String `db:"updated_by" json:"updatedBy" swaggertype:"string"`
	DeletedAt  null.Time   `db:"deleted_at" json:"deletedAt,omitempty" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy  null.String `db:"deleted_by" json:"deletedBy,omitempty" swaggertype:"string"`
}

type TaskParam struct {
	ID         null.Int64  `param:"id" uri:"task_id" db:"id" form:"task_id"`
	IDs        []int64     `param:"ids" uri:"task_ids" db:"id"`
	UserId     null.Int64  `param:"fk_user_id" uri:"user_id" db:"fk_user_id"`
	CategoryID null.Int64  `param:"fk_category_id" uri:"category_id" db:"fk_category_id"`
	Title      null.String `param:"title" db:"title"`
	Priority   null.Int64  `param:"priority" db:"priority"`
	TaskStatus null.String `param:"task_status" db:"task_status"`
	Periodic   null.String `param:"periodic" db:"periodic"`
	DueTime    null.Time   `param:"due_time" db:"due_time"`
	Status     null.Int64  `param:"status" db:"status" swaggertype:"string"`
	PaginationParam
	QueryOption query.Option
}

type CreateTaskParam struct {
	UserId     int64       `db:"fk_user_id" json:"-"`
	CategoryID int64       `db:"fk_category_id" json:"categoryId"`
	Title      string      `db:"title" json:"title"`
	Priority   int64       `db:"priority" json:"priority"`
	TaskStatus string      `db:"task_status" json:"taskStatus"`
	Periodic   string      `db:"periodic" json:"periodic"`
	DueTime    null.Time   `db:"due_time" json:"due_time"`
	CreatedBy  null.String `json:"-" db:"created_by" swaggertype:"string"`
}

type UpdateTaskParam struct {
	UserId     null.Int64  `param:"fk_user_id" db:"fk_user_id" json:"userId"`
	CategoryID null.Int64  `param:"fk_category_id" db:"fk_category_id" json:"categoryId"`
	Title      string      `param:"title" db:"title" json:"title"`
	Priority   int64       `param:"priority" db:"priority" json:"priority"`
	TaskStatus string      `param:"task_status" db:"task_status" json:"taskStatus"`
	Periodic   null.String `db:"periodic" param:"periodic" json:"periodic"`
	DueTime    null.Time   `db:"due_time" json:"dueTime" param:"due_time"`
	Status     null.Int64  `db:"status" param:"status" json:"status" swaggertype:"string"`
	UpdatedAt  null.Time   `db:"updated_at" json:"-" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy  null.String `db:"updated_by" json:"-" swaggertype:"string"`
	DeletedAt  null.Time   `db:"deleted_at" json:"-" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy  null.String `db:"deleted_by" json:"-" swaggertype:"string"`
}
