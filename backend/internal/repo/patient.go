package repo

import (
	"encoding/json"
	"neurotech-assignment/backend/internal/errs"
	"neurotech-assignment/backend/internal/models"
	"sync"
)

type Storage interface {
	Save(data []byte) error
	Load() ([]byte, error)
}

type GUIDGenerator interface {
	GenerateGUID() string
}

type PatientRepo struct {
	storage       Storage
	mu            sync.Mutex
	guidGenerator GUIDGenerator
}

func NewPatientRepo(storage Storage, guidGenerator GUIDGenerator) *PatientRepo {
	return &PatientRepo{
		storage:       storage,
		guidGenerator: guidGenerator,
	}
}

func (pr *PatientRepo) loadPatients() ([]models.Patient, error) {
	op := "PatientRepo.loadPatients"

	data, err := pr.storage.Load()
	if err != nil {
		return nil, errs.WrapError(op, "cannot load data", err)
	}
	var patients []models.Patient

	if len(data) == 0 {
		return patients, nil
	}

	err = json.Unmarshal(data, &patients)
	if err != nil {
		return nil, errs.WrapError(op, "cannot unmarshal data", err)
	}

	return patients, nil
}

func (pr *PatientRepo) savePatients(patients []models.Patient) error {
	op := "PatientRepo.savePatients"

	json, err := json.Marshal(patients)
	if err != nil {
		return errs.WrapError(op, "cannot marshal data", err)
	}
	return pr.storage.Save(json)
}

func (pr *PatientRepo) GetListPatients() ([]models.Patient, error) {
	return pr.loadPatients()
}

func (pr *PatientRepo) NewPatient(patient models.Patient) (string, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	patients, err := pr.GetListPatients()
	if err != nil {
		return "", err
	}

	patient.GUID = pr.guidGenerator.GenerateGUID()
	patients = append(patients, patient)

	return patient.GUID, pr.savePatients(patients)
}

func (pr *PatientRepo) EditPatient(patient models.Patient) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	patients, err := pr.GetListPatients()
	if err != nil {
		return err
	}

	for i, p := range patients {
		if p.GUID == patient.GUID {
			patients[i] = patient
			return pr.savePatients(patients)
		}
	}

	return errs.WrapError("PatientRepo.EditPatient", "error editing patient: patient not found", errs.ErrNoPatient)
}

func (pr *PatientRepo) DelPatient(guid string) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	patients, err := pr.GetListPatients()
	if err != nil {
		return err
	}

	for i, p := range patients {
		if p.GUID == guid {
			patients = append(patients[:i], patients[i+1:]...)
			return pr.savePatients(patients)
		}
	}

	return errs.WrapError("PatientRepo.DelPatient", "error deleting patient: patient not found", errs.ErrNoPatient)
}
