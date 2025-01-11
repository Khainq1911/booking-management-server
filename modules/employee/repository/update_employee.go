package repository

import (
	"context"
	"fmt"
	"manage-system-server/modules/employee/model"
)

func (db *employeeStorage) UpdateEmployRepo(ctx context.Context, id int, updateData model.UpdateEmploy) error {

	result := db.sql.Model(&model.Employee{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated, employee with id %d not found", id)
	}
	return nil
}
