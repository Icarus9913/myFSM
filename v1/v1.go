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
		cmd = strings.TrimSpace(cmd)
		switch state {
		case Locked:
			if cmd == CmdCoin {
				fmt.Println("解锁,请通行")
				state = Unlocked
			} else if cmd == CmdPush {
				fmt.Println("禁止通行,请先解锁")
			} else {
				fmt.Println("命令未知,请重新输入")
			}
		case Unlocked:
			if cmd == CmdCoin {
				fmt.Println("门开着的,别浪费钱了")
			} else if cmd == CmdPush {
				fmt.Println("请通行,通行之后将会关闭")
				state = Locked
			} else {
				fmt.Println("命令未知,请重新输入")
			}
		}
	}
}

func promt(s State) {
	m := map[State]string{
		Locked:   "Locked",
		Unlocked: "Unlocked",
	}
	fmt.Printf("当前的状态是 [%s], 请输入命令: [coin|push]\n", m[s])
}
