package main

import (
	"context"
	"fmt"
	"log"

	"go-programing-essence/Chapter12/12-2/ent"
	"go-programing-essence/Chapter12/12-2/ent/todo"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// スキーマ作成（マイグレーション）
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Todoを作成
	_, err = client.Todo.Create().SetText("Go言語プログラミングエッセンス").Save(context.Background())
	if err != nil {
		log.Fatalf("failed creating a todo: %V", err)
	}

	_, err = client.Todo.Create().SetText("実用Go言語").Save(context.Background())
	if err != nil {
		log.Fatalf("failed creating a todo: %V", err)
	}

	// Todoを列挙（全権選択）
	for _, e := range client.Todo.Query().AllX(context.Background()) {
		fmt.Println(e.Text)
		fmt.Println(e.ID)
	}

	// Todoを列挙（idで選択）
	a, err := client.Todo.Query().Where(todo.ID(2)).Only(context.Background())
	if err != nil {
		fmt.Printf("failed querying todo: %v", err)
		return
	}
	fmt.Println(a.Text)

}
