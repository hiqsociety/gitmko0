package gitmko0

func GetTimeInSeconds() uint64 {
     return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}