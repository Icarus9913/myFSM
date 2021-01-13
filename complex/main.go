package main

import (
	"fmt"
)

//接口
type IFSMState interface {
	Enter()
	Exit()
	CheckTransition(hour int) bool
	Hour() int
}

//state父struct
type FSMState struct{}

//进入状态
func (this *FSMState) Enter() {}

//退出状态
func (this *FSMState) Exit() {}

//状态转移检测
func (this *FSMState) CheckTransaction(hour int) {}

//打坐
type ZazenState struct {
	hour int
	FSMState
}

func NewZazenState() *ZazenState {
	return &ZazenState{hour: 8}
}

func (this *ZazenState) Enter() {
	fmt.Println("开始打坐")
}

func (this *ZazenState) Exit() {
	fmt.Println("退出打坐")
}

func (this *ZazenState) Hour() int {
	return this.hour
}

//状态转移检测
func (this *ZazenState) CheckTransition(hour int) bool {
	if hour == this.hour {
		return true
	}
	return false
}

//工作
type WorkState struct {
	hour int
	FSMState
}

func NewWokerState() *WorkState {
	return &WorkState{hour: 12}
}

func (this *WorkState) Enter() {
	fmt.Println("WorkState: 开始工作")
}

func (this *WorkState) Exit() {
	fmt.Println("WorkState: 退出工作")
}

func (this *WorkState) Hour() int {
	return this.hour
}

//状态转移检测
func (this *WorkState) CheckTransition(hour int) bool {
	if hour == this.hour {
		return true
	}
	return false
}

type FSM struct {
	//持有状态集合
	states map[string]IFSMState
	//当前状态
	current_state IFSMState
	//默认状态
	default_state IFSMState
	//外部输入数据
	input_date int
	//是否初始化
	inited bool
}

//初始化FSM
func (this *FSM) Init() {
	this.Reset()
}

//重置
func (this *FSM) Reset() {
	this.inited = false
}

//添加状态到FSM
func (this *FSM) AddState(key string, state IFSMState) {
	if this.states == nil {
		this.states = make(map[string]IFSMState, 2)
	}
	this.states[key] = state
}

//设置默认的state
func (this *FSM) SetDefaultState(state IFSMState) {
	this.default_state = state
}

//设置输入数据
func (this *FSM) SetInputData(inputData int) {
	this.input_date = inputData
	this.TransitionState()
}

//转移状态
func (this *FSM) TransitionState() {
	nextState := this.default_state
	input_data := this.input_date
	if this.inited {
		for _, v := range this.states {
			if input_data == v.Hour() {
				nextState = v
				break
			}
		}
	}
	if ok := nextState.CheckTransition(this.input_date); ok {
		if this.current_state != nil {
			//退出当前一个状态
			this.current_state.Exit()
		}
		this.current_state = nextState
		this.inited = true
		nextState.Enter()
	}
}



func main() {
	zazenState := NewZazenState()
	wokerState := NewWokerState()
	fsm := new(FSM)
	fsm.AddState("ZazenState", zazenState)
	fsm.AddState("WorkerState", wokerState)
	fsm.SetDefaultState(zazenState)
	fsm.Init()
	fsm.SetInputData(8)
	fsm.SetInputData(12)
	fsm.SetInputData(12)
	fsm.SetInputData(8)
	fsm.SetInputData(12)
}
