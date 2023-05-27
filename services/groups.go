package services

import (
	"fmt"

	"github.com/bekha-io/secrets/db"
	"gorm.io/gorm/clause"
)

type groupsService interface {
	CreateGroup(name string) (*db.Group, error)
	DeleteGroup(name string) error
	GetGroupByField(string, interface{}) (*db.Group, error)
}


func (s *serviceManager) GetGroupByField(fieldName string, fieldValue interface{}) (*db.Group, error) {
	var group *db.Group
	if tx := db.Database.Find(&group, fmt.Sprintf("%v = ?", fieldName), fieldValue); tx.RowsAffected == 0 || tx.Error != nil {
		return nil, NewError("Cannot find matching group. Groups with that parameter found: %v. Error: %v", tx.RowsAffected, tx.Error)
	}
	return group, nil
}


func (s *serviceManager) CreateGroup(name string) (*db.Group, error) {
	var group *db.Group
	if tx := db.Database.Find(&group, "name = ?", name); tx.RowsAffected != 0 || tx.Error != nil {
		return nil, NewError("Cannot create group. Groups with that name found: %v. Error: %v", tx.RowsAffected, tx.Error)
	}

	group = &db.Group{Name: name,}

	err := group.Validate()
	if err != nil {
		return nil, NewError("Cannot create group: %v", err)
	}

	if tx := db.Database.Create(&group); tx.RowsAffected != 1 || tx.Error != nil {
		return nil, NewError("Cannot create group! Error: %v", tx.Error)
	}

	return group, nil
}


func (s *serviceManager) DeleteGroup(name string) error {
	if tx := db.Database.Delete(&db.Group{}, "name = ?", name); tx.RowsAffected == 0 || tx.Error != nil {
		return NewError("Cannot delete group! Groups with that name found: %v. Error: %v", tx.RowsAffected, tx.Error)
	}
	return nil
}


func (s *serviceManager) GetAllGroups() ([]*db.Group, error) {
	var groups []*db.Group
	if tx := db.Database.Preload(clause.Associations).Find(&groups).Order("created_at DESC"); tx.Error != nil {
		return nil, NewError("Cannot get all groups! Error: %v", tx.Error)
	}
	return groups, nil
}