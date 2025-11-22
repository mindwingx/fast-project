package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Std struct {
	reader *bufio.Reader
}

func NewStd() *Std {
	return &Std{reader: bufio.NewReader(os.Stdin)}
}

func (s *Std) Ask(question string, params ...any) string {
	q := "🎯  " + question + " "

	if len(params) > 0 {
		fmt.Printf(q, params...)
	} else {
		fmt.Print(q)
	}

	input, _ := s.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
