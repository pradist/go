package main

import (
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cover struct {
	Id 		int
	Name 	string
}

// var db *sql.DB
var db *sqlx.DB

func main() {
	var err error
	// db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@localhost:1433?database=master")
	db, err = sqlx.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/mysql1")
	if err != nil {
		panic(err)
	}

	// cover := Cover{
	// 	Id: 8,
	// 	Name: "test1",
	// }

	// err = DeleteCover(8)
	// if err != nil {
	// 	panic(err)
	// }

	cover, err := GetCoverX(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cover);
	// for _, cover := range covers {
	// 	fmt.Println(cover);
	// }

	// cover, err := GetCover(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cover);
}

func GetCoversX() ([]Cover, error) {
	query := "select id, name from cover"
	covers := []Cover{}
	err := db.Select(&covers, query)

	if err != nil {
		return nil, err
	}
	return covers, nil
}

func GetCoverX(id int) (*Cover, error) {
	query := "select id, name from cover where id = ?"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}
	return &cover, err
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, cover)
	}
	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	// //MS SQL Server
	// query := "select id, name from cover where id = @id"
	// row := db.QueryRow(query, sql.Named("id", id))
	
	//MY SQL
	query := "select name from cover where id = ?"
	row := db.QueryRow(query, id)

	if err != nil {
		return nil, err
	}
	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}

func AddCover(cover Cover) error {
	
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "insert into cover (name) values (?)"
	result, err := db.Exec(query, cover.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected <= 0 {
		tx.Rollback()
		return errors.New("Can't insert")
	}
	
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateCover(cover Cover) error {	
	query := "update cover set name = ? where id = ?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("Can't Update")
	}
	return nil
}

func DeleteCover(id int) error {	
	query := "delete from cover where id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("Can't delete")
	}
	return nil
}