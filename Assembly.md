// Lab 1
// Activity 1
//Immediate Mode
```
;ORG 0000H
;	MOV A,#52H
;	MOV R0,#26H
;	MOV R2,#32
;	MOV DPTR,#1234H
;	END
```

// Register Mode
```
;ORG 0000H
;	MOV R0,4H
;	MOV R2,3H
;	MOV A,R0
;	ADD A,R2
;	MOV R6,A
;	MOV R3,DPH
;	MOV R4,DPL
;	END
```

// Direct Addresing Mode
```
;ORG 0000H
;	MOV 33H,#35H
;	END
```

// Activity 2
// Direct Addressing Mode
```
;ORG 0000H
;	MOV 40H,#44H
;	MOV 41H,#44H
;	MOV 42H,#44H
;	MOV 43H,#44H
;	END
```

// Register Indirect Addressing Mode
;ORG 0000H
;	MOV R0,#40H
;	MOV R1,#4
;	MOV A,#44H
;	COPY:
;	MOV @R0,A
;	INC R0
;	DJNZ R1,COPY
;	END
		
//1 Activity 3
```
;ORG 0000H
;	MOV DPTR,#200H
;	MOVC A,@A+DPTR
;	MOV R1,A
;	CLR A
;	INC DPTR
;	MOVC A,@A+DPTR
;	MOV R2,A
;	CLR A
;	MOVC A,@DPTR
;	END
```

//LAB 2
//2 1 To write an ALP to transfer data from code ROM space into RAM
```
;ORG 200H
;	DB "GO INDIA"
;ORG 00H
;	MOV R0,#40H
;	MOV R1,#08H
;	MOV DPTR,#200H
;	BACK:
;	MOVC A,@A+DPTR
;	MOV @R0,A
;	INC DPTR
;	INC R0
;	CLR A
;	DJNZ R1,BACK
;END
```

//2 2 To write an ALP to EXCHANGE
```
;ORG 00H
;	MOV 40H,#5H
;	MOV 41H,#6H
;	MOV 42H,#7H
;	MOV 43H,#8H
;	MOV 44H,#9H
;	
;	MOV 60H,#15H
;	MOV 61H,#16H
;	MOV 62H,#17H
;	MOV 63H,#18H
;	MOV 64H,#19H
;	
;	MOV R0,#40H
;	MOV R1,#60H
;	MOV R2,#05H
;	
;	LOOP:
;	MOV A,@R0
;	MOV B,@R1
;	MOV @R1,A
;	MOV @R0,B
;	INC R0
;	INC R1
;	DJNZ R2,LOOP
;	END
```

//2 3 To write an ALP to exchange data from internal RAM location
```
;ORG 00H
;	MOV 40H,#9H
;	MOV 41H,#5H
;	MOV 42H,#1H
;	MOV 43H,#15H
;	MOV 44H,#10H
;	MOV R0,#40H
;	MOV R1,#5H
;	MOV B,#0H
;	COMP:
;	MOV A,@R0
;	CJNE A,B,LOOP
;	LOOP:
;	JC LOOP1
;	MOV B,A
;	LOOP1:
;	INC R0
;	DJNZ R1,COMP
;	MOV P2,B
;	END
```

//2 4 To write an ALP to find the largest element in an array sorted in the internal RAM locations & display the result in port2
```
ORG 00H
	MOV 40H,#03H
	MOV 41H,#01H
	MOV 42H,#05H
	MOV 43H,#02H
	MOV 44H,#04H
	MOV R1,#04H
	AGAIN:
	MOV A,R1
	MOV R2,A
	MOV R0,#40H
	MOV A,@R0
	UP:
	INC R0
	MOV B,@R0
	CLR C
	SUBB A,B
	JC SKIP
	MOV B,@R0
	DEC R0
	MOV A,@R0
	MOV @R0,B
	INC R0
	MOV @R0,A
	SKIP:
	DJNZ R2,UP
	DJNZ R1,AGAIN
	END
```

//2 5 LookUp Table
```
;ORG 00H
;	MOV DPTR,#200H
;	MOV A,#05H
;	MOV R1,A
;	MOV R0,A
;	ADDC A,R1
;	MOV R1,A
;	MOV A,R0
;	MOVC A,@A+DPTR
;	ADDC A,#04H
;	ADDC A,R1
;	MOV R2,A
;ORG 200H
;	DB 00H,01H,04H,09H,10,19H,24H,31H,40H,51H
;END
```

//LAB 3
//3 1 ADD HEX NUMBERS
```
;org 0000H
;mov R0,#40H
;mov R1,#0AH
;mov R2,#00H
;mov R3, #00H
;clr A
;Addition:
;Add A,@R0
;jnc down
;inc R2
;down: inc R0
;djnz R1,Addition
;mov R3,A
;end
```

//3 2 ADD BCD NUMBERS
```
;org 0000H
;mov R0,#60H
;mov R1,#0AH
;mov R2,#00H  //Higher Byte
;mov R3, #00H
;clr A
;Addition:
;Add A,@R0   
;da A
;jnc up
;inc R2
;up: inc r0
;djnz R1,Addition
;mov R3,A	//Lower Byte
;end  // Don't enter Alphabates at i:60h
```

//3 3 ADD MULTI-BYTE BCD NUMBERS
```
;org 0000h
;mov r0,#60h
;mov r1,#40h
;mov r2,#04h
;func:
;mov a,@r0
;inc r0
;addc a,@r0
;inc r0
;mov @r1,a
;inc r1
;djnz r2,func
;end
```

//3 4 EXAMINE MUL & DIV FUNCTION
; Multiplication
```
;org 0000H
;mov A,#4H
;mov B,#5H
;mul AB
;mov r2,a
;mov A,#60H
;mov B,#20H
;div AB
;mov r3,a
;end
```

// LAB 4
//EXAMINE I/O PORT OPERATION USING A SIMULATOR
//TRACE THROUGH A CALL SUBROUTINE USING A SIMULATOR
//EXAMINE THE FLAG BITS OF THE PSW & EXAMINE THE STACK
//4 1
```
;ORG 0000H
;	HERE: MOV P0, #55H
;	MOV P1, #55H
;	MOV P2, #55H
;	ACALL DELAY
;	MOV P0, #0AAH
;	MOV P1, #0AAH
;	MOV P2, #0AAH
;	ACALL DELAY
;	SJMP HERE
;	DELAY : MOV R1, #300
;	BACK : MOV R2, #200
;	AGAIN : DJNZ R2, AGAIN
;	DJNZ R1, BACK
;	RET
;END
```

//4 2
```
ORG 0
MOV A, P1 ; get a byte of data from P1
MOV P0, A ; send value in accumulator to P0
MOV P2, A ; send value in accumulator to P2
MOV R0, A ; copy value in accumulator to R0
MOV R1, A ; copy value in accumulator to R1
MOV R2, A ; copy value in accumulator to R2
```

;END
	
//4 3
```
ORG 00H
	MOV R0, #00H
	MOV A, #92H
	ADD A, #23H
	JNC L1
	INC R0
	L1: ADD A, #66H
	JNC L2
	INC R0
	L2: ADD A, #87H
	JNC L3
	INC R0
	L3: ADD A, #0F5H
	JNC L4
	INC R0
	L4: MOV B,A
	END
```

//4 4
```
ORG 00H
	MOV R0, #15H
	MOV R1, #45H
	MOV R2, #30H
	MOV R3, #27H
	MOV R4, #52H
	PUSH 0
	PUSH 1
	PUSH 2
	PUSH 3
	PUSH 4
	END
```

//4 5
```
ORG 00H
	MOV SP, #0DH
	MOV 0DH, #15H
	MOV 0CH, #17H
	MOV 0BH, #10H
	MOV 0AH, #23H
	MOV 09H, #24H
	MOV 08H, #25H
	
	POP 0
	POP 1
	POP 2
	POP 3
	POP 4
	MOV A, SP
	END
```

//LAB 5
//HEX TO ASCII
//AVERAGE OF A SET OF HEX DATA
//5 1
```
ORG 0000H
	ACALL HEXDEC
	ACALL DECASCII
	HERE: SJMP HERE
	HEXDEC:
	MOV R3, #0FBH
	MOV A,R3
	MOV R0, #30H
	MOV B, #0AH
	DIV AB
	MOV @R0, B
	MOV B, #0AH
	DIV AB
	INC R0
	MOV @R0, B
	INC R0
	MOV @R0, A
	RET
	DECASCII:
	MOV R0, #30H
	MOV R1, #40H
	MOV R2, #04H
	REP: MOV A, @R0
	
	ORL A, #30H
	MOV @R1, A
	INC R0
	INC R1
	DJNZ R2, REP
END
```

//5 2
```
ORG 150H
MYDATA: DB 7FH,3CH,54H,2AH
ORG 00H
MOV DPTR, #150H
MOV R0, #5BH
MOV R2, #4
LOOP: CLR A
MOVC A, @A+DPTR
MOV R1, #3
LOOP1: MOV B, #10
DIV AB
XCH A, B
ADD A, #30H
MOV @R0, A
XCH A, B
DEC R0
DJNZ R1, LOOP1
INC DPTR
djnz r2, loop
END
```	
//5 3
```
ORG 0000H
	MOV DPTR, #200H;
	MOV R0, #60H;
	MOV R7, #12;
	NEXT:
		MOVC A, @A+DPTR;
		MOV @R0, A;
		INC DPTR;
		INC R0;
		CLR A;
	DJNZ R7, NEXT;
	CLR A;
	
	MOV R0, #12;
	MOV R1, #60H;
	MOV A, #00H;
	L1:ADDC A, @R1;
	INC R1;
	DJNZ R0, L1;
	MOV B, #12;
	DIV AB;
	
	ORG 200H
		DB 5H, 6H, 7H, 8H, 9H, 5H, 0AH, 6H, 7H, 9H, 0AH, 1H
END
ORG 00H
SETB P1.0
LCALL DELAY
CLR P1.0
LCALL DELAY
END
```

//5 4
```
ORG 120H
MYDATA: DB '7','5','4','9','3','4','5'
ORG 00H
MOV DPTR, #120H
MOV R0, #40H
MOV R1, #07H
LOOP:
CLR A
MOVC A, @A+DPTR
SUBB A, #30H
MOV @R0, A
INC DPTR
INC R0
DJNZ R1,LOOP
CLR A
MOV R0,#40H
MOV R1,#07H
LOOP1:
ADD A,@R0
INC R0
DJNZ R1,LOOP1
MOV B, #07H
DIV AB
END
```

//LAB 6
//6 1
```
ORG 0000H;
	MOV TMOD, #01;
	MOV TL0, #0F2H;
	HERE:	MOV TH0, #0FFH;
	CPL P2.3;
	ACALL DELAY;
	SJMP HERE;
	DELAY:
	SETB TR0
	AGAIN: JNB TF0, AGAIN
	CLR TR0
	CLR TF0
	RET
	END
```
		
//6 2
```
ORG 0000H
	CLR P1.1
	MOV TMOD,#1
    AGAIN: MOV TL0,#99H
    MOV TH0,#19H
    SETB TR0
    BACK: JNB TF0,BACK
	CLR TF0
    CPL P1.1
    SJMP AGAIN
	END
```
	
//6 3
```
ORG 0000H
MOV TM0D,#01100000B 
MOV TH1,#0
SETB P3.5
AGAIN: SETB TR1
BACK: MOV A,TL1
MOV P2,A
JNB TF1,Back
CLR TR1
CLR TF1
SJMP AGAIN
END
```
	
//LAB 7
//7 1
```
ORG 0000h
	Mov TMOD,#20H
	MOV TH1,#-6
	MOV SCON,#50H
	SETB TR1
	AGAIN:
	MOV SBUF,#"A"
	HERE:
	JNB TI,HERE
	CLR TI
	SJMP AGAIN
;END
```

//7 2
```
ORG 0000h
	MOV    A, PCON
	SETB   ACC.7
	MOV    PCON,A
	Mov TMOD,#20H
	MOV TH1,#-3
	MOV SCON,#50H
	SETB TR1
	
	DAT:
	MOV A,#"N"
	ACALL AGAIN
	MOV A,#"I"
	ACALL AGAIN
	MOV A,#"T"
	ACALL AGAIN
	MOV A,#"R"
	ACALL AGAIN
	SJMP DAT
	AGAIN:
	MOV SBUF,A
	HERE:
	JNB TI,HERE
	CLR TI
	RET
END
```

//LAB 8
;Program :
;************************************************************************
;: Objective : Interfacing 16X2LCD with AT89S8253 in 8-bit mode.  
;: Program : To display text on the LCD screen.  
;: Compiler : KEIL uvision3  
;: Development Board : EASY 8051B, 16X2 LCD.  
;;***********************************************************************
```
org 0000h
mov a,#38h
acall comm
acall delay
mov a,#0eh
acall comm
acall delay
mov a,#01h
acall comm
acall delay
mov a,#06h
acall comm
acall delay
mov a,#80h
acall comm
acall delay
mov a,#'B'
acall data1
acall delay
mov a,#'I'
acall data1
acall delay
mov a,#'S'
acall data1
acall delay
mov a,#'W'
acall data1
acall delay
mov a,#'A'
acall data1
acall delay
comm: mov p0,a
clr p2.2
clr p2.3
setb p2.4
clr p2.4
ret
data1: mov p0,a
setb p2.2
clr p2.3
setb p2.4
clr p2.4
ret
delay: mov r1,#255
ag2:mov r2,#255
ag1:djnz r2,ag1
djnz r1,ag2
ret
end
```
;NOTE: 16x2 LCD should be placed on the connector available for GLCD. While  
;connecting 2x16 lcd just leave two pins from the MSB side and two pins from LSB  
;side place the lcd in the middle position.  

Lab 9 C code

//Lab 10  
// rolling 7 Segment display

Program :
***********************************************************************  
: Objective : Interfacing of Seven Segment Display with AT89S8253  
: Program : Write a program for rolling the display 1234 on the Seven  
Segment Display through right entry mode.  
: Compiler : KEIL uvision3  
: Development Board : EASY 8051B  
NOTE: Switch on SW2.5 (i.e. P1.3) to SW2.8 (i.e. P1.0) on Easy 8051B board.  
```
org 0000h
start: setb p1.0
clr p1.1
clr p1.2
clr p1.3
mov p0,#0f9h
call delay
setb p1.1
clr p1.0
clr p1.2
clr p1.3
mov p0,#0a4h
call delay
setb p1.2
clr p1.0
clr p1.1
clr p1.3
mov p0,#0b0h
call delay
setb p1.3
clr p1.0
clr p1.1
clr p1.2
mov p0,#99h
call delay
sjmp start
delay:
mov r2,#0ffh
load:
mov r1,#0ffh
here: djnz r1,here
djnz r2,load
ret
end
```


Program :
***********************************************************************
: Objective : Interfacing of Seven Segment Display with AT89S8253
: Program : Write a program for rolling the display 1234 on the Seven
Segment Display through right entry mode using push button
P1.0
: Compiler : KEIL uvision3
: Development Board : EASY 8051B
***********************************************************************
```
org 0000h
here1: jb p1.0,here1
start:
setb p1.0
clr p1.1
clr p1.2
clr p1.3
mov p0,#0f9h
call delay
setb p1.1
clr p1.0
clr p1.2
clr p1.3
mov p0,#0a4h
call delay
setb p1.2
clr p1.0
clr p1.1
clr p1.3
mov p0,#0b0h
call delay
setb p1.3
clr p1.0
clr p1.1
clr p1.2
mov p0,#99h
call delay
sjmp start
delay:
mov r2,#0ffh
load:
mov r1,#0ffh
here:
djnz r1,here
djnz r2,load
ret

```	









	
		
