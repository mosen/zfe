package hosts

type HostService interface {
	Find() ([]Host, error)
	Get(id int) (*Host, error)
}

type hostService struct {
	repo Repository
}

func NewService(repo Repository) HostService {
	return &hostService{repo}
}

func (svc *hostService) Find() ([]Host, error) {
	return svc.repo.Find()
}

func (svc *hostService) Get(id int) (*Host, error) {
	return nil, nil
}
