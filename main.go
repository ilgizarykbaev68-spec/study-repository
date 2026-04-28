package main

import (
	"context"
	"fmt"
	simpleconnection "study/feature_postgress/simple_connection"
	"study/feature_postgress/simple_sql"
	//"time"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	//чтобы добавлять задачи
	// if err := simple_sql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Обед2",
	// 	"Покушац надо",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.SelectRows(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("succeed!")

}
