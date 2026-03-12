package camera

type Service struct{
	repo *Repository
}

func NewService(r *Repository)*Service{
	return &Service{repo,r}
}

func(s *Service)Create(req CreateCameraRequest)(*Camera,error){
	c:=&Camera{
		Name: req.Name,
		RTSPUrl: req.RTSPUrl,
		Location: req.Location,
	}

	err:=s.repo.Create(c)
	if err!=nil{
		return nil,err
	}

	return c,nil
}

func(s *Service)GetAll()([]Camera,error){
	return s.repo.GetAll()
}

func(s *Service)GetByID(id string)(*Camera,error){
	return s.repo.GetByID()
}

func(s *Service)Delete(id string)error{
	return s.repo.Delete(id)
}