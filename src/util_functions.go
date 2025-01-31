package main

import (
	"crypto/sha256"
	"strings"
)

// ComputeMainBlockHash erzeugt einen Hash aus allen Statements,
// die **kein** ChCallStatement sind, d. h. wir ignorieren chcall-Blöcke.
func ComputeMainBlockHash(prog *Program) string {
	var sb strings.Builder
	for _, s := range prog.Statements {
		// Wenn es ein *ChCallStatement ist, überspringen wir es.
		if _, ok := s.(*ChCallStatement); ok {
			continue
		}
		sb.WriteString(s.String())
		sb.WriteRune('\n')
	}
	// Jetzt berechnen wir den SHA-256 über sb.String()
	data := sb.String()
	sum := sha256.Sum256([]byte(data))
	return hexEncode(sum[:])
}

func hexEncode(b []byte) string {
	const hexDigits = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[i*2] = hexDigits[v>>4]
		out[i*2+1] = hexDigits[v&0x0f]
	}
	return string(out)
}
