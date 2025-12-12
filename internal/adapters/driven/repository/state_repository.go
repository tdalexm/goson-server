package repository

import (
	"fmt"
	"strconv"

	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type StateRepository struct {
	data map[string][]domain.Record
}

func NewStateRepository(data map[string][]domain.Record) portsdriven.Repository {
	return &StateRepository{data: data}
}

func (sr *StateRepository) List(collection string) ([]domain.Record, error) {
	res, exists := sr.data[collection]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}

	return res, nil
}

func (sr *StateRepository) ListWithFilter(collection string, filter domain.Filter) ([]domain.Record, error) {
	res, exists := sr.data[collection]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}

	var records []domain.Record
	for _, element := range res {
		fieldValue, exist := element[filter.Field]
		if !exist {
			continue
		}

		match, err := filter.Matches(fieldValue)
		if err != nil {
			return nil, err
		}

		if match {
			records = append(records, element)
		}
	}

	return records, nil
}

func (sr *StateRepository) GetByID(collection, id string) (domain.Record, error) {
	res, exists := sr.data[collection]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}

	for _, element := range res {
		if element["id"] == id {
			return element, nil
		}
	}
	return nil, domain.NewAppError(
		domain.ErrCodeNotFound,
		fmt.Sprintf("%s with id '%s' not found", collection, id),
	)
}

func (sr *StateRepository) Create(collection string, record domain.Record) (string, error) {
	var res []domain.Record
	var exists bool
	if res, exists = sr.data[collection]; !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}

	if _, hasID := record["id"]; !hasID {
		newID := sr.generateNextID(res)
		record["id"] = newID
	}

	id, isStr := record["id"].(string)
	if !isStr {
		return "", domain.NewAppError(
			domain.ErrValidation,
			"ID must be a string. Ex: 'id':'25' ",
		)
	}

	for _, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			return "", domain.NewAppError(
				domain.ErrValidation,
				fmt.Sprintf("ID '%s' is duplicated. ID must be unique.", id),
			)
		}
	}

	sr.data[collection] = append(sr.data[collection], record)

	return id, nil
}

func (sr *StateRepository) Update(collection, id string, record domain.Record) (string, error) {
	res, exists := sr.data[collection]
	if !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}

	if _, hasID := record["id"]; hasID {
		return "", domain.NewAppError(
			domain.ErrValidation,
			"Cannot update the ID field",
		)
	}

	for i, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			record["id"] = id
			sr.data[collection][i] = record
			return id, nil
		}
	}

	return "", domain.NewAppError(
		domain.ErrCodeNotFound,
		fmt.Sprintf("%s with ID '%s' not found", collection, id),
	)
}

func (sr *StateRepository) UpdateFields(collection, id string, record domain.Record) (string, error) {
	res, exists := sr.data[collection]
	if !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}

	if _, hasID := record["id"]; hasID {
		return "", domain.NewAppError(
			domain.ErrValidation,
			"Cannot update the ID field",
		)
	}

	foundIndex := -1
	for i, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("%s with ID '%s' not found", collection, id),
		)
	}

	current := res[foundIndex]
	for key, value := range record {
		if key != "id" {
			current[key] = value
		}
	}

	sr.data[collection][foundIndex] = current

	return id, nil
}

func (sr *StateRepository) Delete(collection, id string) (string, error) {
	res, exists := sr.data[collection]
	if !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collection),
		)
	}
	foundIndex := -1
	for i, element := range res {
		if element["id"] == id {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("%s with ID '%s' not found", collection, id),
		)
	}
	sr.data[collection] = append(res[:foundIndex], res[foundIndex+1:]...)
	return id, nil
}

func (sr *StateRepository) generateNextID(collection []domain.Record) string {
	if len(collection) == 0 {
		return "1"
	}

	maxID := 0
	for _, record := range collection {
		if id, ok := record["id"].(int); ok {
			if id > maxID {
				maxID = id
			}
		} else if idStr, ok := record["id"].(string); ok {
			if id, err := strconv.Atoi(idStr); err == nil && id > maxID {
				maxID = id
			}
		}
	}

	nextID := maxID + 1
	return strconv.Itoa(nextID)
}
