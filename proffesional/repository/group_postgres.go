package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"proffesional/model"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) AddGroup(group model.AddGroup) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description) VALUES($1, $2) RETURNING id", groupTable)
	row := r.db.QueryRow(query, group.Name, group.Description)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *GroupPostgres) GetAllGroups() ([]model.AllGroup, error) {
	var group []model.AllGroup
	query := fmt.Sprintf("SELECT * FROM %s", groupTable)
	err := r.db.Select(&group, query)
	if err != nil {
		return []model.AllGroup{}, err
	}
	return group, nil
}

func (r *GroupPostgres) ChangeSomeGroup(id int, groupID model.AddGroup) (model.AddGroup, error) {
	var group model.AddGroup
	query := fmt.Sprintf("UPDATE %s SET description = $1 WHERE id = $2", groupTable)
	r.db.Exec(query, groupID.Description, id)
	query = fmt.Sprintf("SELECT name, description FROM %s WHERE id = $1", groupTable)
	err := r.db.Get(&group, query, id)
	if err != nil {
		return model.AddGroup{}, err
	}
	return group, nil
}

func (r *GroupPostgres) DeleteGroup(groupId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", groupTable)
	_, err := r.db.Exec(query, groupId)
	return err
}
