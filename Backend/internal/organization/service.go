package organization

import (
    "context"
)

type Service struct {
    repo *Repository
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) CreateOrganization(ctx context.Context, org *Organization) error {
    return s.repo.Create(ctx, org)
}

func (s *Service) GetOrganizations(ctx context.Context) ([]Organization, error) {
    return s.repo.GetAll(ctx)
}

func (s *Service) GetOrganizationByID(ctx context.Context, id string) (*Organization, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *Service) DeleteOrganization(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}