import React, { useState, useEffect } from "react";
import { useTable, useSortBy, useFilters } from "react-table";
import ConfirmationModal from "./ConfirmationModal";
import "./PatientsTable.css";

const PatientsTable = ({ data, columns, onEditPatient, onDeletePatient }) => {
  const [confirmationData, setConfirmationData] = useState(null);
  const [filterValue, setFilterValue] = useState("");

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow,
    setFilter,
  } = useTable({ columns, data }, useFilters, useSortBy);

  const handleDeleteClick = (patient) => {
    setConfirmationData(patient);
  };

  const handleConfirmDelete = () => {
    if (confirmationData) {
      onDeletePatient(confirmationData.guid);
      setConfirmationData(null);
    }
  };

  const handleCancelDelete = () => {
    setConfirmationData(null);
  };

  const handleInputChange = (e) => {
    const value = e.target.value || undefined;
    setFilter("fullname", value);
    setFilterValue(value);
  };

  useEffect(() => {
    if (filterValue) {
      setFilter("fullname", filterValue);
    }
  }, [data, filterValue, setFilter]);

  return (
    <>
      <div className="table-header">
        <input
          className="search-input"
          type="text"
          placeholder="Поиск..."
          value={filterValue}
          onChange={handleInputChange}
        />
      </div>
      <table {...getTableProps()} className="-striped -highlight">
        <thead>
          {headerGroups.map((headerGroup) => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map((column) => (
                <th {...column.getHeaderProps(column.getSortByToggleProps())}>
                  {column.render("Header")}
                  <span>
                    {column.isSorted
                      ? column.isSortedDesc
                        ? " 🔽"
                        : " 🔼"
                      : ""}
                  </span>
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {rows.map((row) => {
            prepareRow(row);
            return (
              <tr {...row.getRowProps()}>
                {row.cells.map((cell) => (
                  <td {...cell.getCellProps()}>
                    {cell.column.id !== "actions" ? (
                      cell.render("Cell")
                    ) : (
                      <>
                        <button
                          className="edit-button"
                          onClick={() => onEditPatient(row.original)}
                        >
                          Изменить
                        </button>
                        <button
                          className="delete-button"
                          onClick={() => handleDeleteClick(row.original)}
                        >
                          Удалить
                        </button>
                      </>
                    )}
                  </td>
                ))}
              </tr>
            );
          })}
        </tbody>
      </table>
      <ConfirmationModal
        isOpen={confirmationData !== null}
        onConfirm={handleConfirmDelete}
        onCancel={handleCancelDelete}
      />
    </>
  );
};

export default PatientsTable;
