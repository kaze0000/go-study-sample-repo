package article

import (
	"database/sql"
	"fmt"
)

type Article struct {
	id int
	title string
}

func ReadAll(db *sql.DB) {
	var articles []Article
	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		article  := Article{}
		err := rows.Scan(&article.id, &article.title)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}
	rows.Close()
	fmt.Println(articles)
}
