package server

import (
	"errors"
	"fmt"
	"net/http"

	"neurotech-assignment/backend/internal/errs"
	"neurotech-assignment/backend/internal/models"

	"github.com/gin-gonic/gin"
)

type PatientAPI struct {
	repo PatientRepository
}

type PatientRepository interface {
	GetListPatients() ([]models.Patient, error)
	NewPatient(patient models.Patient) (string, error)
	EditPatient(patient models.Patient) error
	DelPatient(guid string) error
}

func NewPatientAPI(repo PatientRepository) *PatientAPI {
	return &PatientAPI{repo: repo}
}

func (api *PatientAPI) RegisterHandlers(router *gin.Engine) {
	router.GET("/getListPatients", api.getListPatients)
	router.POST("/newPatient", api.newPatient)
	router.POST("/editPatient", api.editPatient)
	router.POST("/delPatient", api.delPatient)
}

func (api *PatientAPI) getListPatients(c *gin.Context) {
	patients, err := api.repo.GetListPatients()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMsg})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func (api *PatientAPI) newPatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		if errors.Is(err, errs.ErrInvalidGender) {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidGenderMsg})
			return
		}
		if errors.Is(err, errs.ErrInvalidDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBirthdayMsg})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidJSONMsg})
		return
	}

	if patient.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoFullNameMsg})
		return
	}

	guid, err := api.repo.NewPatient(patient)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMsg})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"guid": guid})
}

func (api *PatientAPI) editPatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		if errors.Is(err, errs.ErrInvalidGender) {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidGenderMsg})
			return
		}
		if errors.Is(err, errs.ErrInvalidDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBirthdayMsg})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidJSONMsg})
		return
	}

	if patient.GUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoGUIDMsg})
		return
	}
	if patient.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoFullNameMsg})
		return
	}

	if err := api.repo.EditPatient(patient); err != nil {
		if errors.Is(err, errs.ErrNoPatient) {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrNoPatientMsg})
			return
		}

		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMsg})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (api *PatientAPI) delPatient(c *gin.Context) {
	guid := c.Query("guid")

	if guid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoGUIDMsg})
		return
	}

	if err := api.repo.DelPatient(guid); err != nil {
		if errors.Is(err, errs.ErrNoPatient) {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrNoPatientMsg})
			return
		}

		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMsg})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
