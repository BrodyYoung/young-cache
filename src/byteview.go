package src

type ByteView struct {
	bs []byte
}

func (bv ByteView) String() string {
	return string(bv.bs)
}

func (bv ByteView) CloneBytes() []byte {

	bytes := make([]byte, len(bv.bs))

	copy(bytes, bv.bs)
	return bytes
}

func (bv ByteView) Len() int {
	return len(bv.bs)
}
