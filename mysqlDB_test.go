package godb

import "testing"

type User struct {
	Id     int
	Name   string
	Age    int
	Gender string
}

func TestInsert(t *testing.T) {
	r, err := MysqlDB.Exec("INSERT INTO user (name, age, gender) VALUES (?, ?, ?)", "test1", 28, "W")
	if err != nil {
		t.Fatal(err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestUpdate(t *testing.T) {
	r, err := MysqlDB.Exec("UPDATE user SET age = ? WHERE id = ?", 18, 2)
	if err != nil {
		t.Fatal(err)
	}
	row, err := r.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(row)
}

func TestDelete(t *testing.T) {
	r, err := MysqlDB.Exec("DELETE FROM user WHERE id = ?", 3)
	if err != nil {
		t.Fatal(err)
	}
	row, err := r.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(row)
}

func TestSelect(t *testing.T) {
	rows, err := MysqlDB.Query("SELECT * FROM user")
	if err != nil {
		t.Fatal(err)
	}
	user := &User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Gender)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(user)
	}

}
