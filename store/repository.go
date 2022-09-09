package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/funasedaisuke/go-web-application/entity"
	- "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Beginner interface{
	BeginTx(ctx context.Context, opts *sql.TxOptions)(*sql.Tx,error)
}


type Preparer interface{
	Preparex(ctx context.Context,query string)(*sqlx.Stmt,error)
}

type Execer interface{
	ExecContext(ctx context.Context, query string,args ...any)
	NamedExecContext(ctx context.Context,query string,args interface{})(sql.Result,error)
}


type Queryer interface{
	Preparer
	QueryxContext(ctx context.Context, query string, args ...any)(*sqlx.Rows,error)
    QueryxRowx(ctx context.Context, query string, args ...any)*sqlx.Rows
    GetContext(ctx context.Context, dest interface{},query string ,args ...any)error
    SelectContext(ctx context.Context, dest interface{},query string ,args ...any)error
}

var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Queryer = (*sqlx.DB)(nil)
	_ Execer = (*sqlx.DB)(nil)
	_ Execer = (*sqlx.Tx)(nil)
)

type Repository struct{
	Clocker clock.Clocker
}


func New(ctx context.Context, cfg *config.Config)(*sqlx.DB,func(),error){
	db,err := sql.Open("mysql",
	fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser,cfg,DBPassword,
		cfg,DBHost,cfg.DBPort,
		cfg.DBName,
	),
)
if err != nil{
	return nil,nil,err
}
ctx,cancel := context.WithTimeout(ctx,2*time.Second)
defer cancel()
if err := db.PingContext(ctx); err != nil{
	return nil, func()(_ =db.Close()),err
}
xdb := sqlx.NewDb(db,"mysql")
return xdb, func(){_ = db.Close()},nil
}
