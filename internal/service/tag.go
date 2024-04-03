package service

type CountTagRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status bool   `form:"status,default=1" binding:"oneof=0 1"`
}

func (s *Service) CountTag(param *CountTagRequest) (int, error) {
	return s.dao.CountTag(param.Name, param.Status)
}
