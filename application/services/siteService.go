package services

import (
	"errors"
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
)

// SiteService lida com operações relacionadas a sites
type SiteService struct {
	siteRepository      repositories.SiteRepository
	moduleRepository    repositories.ModuleRepository
	componentRepository repositories.ComponentRepository
}

// NewSiteService cria uma nova instância de SiteService
func NewSiteService(siteRepository repositories.SiteRepository, moduleRepository repositories.ModuleRepository, componentRepository repositories.ComponentRepository) *SiteService {
	return &SiteService{
		siteRepository:      siteRepository,
		moduleRepository:    moduleRepository,
		componentRepository: componentRepository,
	}
}

// Create cria um novo site
func (s *SiteService) Create(siteDTO dtos.SiteCreateDTO) (*dtos.SiteResponseDTO, error) {
	site := &entities.Site{
		Name:   siteDTO.SiteName,
		Slug:   siteDTO.SiteSlug,
		UserID: siteDTO.UserID,
	}

	if err := s.siteRepository.Create(site); err != nil {
		return nil, err
	}

	return &dtos.SiteResponseDTO{
		SiteID:   site.ID,
		SiteName: site.Name,
		SiteSlug: site.Slug,
	}, nil
}

// GetBySlug busca um site pelo slug, incluindo módulos e componentes
func (s *SiteService) GetBySlug(slug string) (*dtos.SiteFullResponseDTO, error) {
	site, err := s.siteRepository.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if site == nil {
		return nil, errors.New("site not found")
	}

	modules, err := s.moduleRepository.FindBySiteID(site.ID)
	if err != nil {
		return nil, err
	}

	var modulesWithComponents []dtos.ModuleWithComponentsDTO
	for _, module := range modules {
		components, err := s.componentRepository.FindByModuleID(module.ID)
		if err != nil {
			return nil, err
		}

		var componentDTOs []dtos.ComponentDTO
		for _, c := range components {
			componentSettings := make(map[string]string)
			for _, s := range c.Settings {
				componentSettings[s.Key] = s.Value
			}
			componentDTOs = append(componentDTOs, dtos.ComponentDTO{
				ComponentID:       c.ID,
				ComponentType:     c.Type,
				ComponentName:     c.Name,
				ComponentSettings: componentSettings,
			})
		}

		modulesWithComponents = append(modulesWithComponents, dtos.ModuleWithComponentsDTO{
			ModuleID:   module.ID,
			ModuleName: module.Name,
			ModuleSlug: "/" + module.Slug,
			Components: componentDTOs,
		})
	}

	return &dtos.SiteFullResponseDTO{
		SiteID:   site.ID,
		SiteName: site.Name,
		SiteSlug: site.Slug,
		Modules:  modulesWithComponents,
	}, nil
}
