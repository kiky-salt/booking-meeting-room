package entity

import "time"

type BookingUseCase interface {
	BookMeetingRoom(meetingRoomId int, employeeId int, startDatetime time.Time,
		endDatetime time.Time, picContactInformation string) (*Booking, error)
}
