package handler

type TestRepo interface {
	Get(id int) int
}
