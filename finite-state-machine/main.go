package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

type Door struct {
	// 状态值
	state string
	// 状态变化时触发的动作
	beforeOpenAction  func(*fsm.Event)
	beforeCloseAction func(*fsm.Event)
	openingAction     func(*fsm.Event)
	closingAction     func(*fsm.Event)
	afterOpenAction   func(*fsm.Event)
	afterCloseAction  func(*fsm.Event)
	events            fsm.Events
	callbacks         fsm.Callbacks
}

func initDoor() *Door {
	door := new(Door)
	door.state = "closed"
	door.beforeOpenAction = func(e *fsm.Event) {
		fmt.Println("before open")
	}
	door.beforeCloseAction = func(e *fsm.Event) {
		fmt.Println("before close")
	}
	door.openingAction = func(e *fsm.Event) {
		fmt.Println("opening")
	}
	door.closingAction = func(e *fsm.Event) {
		fmt.Println("closing")
	}
	door.afterOpenAction = func(e *fsm.Event) {
		fmt.Println("after open")
	}
	door.afterCloseAction = func(e *fsm.Event) {
		fmt.Println("after close")
	}
	door.callbacks = fsm.Callbacks{}
	door.callbacks["befor_open"] = door.beforeOpenAction
	door.callbacks["befor_close"] = door.beforeCloseAction
	door.callbacks["enter_open"] = door.openingAction
	door.callbacks["enter_closed"] = door.closingAction
	door.callbacks["after_open"] = door.afterOpenAction
	door.callbacks["after_close"] = door.afterCloseAction

	door.events = fsm.Events{
		{Name: "open", Src: []string{"closed"}, Dst: "open"},
		{Name: "close", Src: []string{"open"}, Dst: "closed"},
	}
	return door
}

func main() {
	door := initDoor()
	doorFSM := fsm.NewFSM(
		door.state,
		door.events,
		door.callbacks,
	)

	err := doorFSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = doorFSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}

}
