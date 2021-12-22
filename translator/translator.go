package translator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Translator struct {
	out          *os.File
	instructions []string
}

func New(inputPath string, outputPath string) *Translator {
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	dat, _ := ioutil.ReadFile(inputPath)

	return &Translator{out: f, instructions: strings.Split(string(dat), "\n")}
}

func (t *Translator) Translate() {
	pc := 0
	t.writeInit()
	for pc < len(t.instructions) {
		inst := t.instructions[pc]
		switch inst {
		case "push":
			nextInst := t.instructions[pc+1]
			t.writePush(nextInst)
			pc++

		case "add", "mul":
			if inst == "add" {
				t.writeAdd()
			} else {
				// não tem implementacao
			}

		case "pop":
			arg := t.instructions[pc+1]
			t.writePop(arg)
			pc++

		case "print":
			//
		}
		pc++
	}
	t.writeHalt()
}

func (t *Translator) write(s string) {
	t.out.WriteString(fmt.Sprintf("%s\n", s))

}

func (t *Translator) writePush(value string) {
	arg, err := strconv.Atoi(value)
	if err != nil { // é uma variavel
		t.write("@" + value)
		t.write("D=M")
		t.write("@SP")
		t.write("A=M")
		t.write("M=D")
		t.write("@SP")
		t.write("M=M+1")
	} else {
		t.write(fmt.Sprintf("@%d", arg))
		t.write("D=A")
		t.write("@SP")
		t.write("A=M")
		t.write("M=D")
		t.write("@SP")
		t.write("M=M+1")
	}

}

func (t *Translator) writeAdd() {
	t.write("@SP // add")
	t.write("AM=M-1")
	t.write("D=M")
	t.write("A=A-1")
	t.write("M=D+M")
}

func (t *Translator) writePop(varname string) {
	t.write("@SP")
	t.write("M=M-1")
	t.write("A=M")
	t.write("D=M")
	t.write("@" + varname)
	t.write("M=D")
}

func (t *Translator) writeInit() {
	t.write("@SP")
	t.write("@256")
	t.write("D=A")
	t.write("@SP")
	t.write("M=D")
}

func (t *Translator) writeHalt() {
	t.write("(LOOP)")
	t.write("@LOOP")
	t.write("0;JMP")
}
