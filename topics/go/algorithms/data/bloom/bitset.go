package bloom

type BitSet struct {
	size  int
	bytes []byte
}

func NewBitSet(size int) *BitSet {
	nBytes := (size + 7) / 8
	bs := BitSet{
		size:  size,
		bytes: make([]byte, nBytes),
	}
	return &bs
}

func (bs *BitSet) Get(i int) int {
	val := bs.bytes[i/8]
	bitNum := i % 8
	return int((val >> bitNum) & 1)
}

func (bs *BitSet) Set(i, val int) {
	byteNum := i / 8
	bitNum := i % 8
	if val == 1 {
		bs.bytes[byteNum] |= (1 << bitNum)
	} else {
		bs.bytes[byteNum] &= ^(1 << bitNum)
	}
}
