package main

import (
	"github.com/briansorahan/sc"
)

func main() {
	// this synthdef should map to
	// Name                           SineTone
	// NumConstants                   2
	// Constants
	//     0                          440
	//     1                          0
	// NumParams                      0
	// NumParamNames                  0
	// NumUgens                       2
	// NumVariants                    0

	// Ugen 0:
	// Name                           SinOsc
	// Rate                           2
	// NumInputs                      2
	// NumOutputs                     1
	// SpecialIndex                   0

	// Input 0:
	// UgenIndex                      -1
	// OutputIndex                    0

	// Input 1:
	// UgenIndex                      -1
	// OutputIndex                    1

	// Output 0:
	// Rate                           2

	// Ugen 1:
	// Name                           Out
	// Rate                           2
	// NumInputs                      2
	// NumOutputs                     0
	// SpecialIndex                   0

	// Input 0:
	// UgenIndex                      -1
	// OutputIndex                    1

	// Input 1:
	// UgenIndex                      0
	// OutputIndex                    0
	//
	sc.NewSynthDef("SineTone", func() {
		// how to handle params to ugen graph func?
		return sc.Ar("Out", 0, sc.Ar("SinOsc", 440))
	}).writeDefFile(os.Getcwd())
}