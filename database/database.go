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
		return err
	}

	DB = db
	return nil
}

func Insert(tableName string, data interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}
	fmt.Println(data)
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a struct")
	}

	dataType := dataValue.Type()
	columns := make([]string, 0, dataType.NumField())
	values := make([]interface{}, 0, dataType.NumField())
	placeholders := make([]string, 0, dataType.NumField())

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			tagParts := strings.Split(dbTag, ",")
			columns = append(columns, tagParts[0])
		} else {
			columns = append(columns, field.Name)
		}
		values = append(values, dataValue.Field(i).Interface())
		placeholders = append(placeholders, "?")
	}
	if err := CreateTable(tableName, data); err != nil {
		panic(err)
	}
	fmt.Print(columns)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err = DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func ReadAll(tableName string, result interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}

	CreateTable(tableName, result)

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := DB.Query(query)
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

func Read(tableName string, conditions map[string]interface{}, result interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}

	whereClause, whereValues := buildWhereClause(conditions)
	CreateTable(tableName, result)
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

	whereClause, whereValues := buildWhereClause(conditions)

	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, whereClause)
	_, err = DB.Exec(query, whereValues...)
	if err != nil {
		return err
	}
	return nil
}

func CreateTable(tableName string, data interface{}) error {
	err := ConnectDatabase()
	if err != nil {
		return err
	}

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a struct")
	}

	dataType := dataValue.Type()

	// Check if the table exists
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// Table doesn't exist, create it
		columns := make([]string, 0, dataType.NumField())
		primaryKeyFound := false
		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)
			columnType := getColumnType(field.Type)
			columnName := field.Name
			if tag := field.Tag.Get("db"); tag != "" {
				columnName = strings.Split(tag, ",")[0]
			}
			if tag := field.Tag.Get("primarykey"); tag != "" {
				if strings.Contains(tag, "autoincrement") {
					columns = append(columns, fmt.Sprintf("%s %s PRIMARY KEY AUTOINCREMENT", columnName, columnType))
				} else {
					columns = append(columns, fmt.Sprintf("%s %s PRIMARY KEY", columnName, columnType))
				}
				primaryKeyFound = true
			} else if tag := field.Tag.Get("foreignkey"); tag != "" {
				columns = append(columns, fmt.Sprintf("%s %s", columnName, columnType))
			} else {
				columns = append(columns, fmt.Sprintf("%s %s", columnName, columnType))
			}
		}

		if !primaryKeyFound {
			return fmt.Errorf("a primary key must be defined for the table")
		}

		query := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, strings.Join(columns, ", "))
		_, err = DB.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func getColumnType(fieldType reflect.Type) string {
	switch fieldType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "INTEGER"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.String:
		return "TEXT"
	case reflect.Bool:
		return "BOOLEAN"
	default:
		return "TEXT"
	}
}
