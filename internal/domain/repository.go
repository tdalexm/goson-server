package domain

type Repository interface {
	List(resource string) ([]Record, error)
	ListWithFilter(resource string, filter Filter) ([]Record, error)
	GetByID(resource, id string) (Record, error)
	Create(resource string, record Record) (string, error)
	Update(resource, id string, record Record) (string, error)
	UpdateFields(resource, id string, record Record) (string, error)
	Delete(resource, id string) (string, error)
}
