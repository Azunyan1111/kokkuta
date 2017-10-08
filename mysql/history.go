package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Kokkuta struct {
	ID   int64
	Body string
	Time string
	Good int64
}

func SetHistory(s string) {
	db, err := sql.Open("mysql", "root:@/kokkuta")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // 関数がリターンする直前に呼び出される
	_, err = db.Query(fmt.Sprintf("insert into history (body) VALUES ('%s');", s))
	if err != nil {
		fmt.Println(err)
	}
}

func AddGood(id string) {
	db, err := sql.Open("mysql", "root:@/kokkuta")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // 関数がリターンする直前に呼び出される
	_, err = db.Query(fmt.Sprintf("update history set good = good + 1 WHERE id=%s", id))
	if err != nil {
		fmt.Println(err)
	}
}

func GetHistory() []Kokkuta {
	db, err := sql.Open("mysql", "root:@/kokkuta")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM history ORDER BY id DESC LIMIT %d", 5))
	if err != nil {
		fmt.Println(err)
	}

	history := make([]Kokkuta, 0)

	for rows.Next() {
		var id_ int64
		var body_ string
		var time_ string
		var good_ int64
		if err := rows.Scan(&id_, &body_, &time_, &good_); err != nil {
			fmt.Println(err)
		}

		var historyData Kokkuta
		historyData.ID = id_
		historyData.Body = body_
		historyData.Time = time_
		historyData.Good = good_
		history = append(history, historyData)
	}
	return history
}
