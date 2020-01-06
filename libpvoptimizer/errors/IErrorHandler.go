package errors

// IErrorHandler handles when errors occur
type IErrorHandler interface {
	// Handle the error
	Handle(e Error);
}

// NoImpl reports a not implemented error
func NoImpl(handler IErrorHandler, fn string) {
	handler.Handle(Error {
		Arguments: []string { fn },
		Code: NotImplemented,
		LineNo: -1,
		CharNo: -1,
		FileName: "",
	})
}
