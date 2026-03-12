package organization

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrganization(org *Organization) error {
	orgs, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	for _, o := range orgs {
		if o.Name == org.Name {
			return errors.New("organization name must be unique")
		}
	}

	return s.repo.Create(org)
}

func (s *Service) GetOrganizations() ([]Organization, error) {
	return s.repo.GetAll()
}

func (s *Service) GetOrganizationByID(id string) (*Organization, error) {
	return s.repo.GetByID(id)
}

func (s *Service) DeleteOrganization(id string) error {
	return s.repo.Delete(id)
}