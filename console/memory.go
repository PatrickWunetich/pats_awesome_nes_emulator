package console

type Word uint8

type Memory []Word

func initMemory() *Memory {
	var memory = make(Memory, 0x10000)
	return &memory
}

func (mem *Memory) read(position uint16) Word {
	return (*mem)[position]
}
