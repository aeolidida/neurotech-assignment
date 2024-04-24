import React, { useState, useEffect } from "react";

const PatientForm = ({ mode, patient, onSubmit, onCancel }) => {
  const [formData, setFormData] = useState({
    fullname: "",
    birthday: "",
    gender: null,
  });
  const [errors, setErrors] = useState({});

  useEffect(() => {
    if (mode === "edit" && patient) {
      setFormData({
        guid: patient.guid,
        fullname: patient.fullname,
        birthday: patient.birthday,
        gender: patient.gender !== null ? parseInt(patient.gender, 10) : null,
      });
    } else {
      setFormData({
        fullname: "",
        birthday: "",
        gender: null,
      });
    }
  }, [mode, patient]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]:
        name === "gender" ? (value !== "" ? parseInt(value, 10) : null) : value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const validationErrors = {};

    if (!formData.fullname.trim()) {
      validationErrors.fullname = 'Поле "Полное имя" не должно быть пустым';
    }
    if (!formData.birthday) {
      validationErrors.birthday = 'Поле "Дата рождения" не должно быть пустым';
    }
    if (formData.gender === null) {
      validationErrors.gender = 'Поле "Пол" не должно быть пустым';
    }

    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      return;
    }

    onSubmit(formData);
    setErrors({});
  };
  return (
    <form onSubmit={handleSubmit} className="patient-form">
      <div className="form-inputs-container">
        {mode === "edit" && (
          <label className="form-label">
            GUID: <span>{patient.guid}</span>
          </label>
        )}
        <label className="form-label">Полное имя:</label>
        <input
          type="text"
          name="fullname"
          value={formData.fullname}
          onChange={handleChange}
          className="form-input"
        />
        <div className="form-error">
          {errors.fullname && <span className="error">{errors.fullname}</span>}
        </div>
        <label className="form-label">Дата рождения:</label>
        <input
          type="date"
          name="birthday"
          value={formData.birthday}
          onChange={handleChange}
          className="form-input"
        />
        <div className="form-error">
          {errors.birthday && <span className="error">{errors.birthday}</span>}
        </div>
        <label className="form-label">Пол:</label>
        <select
          name="gender"
          value={formData.gender !== null ? formData.gender : ""}
          onChange={handleChange}
          className="form-input"
        >
          <option value="">Выберите пол</option>
          <option value="0">Мужской</option>
          <option value="1">Женский</option>
        </select>
        <div className="form-error">
          {errors.gender && <span className="error">{errors.gender}</span>}
        </div>
      </div>
      <div className="form-actions">
        <button type="submit" className="submit-button">
          {mode === "add" ? "Добавить" : "Сохранить"}
        </button>
        <button type="button" onClick={onCancel} className="cancel-button">
          Отмена
        </button>
      </div>
    </form>
  );
};

export default PatientForm;
