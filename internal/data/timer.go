package data

import "kratos-timer/internal/biz"

type timerRepo struct {
	data *Data
}

func NewTimerRepo(data *Data) biz.TimerRepo {
	return &timerRepo{data: data}
}
