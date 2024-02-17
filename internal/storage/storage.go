package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GZ91/MyBlog/internal/app/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

type Storage struct {
	logger *zap.Logger
	config *config.Config
	db     *sql.DB
}

func New(logger *zap.Logger, config *config.Config) *Storage {
	return &Storage{
		logger: logger,
		config: config,
	}
}

func (s *Storage) Up(ctx context.Context) error {
	lineConnect := s.config.GetLineDB()
	if lineConnect == "" {
		return errors.New("not line connect")
	}
	db, err := sql.Open("pgx", lineConnect)
	if err != nil {
		return err
	}
	s.db = db
	err = s.createTable(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) createTable(ctx context.Context) error {
	con, err := s.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer con.Close()
	err = s.createTableUsers(con, ctx)
	if err != nil {
		return err
	}
	err = s.createTableArticles(con, ctx)
	if err != nil {
		return err
	}
	err = s.createTableInputFixation(con, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) createTableUsers(con *sql.Conn, ctx context.Context) error {
	_, err := con.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users 
	(
		id serial PRIMARY KEY,
		userID VARCHAR(45),
		login VARCHAR(250) NOT NULL,
		password VARCHAR(250)  NOT NULL
	);`)

	return err
}

func (s *Storage) createTableArticles(con *sql.Conn, ctx context.Context) error {
	_, err := con.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS articles 
	(
		id serial PRIMARY KEY,
		time timestamp without time zone DEFAULT CURRENT_TIMESTAMP, 
		name VARCHAR(250) NOT NULL,
		text TEXT NOT NULL
	);`)
	return err
}

func (s *Storage) createTableInputFixation(con *sql.Conn, ctx context.Context) error {
	_, err := con.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS input_fixation 
	(
		id serial PRIMARY KEY,
		time timestamp without time zone DEFAULT CURRENT_TIMESTAMP, 
		login VARCHAR(250) NOT NULL
	);`)
	return err
}
