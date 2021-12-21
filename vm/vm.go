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

func writePush(value string) {
	arg, err := strconv.Atoi(value)
	if err != nil { // é uma variavel
		fmt.Println("@" + value)
		fmt.Println("D=M")
		fmt.Println("@SP")
		fmt.Println("A=M")
		fmt.Println("M=D")
		fmt.Println("@SP")
		fmt.Println("M=M+1")
	} else {
		fmt.Println(fmt.Sprintf("@%d", arg))
		fmt.Println("D=A")
		fmt.Println("@SP")
		fmt.Println("A=M")
		fmt.Println("M=D")
		fmt.Println("@SP")
		fmt.Println("M=M+1")
	}

}

func writeAdd() {
	fmt.Println("@SP // add")
	fmt.Println("AM=M-1")
	fmt.Println("D=M")
	fmt.Println("A=A-1")
	fmt.Println("M=D+M")

}

func writePop(varname string) {
	fmt.Println("@SP")
	fmt.Println("M=M-1")
	fmt.Println("A=M")
	fmt.Println("D=M")
	fmt.Println("@" + varname)
	fmt.Println("M=D")
}

func writeInit() {
	fmt.Println("@SP")
	fmt.Println("@256")
	fmt.Println("D=A")
	fmt.Println("@SP")
	fmt.Println("M=D")
}

func writeHalt() {
	fmt.Println("(LOOP)")
	fmt.Println("@LOOP")
	fmt.Println("0;JMP")
}

func Translator(path string) {
	dat, _ := ioutil.ReadFile(path)
	instructions := strings.Split(string(dat), "\n")
	vm := New(instructions)
	writeInit()
	for vm.pc < len(vm.instructions) {
		inst := vm.instructions[vm.pc]
		switch inst {
		case "push":
			nextInst := vm.instructions[vm.pc+1]
			writePush(nextInst)
			vm.pc++

		case "add", "mul":
			if inst == "add" {
				writeAdd()
			} else {
				// não tem implementacao
			}

		case "pop":
			arg := vm.instructions[vm.pc+1]
			writePop(arg)
			vm.pc++

		case "print":
			//
		}
		vm.pc++
	}
	writeHalt()
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
		/*
			fmt.Print("Stack:")
			fmt.Println(vm.stack)
			fmt.Print("Symbol table:")
			fmt.Println(vm.st)
		*/
	}
}
