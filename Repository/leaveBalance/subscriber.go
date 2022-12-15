package leaveBalance

import (
	"EmployeeService/Config"
	"EmployeeService/Controller/Dto"
	"database/sql"
)

func (r NewSubscriberLeaveBalanceRepository) LeaveBalanceUpdate(params Dto.SubscribeTransferDTO) (tx *sql.Tx, err error) {

	query := `UPDATE public.leave_balances SET 
						employee_code=$1, 
						company_name=$2,
						location_name=$3, 
						department_name=$4 
					where employee_id = $5 `
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		query,
		params.NewEmployeeCode,
		params.NewCompanyName,
		params.NewLocationName,
		params.NewDepartmentName,
		params.EmployeeId,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r NewSubscriberLeaveBalanceRepository) LeaveBalanceApproved(params Dto.LeavesConfirmation) (tx *sql.Tx, err error) {

	query := `UPDATE public.leave_balances SET current_balance = current_balance - $1 where employee_id = $2`
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		query,
		params.CurrentBalance,
		params.EmployeeId,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		return
	}
	return
}
