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

		var componentResponseDTOs []dtos.ComponentResponseDTO
		for _, c := range components {
			componentSettings := make(map[string]interface{})
			for _, s := range c.Settings {
				componentSettings[s.Key] = s.Value
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
					ComponentItemLink:         item.ComponentItemLink,
				})
			}
			componentSettings["items"] = itemsDTO // Obter o código do tipo de componente
			typeCode := ""
			if c.Type != nil {
				typeCode = c.Type.Code
			}

			componentResponseDTOs = append(componentResponseDTOs, dtos.ComponentResponseDTO{
				ComponentID:       c.ID,
				ComponentTypeId:   c.TypeId,
				ComponentTypeCode: typeCode,
				ComponentName:     c.Name,
				UserId:            c.UserID,
				ComponentSettings: componentSettings,
			})
		}

		// Adicionando o módulo e seus componentes ao DTO
		modulesWithComponents = append(modulesWithComponents, dtos.ModuleWithComponentsDTO{
			ModuleID:          module.ID,
			ModuleName:        module.Name,
			ModuleSlug:        "/" + module.Slug,
			ModuleDescription: module.Description,
			ModuleOrder:       module.Order,
			SiteID:            module.SiteID,
			ModuleActive:      module.ModuleActive, // Propaga o campo para o DTO
			Components:        componentResponseDTOs,
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

// GetSitesByUser busca todos os sites de um usuário pelo user_id
func (s *SiteService) GetSitesByUser(userId string) ([]dtos.SiteResponseDTO, error) {
	sites, err := s.siteRepository.FindByUserID(userId)
	if err != nil {
		return nil, err
	}
	var resp []dtos.SiteResponseDTO
	for _, site := range sites {
		resp = append(resp, dtos.SiteResponseDTO{
			SiteID:         site.ID,
			SiteName:       site.Name,
			SiteSlug:       site.Slug,
			SiteIconWindow: site.SiteIconWindow,
			UserID:         site.UserID,
		})
	}
	return resp, nil
}

// Update atualiza um site existente (parcial)
func (s *SiteService) Update(siteID uint, updateDTO dtos.SiteUpdateDTO) (*dtos.SiteResponseDTO, error) {
	site, err := s.siteRepository.FindByID(siteID)
	if err != nil {
		return nil, err
	}
	if site == nil {
		return nil, errors.New("site não encontrado")
	}
	if updateDTO.SiteName != nil {
		site.Name = *updateDTO.SiteName
	}
	if updateDTO.SiteSlug != nil {
		site.Slug = *updateDTO.SiteSlug
	}
	if updateDTO.UserID != nil {
		site.UserID = *updateDTO.UserID
	}
	if updateDTO.SiteIconWindow != nil {
		site.SiteIconWindow = *updateDTO.SiteIconWindow
	}
	if err := s.siteRepository.Update(site); err != nil {
		return nil, err
	}
	return &dtos.SiteResponseDTO{
		SiteID:         site.ID,
		SiteName:       site.Name,
		SiteSlug:       site.Slug,
		UserID:         site.UserID,
		SiteIconWindow: site.SiteIconWindow,
	}, nil
}
