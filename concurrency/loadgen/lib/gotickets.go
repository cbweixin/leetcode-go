package lib

import (
	"errors"
	"fmt"
)

type GoTickets interface {
	// take on ticket
	Take()
	// return one ticket
	Return()
	// ticket pool is active or not
	Active() bool
	// total amount of tickets
	Total() uint32
	// the remaining tickets
	Remaider() uint32
}

type myGotickets struct {
	// total tickets
	total uint32
	// ticket pool
	ticketCh chan struct{}
	// itcket is active or not
	active bool
}

func (gt *myGotickets) Take() {
	<-gt.ticketCh
}

func (gt *myGotickets) Return() {
	gt.ticketCh <- struct{}{}
}

func (gt *myGotickets) Active() bool {
	return gt.active
}

func (gt *myGotickets) Total() uint32 {
	return gt.total
}

func (gt *myGotickets) Remaider() uint32 {
	return uint32(len(gt.ticketCh))
}

// create a new goroutine ticket pool
func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGotickets{}
	if !gt.init(total) {
		errMsg :=
			fmt.Sprintf("The goroutine ticket pool can NOT be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	// rturn gt's address
	return &gt, nil
}

// gt pointer
func (gt *myGotickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}

	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true

	return true

}
