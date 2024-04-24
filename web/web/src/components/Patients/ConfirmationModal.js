import React from "react";

const ConfirmationModal = ({ isOpen, onConfirm, onCancel }) => {
  return isOpen ? (
    <div className="modal">
      <div className="modal-content">
        <div className="modal-header">
          <h2>Подтверждение удаления</h2>
        </div>
        <div className="modal-body">
          <p>Вы уверены, что хотите удалить этого пациента?</p>
        </div>
        <div className="modal-footer">
          <button className="cancel-button" onClick={onCancel}>
            Отмена
          </button>
          <button className="delete-button" onClick={onConfirm}>
            Да, удалить
          </button>
        </div>
      </div>
    </div>
  ) : null;
};

export default ConfirmationModal;
