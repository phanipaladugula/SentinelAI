package organization

import "time"

type Organization struct{
	ID string `db:"id" json:"id"`
	Name string `db:"name" json"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

