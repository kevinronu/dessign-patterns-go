package originator

type Originator interface {
	State() State
	SetState(state State)
}
