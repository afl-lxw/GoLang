package toDo

//bufio 是 Buffered I/O缩写.
//由于net.Conn类型实现了接口类型io.Reader中的Read接口，所以该接口的类型的一个实现类型。
//因此，我们可以使用bufio.NewReader函数来包装变量conn,
//像这样： reader:=bufio.NewReader(conn)
//可以调用reader变量的ReadBytes("\n")来接受一个byte类型的参数.
//	当然很多情况下并不是查找一个但直接字符那么简单。
//比如,http协议中规定,消息头部的信息的末尾是两个空行,即是字符串"\r\n\r\n",
//writer:=bufio.NewWriter(conn)来写入数据

func TcpMain() {
	//var dataBuffer bytes.Buffer
	//b:=make([byte,10])
	//
	//for {
	//
	//n,err:=conn.Read(b)
	//if err!=nil{
	//if err==io.EOF{
	//fmt.Println("The connection is closed.")
	//conn.Close()
	//} else{
	//fmt.Printf("Read Error:%s\n",err)
	//}
	//break
	//}
	//dataBuffer.Write(b[:n])
	//
	//}

}
