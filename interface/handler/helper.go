package handler

type mysqlRepo struct {
	repo TestRepo
	Val  int
}

func NewTestRepo() TestRepo {
	return &mysqlRepo{
		Val: 10,
	}
}

func (m *mysqlRepo) Get(id int) int {
	val := m.Val
	return val
}
