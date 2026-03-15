package camera

import (
    "context"
    "errors"
    "strings"
)

type Service struct {
    repo *Repository
}

func NewService(r *Repository) *Service {
    return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, req CreateCameraRequest) (*Camera, error) {
    if strings.TrimSpace(req.RTSPUrl) == "" {
        return nil, errors.New("rtsp_url is required")
    }

    c := &Camera{
        Name:     req.Name,
        RTSPUrl:  req.RTSPUrl,
        Location: req.Location,
        OrganizationID: req.OrganizationID,
    }

    // Explicitly passing ctx to repo
    err := s.repo.Create(ctx, c)
    if err != nil {
        return nil, err
    }
    return c, nil
}

func (s *Service) GetAll(ctx context.Context) ([]Camera, error) {
    return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*Camera, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}