package camera

import "github.com/jmoiron/sqlx"

type Repository struct{
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB)*Repository{
	return &Repository{db:db}
}

func(r *Repository)Create(c *Camera)error{
	query:=`INSERT INTO cameras(id,name,rtsp_url,location,status) VALUES(uuid_generate_v4(),$1,$2,$3,'offline') RETURNING id,created_at`
	return r.db.QueryRowx(query,c.Name,c.RTSPUrl,c.Location).Scan(&c.ID, &c.CreatedAt)
}

func(r *Repository)GetAll()([]Camera,error){
	var c []Camera
	err:=r.db.Select(&c,"SELECT * FROM cameras")
	return c,err
}


func(r *Repository)GetByID(id string)(*Camera,error){
	var c Camera
	err := r.db.Get(&c,"SELECT * FROM cameras WHERE id=$1",id)

	if err!= nil{
		return nil,err
	}
	return &c,nil
}

func(r *Repository)Delete(id string)error{
	_,err:=r.db.Exec("DELETE FROM cameras WHERE id =$1",id)
	return err
}