package service

import (
	"fmt"
	"strconv"
)

func (s *Service) DeleteUser(idS string) error {
	id, err := strconv.Atoi(idS)
	if err != nil {
		wrappedErr := fmt.Errorf("cannot convert id to int, Error: %v", err.Error())
		return wrappedErr
	}
	return s.Repository.DeleteUser(int64(id))
}
