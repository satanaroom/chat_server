package chat

var _ Service = (*service)(nil)

type Service interface {
}

type service struct {
}

func NewService() *service {
	return &service{}
}
