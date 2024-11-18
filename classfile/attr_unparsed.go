package classfile

type UnParsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (self *UnParsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
