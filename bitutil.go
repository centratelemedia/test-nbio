package main

type ByteUtil struct {
}

func (b *ByteUtil) isSet(bin int, pos int) bool {
	return ((bin >> pos) & 1) != 0
}

/*func (b *ByteUtil) check(number int64, index int) bool {
	return (number & (1 << index)) != 0
}

func (b *ByteUtil) Between(number int, from int, to int) int {
	return (number >> from) & ((1<<to - from) - 1)
}

func (b *ByteUtil) From(number int, from int) int {
	return number >> from
}

func (b *ByteUtil) To(number int, to int) int {
	return b.Between(number, 0, to)
}

func (b *ByteUtil) Between(number int64, from int, to int) int64 {
	return (number>>from)&(1<<to-from) - 1
}

func (b *ByteUtil) From(number int64, from int) int64 {
	return number >> from
}

func (b *ByteUtil) To(number int64, to int) int64 {
	return b.Between(number, 0, to)
}*/

func BitisSet(bin int, pos int) bool {
	return ((bin >> pos) & 1) != 0
}
