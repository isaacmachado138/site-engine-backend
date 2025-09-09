package services

import (
	"errors"
	"myapp/application/dtos"
	"myapp/application/interfaces/repositories"
	"myapp/domain/entities"
	"strings"
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
		Description:    siteDTO.SiteDescription,
		CityID:         siteDTO.CityID,
		Keywords:       siteDTO.SiteKeywords,
		PhoneWhatsapp:  siteDTO.SitePhoneWhatsapp,
		Phone:          siteDTO.SitePhone,
		UserID:         siteDTO.UserID,
		SiteIconWindow: siteDTO.SiteIconWindow,
		Instagram:      siteDTO.SiteInstagram,
		Facebook:       siteDTO.SiteFacebook,
	}

	if err := s.siteRepository.Create(site); err != nil {
		return nil, err
	}
	// Buscar o nome da cidade se CityID foi fornecido
	cityName := ""
	if site.CityID != nil && site.City != nil {
		cityName = site.City.CityName
	}

	return &dtos.SiteResponseDTO{
		SiteID:            site.ID,
		SiteName:          site.Name,
		SiteSlug:          site.Slug,
		SiteDescription:   site.Description,
		CityID:            site.CityID,
		CityName:          cityName,
		SiteHasWebsite:    site.HasWebsite,
		SiteKeywords:      site.Keywords,
		SitePhoneWhatsapp: site.PhoneWhatsapp,
		SitePhone:         site.Phone,
		UserID:            site.UserID,
		SiteIconWindow:    site.SiteIconWindow,
		SiteInstagram:     site.Instagram,
		SiteFacebook:      site.Facebook,
	}, nil
}

// GetBySlug busca um site pelo slug, incluindo módulos e componentes. Se onlyActive=1, retorna apenas módulos/componentes ativos.
func (s *SiteService) GetBySlug(slug string, onlyActive ...int) (*dtos.SiteFullResponseDTO, error) {
	activeOnly := 0
	if len(onlyActive) > 0 {
		activeOnly = onlyActive[0]
	}

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
		if activeOnly == 1 && module.ModuleActive != 1 {
			continue
		}
		var components []entities.Component
		if activeOnly == 1 {
			components, err = s.componentRepository.FindByModuleIDWithActive(module.ID, true)
		} else {
			components, err = s.componentRepository.FindByModuleID(module.ID)
		}
		if err != nil {
			return nil, err
		}
		var componentResponseDTOs []dtos.ComponentResponseDTO
		for _, c := range components {
			componentSettings := make(map[string]interface{})
			for _, s := range c.Settings {
				// Para componentes regulares, também filtrar valores vazios
				if s.Value != "" && len(strings.TrimSpace(s.Value)) > 0 {
					componentSettings[s.Key] = s.Value
				}
			}
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
			// Só adicionar items se realmente existirem
			if len(itemsDTO) > 0 {
				componentSettings["items"] = itemsDTO
			}
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
	} // Buscar componente único do tipo navbar
	navbarComponent, _ := s.componentRepository.FindUniqueBySiteAndTypeCodeLike(site.ID, "navbar")
	var navbarDTO *dtos.ComponentDTO
	if navbarComponent != nil {
		settings := make(map[string]interface{})
		for _, s := range navbarComponent.Settings {
			// Para componentes únicos, só incluir settings com valor não vazio
			if s.Value != "" && len(strings.TrimSpace(s.Value)) > 0 {
				settings[s.Key] = s.Value
			}
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
	} // Buscar componente único do tipo footer
	footerComponent, _ := s.componentRepository.FindUniqueBySiteAndTypeCodeLike(site.ID, "footer")
	var footerDTO *dtos.ComponentDTO
	if footerComponent != nil {
		settings := make(map[string]interface{})
		for _, s := range footerComponent.Settings {
			// Para componentes únicos, só incluir settings com valor não vazio
			if s.Value != "" && len(strings.TrimSpace(s.Value)) > 0 {
				settings[s.Key] = s.Value
			}
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
	cityName := ""
	if site.CityID != nil && site.City != nil {
		cityName = site.City.CityName
	}

	return &dtos.SiteFullResponseDTO{
		SiteID:            site.ID,
		SiteName:          site.Name,
		SiteSlug:          site.Slug,
		SiteDescription:   site.Description,
		CityID:            site.CityID,
		CityName:          cityName,
		SiteHasWebsite:    site.HasWebsite,
		SiteKeywords:      site.Keywords,
		SitePhoneWhatsapp: site.PhoneWhatsapp,
		SitePhone:         site.Phone,
		SiteIconWindow:    site.SiteIconWindow,
		SiteInstagram:     site.Instagram,
		SiteFacebook:      site.Facebook,
		Modules:           modulesWithComponents,
		Navbar:            navbarDTO,
		Footer:            footerDTO,
	}, nil
}

// GetSitesByUser busca todos os sites de um usuário pelo user_id
func (s *SiteService) GetSitesByUser(userId string) ([]dtos.SiteResponseDTO, error) {
	sites, err := s.siteRepository.FindByUserID(userId)
	if err != nil {
		return []dtos.SiteResponseDTO{}, err
	}

	// Inicializar como array vazio para garantir que retorna [] ao invés de null
	resp := make([]dtos.SiteResponseDTO, 0)

	for _, site := range sites {
		cityName := ""
		if site.CityID != nil && site.City != nil {
			cityName = site.City.CityName
		}

		resp = append(resp, dtos.SiteResponseDTO{
			SiteID:            site.ID,
			SiteName:          site.Name,
			SiteSlug:          site.Slug,
			SiteDescription:   site.Description,
			CityID:            site.CityID,
			CityName:          cityName,
			SiteHasWebsite:    site.HasWebsite,
			SiteKeywords:      site.Keywords,
			SitePhoneWhatsapp: site.PhoneWhatsapp,
			SitePhone:         site.Phone,
			SiteIconWindow:    site.SiteIconWindow,
			SiteInstagram:     site.Instagram,
			SiteFacebook:      site.Facebook,
			UserID:            site.UserID,
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
	if updateDTO.SiteDescription != nil {
		site.Description = *updateDTO.SiteDescription
	}
	if updateDTO.CityID != nil {
		site.CityID = updateDTO.CityID
	}
	if updateDTO.SiteKeywords != nil {
		site.Keywords = *updateDTO.SiteKeywords
	}
	if updateDTO.SitePhoneWhatsapp != nil {
		site.PhoneWhatsapp = *updateDTO.SitePhoneWhatsapp
	}
	if updateDTO.SitePhone != nil {
		site.Phone = *updateDTO.SitePhone
	}
	if updateDTO.UserID != nil {
		site.UserID = *updateDTO.UserID
	}
	if updateDTO.SiteIconWindow != nil {
		site.SiteIconWindow = *updateDTO.SiteIconWindow
	}
	if updateDTO.SiteInstagram != nil {
		site.Instagram = *updateDTO.SiteInstagram
	}
	if updateDTO.SiteFacebook != nil {
		site.Facebook = *updateDTO.SiteFacebook
	}
	if err := s.siteRepository.Update(site); err != nil {
		return nil, err
	}

	cityName := ""
	if site.CityID != nil && site.City != nil {
		cityName = site.City.CityName
	}

	return &dtos.SiteResponseDTO{
		SiteID:            site.ID,
		SiteName:          site.Name,
		SiteSlug:          site.Slug,
		SiteDescription:   site.Description,
		CityID:            site.CityID,
		CityName:          cityName,
		SiteHasWebsite:    site.HasWebsite,
		SiteKeywords:      site.Keywords,
		SitePhoneWhatsapp: site.PhoneWhatsapp,
		SitePhone:         site.Phone,
		UserID:            site.UserID,
		SiteIconWindow:    site.SiteIconWindow,
		SiteInstagram:     site.Instagram,
		SiteFacebook:      site.Facebook,
	}, nil
}

// VerifyOwnership verifica se um site pertence a um usuário específico
func (s *SiteService) VerifyOwnership(siteID uint, userID uint) error {
	site, err := s.siteRepository.FindByID(siteID)
	if err != nil {
		return err
	}
	if site == nil {
		return errors.New("site não encontrado")
	}
	if site.UserID != userID {
		return errors.New("este site não pertence ao usuário logado")
	}
	return nil
}

// GetSitesWithFilters busca sites usando filtros genéricos
func (s *SiteService) GetSitesWithFilters(filters repositories.SiteFilters) ([]dtos.SiteResponseDTO, error) {
	sites, err := s.siteRepository.FindWithFilters(filters)
	if err != nil {
		return []dtos.SiteResponseDTO{}, err
	}

	// Inicializar como array vazio para garantir que retorna [] ao invés de null
	resp := make([]dtos.SiteResponseDTO, 0)

	for _, site := range sites {
		cityName := ""
		if site.CityID != nil && site.City != nil {
			cityName = site.City.CityName
		}

		resp = append(resp, dtos.SiteResponseDTO{
			SiteID:            site.ID,
			SiteName:          site.Name,
			SiteSlug:          site.Slug,
			SiteDescription:   site.Description,
			CityID:            site.CityID,
			CityName:          cityName,
			SiteHasWebsite:    site.HasWebsite,
			SiteKeywords:      site.Keywords,
			SitePhoneWhatsapp: site.PhoneWhatsapp,
			SitePhone:         site.Phone,
			SiteIconWindow:    site.SiteIconWindow,
			SiteInstagram:     site.Instagram,
			SiteFacebook:      site.Facebook,
			UserID:            site.UserID,
		})
	}
	return resp, nil
}
