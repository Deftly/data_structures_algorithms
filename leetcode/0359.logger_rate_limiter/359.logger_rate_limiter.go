package data_structures_algorithms

type Logger struct {
	message map[string]int
}

func Constructor() Logger {
	return Logger{
		message: make(map[string]int),
	}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if ts, ok := l.message[message]; ok && timestamp < ts {
		return false
	}
	l.message[message] = timestamp + 10
	return true
}
