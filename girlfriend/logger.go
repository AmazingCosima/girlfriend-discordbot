package girlfriend

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {

}

func (logger *Logger) Log(text string) {
	var timestamp *strings.Builder
	timestamp = new(strings.Builder)
	snapshot := time.Now()
	timestamp.WriteString(snapshot.Format("2006"))
	timestamp.WriteString("-")
	timestamp.WriteString(snapshot.Format("01"))
	timestamp.WriteString("-")
	timestamp.WriteString(snapshot.Format("02"))
	timestamp.WriteString(" ")
	timestamp.WriteString(snapshot.Format("15"))
	timestamp.WriteString(":")
	timestamp.WriteString(snapshot.Format("04"))
	timestamp.WriteString(":")
	timestamp.WriteString(snapshot.Format("05.000"))
	timestamp.WriteString(" -> ")
	fmt.Println(timestamp.String() + text)
}
