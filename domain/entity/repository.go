package entity

import "time"

type BookingRepository interface {
	Save(booking *Booking) error
	GetByDateAndMeetingRoom(meetingRoomId int, startDatetime time.Time, endDatetime time.Time) (*Booking, error)
}

type MeetingRoomRepository interface {
	GetById(id int) (*MeetingRoom, error)
}
