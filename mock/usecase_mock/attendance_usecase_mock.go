package usecase_mock

import (
	"enigmaCamp.com/instructor_led/model"
	"github.com/stretchr/testify/mock"
)

type AttendanceUsecaseMock struct {
	mock.Mock
}

func (m *AttendanceUsecaseMock) GetAttendance(id string) (model.Attendance, error) {
	args := m.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (m *AttendanceUsecaseMock) ListAttendances() ([]model.Attendance, error) {
	args := m.Called()
	return args.Get(0).([]model.Attendance), args.Error(1)
}

func (m *AttendanceUsecaseMock) AddAttendance(user_id string, schedule_id string) (model.Attendance, error) {
	args := m.Called(user_id, schedule_id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (m *AttendanceUsecaseMock) DeleteAttandace(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
