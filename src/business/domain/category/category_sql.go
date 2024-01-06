package category

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/query"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

func (c *category) createSQLCategory(tx sql.CommandTx, v entity.CreateCategoryParam) (sql.CommandTx, entity.Category, error) {
	category := entity.Category{}

	res, err := tx.NamedExec("iCreateCategory", createCategory, v)
	if err != nil {
		return tx, category, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil || rowCount < 1 {
		return tx, category, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no rows affected")
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return tx, category, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	category.ID = lastID

	return tx, category, nil
}

func (c *category) getSQLCategory(ctx context.Context, params entity.CategoryParam) (entity.Category, error) {
	category := entity.Category{}

	key, err := c.json.Marshal(params)
	if err != nil {
		return category, nil
	}

	cachedEntry, err := c.getCache(ctx, fmt.Sprintf(getCategoryByIdKey, string(key)))
	switch {
	case errors.Is(err, redis.Nil):
		c.log.Info(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
	case err != nil:
		c.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	default:
		return cachedEntry, nil
	}

	qb := query.NewSQLQueryBuilder(c.db, "param", "db", &params.QueryOption)
	queryExt, queryArgs, _, _, err := qb.Build(&params)
	if err != nil {
		return category, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := c.db.Follower().QueryRow(ctx, "rCategoryByID", getCategory+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return category, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return category, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err := row.StructScan(&category); err != nil && !errors.Is(err, sql.ErrNotFound) {
		return category, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return category, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err = c.upsertCache(ctx, fmt.Sprintf(getCategoryByIdKey, string(key)), category, time.Minute); err != nil {
		c.log.Error(ctx, err)
	}

	return category, nil
}

func (c *category) getSQLCategoryList(ctx context.Context, params entity.CategoryParam) ([]entity.Category, *entity.Pagination, error) {
	categories := []entity.Category{}

	qb := query.NewSQLQueryBuilder(c.db, "param", "db", &params.QueryOption)
	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&params)
	if err != nil {
		return categories, nil, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := c.db.Follower().Query(ctx, "rListCategory", getCategory+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return categories, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		temp := entity.Category{}
		if err := rows.StructScan(&temp); err != nil {
			c.log.Error(ctx, errors.NewWithCode(codes.CodeSQLRowScan, err.Error()))
			continue
		}
		categories = append(categories, temp)
	}

	pg := entity.Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(len(categories)),
	}

	if len(categories) > 0 && !params.QueryOption.DisableLimit && params.IncludePagination {
		if err := c.db.Follower().Get(ctx, "cCategory", readCategoryCount+countExt, &pg.TotalElements, countArgs...); err != nil {
			return categories, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(params.Limit)

	return categories, &pg, nil
}

func (c *category) updateSQLCategory(ctx context.Context, updateParam entity.UpdateCategoryParam, selectParam entity.CategoryParam) error {
	c.log.Debug(ctx, fmt.Sprintf("update category by: %v", selectParam))

	qb := query.NewSQLQueryBuilder(c.db, "param", "db", &selectParam.QueryOption)

	var err error
	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &selectParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	_, err = c.db.Leader().Exec(ctx, "uCategory", updateCategory+queryUpdate, args...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("successfully updated category: %v", updateParam))

	if err := c.deleteCache(ctx); err != nil {
		c.log.Error(ctx, err)
	}

	return nil
}
