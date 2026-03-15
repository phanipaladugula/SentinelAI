package camera

import (
	"github.com/jmoiron/sqlx"
	"context"
	"time"
)

type Repository struct{
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB)*Repository{
	return &Repository{db:db}
}

func(r *Repository)Create(ctx context.Context,c *Camera)error{
	query:=`INSERT INTO cameras(id,name,rtsp_url,location,status) VALUES(uuid_generate_v4(),$1,$2,$3,'offline') RETURNING id,created_at`
	return r.db.QueryRowxContext(ctx,query,c.Name,c.RTSPUrl,c.Location).Scan(&c.ID, &c.CreatedAt)
}

func(r *Repository)GetAll(ctx context.Context)([]Camera,error){
	var c []Camera
	err:=r.db.SelectContext(ctx,&c,"SELECT * FROM cameras")
	return c,err
}


func(r *Repository)GetByID(ctx context.Context,id string)(*Camera,error){
	var c Camera
	err := r.db.GetContext(ctx,&c,"SELECT * FROM cameras WHERE id=$1",id)

	if err!= nil{
		return nil,err
	}
	return &c,nil
}

func(r *Repository)Delete(id string)error{
	_,err:=r.db.Exec("DELETE FROM cameras WHERE id =$1",id)
	return err
}

func (r *Repository)UpdateStatus(ctx context.Context,id string,status string)error{
	query := `UPDATE cameras SET status = $1 WHERE ID = $2`
	 _,err:=r.db.ExecContext(ctx,query,status,id)
	 return err
}