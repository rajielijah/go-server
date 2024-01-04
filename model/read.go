package model

import "github.com/rajielijah/go-server/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("select * from Todos")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
func ReadbyName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("select * from Todos where name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
