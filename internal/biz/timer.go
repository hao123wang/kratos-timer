package biz

type TimerRepo interface{}

type TimerUsecase struct {
	repo TimerRepo
}

func NewTimerUsecase(repo TimerRepo) *TimerUsecase {
	return &TimerUsecase{repo: repo}
}
