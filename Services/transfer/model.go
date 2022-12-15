package Services

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Repository/transfer"
	"database/sql"
)

var (
	Transfer transferInterface = &transferStruct{}
)

type transferInterface interface {
	SaveTransfer(params Dto.TransferDataDTO) (*sql.Tx, transfer.Transfers, error)
	DetailTransfer(ID int64) (transfer.Transfers, error)
	DataTransfer(params Dto.TransferParams) (transfer.TransfersData, error)
	ValidationDuplicateDataTransfer(TransferCode string) (bool, error)
	HardDeleteTransfer(ID int64) (bool, error)
}

type transferStruct struct {
}
