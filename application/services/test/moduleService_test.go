package services

import (
	"testing"

	"myapp/application/dtos"
	"myapp/application/services"
	"myapp/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockModuleRepository é um mock do repositório de módulos
type MockModuleRepository struct {
	mock.Mock
}

// FindBySiteID is a mock implementation of the FindBySiteID method
func (m *MockModuleRepository) FindBySiteID(siteID uint) ([]entities.Module, error) {
	args := m.Called(siteID)
	return args.Get(0).([]entities.Module), args.Error(1)
}

func (m *MockModuleRepository) Create(module *entities.Module) error {
	args := m.Called(module)
	return args.Error(0)
}

func TestModuleService_Create(t *testing.T) {
	repo := new(MockModuleRepository)
	service := services.NewModuleService(repo)

	repo.On("Create", mock.Anything).Return(nil)

	moduleDTO := dtos.ModuleCreateDTO{
		ModuleName:        "Module 1",
		ModuleSlug:        "module-1",
		ModuleDescription: "Description of Module 1",
		SiteID:            1,
	}

	result, err := service.Create(moduleDTO)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	repo.AssertCalled(t, "Create", mock.Anything)
}
