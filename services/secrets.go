package services

import (
	"fmt"

	"github.com/bekha-io/secrets/db"
)


type secretsService interface {
	CreateSecret() (*db.Secret, error)
	DeleteSecret() error
}

func (s *serviceManager) CreateSecret(name string, value string, groupName string, isHidden bool) (*db.Secret, error) {	
	group, err := s.GetGroupByField("name", groupName)
	if err != nil {
		return nil, err
	}
	
	var secret *db.Secret
	if tx := db.Database.Find(&secret, "name = ? AND group_id = ?", name, group.ID); tx.RowsAffected != 0 || tx.Error != nil {
		return nil, NewError("Cannot create secret. Secret names should be unique. Secrets with that name in group %v found: %v. Error: %v", group.Name, tx.RowsAffected, tx.Error)
	}

	secret = &db.Secret{
		Name: name,
		Value: value,
		GroupID: group.ID,
		IsHidden: isHidden,
	}

	if tx := db.Database.Create(&secret); tx.RowsAffected == 0 || tx.Error != nil {
		return nil, NewError("Cannot create secret. Error: %v", tx.Error)
	}

	return secret, nil
}


func (s *serviceManager) DeleteSecret(name string, groupName string) error {
	var query string
	if groupName == "" {
		query = fmt.Sprintf("DELETE FROM secrets WHERE name = '%v", name)
	} else {
		group, _ := s.GetGroupByField("name", groupName)
		query = fmt.Sprintf("DELETE FROM secrets WHERE group_id = %v", group.ID)
	}

	if tx := db.Database.Exec(query); tx.RowsAffected == 0 || tx.Error != nil {
		return NewError("Cannot delete group! Groups with that name found: %v. Error: %v", tx.RowsAffected, tx.Error)
	}
	return nil
}


func (s *serviceManager) GetAllSecrets() ([]*db.Secret, error) {
	var secrets []*db.Secret
	if tx := db.Database.Find(&secrets).Preload("group").Order("created_at DESC"); tx.Error != nil {
		return nil, NewError("Cannot get all groups! Error: %v", tx.Error)
	}
	return secrets, nil
}


func (s *serviceManager) GetGroupSecrets(groupName string) ([]*db.Secret, error) {
	group, err := s.GetGroupByField("name", groupName)
	if err != nil {
		return nil, err
	}

	var secrets []*db.Secret
	if tx := db.Database.Find(&secrets, "group_id = ?", group.ID); tx.Error != nil {
		return nil, NewError("Cannot get group's secrets. Error: %v", err)
	}
	return secrets, nil
}