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
		Name:           siteDTO.SiteName,
		Slug:           siteDTO.SiteSlug,
		UserID:         siteDTO.UserID,
		SiteIconWindow: siteDTO.SiteIconWindow,
	}

	if err := s.siteRepository.Create(site); err != nil {
		return nil, err
	}

	return &dtos.SiteResponseDTO{
		SiteID:         site.ID,
		SiteName:       site.Name,
		SiteSlug:       site.Slug,
		SiteIconWindow: site.SiteIconWindow,
	}, nil
}

// GetBySlug busca um site pelo slug, incluindo módulos e componentes
func (s *SiteService) GetBySlug(slug string) (*dtos.SiteFullResponseDTO, error) {

	// Buscando o site pelo slug
	site, err := s.siteRepository.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if site == nil {
		return nil, errors.New("site not found")
	}

	// Buscando os módulos associados ao site
	modules, err := s.moduleRepository.FindBySiteID(site.ID)
	if err != nil {
		return nil, err
	}

	// Buscando os componentes associados a cada módulo
	var modulesWithComponents []dtos.ModuleWithComponentsDTO
	for _, module := range modules {
		components, err := s.componentRepository.FindByModuleID(module.ID)
		if err != nil {
			return nil, err
		}

		var componentDTOs []dtos.ComponentDTO
		for _, c := range components {
			componentSettings := make(map[string]interface{})
			for _, s := range c.Settings {
				componentSettings[s.Key] = s.Value
			}
			componentTypeCode := ""
			if c.Type != nil {
				componentTypeCode = c.Type.Code
			}
			// Adicionando o tipo de componente ao DTO
			var itemsDTO []dtos.ComponentItemDTO
			for _, item := range c.Items {
				itemsDTO = append(itemsDTO, dtos.ComponentItemDTO{
					ComponentItemID:           item.ComponentItemID,
					ComponentID:               item.ComponentID,
					ComponentItemTitle:        item.ComponentItemTitle,
					ComponentItemSubtitle:     item.ComponentItemSubtitle,
					ComponentItemSubtitleType: item.ComponentItemSubtitleType,
					ComponentItemText:         item.ComponentItemText,
					ComponentItemImage:        item.ComponentItemImage,
					ComponentItemOrder:        item.ComponentItemOrder,
				})
			}
			componentSettings["items"] = itemsDTO
			componentDTOs = append(componentDTOs, dtos.ComponentDTO{
				ComponentID:       c.ID,
				ComponentTypeId:   c.TypeId,
				ComponentTypeCode: componentTypeCode,
				ComponentName:     c.Name,
				ComponentSettings: componentSettings,
			})
		}

		// Adicionando o módulo e seus componentes ao DTO
		modulesWithComponents = append(modulesWithComponents, dtos.ModuleWithComponentsDTO{
			ModuleID:   module.ID,
			ModuleName: module.Name,
			ModuleSlug: "/" + module.Slug,
			Components: componentDTOs,
		})
	}

	// Buscar componente único do tipo navbar
	navbarComponent, _ := s.componentRepository.FindUniqueBySiteAndTypeCodeLike(site.ID, "navbar")
	var navbarDTO *dtos.ComponentDTO
	if navbarComponent != nil {
		settings := make(map[string]interface{})
		for _, s := range navbarComponent.Settings {
			settings[s.Key] = s.Value
		}
		typeCode := ""
		if navbarComponent.Type != nil {
			typeCode = navbarComponent.Type.Code
		}
		navbarDTO = &dtos.ComponentDTO{
			ComponentID:       navbarComponent.ID,
			ComponentTypeId:   navbarComponent.TypeId,
			ComponentTypeCode: typeCode,
			ComponentName:     navbarComponent.Name,
			ComponentSettings: settings,
		}
	}

	// Buscar componente único do tipo footer
	footerComponent, _ := s.componentRepository.FindUniqueBySiteAndTypeCodeLike(site.ID, "footer")
	var footerDTO *dtos.ComponentDTO
	if footerComponent != nil {
		settings := make(map[string]interface{})
		for _, s := range footerComponent.Settings {
			settings[s.Key] = s.Value
		}
		typeCode := ""
		if footerComponent.Type != nil {
			typeCode = footerComponent.Type.Code
		}
		footerDTO = &dtos.ComponentDTO{
			ComponentID:       footerComponent.ID,
			ComponentTypeId:   footerComponent.TypeId,
			ComponentTypeCode: typeCode,
			ComponentName:     footerComponent.Name,
			ComponentSettings: settings,
		}
	}

	// Retornando o DTO completo do site
	return &dtos.SiteFullResponseDTO{
		SiteID:         site.ID,
		SiteName:       site.Name,
		SiteSlug:       site.Slug,
		SiteIconWindow: site.SiteIconWindow,
		Modules:        modulesWithComponents,
		Navbar:         navbarDTO,
		Footer:         footerDTO,
	}, nil
}
