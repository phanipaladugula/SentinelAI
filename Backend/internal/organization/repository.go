package organization

import(
	"github.com/jmoiron/sqlx"
)

type Repository struct{
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB)*Repository{
	return &Repository{db:db}
}

func (r *Repository) Create(org *Organization) error {
	query := `
	INSERT INTO organizations (name)
	VALUES ($1)
	RETURNING id, created_at
	`

	return r.db.QueryRowx(query, org.Name).
		Scan(&org.ID, &org.CreatedAt)
}

func (r *Repository) GetAll() ([]Organization, error) {
	var orgs []Organization

	query := `SELECT * FROM organizations`

	err := r.db.Select(&orgs, query)

	return orgs, err
}


func (r *Repository) GetByID(id string) (*Organization, error) {
	var org Organization

	query := `SELECT * FROM organizations WHERE id=$1`

	err := r.db.Get(&org, query, id)

	return &org, err
}


func (r *Repository) Delete(id string) error {
	query := `DELETE FROM organizations WHERE id=$1`

	_, err := r.db.Exec(query, id)

	return err
}