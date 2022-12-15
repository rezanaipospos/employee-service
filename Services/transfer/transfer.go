package Services

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Repository/transfer"
	"database/sql"
)

func (c transferStruct) SaveTransfer(params Dto.TransferDataDTO) (*sql.Tx, transfer.Transfers, error) {
	repo := &transfer.NewTransferRepository{}
	return repo.SaveTransfer(params)
}

func (c transferStruct) DetailTransfer(ID int64) (transfer.Transfers, error) {
	repo := &transfer.NewTransferRepository{}
	return repo.DetailTransfer(ID)
}

func (c transferStruct) DataTransfer(params Dto.TransferParams) (transfer.TransfersData, error) {
	repo := &transfer.NewTransferRepository{}
	return repo.DataTransfer(params)
}

func (c transferStruct) ValidationDuplicateDataTransfer(TransferCode string) (bool, error) {
	repo := &transfer.NewTransferRepository{}
	return repo.ValidationDuplicateDataTransfer(TransferCode)
}

func (c transferStruct) HardDeleteTransfer(ID int64) (bool, error) {
	repo := &transfer.NewTransferRepository{}
	return repo.HardDeleteTransfer(ID)
}
