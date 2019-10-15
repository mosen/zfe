package hosts

// The Host service implements all Host based business logic. Templates are actually the same database type so
// the same hostsRepository is used
type HostService interface {
	Find() ([]Host, error)
	Get(id int) (*Host, error)
}

type hostService struct {
	repo HostsRepository
}

func NewHostService(repo HostsRepository) HostService {
	return &hostService{repo}
}

func (svc *hostService) Find() ([]Host, error) {
	return svc.repo.Find()
}

func (svc *hostService) Get(id int) (*Host, error) {
	return nil, nil
}

type TemplateService interface {
	Find() ([]Template, error)
	Get(id int) (*Template, error)
}

type templateService struct {
	repo TemplatesRepository
}

func NewTemplateService(repo TemplatesRepository) TemplateService {
	return &templateService{repo}
}

func (svc *templateService) Find() ([]Template, error) {
	hosts, err := svc.repo.Find()
	if err != nil {
		return nil, err
	}

	return []Template(hosts), nil
}

func (svc *templateService) Get(id int) (*Template, error) {
	return nil, nil
}
