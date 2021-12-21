
@10 // push constant 10
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP // pop static 0
M=M-1
A=M
D=M
@x
M=D
@x // push static 0
D=M
@SP
A=M
M=D
@SP
M=M+1
@5 // push constant 5
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP // add
AM=M-1
D=M
A=A-1
M=D+M
@SP // pop static 1
M=M-1
A=M
D=M
@y
M=D