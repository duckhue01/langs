package main

type (
	Persistence interface {
		GetUser(string) string
		ListUser(string) []string
	}

	Service struct {
		per Persistence
	}
)

func (s *Service) GetUser(name string) string {
	return s.per.GetUser(name)

}
