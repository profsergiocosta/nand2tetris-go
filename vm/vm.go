package vm

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type VM struct {
	pc           int
	stack        []int
	st           map[string]int
	instructions []string
}

func New(inst []string) *VM {

	return &VM{pc: 0, stack: nil, st: make(map[string]int), instructions: inst}
}

func Interpret(path string) {
	dat, _ := ioutil.ReadFile(path)
	instructions := strings.Split(string(dat), "\n")
	vm := New(instructions)
	vm.Run()
}

func (vm *VM) Run() {
	for vm.pc < len(vm.instructions) {
		inst := vm.instructions[vm.pc]
		switch inst {
		case "push":
			nextInst := vm.instructions[vm.pc+1]
			arg, err := strconv.Atoi(nextInst)
			if err != nil {
				arg = vm.st[nextInst]
			}
			vm.stack = append(vm.stack, arg)
			vm.pc++
		case "add", "mul":
			sp := len(vm.stack)
			arg1 := vm.stack[sp-1]
			arg2 := vm.stack[sp-2]
			var res int
			if inst == "add" {
				res = arg1 + arg2
			} else {
				res = arg1 * arg2
			}
			vm.stack = vm.stack[:sp-2]
			vm.stack = append(vm.stack, res)

		case "pop":
			arg1 := vm.instructions[vm.pc+1]
			sp := len(vm.stack)
			arg2 := vm.stack[sp-1]
			vm.st[arg1] = arg2
			vm.stack = vm.stack[:sp-1]
			vm.pc++

		}
		vm.pc++
		fmt.Print("Stack:")
		fmt.Println(vm.stack)
		fmt.Print("Symbol table:")
		fmt.Println(vm.st)
	}
}
