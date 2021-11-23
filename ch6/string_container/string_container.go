package stringcontainer

//字符串容器
type StringContainer struct {
	container string //存放的字符串
	byteArray []byte //byte形式的字符串
	index     int    //当前读取到的位置
}

func (sc *StringContainer) SetStr(str string) {
	sc.container = str
	sc.byteArray = []byte(str)
	sc.index = 0
}

//从字符串容器中读取一行数据
//param container 存放读取到的数据
//return length 读取到的字节数
//return err 错误
func (sc *StringContainer) Read(container []byte) (length int, err error) {
	for {
		if sc.index >= len(sc.byteArray) {
			return
		}
		char := sc.byteArray[sc.index]
		container = append(container, char)
		sc.index++
		length++
		if char == '\n' {
			return
		}
	}
}
