package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./cisl-database.db")
	if err != nil {
		fmt.Println(err)
		return err
	}
	DB = db
	return nil
}

func Insert(tableName string, data map[string]interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}
	defer DB.Close()

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
	placeholders := make([]string, 0, len(data))
	for k, v := range data {
		columns = append(columns, k)
		values = append(values, v)
		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err = DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func Read(tableName string, conditions map[string]interface{}, result interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}
	defer DB.Close()

	whereClause, whereValues := buildWhereClause(conditions)

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, whereClause)
	rows, err := DB.Query(query, whereValues...)
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result argument must be a pointer to a slice")
	}

	sliceValue := resultValue.Elem()
	elementType := sliceValue.Type().Elem()

	for rows.Next() {
		newElement := reflect.New(elementType).Interface()
		rowPtr := make([]interface{}, len(cols))
		for i := range rowPtr {
			rowPtr[i] = reflect.ValueOf(newElement).Elem().Field(i).Addr().Interface()
		}

		err := rows.Scan(rowPtr...)
		if err != nil {
			return err
		}

		sliceValue.Set(reflect.Append(sliceValue, reflect.ValueOf(newElement).Elem()))
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func Update(tableName string, data map[string]interface{}, conditions map[string]interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}
	defer DB.Close()

	setClause, setValues := buildSetClause(data)
	whereClause, whereValues := buildWhereClause(conditions)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, setClause, whereClause)
	values := append(setValues, whereValues...)
	_, err = DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func buildWhereClause(conditions map[string]interface{}) (string, []interface{}) {
	whereClause := make([]string, 0, len(conditions))
	values := make([]interface{}, 0, len(conditions))
	for k, v := range conditions {
		whereClause = append(whereClause, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}
	return strings.Join(whereClause, " AND "), values
}

func buildSetClause(data map[string]interface{}) (string, []interface{}) {
	setClause := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
	for k, v := range data {
		setClause = append(setClause, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}
	return strings.Join(setClause, ", "), values
}

func Delete(tableName string, conditions map[string]interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}
	defer DB.Close()

	whereClause, whereValues := buildWhereClause(conditions)

	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, whereClause)
	_, err = DB.Exec(query, whereValues...)
	if err != nil {
		return err
	}
	return nil
}
