package repository

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/util"

	"gorm.io/gorm"
)

type TrainingRepository struct {
	db *gorm.DB
}

func NewTrainingRepository() *TrainingRepository {
	return &TrainingRepository{
		db: config.GetDB(),
	}
}

func (repo *TrainingRepository) Create(dbObj *model.TrainingModel) (*model.TrainingModel, error) {
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New("failed to create new training")
	}
	return dbObj, nil
}

func (repo *TrainingRepository) Update(id uint, dbObj *model.TrainingModel) (*model.TrainingModel, error) {
	result := repo.db.Where("id = ?", id).Updates(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("failed to update training with ID: ", id))
	}
	return dbObj, nil
}

func (repo *TrainingRepository) GetById(id uint) (*model.TrainingModel, error) {
	dbObj := &model.TrainingModel{}
	result := repo.db.Where("id = ?", id).Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("failed to find training with ID: ", id))
	}
	return dbObj, nil
}

func (repo *TrainingRepository) GetList(pagedData *response.PaginationData) (*response.PaginationData, error) {
	dbObjs := []*model.TrainingModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New("failed to get training list")
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *TrainingRepository) Delete(id uint) error {
	result := repo.db.Where("id_training = ?", id).Delete(&model.EmployeeTrainingModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("failed to delete karyawan training with training ID: ", id))
	}

	result = repo.db.Where("id = ?", id).Delete(&model.TrainingModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("failed to delete training with ID: ", id))
	}
	return nil
}