import React from "react";
import "./ErrorModal.css";

const ErrorModal = ({ message, onClose }) => {
  return (
    <div className="error-modal">
      <div className="error-modal-content">
        <div className="error-modal-header">
          <h2>Ошибка</h2>
        </div>
        <div className="error-modal-body">
          <p>{message}</p>
        </div>
        <div className="error-modal-footer">
          <button onClick={onClose}>OK</button>
        </div>
      </div>
    </div>
  );
};

export default ErrorModal;
