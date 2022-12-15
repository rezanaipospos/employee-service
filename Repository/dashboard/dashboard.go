package dashboard

import (
	"EmployeeService/Config"
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	queryTotalReligionSummary            = `SELECT religion, COUNT(religion) FROM public.employees WHERE active_status = true AND deleted = false GROUP BY religion`
	queryTotalWorkStatusSummary          = `SELECT work_status, COUNT(work_status) FROM public.employees WHERE active_status = true AND deleted = false GROUP BY work_status`
	queryTotalWillExpireEmployeeContract = `SELECT COALESCE((count_contract + count_training) * 100 / NULLIF(total_work_status, 0)::FLOAT, 0) AS percentage, count_contract, count_training FROM (
												SELECT 
												(SELECT COUNT(work_status) FROM public.employees WHERE active_status = true AND deleted = false) AS total_work_status, 
												(SELECT COUNT(work_status) FROM public.employees WHERE LOWER(work_status) = 'kontrak' AND contract_end BETWEEN (CURRENT_DATE + INTERVAL '1 day')::DATE AND (CURRENT_DATE + INTERVAL '2 weeks')::DATE AND (active_status IS NULL OR active_status = true) AND deleted = false) AS count_contract,
												(SELECT COUNT(work_status) FROM public.employees WHERE LOWER(work_status) = 'training' AND contract_end BETWEEN (CURRENT_DATE + INTERVAL '1 day')::DATE AND (CURRENT_DATE + INTERVAL '2 weeks')::DATE AND (active_status IS NULL OR active_status = true) AND deleted = false) AS count_training
											) AS multiple_select`
	queryTotalEmployeeByLengthOfWork = `SELECT total_male + total_female AS total_employees, total_male, total_female FROM (
											SELECT 
											(SELECT COUNT(gender) FROM public.employees WHERE LOWER(gender) = 'male' AND (join_date BETWEEN CURRENT_DATE - $1::INTERVAL AND CURRENT_DATE) AND work_status NOT ILIKE '%Resign%' AND (active_status IS NULL OR active_status = true) AND deleted = false) AS total_male,
											(SELECT COUNT(gender) FROM public.employees WHERE LOWER(gender) = 'female' AND (join_date BETWEEN CURRENT_DATE - $1::INTERVAL AND CURRENT_DATE) AND work_status NOT ILIKE '%Resign%' AND (active_status IS NULL OR active_status = true) AND deleted = false) AS total_female
										) AS multiple_select`
)

func (c NewDashboardRepository) NewEmployeeData() (result []NewEmployeeData, err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "join_date", Value: -1}}).SetLimit(10)
	cur, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = cur.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (c NewDashboardRepository) TotalReligionSummary() (result []TotalReligionSummary, err error) {
	var totalReligionSummaryData TotalReligionSummary

	connection := Config.DATABASE_MAIN.Get()

	rTotalReligionSummaryData := reflect.ValueOf(&totalReligionSummaryData).Elem()
	columns := make([]interface{}, 0)
	for i := 0; i < rTotalReligionSummaryData.NumField(); i++ {
		columns = append(columns, rTotalReligionSummaryData.Field(i).Addr().Interface())
	}

	rows, err := connection.Query(queryTotalReligionSummary)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(columns...); err != nil {
			return
		}
		result = append(result, totalReligionSummaryData)
	}
	return
}

func (c NewDashboardRepository) TotalWorkStatusSummary() (result []TotalWorkStatusSummary, err error) {
	var totalWorkStatusSummaryData TotalWorkStatusSummary

	connection := Config.DATABASE_MAIN.Get()

	rTotalWorkStatusSummary := reflect.ValueOf(&totalWorkStatusSummaryData).Elem()
	columns := make([]interface{}, 0)
	for i := 0; i < rTotalWorkStatusSummary.NumField(); i++ {
		columns = append(columns, rTotalWorkStatusSummary.Field(i).Addr().Interface())
	}

	rows, err := connection.Query(queryTotalWorkStatusSummary)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(columns...); err != nil {
			return
		}
		result = append(result, totalWorkStatusSummaryData)
	}
	return
}

func (c NewDashboardRepository) TotalWillExpireEmployeeContract() (result TotalWillExpireEmployeeContract, err error) {
	connection := Config.DATABASE_MAIN.Get()

	rTotalWillExpireEmployeeContract := reflect.ValueOf(&result).Elem()
	columns := make([]interface{}, 0)
	for i := 0; i < rTotalWillExpireEmployeeContract.NumField(); i++ {
		columns = append(columns, rTotalWillExpireEmployeeContract.Field(i).Addr().Interface())
	}

	row := connection.QueryRow(queryTotalWillExpireEmployeeContract)

	if err = row.Scan(columns...); err != nil {
		return
	}
	return
}

func (c NewDashboardRepository) TotalEmployeeByLengthOfWork(numberOfYear int64) (result TotalEmployeeByLengthOfWork, err error) {
	var args []interface{}
	connection := Config.DATABASE_MAIN.Get()

	rTotalEmployeeByLengthOfWork := reflect.ValueOf(&result).Elem()
	columns := make([]interface{}, 0)
	for i := 0; i < rTotalEmployeeByLengthOfWork.NumField(); i++ {
		columns = append(columns, rTotalEmployeeByLengthOfWork.Field(i).Addr().Interface())
	}

	args = append(args, fmt.Sprintf("%d years", numberOfYear))
	row := connection.QueryRow(queryTotalEmployeeByLengthOfWork, args...)

	if err = row.Scan(columns...); err != nil {
		return
	}
	return
}
