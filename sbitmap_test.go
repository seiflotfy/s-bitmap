package sbitmap

import (
	"bufio"
	"math"
	"os"
	"testing"
)

func TestDefaultSbitmap(t *testing.T) {
	x := NewDefault()

	fd, err := os.Open("/usr/share/dict/words")
	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(fd)

	i := 0.0
	for scanner.Scan() {
		s := []byte(scanner.Text())
		x.Update(s)
		i++
	}

	errRate := math.Abs(1 - (x.Estimate() / i))

	if errRate > 0.01 {
		t.Error("Expected error rate <= 0.08, got", errRate)
	}
}
