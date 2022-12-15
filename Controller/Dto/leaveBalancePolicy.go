package Dto

import "time"

type LeaveBalancePolicyDTO struct {
	ID                              int64       `json:"id" field:"required"`
	CompanyId                       int64       `json:"companyId"`
	CompanyName                     string      `json:"companyName"`
	AutoCutLeaveWeekly              bool        `json:"autoCutLeaveWeekly"`
	LeaveBalanceAccumulation        bool        `json:"leaveBalanceAccumulation"`
	LeaveBalanceBonusByLenghtOfWork bool        `json:"leaveBalanceBonusByLenghtOfWork"`
	LeaveBalanceBonusList           []BonusList `json:"leaveBalanceBonusList"`
	CreatedBy                       string      `json:"createdBy"`
	CreatedTime                     time.Time   `json:"createdTime"`
	ModifiedBy                      string      `json:"modifiedBy"`
	ModifiedTime                    time.Time   `json:"modifiedTime"`
}

type BonusList struct {
	WorkingPeriodStart int64 `json:"workingPeriodStart"`
	WorkingPeriodEnd   int64 `json:"workingPeriodEnd"`
	Reward             int64 `json:"reward"`
}
