package Constant

import "net/http"

type InternalSuccess int

const (
	StatusOKJson            InternalSuccess = 2000
	StatusOKAccountApproved InternalSuccess = 2001
	StatusOKValidationOK    InternalSuccess = 2002
	StatusOKNoRowAffected   InternalSuccess = 2003
)

type SuccessInfo struct {
	HttpCode    int
	Description string
	Title       string
}

var successInfo = map[InternalSuccess]SuccessInfo{
	StatusOKJson: {
		HttpCode:    http.StatusOK,
		Description: "Success",
		Title:       http.StatusText(http.StatusOK),
	},
	StatusOKAccountApproved: {
		HttpCode:    http.StatusOK,
		Description: "Account Approved",
		Title:       http.StatusText(http.StatusOK),
	},
	StatusOKValidationOK: {
		HttpCode:    http.StatusOK,
		Description: "Validation Successful",
		Title:       http.StatusText(http.StatusOK),
	},
	StatusOKNoRowAffected: {
		HttpCode:    http.StatusOK,
		Description: "No Row Affected",
		Title:       http.StatusText(http.StatusOK),
	},
}

func (i InternalSuccess) Info() SuccessInfo {
	return successInfo[i]
}
