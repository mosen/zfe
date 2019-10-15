package hosts

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	db "github.com/mosen/zfe/pkg/database"
)

// This should roughly track the behaviour of frontends/php/include/classes/api/services/CHost.php

type HostsRepository interface {
	Find(options ...func(*db.FindOptions) error) ([]Host, error)
}

type hostsRepository struct {
	db *sqlx.DB
}

func NewHostsRepository(conn *sqlx.DB) HostsRepository {
	return &hostsRepository{conn}
}

func (r *hostsRepository) Find(options ...func(*db.FindOptions) error) ([]Host, error) {
	var err error
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Zabbix Frontend filters by flags to exclude prototypes and rules(?)
	stmt := psql.Select("*").From("hosts").Where(
		sq.Eq{"flags": []int{ZbxFlagDiscoveryNormal, ZbxFlagDiscoveryCreated}}).Where(
		sq.NotEq{"status": HostStatusTemplate})

	findOptions := &db.FindOptions{Limit: 100}
	for _, op := range options {
		err := op(findOptions)
		if err != nil {
			return nil, err
		}
	}

	stmt = stmt.Limit(findOptions.Limit)

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

type TemplatesRepository interface {
	Find(options ...func(*db.FindOptions) error) ([]Template, error)
}

type templatesRepository struct {
	db *sqlx.DB
}

func NewTemplatesRepository(conn *sqlx.DB) TemplatesRepository {
	return &templatesRepository{conn}
}

func (r *templatesRepository) Find(options ...func(*db.FindOptions) error) ([]Template, error) {
	var err error
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Zabbix Frontend filters by flags to exclude prototypes and rules(?)
	stmt := psql.Select("*").From("hosts").Where(
		sq.Eq{"status": HostStatusTemplate})

	findOptions := &db.FindOptions{Limit: 100}
	for _, op := range options {
		err := op(findOptions)
		if err != nil {
			return nil, err
		}
	}

	stmt = stmt.Limit(findOptions.Limit)

	query, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	var templates []Template
	if err = r.db.Select(&templates, query, args...); err != nil {
		return nil, err
	}

	return templates, nil
}
