package main

import "fmt"
import "github.com/looplab/fsm"

func main() {

	fsm := fsm.NewFSM(
		"green",
		fsm.Events{
			{Name: "warn", Src: []string{"green"}, Dst: "yellow"},
			{Name: "panic", Src: []string{"yellow"}, Dst: "red"},
			{Name: "panic", Src: []string{"green"}, Dst: "red"},
			{Name: "calm", Src: []string{"red"}, Dst: "yellow"},
			{Name: "clear", Src: []string{"yellow"}, Dst: "green"},
		},
		// 状态装换一次调用的回调. 状态转换一次 执行一次
		fsm.Callbacks{
			"before_warn": func(e *fsm.Event) {
				fmt.Println("before_warn event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"before_event": func(e *fsm.Event) {
				fmt.Println("before_event event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"leave_green": func(e *fsm.Event) {
				fmt.Println("leave_green event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"leave_state": func(e *fsm.Event) {
				fmt.Println("leave_state event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"enter_yellow": func(e *fsm.Event) {
				fmt.Println("enter_yellow event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"enter_state": func(e *fsm.Event) {
				fmt.Println("enter_state event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"after_warn": func(e *fsm.Event) {
				fmt.Println("after_warn event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
			"before_yellow": func(e *fsm.Event) {
				fmt.Println("before_yellow event=%v,src=%v,dst=%v", e.Event, e.Src, e.Dst)
			},
		},
	)
	fmt.Println(fsm.Current())
	err := fsm.Event("warn")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fsm.Current())
	err = fsm.Event("panic")
	fmt.Println(fsm.Current())

}
