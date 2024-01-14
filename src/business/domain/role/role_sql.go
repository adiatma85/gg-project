package role

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

func (r *role) createSQLRole(tx sql.CommandTx, v entity.CreateRoleParam) (sql.CommandTx, entity.Role, error) {
	role := entity.Role{}

	res, err := tx.NamedExec("iCreateRole", createRole, v)
	if err != nil {
		return tx, role, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil || rowCount < 1 {
		return tx, role, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no rows affected")
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return tx, role, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	role.ID = lastID

	return tx, role, nil
}

func (r *role) getSQLRole(ctx context.Context, params entity.RoleParam) (entity.Role, error) {
	result := entity.Role{}

	key, err := r.json.Marshal(params)
	if err != nil {
		return result, nil
	}

	cachedEntry, err := r.getCache(ctx, fmt.Sprintf(getRoleByIdKey, string(key)))
	switch {
	case errors.Is(err, redis.Nil):
		r.log.Info(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
	case err != nil:
		r.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	default:
		return cachedEntry, nil
	}

	qb := query.NewSQLQueryBuilder(r.db, "param", "db", &params.QueryOption)
	queryExt, queryArgs, _, _, err := qb.Build(&params)
	if err != nil {
		return result, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := r.db.Follower().QueryRow(ctx, "rRoleByID", getRole+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return result, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return result, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err := row.StructScan(&result); err != nil && !errors.Is(err, sql.ErrNotFound) {
		return result, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return result, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err = r.upsertCache(ctx, fmt.Sprintf(getRoleByIdKey, string(key)), result, time.Minute); err != nil {
		r.log.Error(ctx, err)
	}

	return result, nil
}

func (r *role) getSQLRoleList(ctx context.Context, params entity.RoleParam) ([]entity.Role, *entity.Pagination, error) {
	results := []entity.Role{}

	qb := query.NewSQLQueryBuilder(r.db, "param", "db", &params.QueryOption)
	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&params)
	if err != nil {
		return results, nil, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := r.db.Follower().Query(ctx, "rListRole", getRole+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return results, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		temp := entity.Role{}
		if err := rows.StructScan(&temp); err != nil {
			r.log.Error(ctx, errors.NewWithCode(codes.CodeSQLRowScan, err.Error()))
			continue
		}
		results = append(results, temp)
	}

	pg := entity.Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(len(results)),
	}

	if len(results) > 0 && !params.QueryOption.DisableLimit && params.IncludePagination {
		if err := r.db.Follower().Get(ctx, "cRole", readRoleCount+countExt, &pg.TotalElements, countArgs...); err != nil {
			return results, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(params.Limit)

	return results, &pg, nil
}

func (r *role) updateSQLRole(ctx context.Context, updateParam entity.UpdateRoleParam, selectParam entity.RoleParam) error {
	r.log.Debug(ctx, fmt.Sprintf("update role by: %v", selectParam))

	qb := query.NewSQLQueryBuilder(r.db, "param", "db", &selectParam.QueryOption)

	var err error
	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &selectParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	_, err = r.db.Leader().Exec(ctx, "uRole", updateRole+queryUpdate, args...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	r.log.Debug(ctx, fmt.Sprintf("successfully updated role: %v", updateParam))

	if err := r.deleteCache(ctx); err != nil {
		r.log.Error(ctx, err)
	}

	return nil
}
