package portsdriver

type DeleteService interface {
	Execute(record, id string) (string, error)
}
