import React from "react";
import PatientForm from "./PatientForm";

const PatientModal = ({ mode, patient, onSubmit, onCancel }) => {
  const handleSubmit = (formData) => {
    onSubmit(formData);
  };

  const handleCancel = () => {
    onCancel();
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <div className="modal-header">
          <h2>
            {mode === "add" ? "Добавить пациента" : "Редактировать пациента"}
          </h2>
        </div>
        <div className="modal-body">
          <PatientForm
            mode={mode}
            patient={patient}
            onSubmit={handleSubmit}
            onCancel={handleCancel}
          />
        </div>
      </div>
    </div>
  );
};

export default PatientModal;
