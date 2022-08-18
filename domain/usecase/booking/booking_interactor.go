package booking

import (
	"book-meeting-hotel/domain/entity"
	"errors"
	"time"
)

type BookingInteractor struct {
	repo            entity.BookingRepository
	meetingRoomRepo entity.MeetingRoomRepository
}

func NewBookingInteractor(repo entity.BookingRepository, meetingRoomRepo entity.MeetingRoomRepository) BookingInteractor {
	return BookingInteractor{repo: repo, meetingRoomRepo: meetingRoomRepo}
}

func (i BookingInteractor) BookMeetingRoom(meetingRoomId int, employeeId int, startDatetime time.Time,
	endDatetime time.Time, picContactInformation string) (*entity.Booking, error) {
	meetingRoom, err := i.meetingRoomRepo.GetById(meetingRoomId)

	if err != nil {
		return nil, err
	}

	if meetingRoom == nil {
		return nil, errors.New("Meeting room not found")
	}

	existBooking, err := i.repo.GetByDateAndMeetingRoom(meetingRoomId, startDatetime, endDatetime)
	if err != nil {
		return nil, err
	}
	if existBooking != nil {
		return nil, ErrMeetingRoomAlreadyBooked
	}

	newbooking, err := entity.NewBooking(employeeId, *meetingRoom, "invoice.jpg",
		picContactInformation, startDatetime, endDatetime)
	if err != nil {
		return nil, err
	}

	err = i.repo.Save(newbooking)
	if err != nil {
		return nil, err
	}

	return newbooking, err
}
