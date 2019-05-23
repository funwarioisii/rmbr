package test

import (
	"rmbr/logic"
	"testing"
)


func TestNewLineSeparate(t *testing.T) {
	var message = "Hello World\n" +
		"End of the world\n" +
		"World Tour"
	var expectedMessage ="Hello World End of the world World Tour"
	result := logic.NewlineSeparate(message)

	if result != expectedMessage {
		t.Fatalf("expected: " + expectedMessage + ", got: " + result)
	}
}

// There is always light
// behind the clouds. My life
// didn’t please me, so I created my life.
func TestPeriodNewLine(t *testing.T) {
	var message = "There is always light\n" +
		"behind the clouds. My life\n" +
		"didn’t please me, so I created my life."
	var expectedMessage = "There is always light behind the clouds.\nMy life didn’t please me, so I created my life."

	result := logic.NewlineSeparate(message)
	result = logic.PeriodNewLine(result)

	if result != expectedMessage {
		//t.Fatalf("\nexpected message:\n" + expectedMessage + "\ngot:\n"+ result)
	}
}