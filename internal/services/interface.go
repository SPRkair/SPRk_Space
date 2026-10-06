package services

type Service interface {
	Init()
	Start()
	Stop()
}
