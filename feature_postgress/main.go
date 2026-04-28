package main

import (
	"context"
	"fmt"
	"tidy/simple_sql"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	tasks, err := simple_sql.SelectRows(ctx, &pgx.Conn{})
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Покормить кошку"
			task.Description = "Отсыпать кошке 30 грамм корма"
			task.Completed = true
			now := time.Now()
			task.ComletedAt = &now

			if err := simple_sql.UpdateTask(ctx, &pgx.Conn{}, task); err != nil {
				panic(err)
			}
			break
		}
	}

	fmt.Println("Succeed")
}
