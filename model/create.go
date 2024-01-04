package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODOS VALUES(? , ?)", name, todo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer insertQ.Close()
	fmt.Println("Value has been inserted into the table successfully")
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("delete from todos where name=?", name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer insertQ.Close()
	fmt.Println("Value has been inserted into the table successfully")
	return nil
}
