package hosts

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// This should roughly track the behaviour of frontends/php/include/classes/api/services/CHost.php

type Repository interface {
	Find(options ...func(*FindOptions) error) ([]Host, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return &repository{conn}
}

func (r *repository) Find(options ...func(*FindOptions) error) ([]Host, error) {
	var err error
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Zabbix Frontend filters by flags to exclude prototypes and rules(?)
	stmt := psql.Select("*").From("hosts") //.Where(sq.Eq{"flags": []int{ZbxFlagDiscoveryNormal, ZbxFlagDiscoveryCreated}})

	findOptions := &FindOptions{limit: 0}
	for _, op := range options {
		err := op(findOptions)
		if err != nil {
			return nil, err
		}
	}

	// stmt = stmt.Limit(findOptions.limit)

	query, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	var hosts []Host
	if err = r.db.Select(&hosts, query, args...); err != nil {
		return nil, err
	}

	return hosts, nil
}

type FindOptions struct {
	// Limit results returned to this number
	limit uint64

	// Include relationships specified by name
	include []string
}

func Limit(count uint64) func(*FindOptions) error {
	return func(opts *FindOptions) error {
		if count > 0 {
			opts.limit = count
		} else {
			opts.limit = 100
		}
		return nil
	}
}
