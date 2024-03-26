package logger

import "testing"

func TestToggleColorfulLogger(t *testing.T) {
	ToggleColorfulLogger()
	if !useLogurs {
		t.Errorf("Expected useLogurs to be true, got false")
	}

	if logrusLogger == nil {
		t.Errorf("Expected logrusLogger to be not nil")
	}

	ToggleColorfulLogger()
	if useLogurs {
		t.Errorf("Expected useLogurs to be false, got true")
	}

	if logrusLogger != nil {
		t.Errorf("Expected logrusLogger to be nil")
	}
}
