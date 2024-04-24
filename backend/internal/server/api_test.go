package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"neurotech-assignment/backend/internal/errs"
	"neurotech-assignment/backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetListPatients() ([]models.Patient, error) {
	args := m.Called()
	return args.Get(0).([]models.Patient), args.Error(1)
}

func (m *MockRepo) NewPatient(patient models.Patient) (string, error) {
	args := m.Called(patient)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockRepo) EditPatient(patient models.Patient) error {
	args := m.Called(patient)
	return args.Error(0)
}

func (m *MockRepo) DelPatient(guid string) error {
	args := m.Called(guid)
	return args.Error(0)
}

type PatientAPISuite struct {
	suite.Suite
	mockRepo   *MockRepo
	patientAPI *PatientAPI
	ginEngine  *gin.Engine
}

func (suite *PatientAPISuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockRepo = new(MockRepo)
	suite.patientAPI = NewPatientAPI(suite.mockRepo)
	suite.ginEngine = gin.New()
	suite.patientAPI.RegisterHandlers(suite.ginEngine)
}

func (suite *PatientAPISuite) TestGetListPatientsSuccess() {
	repoReturn := []models.Patient{
		{
			FullName: "John Doe",
			Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			Gender:   0,
			GUID:     "1",
		},
		{
			FullName: "Jane Doe",
			Birthday: models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			Gender:   1,
			GUID:     "2",
		},
	}

	suite.mockRepo.On("GetListPatients").Return(repoReturn, nil)

	req, err := http.NewRequest(http.MethodGet, "/getListPatients", nil)
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()
	suite.ginEngine.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var patients []models.Patient
	err = json.NewDecoder(w.Body).Decode(&patients)
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), repoReturn, patients)
}

func (suite *PatientAPISuite) TestGetListPatientsInternalServerError() {
	suite.mockRepo.On("GetListPatients").Return([]models.Patient{}, errors.New("some internal error"))

	req, err := http.NewRequest(http.MethodGet, "/getListPatients", nil)
	assert.NoError(suite.T(), err)

	w := httptest.NewRecorder()
	suite.ginEngine.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	var responseBody map[string]string
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), InternalServerErrorMsg, responseBody["error"])
}

func (suite *PatientAPISuite) TestNewPatient() {
	tests := []struct {
		name           string
		requestBody    gin.H
		repoErr        error
		expectedStatus int
		expectedBody   gin.H
	}{
		{
			"Success",
			gin.H{
				"FullName": "John",
				"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				"Gender":   models.Male,
			},
			nil,
			http.StatusCreated,
			gin.H{
				"FullName": "John",
				"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				"Gender":   models.Male,
			},
		},
		{
			"InvalidGender",
			gin.H{
				"FullName": "John",
				"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				"Gender":   5,
			},
			nil,
			http.StatusBadRequest,
			gin.H{"error": ErrInvalidGenderMsg},
		},
		{
			"NoFullName",
			gin.H{
				"FullName": "",
				"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				"Gender":   models.Male,
			},
			nil,
			http.StatusBadRequest,
			gin.H{"error": ErrNoFullNameMsg},
		},
		{
			"RepoError",
			gin.H{
				"FullName": "John",
				"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				"Gender":   models.Male,
			},
			errors.New("repository error"),
			http.StatusInternalServerError,
			gin.H{"error": InternalServerErrorMsg},
		},
		{
			"NoBirthday",
			gin.H{
				"FullName": "John",
				"Birthday": nil,
				"Gender":   models.Male,
			},
			nil,
			http.StatusBadRequest,
			gin.H{"error": ErrInvalidBirthdayMsg},
		},
	}

	for _, tt := range tests {
		expectedGuid := "expected-guid"
		suite.mockRepo.ExpectedCalls = nil
		suite.mockRepo.On("NewPatient", mock.Anything).Return(expectedGuid, tt.repoErr).Once()
		w := httptest.NewRecorder()
		var reqBody io.Reader

		body, _ := json.Marshal(tt.requestBody)
		reqBody = bytes.NewBuffer(body)

		req, _ := http.NewRequest("POST", "/newPatient", reqBody)
		suite.ginEngine.ServeHTTP(w, req)

		suite.Equalf(tt.expectedStatus, w.Code, "%s: status code mismatch", tt.name)
		if tt.expectedStatus == http.StatusCreated {
			var patient models.Patient
			json.Unmarshal(w.Body.Bytes(), &patient)
			suite.Equalf(expectedGuid, patient.GUID, "%s: guid mismatch", tt.name)
		} else {
			var gotBody gin.H
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			suite.Equalf(tt.expectedBody, gotBody, "%s: response body mismatch", tt.name)
		}
	}
}

func (suite *PatientAPISuite) TestEditPatient() {
	tests := []struct {
		name           string
		requestBody    gin.H
		repoErr        error
		expectedStatus int
		expectedBody   gin.H
	}{
		{"Success", gin.H{
			"GUID":     "1",
			"FullName": "John",
			"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			"Gender":   models.Male,
		}, nil, http.StatusOK, gin.H{"FullName": "John", "Gender": models.Male}},
		{"NoGUID", gin.H{
			"FullName": "John",
			"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			"Gender":   models.Male,
		}, nil, http.StatusBadRequest, gin.H{"error": ErrNoGUIDMsg}},
		{"InvalidGender", gin.H{
			"GUID":     "1",
			"FullName": "John",
			"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			"Gender":   "invalid",
		}, nil, http.StatusBadRequest, gin.H{"error": ErrInvalidGenderMsg}},
		{"NoFullName", gin.H{
			"GUID":     "1",
			"FullName": "",
			"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			"Gender":   models.Male,
		}, nil, http.StatusBadRequest, gin.H{"error": ErrNoFullNameMsg}},
		{"RepoError", gin.H{
			"GUID":     "1",
			"FullName": "John",
			"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			"Gender":   models.Male,
		}, errors.New("repository error"), http.StatusInternalServerError, gin.H{"error": InternalServerErrorMsg}},
		{"PatientNotFound", gin.H{
			"GUID":     "1",
			"FullName": "John",
			"Birthday": models.Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			"Gender":   models.Male,
		}, errs.ErrNoPatient, http.StatusNotFound, gin.H{"error": ErrNoPatientMsg}},
	}

	for _, tt := range tests {
		suite.mockRepo.ExpectedCalls = nil
		suite.mockRepo.On("EditPatient", mock.Anything).Return(tt.repoErr)

		w := httptest.NewRecorder()
		reqBody, _ := json.Marshal(tt.requestBody)
		req, _ := http.NewRequest("POST", "/editPatient", bytes.NewBuffer(reqBody))
		suite.ginEngine.ServeHTTP(w, req)

		suite.Equal(tt.expectedStatus, w.Code, fmt.Sprintf("%s: status code mismatch", tt.name))

		if tt.expectedStatus == http.StatusOK {
			var gotPatient models.Patient
			json.Unmarshal(w.Body.Bytes(), &gotPatient)
			suite.Equal(tt.requestBody["FullName"], gotPatient.FullName, fmt.Sprintf("%s: FullName mismatch", tt.name))
			suite.Equalf(tt.requestBody["Birthday"], gotPatient.Birthday, "%s: Birthday mismatch", tt.name)
			suite.Equal(tt.requestBody["Gender"], gotPatient.Gender, fmt.Sprintf("%s: Gender mismatch", tt.name))
		} else {
			var gotBody gin.H
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			suite.Equal(tt.expectedBody, gotBody, fmt.Sprintf("%s: response body mismatch", tt.name))
		}
	}
}

func (suite *PatientAPISuite) TestDelPatient() {
	tests := []struct {
		name           string
		guid           string
		repoErr        error
		expectedStatus int
		expectedBody   gin.H
	}{
		{"Success", "1", nil, http.StatusNoContent, nil},
		{"NoGUID", "", nil, http.StatusBadRequest, gin.H{"error": ErrNoGUIDMsg}},
		{"RepoError", "1", errors.New("repository error"), http.StatusInternalServerError, gin.H{"error": InternalServerErrorMsg}},
		{"PatientNotFound", "1", errs.ErrNoPatient, http.StatusNotFound, gin.H{"error": ErrNoPatientMsg}},
	}

	for _, tt := range tests {
		suite.mockRepo.ExpectedCalls = nil
		suite.mockRepo.On("DelPatient", tt.guid).Return(tt.repoErr)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/delPatient?guid="+tt.guid, nil)
		suite.ginEngine.ServeHTTP(w, req)

		suite.Equal(tt.expectedStatus, w.Code, fmt.Sprintf("%s: status code mismatch", tt.name))

		if tt.expectedStatus != http.StatusNoContent {
			var gotBody gin.H
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			suite.Equal(tt.expectedBody, gotBody, fmt.Sprintf("%s: response body mismatch", tt.name))
		}
	}
}

func TestPatientAPISuite(t *testing.T) {
	suite.Run(t, new(PatientAPISuite))
}
