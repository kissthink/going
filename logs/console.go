package logs

import (
	"log"
	"os"
)

//30 black		黑色

//31 red		红色
//32 green		绿色
//33 yellow		黄色
//34 blue		蓝色
//35 magenta	洋红
//36 cyan		蓝绿色
//37 white		白色


//LevelDebug 	 	//= "Debug"		白色		37
//LevelInfo 		//= "Info"		蓝绿色	36
//LevelWarn 		//= "Warn"    	洋红  	35
//LevelPanic		//= "Panic"   	蓝色  	34
//LevelFatal		//= "Fatal"   	红色  	31


type color func(string) string

func newColor(c string) color {
	return func(t string) string {
		return "\033[1;"+ c +"m" + t + "\033[0m"
	}
}

var colors = []color{
	newColor("37"),
	newColor("36"),
	newColor("35"),
	newColor("34"),
	newColor("31"),
}

type Console struct {
	logger *log.Logger
	level int
}

func NewConsole() *Console {
	var console = &Console{}
	console.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return console
}

func (this *Console)WriteMessage(level int, file string, line int, prefix string, msg string) {
	if level > this.level || level > len(colors) {
		return
	}
	this.logger.Printf("%s[%d] %s %s", file, line, colors[level](prefix), msg)
}