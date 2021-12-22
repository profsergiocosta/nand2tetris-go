package vm

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

func Open(filename string) *VM {
	dat, _ := ioutil.ReadFile(filename)
	instructions := strings.Split(string(dat), "\n")
	return New(instructions)
}

func (vm *VM) stackPop() int {
	sp := len(vm.stack)
	value := vm.stack[sp-1]
	vm.stack = vm.stack[:sp-1]
	return value
}

func (vm *VM) stackPush(value int) {
	vm.stack = append(vm.stack, value)
}

func (vm *VM) Save(outputPath string) {
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)

	for _, inst := range vm.instructions {
		w.WriteString(fmt.Sprintf("%s\n", inst))
	}
	w.Flush()
	f.Close()
}

func (vm *VM) Run() {
	for vm.pc < len(vm.instructions) {
		inst := vm.instructions[vm.pc]
		switch inst {
		case "push":
			nextInst := vm.instructions[vm.pc+1]
			arg, err := strconv.Atoi(nextInst)
			if err != nil { // é uma variavel
				arg = vm.st[nextInst]
			}
			vm.stackPush(arg)
			vm.pc++

		case "add", "mul":
			arg1 := vm.stackPop()
			arg2 := vm.stackPop()
			if inst == "add" {
				vm.stackPush(arg1 + arg2)
			} else {
				vm.stackPush(arg1 * arg2)
			}

		case "pop":
			arg := vm.instructions[vm.pc+1]
			vm.st[arg] = vm.stackPop()
			vm.pc++

		case "print":
			fmt.Println(vm.stackPop())
		}
		vm.pc++
	}
}
