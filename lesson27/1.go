package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "classwork"
	password = "1234"
)

type Football struct {
	ID          string
	Name        string
	PlayerNames string
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`insert into football_club(id, name) values
	($1, $2),
	($3,$4),
	($5,$6),
	($7,$8)
	`,
		uuid.NewString(), "Barcelona",
		uuid.NewString(), "Real Madrid",
		uuid.NewString(), "Man city",
		uuid.NewString(), "Bayern",
	)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`insert into football_player(id, club_id,name) values
	($1, $2, $3),
	($4, $5, $6),
	($7, $8, $9),
	($10, $11, $12),
	($13, $14, $15),
	($16, $17, $18),
	($19, $20, $21),
	($22, $23, $24),
	($25, $26, $27),
	($28, $29, $30)
	`,
		uuid.NewString(), "c61ad80d-699c-4c4f-965f-1c71c25cbf67", "Messi",
		uuid.NewString(), "87b5fbdc-eb45-4d0e-8ebf-ec0d9e16c6ee", "Ronaldo",
		uuid.NewString(), "c61ad80d-699c-4c4f-965f-1c71c25cbf67", "Neymar",
		uuid.NewString(), "d59518ad-c84d-4d42-b303-4cb7497a4a12", "Sane",
		uuid.NewString(), "145342bc-1a6b-4f03-80dc-849e747b37d7", "Aguero",
		uuid.NewString(), "c61ad80d-699c-4c4f-965f-1c71c25cbf67", "Xavi",
		uuid.NewString(), "87b5fbdc-eb45-4d0e-8ebf-ec0d9e16c6ee", "Ramos",
		uuid.NewString(), "87b5fbdc-eb45-4d0e-8ebf-ec0d9e16c6ee", "Raul",
		uuid.NewString(), "d59518ad-c84d-4d42-b303-4cb7497a4a12", "Muller",
		uuid.NewString(), "145342bc-1a6b-4f03-80dc-849e747b37d7", "Haland",
	)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`update football_club
	set count_players=(SELECT COUNT(p.name)
		FROM football_player AS p
		WHERE p.club_id = football_club.id
		GROUP BY p.club_id)
	`)

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(`SELECT c.id,c.name,p.name as player_name
	FROM football_club as c 
	join football_player as p
	on p.club_id=c.id
	`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var football Football

		err := rows.Scan(&football.ID, &football.Name, &football.PlayerNames)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		fmt.Printf("id: %s\nCLub: %s\nPlayer name: %s\n\n", football.ID, football.Name, football.PlayerNames)
	}

	fmt.Println("Successfully connected!")
}

