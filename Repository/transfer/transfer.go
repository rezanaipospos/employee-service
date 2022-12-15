package transfer

import (
	"EmployeeService/Config"
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Convert"
	validator "EmployeeService/Library/Helper/Validator"
	"database/sql"
	"fmt"
)

const (
	defaultTransferSortColumn = `id`
	defaultTransferSortOrder  = `DESC`
	defaultTransferPageNumber = 0
	defaultTransferPageSize   = 10

	queryInsertTransfer = `INSERT INTO
							public.employee_transfers
								(transfer_code,transfer_date,employee_id,employee_code,machine_id,machine_name,finger_print_id,old_employee_code,old_company_name,old_location_name,old_department_name,old_section_name,old_position_name,old_company_location_code,new_employee_code,new_company_name,new_location_name,new_department_name,new_section_name,new_position_name,new_company_location_code,reason,deleted, created_by, created_time)
							VALUES 
								($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25) 
							RETURNING *`

	queryTransferData = `SELECT A.id, A.transfer_code, A.transfer_date, A.employee_id, A.employee_code, B.name, A.machine_id, A.machine_name, A.finger_print_id, A.old_employee_code, A.old_company_name, A.old_location_name, A.old_department_name, A.old_section_name, A.old_position_name, A.old_company_location_code, A.new_employee_code, A.new_company_name, A.new_location_name, A.new_department_name, A.new_section_name, A.new_position_name, A.new_company_location_code, A.reason, A.created_by, A.created_time, A.modified_by, A.modified_time, COUNT(*) OVER() AS TotalRecords, CASE WHEN ( CEILING(COUNT(*) OVER() / CAST(%d AS FLOAT)) = (%d + 1) ) THEN CAST(1 AS BIT) ELSE CAST(0 AS BIT) END AS HasReachMax from public.employee_transfers A join public.employees B on A.employee_id = b.id`

	queryTransferDetail = `SELECT id,transfer_code,transfer_date,employee_id,employee_code,machine_id,machine_name,finger_print_id,old_employee_code,old_company_name,old_location_name,old_department_name,old_section_name,old_position_name,old_company_location_code,new_employee_code,new_company_name,new_location_name,new_department_name,new_section_name,new_position_name,new_company_location_code,reason,deleted,created_by,created_time,modified_by,modified_time,deleted_by,deleted_time FROM public.employee_transfers WHERE id = $1 AND deleted=FALSE`

	queryValidationTransfer = `SELECT EXISTS(SELECT * FROM public.employee_transfers WHERE transfer_code = $1 AND deleted=FALSE LIMIT 1 )`

	queryHardDeleteTransfer = `DELETE FROM public.employee_transfers WHERE id=$1`
)

func (c NewTransferRepository) setDefaultParams(params Dto.TransferParams) Dto.TransferParams {
	if params.SortColumn == "" {
		params.SortColumn = defaultTransferSortColumn
	}
	if params.SortOrder == "" {
		params.SortOrder = defaultTransferSortOrder
	}
	if params.PageNumber == 0 {
		params.PageNumber = defaultTransferPageNumber
	}
	if params.PageSize == 0 {
		params.PageSize = defaultTransferPageSize
	}
	return params
}

func (c NewTransferRepository) DataTransfer(params Dto.TransferParams) (result TransfersData, err error) {

	var dataTf Transfers
	var whereClause, orderBy, Limit string
	var args []interface{}
	var hasReachMax bool
	var totalRecords int

	connection := Config.DATABASE_MAIN.Get()
	params = c.setDefaultParams(params)
	params.SortOrder = validator.ValidateSortOrder(params.SortOrder, defaultTransferSortOrder)
	params.SortColumn = validator.ValidateSortColumn(allowedFields, params.SortColumn, defaultTransferSortColumn)
	query := fmt.Sprintf(queryTransferData, params.PageSize, params.PageNumber)

	whereClause = ` WHERE A.deleted=FALSE`
	whereClause += ` AND A.transfer_code ILIKE $1`
	whereClause += ` AND A.employee_code ILIKE $2`
	whereClause += ` AND B.name ILIKE $3`

	orderBy = fmt.Sprintf(` ORDER BY %s %s `, Convert.ToSnakeCase(params.SortColumn), params.SortOrder)
	Limit = fmt.Sprintf(` LIMIT %d OFFSET %d * %d `, params.PageSize, params.PageSize, params.PageNumber)

	args = append(args, fmt.Sprintf("%%%s%%", params.TransferCode)) //$1
	args = append(args, fmt.Sprintf("%%%s%%", params.EmployeeCode)) //$2
	args = append(args, fmt.Sprintf("%%%s%%", params.EmployeeName)) //$3

	query += whereClause + orderBy + Limit
	row, err := connection.Query(query, args...)
	if err != nil {
		return
	}
	defer row.Close()

	resultData := make([]Transfers, 0)
	for row.Next() {
		if err = row.Scan(
			&dataTf.ID,
			&dataTf.TransferCode,
			&dataTf.TransferDate,
			&dataTf.EmployeeId,
			&dataTf.EmployeeCode,
			&dataTf.EmployeeName,
			&dataTf.MachineId,
			&dataTf.MachineName,
			&dataTf.FingerPrintId,
			&dataTf.OldEmployeeCode,
			&dataTf.OldCompanyName,
			&dataTf.OldLocationName,
			&dataTf.OldDepartmentName,
			&dataTf.OldSectionName,
			&dataTf.OldPositionName,
			&dataTf.OldCompanyLocationCode,
			&dataTf.NewEmployeeCode,
			&dataTf.NewCompanyName,
			&dataTf.NewLocationName,
			&dataTf.NewDepartmentName,
			&dataTf.NewSectionName,
			&dataTf.NewPositionName,
			&dataTf.NewCompanyLocationCode,
			&dataTf.Reason,
			&dataTf.CreatedBy,
			&dataTf.CreatedTime,
			&dataTf.ModifiedBy,
			&dataTf.ModifiedTime,
			&totalRecords,
			&hasReachMax); err != nil {
			return
		}
		resultData = append(resultData, dataTf)
	}

	if len(resultData) == 0 {
		hasReachMax = true
	}

	result = TransfersData{
		RecordsTotal: totalRecords,
		HasReachMax:  hasReachMax,
		Data:         resultData,
	}
	return
}

func (r NewTransferRepository) SaveTransfer(params Dto.TransferDataDTO) (tx *sql.Tx, transfer Transfers, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryInsertTransfer,
		params.TransferCode,
		params.TransferDate,
		params.EmployeeId,
		params.EmployeeCode,
		params.MachineId,
		params.MachineName,
		params.FingerPrintId,
		params.OldEmployeeCode,
		params.OldCompanyName,
		params.OldLocationName,
		params.OldDepartmentName,
		params.OldSectionName,
		params.OldPositionName,
		params.OldCompanyLocationCode,
		params.NewEmployeeCode,
		params.NewCompanyName,
		params.NewLocationName,
		params.NewDepartmentName,
		params.NewSectionName,
		params.NewPositionName,
		params.NewCompanyLocationCode,
		params.Reason,
		params.Deleted,
		params.CreatedBy,
		params.CreatedTime,
	)

	if err = row.Scan(
		&transfer.ID,
		&transfer.TransferCode,
		&transfer.TransferDate,
		&transfer.EmployeeId,
		&transfer.EmployeeCode,
		&transfer.MachineId,
		&transfer.MachineName,
		&transfer.FingerPrintId,
		&transfer.OldEmployeeCode,
		&transfer.OldCompanyName,
		&transfer.OldLocationName,
		&transfer.OldDepartmentName,
		&transfer.OldSectionName,
		&transfer.OldPositionName,
		&transfer.OldCompanyLocationCode,
		&transfer.NewEmployeeCode,
		&transfer.NewCompanyName,
		&transfer.NewLocationName,
		&transfer.NewDepartmentName,
		&transfer.NewSectionName,
		&transfer.NewPositionName,
		&transfer.NewCompanyLocationCode,
		&transfer.Reason,
		&transfer.Deleted,
		&transfer.CreatedBy,
		&transfer.CreatedTime,
		&transfer.ModifiedBy,
		&transfer.ModifiedTime,
		&transfer.DeletedBy,
		&transfer.DeletedTime); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (c NewTransferRepository) DetailTransfer(ID int64) (result Transfers, err error) {

	connection := Config.DATABASE_MAIN.Get()
	row := connection.QueryRow(queryTransferDetail, ID)

	if row.Err() != nil {
		err = row.Err()
		return
	}

	if err = row.Scan(
		&result.ID,
		&result.TransferCode,
		&result.TransferDate,
		&result.EmployeeId,
		&result.EmployeeCode,
		&result.MachineId,
		&result.MachineName,
		&result.FingerPrintId,
		&result.OldEmployeeCode,
		&result.OldCompanyName,
		&result.OldLocationName,
		&result.OldDepartmentName,
		&result.OldSectionName,
		&result.OldPositionName,
		&result.OldCompanyLocationCode,
		&result.NewEmployeeCode,
		&result.NewCompanyName,
		&result.NewLocationName,
		&result.NewDepartmentName,
		&result.NewSectionName,
		&result.NewPositionName,
		&result.NewCompanyLocationCode,
		&result.Reason,
		&result.Deleted,
		&result.CreatedBy,
		&result.CreatedTime,
		&result.ModifiedBy,
		&result.ModifiedTime,
		&result.DeletedBy,
		&result.DeletedTime); err != nil {
		return
	}
	return
}

func (c NewTransferRepository) ValidationDuplicateDataTransfer(TransferCode string) (exists bool, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryValidationTransfer, TransferCode).Scan(&exists)
	return
}

func (c NewTransferRepository) HardDeleteTransfer(ID int64) (result bool, err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(queryHardDeleteTransfer, ID)
	if err = row.Err(); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}
	tx.Commit()
	return
}
