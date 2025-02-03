package console

type Cpu struct {
	memory    *Memory
	registers *Registers
}

type Registers struct {
	A  uint8
	X  uint8
	Y  uint8
	PC uint16
	S  uint8
	P  uint8
}

func initCpu() *Cpu {
	var cpu = Cpu{memory: initMemory(), registers: &Registers{}}
	return &cpu
}

// return cycles the cpu op took
func (cpu *Cpu) step() int {

	var op = cpu.memory.read(cpu.registers.PC)
	cpu.registers.PC++
	return 0
}
