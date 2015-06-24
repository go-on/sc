package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestFSinOsc(t *testing.T) {
	def := NewSynthdef("FSinOscExample", func(p *Params) Ugen {
		line := XLine{C(4), C(401), C(8), 0}.Rate(KR)
		sin1 := FSinOsc{line, C(0)}.Rate(AR).MulAdd(C(200), C(800))
		sin2 := FSinOsc{Freq:sin1}.Rate(AR).Mul(C(0.2))
		bus := C(0)
		return Out{bus, sin2}.Rate(AR)
	})
	compareAndWrite(t, "FSinOscExample", def)
}