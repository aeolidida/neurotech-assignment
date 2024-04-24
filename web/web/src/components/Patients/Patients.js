import React, { useState, useEffect } from "react";
import PatientModal from "./PatientModal";
import {
  getPatients,
  createPatient,
  updatePatient,
  deletePatient,
} from "../../utils/api";
import ErrorModal from "../ErrorModal/ErrorModal";
import PatientsTable from "./PatientsTable";
import "./Patients.css";

const Patients = () => {
  const [patients, setPatients] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [modalMode, setModalMode] = useState("add");
  const [selectedPatient, setSelectedPatient] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");

  const fetchPatients = async () => {
    try {
      const patientsData = await getPatients();
      if (patientsData != null && patientsData.length > 0) {
        setPatients(patientsData);
      }
    } catch (error) {
      setErrorMessage("Ошибка при получении данных пациентов");
    }
  };

  useEffect(() => {
    fetchPatients();
  }, []);

  const handleAddPatient = async (patientData) => {
    try {
      const newPatient = await createPatient(patientData);
      setPatients([...patients, newPatient]);
      setShowModal(false);
    } catch (error) {
      setErrorMessage("Ошибка при добавлении пациента");
    }
  };

  const handleEditPatient = async (patientData) => {
    try {
      const updatedPatient = await updatePatient(patientData);
      const updatedPatients = patients.map((patient) =>
        patient.guid === updatedPatient.guid ? updatedPatient : patient
      );
      setPatients(updatedPatients);
      setShowModal(false);
    } catch (error) {
      setErrorMessage("Ошибка при обновлении данных пациента");
    }
  };

  const handleDeletePatient = async (patientGuid) => {
    try {
      await deletePatient(patientGuid);
      const updatedPatients = patients.filter(
        (patient) => patient.guid !== patientGuid
      );
      setPatients(updatedPatients);
    } catch (error) {
      setErrorMessage("Ошибка при удалении пациента");
    }
  };

  const handleAddPatientClick = () => {
    setSelectedPatient(null);
    setModalMode("add");
    setShowModal(true);
  };

  const handleEditPatientClick = (patient) => {
    setSelectedPatient(patient);
    setModalMode("edit");
    setShowModal(true);
  };

  const handleSubmit = (patientData) => {
    if (modalMode === "add") {
      handleAddPatient(patientData);
    } else {
      handleEditPatient(patientData);
    }
  };

  const columns = React.useMemo(
    () => [
      { Header: "GUID", accessor: "guid" },
      { Header: "Полное имя", accessor: "fullname" },
      { Header: "Дата рождения", accessor: "birthday" },
      {
        Header: "Пол",
        accessor: "gender",
        Cell: ({ value }) => (value === 0 ? "Мужской" : "Женский"),
      },
      ,
      {
        Header: "Действия",
        id: "actions",
      },
    ],
    []
  );

  return (
    <div className="container">
      <h1 className="container-h1">Таблица пациентов</h1>
      <div className="button-container">
        <button className="add-button" onClick={handleAddPatientClick}>
          Добавить
        </button>
        <button className="refresh-button" onClick={() => fetchPatients()}>
          Обновить
        </button>
      </div>
      <PatientsTable
        data={patients}
        columns={columns}
        onEditPatient={handleEditPatientClick}
        onDeletePatient={handleDeletePatient}
      />
      {showModal && (
        <PatientModal
          mode={modalMode}
          patient={selectedPatient}
          onSubmit={handleSubmit}
          onCancel={() => setShowModal(false)}
        />
      )}
      {errorMessage && (
        <ErrorModal
          message={errorMessage}
          onClose={() => setErrorMessage("")}
        />
      )}
    </div>
  );
};

export default Patients;
