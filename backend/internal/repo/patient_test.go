package repo

import (
	"encoding/json"
	"testing"
	"time"

	"neurotech-assignment/backend/internal/errs"
	"neurotech-assignment/backend/internal/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockGUIDGenerator struct {
	mock.Mock
}

func (m *MockGUIDGenerator) GenerateGUID() string {
	args := m.Called()
	return args.String(0)
}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Save(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockStorage) Load() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

type PatientRepoSuite struct {
	suite.Suite
	mockStorage       *MockStorage
	mockGUIDGenerator *MockGUIDGenerator
	patientRepo       *PatientRepo
}

func (suite *PatientRepoSuite) SetupTest() {
	suite.mockStorage = new(MockStorage)
	suite.mockGUIDGenerator = new(MockGUIDGenerator)
	suite.patientRepo = NewPatientRepo(suite.mockStorage, suite.mockGUIDGenerator)
}

func (suite *PatientRepoSuite) TestGetListPatients() {
	suite.mockGUIDGenerator.ExpectedCalls = nil
	suite.mockStorage.ExpectedCalls = nil

	expectedGUID1 := "expected-guid-1"
	expectedGUID2 := "expected-guid-2"

	suite.mockGUIDGenerator.On("GenerateGUID").Return(expectedGUID1, expectedGUID2)

	testPatients := []models.Patient{
		{
			FullName: "John",
			Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			Gender:   0,
			GUID:     expectedGUID1,
		},
		{
			FullName: "Jane",
			Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			Gender:   1,
			GUID:     expectedGUID2,
		},
	}

	jsonData, _ := json.Marshal(testPatients)
	suite.mockStorage.On("Load").Return(jsonData, nil)

	patients, err := suite.patientRepo.GetListPatients()
	suite.NoError(err)

	suite.ElementsMatch(testPatients, patients)
	suite.mockStorage.AssertCalled(suite.T(), "Load")
}

func (suite *PatientRepoSuite) TestNewPatient() {
	testPatient := models.Patient{
		FullName: "John",
		Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
		Gender:   0,
	}

	suite.mockGUIDGenerator.ExpectedCalls = nil
	suite.mockStorage.ExpectedCalls = nil

	expectedGUID := "expected-guid"
	suite.mockGUIDGenerator.On("GenerateGUID").Return(expectedGUID)
	suite.mockStorage.On("Load").Return([]byte{}, nil)

	testPatient.GUID = expectedGUID
	expectedData, _ := json.Marshal([]models.Patient{testPatient})
	suite.mockStorage.On("Save", expectedData).Return(nil)

	guid, err := suite.patientRepo.NewPatient(testPatient)
	suite.NoError(err)
	suite.Equal(expectedGUID, guid)
}

func (suite *PatientRepoSuite) TestEditPatient() {
	editedPatient := models.Patient{
		GUID:     "existing-guid",
		FullName: "Doe",
		Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
		Gender:   0,
	}

	existingPatients := []models.Patient{
		{
			GUID:     "existing-guid",
			FullName: "John",
			Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			Gender:   0,
		},
	}

	suite.mockStorage.ExpectedCalls = nil

	jsonData, _ := json.Marshal(existingPatients)
	suite.mockStorage.On("Load").Return(jsonData, nil)

	expectedPatients := []models.Patient{editedPatient}
	expectedData, _ := json.Marshal(expectedPatients)
	suite.mockStorage.On("Save", expectedData).Return(nil)

	err := suite.patientRepo.EditPatient(editedPatient)
	suite.NoError(err)

	suite.mockStorage.AssertCalled(suite.T(), "Load")
	suite.mockStorage.AssertCalled(suite.T(), "Save", expectedData)
}

func (suite *PatientRepoSuite) TestEditPatient_ErrorNoPatient() {
	editedPatient := models.Patient{
		FullName: "Doe",
		Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
		Gender:   0,
	}

	suite.mockGUIDGenerator.ExpectedCalls = nil
	suite.mockGUIDGenerator.On("GenerateGUID").Return("unused-guid")

	suite.mockStorage.ExpectedCalls = nil
	suite.mockStorage.On("Load").Return([]byte{}, nil)

	err := suite.patientRepo.EditPatient(editedPatient)
	suite.ErrorIs(err, errs.ErrNoPatient)

	suite.mockStorage.AssertCalled(suite.T(), "Load")
}
func (suite *PatientRepoSuite) TestDelPatient() {
	testPatient := models.Patient{
		FullName: "John",
		Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
		Gender:   0,
	}

	expectedGUID := "expected-guid"
	suite.mockGUIDGenerator.ExpectedCalls = nil
	suite.mockGUIDGenerator.On("GenerateGUID").Return(expectedGUID)

	existingPatients := []models.Patient{
		{
			FullName: testPatient.FullName,
			Birthday: testPatient.Birthday,
			Gender:   testPatient.Gender,
			GUID:     expectedGUID,
		},
	}

	suite.mockStorage.ExpectedCalls = nil

	jsonData, _ := json.Marshal(existingPatients)
	suite.mockStorage.On("Load").Return(jsonData, nil)

	expectedPatients := []models.Patient{}
	expectedData, _ := json.Marshal(expectedPatients)
	suite.mockStorage.On("Save", expectedData).Return(nil)

	err := suite.patientRepo.DelPatient(expectedGUID)
	suite.NoError(err)

	suite.mockStorage.AssertCalled(suite.T(), "Load")
	suite.mockStorage.AssertCalled(suite.T(), "Save", expectedData)
}

func (suite *PatientRepoSuite) TestDelPatient_ErrorNoPatient() {
	suite.mockGUIDGenerator.ExpectedCalls = nil
	suite.mockGUIDGenerator.On("GenerateGUID").Return("unused-guid")

	suite.mockStorage.ExpectedCalls = nil
	suite.mockStorage.On("Load").Return([]byte{}, nil)

	err := suite.patientRepo.DelPatient("non-existing-guid")
	suite.ErrorIs(err, errs.ErrNoPatient)

	suite.mockStorage.AssertCalled(suite.T(), "Load")
}

func TestPatientRepoSuite(t *testing.T) {
	suite.Run(t, new(PatientRepoSuite))
}
