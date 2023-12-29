package entity

import (
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
)

type Category struct {
	ID        int64       `db:"id" json:"id"`
	Name      string      `db:"name" json:"name"`
	Status    null.Int64  `db:"status" json:"status" swaggertype:"integer"`
	CreatedAt null.Time   `db:"created_at" json:"createdAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CreatedBy null.String `db:"created_by" json:"createdBy" swaggertype:"string"`
	UpdatedAt null.Time   `db:"updated_at" json:"updatedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy null.String `db:"updated_by" json:"updatedBy" swaggertype:"string"`
	DeletedAt null.Time   `db:"deleted_at" json:"deletedAt,omitempty" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy null.String `db:"deleted_by" json:"deletedBy,omitempty" swaggertype:"string"`
}

type CategoryParam struct {
	ID   null.Int64  `param:"id" uri:"category_id" db:"id" form:"id"`
	IDs  []int64     `param:"ids" uri:"category_ids" db:"id" form:"categoryIds"`
	Name null.String `param:"name" db:"name"`
	PaginationParam
	QueryOption query.Option
}

type CreateCategoryParam struct {
	Name      string      `db:"name" json:"name"`
	CreatedBy null.String `json:"-" db:"created_by" swaggertype:"string"`
}

type UpdateCategoryParam struct {
	Name      string      `param:"name" db:"name" json:"name"`
	Status    null.Int64  `param:"status" db:"status" json:"-" swaggertype:"integer"`
	UpdatedAt null.Time   `param:"updated_at" db:"updated_at" json:"-" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy null.String `param:"updated_by" db:"updated_by" json:"-" swaggertype:"string"`
	DeletedAt null.Time   `param:"deleted_at" db:"deleted_at" json:"-" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy null.String `param:"deleted_by" db:"deleted_by" json:"-" swaggertype:"string"`
}
