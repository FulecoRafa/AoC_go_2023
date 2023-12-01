package lib

func MayPanicBool(ok bool) {
    if !ok {
        panic("Not ok")
    }
}

func MayPanic(err error) {
    if err != nil {
        panic(err)
    }
}
