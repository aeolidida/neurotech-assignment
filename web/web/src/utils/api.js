import { apiClient } from './apiClient';

export const getPatients = async () => {
  const patients = await apiClient('getListPatients', 'GET');
  console.log(patients)
  return patients;
};

export const createPatient = async (patientData) => {
  const response = await apiClient('newPatient', 'POST', patientData);
  patientData["guid"] = response["guid"];
  return patientData;
};

export const updatePatient = async (patientData) => {
  const updatedPatient = await apiClient('editPatient', 'POST', patientData);
  return updatedPatient;
};

export const deletePatient = async (patientGuid) => {
  await apiClient(`delPatient?guid=${patientGuid}`, 'POST');
};