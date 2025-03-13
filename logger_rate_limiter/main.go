package main

func main() {
	obj := Constructor()
	println(obj.ShouldPrintMessage(1, "foo")) //true
	println(obj.ShouldPrintMessage(2, "bar")) //true
	println(obj.ShouldPrintMessage(3, "foo")) //false
}

type Logger struct {
	unit  map[string]int
	delay int
}

func Constructor() Logger {
	return Logger{
		unit:  make(map[string]int),
		delay: 10,
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if t, ok := this.unit[message]; !ok {
		this.unit[message] = timestamp + this.delay
		return true
	} else {
		if timestamp >= t {
			this.unit[message] = timestamp + this.delay
			return true
		} else {
			return false
		}
	}
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
