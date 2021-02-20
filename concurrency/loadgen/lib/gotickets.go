package lib

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
