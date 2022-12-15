package LeaveBalancePolicy

import (
	"EmployeeService/Config"
	"EmployeeService/Controller/Dto"
	"fmt"
)

const (
	querySelectLeaveBalancePolicy = `Select * from Public.leave_balances_policy`

	queryAddLeaveBalancePolicy = `Insert Into Public.leave_balances_policy (company_id,company_name,auto_cut_leave_weekly,leave_balance_accumulation,leave_balance_bonus_by_lenght_of_work,leave_balance_bonus_list,created_by,created_time) values ($1,$2,$3,$4,$5,$6,$7,$8)`

	queryUpdateLeaveBalancePolicy = `Update Public.leave_balances_policy SET
										auto_cut_leave_weekly=$1,
										leave_balance_accumulation=$2,
										leave_balance_bonus_by_lenght_of_work=$3,
										leave_balance_bonus_list=$4,
										modified_by = $5,
										modified_time=$6 where id=$7`

	queryValidationCompany = `Select Exists(Select * From leave_balances_policy Where company_id=$1)`

	queryCheckValidationCompanyName = `Select id,company_id,company_name,auto_cut_leave_weekly,leave_balance_accumulation,leave_balance_bonus_by_lenght_of_work,leave_balance_bonus_list,created_by,created_time,modified_by,modified_time From leave_balances_policy Where company_name=$1 Limit 1`
)

func (r NewLeaveBalancePolicyRepository) SelectLeaveBalancePolicy(params Dto.LeaveBalancePolicyDTO) (result []LeaveBalancePolicy, err error) {

	var args []interface{}
	var whereClause string

	connection := Config.DATABASE_MAIN.Get()
	query := querySelectLeaveBalancePolicy
	whereClause = `WHERE company_name ILIKE $1`
	args = append(args, fmt.Sprintf("%%%s%%", params.CompanyName)) //$1
	query += whereClause
	row, err := connection.Query(query, args...)

	if row.Err() != nil {
		err = row.Err()
		return
	}
	defer row.Close()

	var lbp LeaveBalancePolicy
	resultData := make([]LeaveBalancePolicy, 0)
	for row.Next() {
		if err = row.Scan(
			&lbp.ID,
			&lbp.CompanyId,
			&lbp.CompanyName,
			&lbp.AutoCutLeaveWeekly,
			&lbp.LeaveBalanceAccumulation,
			&lbp.LeaveBalanceBonusByLenghtOfWork,
			&lbp.LeaveBalanceBonusList,
			&lbp.CreatedBy,
			&lbp.CreatedTime,
			&lbp.ModifiedBy,
			&lbp.ModifiedTime); err != nil {
			return
		}
		resultData = append(resultData, lbp)
	}

	result = resultData
	return
}

func (r NewLeaveBalancePolicyRepository) SaveLeaveBalancePolicy(leaveBalanceBonusList string, params Dto.LeaveBalancePolicyDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryAddLeaveBalancePolicy,
		params.CompanyId,
		params.CompanyName,
		params.AutoCutLeaveWeekly,
		params.LeaveBalanceAccumulation,
		params.LeaveBalanceBonusByLenghtOfWork,
		leaveBalanceBonusList,
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

func (r NewLeaveBalancePolicyRepository) UpdateLeaveBalancePolicy(leaveBalanceBonusList string, params Dto.LeaveBalancePolicyDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryUpdateLeaveBalancePolicy,
		params.AutoCutLeaveWeekly,
		params.LeaveBalanceAccumulation,
		params.LeaveBalanceBonusByLenghtOfWork,
		leaveBalanceBonusList,
		params.ModifiedBy,
		params.ModifiedTime,
		params.ID,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		tx.Commit()
		return
	}

	tx.Commit()
	return
}

func (c NewLeaveBalancePolicyRepository) ValidationDuplicateCompany(companyId int64) (exists bool, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryValidationCompany, companyId).Scan(&exists)
	return
}

func (c NewLeaveBalancePolicyRepository) CheckValidationCompanyName(companyName string) (lbp LeaveBalancePolicy, err error) {

	connection := Config.DATABASE_MAIN.Get()
	row := connection.QueryRow(queryCheckValidationCompanyName, companyName)
	if err = row.Scan(
		&lbp.ID,
		&lbp.CompanyId,
		&lbp.CompanyName,
		&lbp.AutoCutLeaveWeekly,
		&lbp.LeaveBalanceAccumulation,
		&lbp.LeaveBalanceBonusByLenghtOfWork,
		&lbp.LeaveBalanceBonusList,
		&lbp.CreatedBy,
		&lbp.CreatedTime,
		&lbp.ModifiedBy,
		&lbp.ModifiedTime); err != nil {
		return
	}
	return
}
