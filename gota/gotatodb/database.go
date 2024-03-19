package gotadb

import (
	"context"
	"errors"
	"fmt"
	slog "framework-iris/gota/commons/log"
	serveries "framework-iris/gota/config"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type gotaDB struct {
	db     *gorm.DB
	name   string //db name
	dbType string //db 类型 mysql sqlite
}

func (slf *gotaDB) GormDB() *gorm.DB {
	return slf.db
}

func (slf *gotaDB) Name() string {
	return slf.name
}
func (slf *gotaDB) StartPgsql(dbConfig serveries.DataBaseConfig) (err error) {
	if slf.db != nil {
		return errors.New("db already open")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbConfig.Addr, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port, dbConfig.Loc)
	if dbConfig.Silent == true {
		slf.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, Logger: nil})
		if err != nil {
			slog.Slog.InfoF(context.Background(), "conn database error %s", err)
			return err
		}
	} else {
		slf.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, Logger: &slog.Gorm})
		if err != nil {
			slog.Slog.InfoF(context.Background(), "conn database error %s", err)
			return err
		}
	}
	slf.name = dbConfig.Name
	db, err := slf.db.DB()
	if err != nil {
		slog.Slog.InfoF(context.Background(), "conn slf.db.DB() error %s", err)
		return err
	}
	db.SetConnMaxLifetime(dbConfig.MaxLifeTime * time.Millisecond)
	db.SetConnMaxIdleTime(dbConfig.MaxIdleTime * time.Millisecond)
	db.SetMaxOpenConns(dbConfig.MaxConn)
	db.SetMaxIdleConns(dbConfig.IdleConn)
	return nil
}
func (slf *gotaDB) StartSqlite(dbConfig serveries.DataBaseConfig) error {
	if slf.db != nil {
		return errors.New("db already open")
	}
	slf.name = dbConfig.Name
	var err error
	if dbConfig.Silent == true {
		slf.db, err = gorm.Open(sqlite.Open(dbConfig.DBFilePath), &gorm.Config{PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}, Logger: nil})
		if err != nil {
			slog.Slog.InfoF(context.Background(), "conn database error %s", err)
			return err
		}
	} else {
		slf.db, err = gorm.Open(sqlite.Open(dbConfig.DBFilePath), &gorm.Config{PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}, Logger: &slog.Gorm})
		if err != nil {
			slog.Slog.InfoF(context.Background(), "conn database error %s", err)
			return err
		}
	}
	return nil
}
func (slf *gotaDB) StartMysql(dbConfig serveries.DataBaseConfig) error {
	if slf.db != nil {
		return errors.New("db already open")
	}
	slf.name = dbConfig.Name
	var err error

	if len(dbConfig.Loc) <= 0 {
		dbConfig.Loc = "Local"
	}

	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Addr,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.Charset,
		dbConfig.Loc,
	)
	if dbConfig.Silent == true {
		slf.db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}, Logger: nil})
		if err != nil {
			slog.Slog.InfoF(context.Background(), "conn database error %s", err)
			return err
		}
	} else {
		slf.db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}, Logger: &slog.Gorm})
		if err != nil {
			slog.Slog.InfoF(context.Background(), "conn database error %s", err)
			return err
		}
	}

	db, err := slf.db.DB()
	if err != nil {
		slog.Slog.InfoF(context.Background(), "conn slf.db.DB() error %s", err)
		return err
	}
	db.SetConnMaxLifetime(dbConfig.MaxLifeTime * time.Millisecond)
	db.SetConnMaxIdleTime(dbConfig.MaxIdleTime * time.Millisecond)
	db.SetMaxOpenConns(dbConfig.MaxConn)
	db.SetMaxIdleConns(dbConfig.IdleConn)
	return nil
}

func (slf *gotaDB) StopDb() error {
	if slf.db != nil {
		db, err := slf.db.DB()
		if err != nil {
			slf.db = nil
		} else {
			_ = db.Close()
		}
		return err
	} else {
		return errors.New("db is nil")
	}
}
