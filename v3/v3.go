package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//旋转门状态
type State uint32

const (
	Locked State = iota
	Unlocked
)

//相关命令
const (
	CmdCoin = "coin"
	CmdPush = "push"
)

func main() {
	state := Locked
	promt(state)
	reader := bufio.NewReader(os.Stdin)

	for {
		cmd, err := reader.ReadString('\n')
		if nil != err {
			log.Fatal(err)
		}

		//获取状态转换表中的值
		tupple := CommandStateTupple{Command: strings.TrimSpace(cmd), State: state}
		if f := StateTransitionTable[tupple]; nil == f {
			fmt.Println("未知命令,请重新输入")
		} else {
			f(&state)
		}
	}
}

//CommandStateTupple 用于存放状态转换表的结构体
type CommandStateTupple struct {
	Command string
	State   State
}

//TransitionFunc 状态转移方程
type TransitionFunc func(state *State)

//StateTransitionTable 状态转换表
var StateTransitionTable = map[CommandStateTupple]TransitionFunc{
	{Command: CmdCoin, State: Locked}: func(state *State) {
		fmt.Println("已解锁,请通行")
		*state = Unlocked
	},
	{Command: CmdPush, State: Locked}: func(state *State) {
		fmt.Println("禁止通行,请先解锁")
	},
	{Command: CmdCoin, State: Unlocked}: func(state *State) {
		fmt.Println("兄dei,已经解锁了,别浪费钱了")
	},
	{Command: CmdPush, State: Unlocked}: func(state *State) {
		fmt.Println("请尽快通行,通行后将自动上锁")
		*state = Locked
	},
}

func promt(s State) {
	m := map[State]string{
		Locked:   "Locked",
		Unlocked: "Unlocked",
	}
	fmt.Printf("当前的状态是 [%s], 请输入命令: [coin|push]\n", m[s])
}
