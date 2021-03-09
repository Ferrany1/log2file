package logger

type FileLogger interface {
	FormatError(messageText string) (err error)

	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Print(v ...interface{})

	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatal(v ...interface{})

	Panicln(v ...interface{})
	Panicf(format string, v ...interface{})
	Panic(v ...interface{})
}
