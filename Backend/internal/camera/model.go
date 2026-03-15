package camera

import "time"

type Camera struct{
	ID string	`db:"id" json:"id"`
	Name string	 `db:"name" json:"name"`
	RTSPUrl string  `db:"rtsp_url" json:"rtsp_url"`
	Location string  `db:"location" json:"location"`
	Status string  `db:"status" json:"status"`
	OrganizationID string    `json:"organization_id" db:"organization_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}

type CreateCameraRequest struct{
	Name string `json:"name"`
	RTSPUrl string `json:"rtsp_url"`
	Location string `json:"location"`
	OrganizationID string `json:"organization_id"`
}