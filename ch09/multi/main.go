package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
}

type myWriterCloser struct {
	MyWriter // interface也是一个类型 想放入一个写文件的实现 也可以放入一个写数据库的实现
}

func (w *writerCloser) Write(s string) error {
	fmt.Println("write string:" + s)
	return nil
}

func (w *writerCloser) Close() error {
	fmt.Println("close writerCloser")
	return nil
}

type fileWriter struct {
	filePath string
}

type dataWriter struct {
	port     int
	host     string
	username string
	password string
}

func (fw *fileWriter) Write(s string) error {
	fmt.Println("write string:" + s + " to file")
	return nil
}

func (dw *dataWriter) Write(s string) error {
	fmt.Println("write string:" + s + " to database")
	return nil
}

func main() {
	var w MyWriter = &writerCloser{}
	w.Write("hello")
	var c MyCloser = &writerCloser{}
	c.Close()

	var fw MyWriter = &myWriterCloser{
		&fileWriter{},
	}
	fw.Write("hello")
	var dw MyWriter = &myWriterCloser{
		&dataWriter{},
	}
	dw.Write("hello")
}
