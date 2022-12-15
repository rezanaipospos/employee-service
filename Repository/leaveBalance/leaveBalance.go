package leaveBalance

import (
	"EmployeeService/Config"
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Convert"
	validator "EmployeeService/Library/Helper/Validator"
	"fmt"
)

const (
	defaultLeaveBalanceSortColumn = `employee_code`
	defaultLeaveBalanceSortOrder  = `DESC`
	defaultLeaveBalancePageNumber = 0
	defaultLeaveBalancePageSize   = 10

	queryLeaveBalanceData = `SELECT lb.id,lb.employee_id, lb.employee_code, em.name, lb.company_name, lb.location_name, lb.department_name, lb.start_balances, lb.increase_balance, lb.decrease_balance, lb.last_period_balance, lb.current_balance, lb.is_active, lb.period, lb.join_date, lb.expired_date, lb.deleted, lb.created_by, lb.created_time, lb.modified_by, lb.modified_time, lb.deleted_by, lb.deleted_time, COUNT(*) OVER() AS TotalRecords, CASE WHEN ( CEILING(COUNT(*) OVER() / CAST(%d AS FLOAT)) = (%d + 1) ) THEN CAST(1 AS BIT) ELSE CAST(0 AS BIT) END AS HasReachMax from public.leave_balances As lb INNER JOIN public.employees as em ON lb.employee_id = em.id `

	queryInsertLeaveBalance = `INSERT INTO
							public.leave_balances
								(employee_id,employee_code,company_name,location_name,department_name,start_balances,increase_balance,decrease_balance,last_period_balance,current_balance,is_active,period,join_date,expired_date, deleted, created_by, created_time)
							VALUES 
								($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17) `

	queryInsertLeaveBalanceAdjusment = `INSERT INTO
						public.leave_balances_adjustment
							(employee_id,employee_code,start_date,end_date,type,quantity,reason,deleted,created_by,created_time,leave_id)
						VALUES 
							($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id,employee_id,employee_code,start_date,end_date,type,quantity,reason,deleted,created_by,created_time`

	queryUpdateLeaveBalance = ` UPDATE public.leave_balances SET 
									increase_balance=$1, 
									decrease_balance=$2,
									current_balance=$3 
								where employee_id = $4 `

	queryLeaveBalanceDetail = `select lb.id, lb.employee_id, lb.employee_code, em.name, lb.company_name, lb.location_name, lb.department_name, lb.start_balances, lb.increase_balance, lb.decrease_balance, lb.last_period_balance, lb.current_balance, lb.is_active, lb.period, lb.join_date, lb.expired_date from public.leave_balances As lb JOIN public.employees as em ON lb.employee_id = em.id WHERE lb.employee_id = $1 AND lb.deleted = FALSE`

	queryValidationLeaveBalance = `SELECT EXISTS(SELECT * FROM public.leave_balances WHERE employee_id = $1 AND deleted=FALSE LIMIT 1 )`

	queryDetailLeaveBalanceAdjustment = `select lba.id, lba.employee_id, lba.employee_code, lba.start_date, lba.end_date, lba.type, lba.quantity, lba.reason, lba.deleted, lba.created_by, lba.created_time from public.leave_balances_adjustment as lba JOIN leave_balances as lb on lba.employee_id = lb.employee_id WHERE EXTRACT(YEAR FROM join_date) = $1 AND lba.employee_id = $2 AND start_date BETWEEN lb.join_date + ( $3 - EXTRACT( YEAR FROM lb.join_date ) ||' years') :: interval AND lb.expired_date`

	queryHardDeleteLeaveBalance = `DELETE FROM public.leave_balances WHERE employee_id = $1`

	//query Scheduler Reset Leave Balance
	querySelectLeaveBalanceActive = `Select id, employee_id, employee_code, company_name, location_name, department_name, start_balances, increase_balance, decrease_balance, last_period_balance, current_balance, is_active, period, join_date, expired_date from leave_balances where expired_date = now() :: date AND is_active = true And deleted = false`
	queryUpdateLeaveBalanceActive = `Update public.leave_balances SET is_active=false where id=$1`
)

func (c NewLeaveBalanceRepository) setDefaultParams(params Dto.LeaveBalanceParams) Dto.LeaveBalanceParams {
	if params.SortColumn == "" {
		params.SortColumn = defaultLeaveBalanceSortColumn
	}
	if params.SortOrder == "" {
		params.SortOrder = defaultLeaveBalanceSortOrder
	}
	if params.PageNumber == 0 {
		params.PageNumber = defaultLeaveBalancePageNumber
	}
	if params.PageSize == 0 {
		params.PageSize = defaultLeaveBalancePageSize
	}
	return params
}

func (r NewLeaveBalanceRepository) SaveLeaveBalanceAdjusment(params Dto.LeaveBalanceAdjustmentDTO) (cuti LeaveBalanceAdjustment, err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryInsertLeaveBalanceAdjusment,
		params.EmployeeId,
		params.EmployeeCode,
		params.StartDate,
		params.EndDate,
		params.Type,
		params.Quantity,
		params.Reason,
		params.Deleted,
		params.CreatedBy,
		params.CreatedTime,
		params.LeaveId,
	)

	if err = row.Scan(
		&cuti.Id,
		&cuti.EmployeeId,
		&cuti.EmployeeCode,
		&cuti.StartDate,
		&cuti.EndDate,
		&cuti.Type,
		&cuti.Quantity,
		&cuti.Reason,
		&cuti.Deleted,
		&cuti.CreatedBy,
		&cuti.CreatedTime); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}

	tx.Commit()
	return
}

func (r NewLeaveBalanceRepository) SelectEmployeeLeaveBalance(params Dto.LeaveBalanceDataDTO) (cuti LeaveBalance, err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	query := ` SELECT employee_id,start_balances,increase_balance,decrease_balance,last_period_balance,current_balance,is_active 
	FROM public.leave_balances WHERE employee_id = $1 `

	rows := tx.QueryRow(
		query,
		params.EmployeeId,
	)

	if err = rows.Scan(
		&cuti.EmployeeId,
		&cuti.StartBalances,
		&cuti.IncreaseBalance,
		&cuti.DecreaseBalance,
		&cuti.LastPeriodBalance,
		&cuti.CurrentBalance,
		&cuti.IsActive); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}

	tx.Commit()
	return
}

func (r NewLeaveBalanceRepository) UpdateLeaveBalance(params Dto.LeaveBalanceDataDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryUpdateLeaveBalance,
		params.IncreaseBalance,
		params.DecreaseBalance,
		params.CurrentBalance,
		params.EmployeeId,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}

	tx.Commit()
	return
}

func (c NewLeaveBalanceRepository) DataLeaveBalance(params Dto.LeaveBalanceParams) (result LeaveBalanceData, err error) {

	var saldo LeaveBalance
	var whereClause, orderBy, Limit string
	var args []interface{}
	var hasReachMax bool
	var totalRecords int

	connection := Config.DATABASE_MAIN.Get()
	params = c.setDefaultParams(params)
	params.SortOrder = validator.ValidateSortOrder(params.SortOrder, defaultLeaveBalanceSortOrder)
	params.SortColumn = validator.ValidateSortColumn(allowedFields, params.SortColumn, defaultLeaveBalanceSortColumn)
	query := fmt.Sprintf(queryLeaveBalanceData, params.PageSize, params.PageNumber)

	whereClause = ` WHERE lb.deleted=FALSE`
	whereClause += ` AND company_name ILIKE $1`
	whereClause += ` AND location_name ILIKE $2`
	whereClause += ` AND department_name ILIKE $3`
	orderBy = fmt.Sprintf(` ORDER BY %s %s `, Convert.ToSnakeCase(params.SortColumn), params.SortOrder)
	Limit = fmt.Sprintf(` LIMIT %d OFFSET %d * %d `, params.PageSize, params.PageSize, params.PageNumber)

	args = append(args, fmt.Sprintf("%%%s%%", params.CompanyName))    //$1
	args = append(args, fmt.Sprintf("%%%s%%", params.LocationName))   //$2
	args = append(args, fmt.Sprintf("%%%s%%", params.DepartmentName)) //$3

	query += whereClause + orderBy + Limit
	row, err := connection.Query(query, args...)
	if err != nil {
		return
	}
	defer row.Close()

	resultData := make([]LeaveBalance, 0)
	for row.Next() {
		if err = row.Scan(
			&saldo.ID,
			&saldo.EmployeeId,
			&saldo.EmployeeCode,
			&saldo.CompanyName,
			&saldo.LocationName,
			&saldo.DepartmentName,
			&saldo.StartBalances,
			&saldo.IncreaseBalance,
			&saldo.DecreaseBalance,
			&saldo.LastPeriodBalance,
			&saldo.CurrentBalance,
			&saldo.IsActive,
			&saldo.Period,
			&saldo.JoinDate,
			&saldo.ExpiredDate,
			&saldo.Deleted,
			&saldo.CreatedBy,
			&saldo.CreatedTime,
			&saldo.ModifiedBy,
			&saldo.ModifiedTime,
			&saldo.DeletedBy,
			&saldo.DeletedTime,
			&totalRecords,
			&hasReachMax); err != nil {
			return
		}
		resultData = append(resultData, saldo)
	}

	if len(resultData) == 0 {
		hasReachMax = true
	}

	result = LeaveBalanceData{
		RecordsTotal: totalRecords,
		HasReachMax:  hasReachMax,
		Data:         resultData,
	}
	return
}

func (c NewLeaveBalanceRepository) DetailLeaveBalance(EmployeeId int64) (result DetailLeaveBalance, err error) {

	connection := Config.DATABASE_MAIN.Get()
	row := connection.QueryRow(queryLeaveBalanceDetail, EmployeeId)

	if row.Err() != nil {
		err = row.Err()
		return
	}

	if err = row.Scan(
		&result.ID,
		&result.EmployeeId,
		&result.EmployeeCode,
		&result.EmployeeName,
		&result.CompanyName,
		&result.LocationName,
		&result.DepartmentName,
		&result.StartBalances,
		&result.IncreaseBalance,
		&result.DecreaseBalance,
		&result.LastPeriodBalance,
		&result.CurrentBalance,
		&result.IsActive,
		&result.Period,
		&result.JoinDate,
		&result.ExpiredDate); err != nil {
		return
	}
	return
}

func (c NewLeaveBalanceRepository) DetailLeaveBalanceAdjustment(Tahun, EmployeeId int64) (result LeaveBalanceAdjustmentData, err error) {
	var dataAdjust LeaveBalanceAdjustment
	connection := Config.DATABASE_MAIN.Get()
	row, err := connection.Query(queryDetailLeaveBalanceAdjustment, Tahun, EmployeeId, Tahun)
	if row.Err() != nil {
		err = row.Err()
		return
	}
	defer row.Close()

	resultData := make([]LeaveBalanceAdjustment, 0)
	for row.Next() {
		if err = row.Scan(
			&dataAdjust.Id,
			&dataAdjust.EmployeeId,
			&dataAdjust.EmployeeCode,
			&dataAdjust.StartDate,
			&dataAdjust.EndDate,
			&dataAdjust.Type,
			&dataAdjust.Quantity,
			&dataAdjust.Reason,
			&dataAdjust.Deleted,
			&dataAdjust.CreatedBy,
			&dataAdjust.CreatedTime); err != nil {
			return
		}
		resultData = append(resultData, dataAdjust)
	}
	result = LeaveBalanceAdjustmentData{
		Data: resultData,
	}
	return
}

func (c NewLeaveBalanceRepository) ValidationDuplicateDataLeaveBalance(EmployeeId int64) (exists bool, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryValidationLeaveBalance, EmployeeId).Scan(&exists)
	return
}

func (c NewLeaveBalanceRepository) HardDeleteLeaveBalance(EmployeeId int64) (result bool, err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(queryHardDeleteLeaveBalance, EmployeeId)
	if err = row.Err(); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}
	tx.Commit()
	return
}

func (r NewLeaveBalanceRepository) SelectLeaveBalanceActive() (cuti []ResetLeaveBalance, err error) {

	connection := Config.DATABASE_MAIN.Get()
	var result ResetLeaveBalance
	var args []interface{}

	rows, err := connection.Query(querySelectLeaveBalanceActive, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&result.ID,
			&result.EmployeeId,
			&result.EmployeeCode,
			&result.CompanyName,
			&result.LocationName,
			&result.DepartmentName,
			&result.StartBalances,
			&result.IncreaseBalance,
			&result.DecreaseBalance,
			&result.LastPeriodBalance,
			&result.CurrentBalance,
			&result.IsActive,
			&result.Period,
			&result.JoinDate,
			&result.ExpiredDate,
		); err != nil {
			return
		}
		cuti = append(cuti, result)
	}

	return
}

func (r NewLeaveBalanceRepository) UpdateLeaveBalanceActive(ID int64) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryUpdateLeaveBalanceActive,
		ID,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}

	tx.Commit()
	return
}

// Dieksekusi setelah EmployeeAdd dijalankan pada service employee
func (r NewLeaveBalanceRepository) SaveLeaveBalance(params Dto.LeaveBalanceDataDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryInsertLeaveBalance,
		params.EmployeeId,
		params.EmployeeCode,
		params.CompanyName,
		params.LocationName,
		params.DepartmentName,
		params.StartBalances,
		params.IncreaseBalance,
		params.DecreaseBalance,
		params.LastPeriodBalance,
		params.CurrentBalance,
		params.IsActive,
		params.Period,
		params.JoinDate,
		params.ExpiredDate,
		params.Deleted,
		params.CreatedBy,
		params.CreatedTime,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}

	tx.Commit()
	return
}
