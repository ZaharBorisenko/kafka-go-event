package main

import (
	"analytics-service/internal/repository"
	"context"
	"log"
)

func main() {
	conn, err := repository.ClickHouseConnect()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "SELECT name, toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, uuid string
		if err := rows.Scan(&name, &uuid); err != nil {
			log.Fatal(err)
		}
		log.Printf("name: %s, uuid: %s", name, uuid)
	}

	// NOTE: Do not skip rows.Err() check
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
