package main

import (
	"database/sql"
	"fmt"
	"time"
)

type DbManager struct {
	Config *Config
	Db     *sql.DB
}

func NewDbManager() *DbManager {
	return &DbManager{}
}

func (d *DbManager) Init() error {
	var err error
	d.Db, err = sql.Open("mysql", fmt.Sprintln("%s:%s@/%s", d.Config.DbUser, d.Config.DbPassword, d.Config.DBName))
	if err != nil {
		return err
	}
	// See "Important settings" section.
	d.Db.SetConnMaxLifetime(time.Minute * 3)
	d.Db.SetMaxOpenConns(10)
	d.Db.SetMaxIdleConns(10)
	return nil
}
