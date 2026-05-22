package main

import (
	"context"
	"fmt"
	"log"
	"tidy/simple_sql"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	// 1. Укажите настройки подключения к вашей базе данных
	// Формат: postgres://логин:пароль@хост:порт/имя_базы
	connStr := "postgres://postgres:salamandra@localhost:5432/postgres"

	// 2. Устанавливаем реальное соединение
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	// Не забываем закрыть соединение при выходе из main
	defer conn.Close(ctx)

	// 3. Передаем РЕАЛЬНОЕ соединение conn (а не &pgx.Conn{})
	tasks, err := simple_sql.SelectRows(ctx, conn)
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

			// Передаем то же самое рабочее соединение в UpdateTask
			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}

	fmt.Println("Succeed")
}
