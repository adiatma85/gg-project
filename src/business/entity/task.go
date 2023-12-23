package entity

import (
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
)

type Task struct {
	ID         int64     `db:"id" json:"id"`
	FkUserId   int64     `db:"fk_user_id" json:"userId"`
	Title      string    `db:"title" json:"title"`
	Priority   int64     `db:"priority" json:"priority"`
	TaskStatus string    `db:"task_status" json:"taskStatus"`
	Periodic   string    `db:"periodic" json:"periodic"`
	DueTime    null.Time `db:"due_time" json:"dueTime"`
	Status     int64     `db:"status" json:"status"`
	CreatedAt  null.Time `db:"created_at" json:"createdAt"`
	CreatedBy  string    `db:"created_by" json:"createdBy"`
	UpdatedAt  null.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy  string    `db:"updated_by" json:"updatedBy"`
	DeletedAt  null.Time `db:"deleted_at" json:"deletedAt"`
	DeletedBy  string    `db:"deleted_by" json:"deletedBy"`
}

type TaskParam struct {
	ID         null.Int64  `param:"id" uri:"task_id" db:"id" form:"id"`
	IDs        []int64     `param:"ids" uri:"task_ids" db:"id"`
	FkUserId   null.Int64  `param:"fk_user_id" uri:"user_id" db:"fk_user_id"`
	Title      null.String `param:"title" db:"title"`
	Priority   null.Int64  `param:"priority" db:"priority"`
	TaskStatus null.String `param:"task_status" db:"task_status"`
	Periodic   null.String `param:"periodic" db:"periodic"`
	DueTime    null.Time   `param:"due_time" db:"due_time"`
	Status     null.Int64  `param:"status" db:"status"`
	PaginationParam
	QueryOption query.Option
}

type CreateTaskParam struct {
	Title      string      `db:"title" json:"title"`
	Priority   int64       `db:"priority" json:"priority"`
	TaskStatus string      `db:"task_status" json:"taskStatus"`
	Periodic   null.String `db:"periodic" json:"periodic"`
	DueTime    null.Time   `db:"due_time" json:"due_time"`
	Status     null.Int64  `db:"status" json:"status"`
}

type UpdateTaskParam struct {
	Title      string      `param:"title" db:"title" json:"title"`
	Priority   int64       `param:"priority" db:"priority" json:"priority"`
	TaskStatus string      `param:"task_status" db:"task_status" json:"taskStatus"`
	Periodic   null.String `db:"periodic" param:"periodic" json:"periodic"`
	DueTime    null.Time   `db:"due_time" json:"dueTime" param: "due_time"`
	Status     null.Int64  `db:"status" param:"status" json:"status"`
}
