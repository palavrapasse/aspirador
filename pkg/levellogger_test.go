package pkg

import (
	"os"
	"testing"
)

func TestNewLevelLoggerWithoutLevelsReturnsLevelLoggerWithLenght4(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{})
	size := len(result)

	if size != 4 {
		t.Fatalf("No levels provided so it should return a map with 4 elements, but it has %d", size)
	}
}

func TestNewLevelLoggerWithTRACELevelReturnsLevelLoggerWithTRACELevel(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{TRACE})

	_, ok := result[TRACE]

	if !ok {
		t.Fatalf("No levels provided so it should return a map with TRACE level")
	}
}

func TestNewLevelLoggerWithTRACELevelReturnsLevelLoggerWithLength1(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{TRACE})
	size := len(result)

	if size != 1 {
		t.Fatalf("One level provided so it should return a map with 1 elements, but it has %d", size)
	}
}

func TestNewLevelLoggerWithINFOLevelReturnsLevelLoggerWithINFOLevel(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{INFO})

	_, ok := result[INFO]

	if !ok {
		t.Fatalf("No levels provided so it should return a map with INFO level")
	}
}

func TestNewLevelLoggerWithINFOLevelReturnsLevelLoggerWithLength1(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{INFO})
	size := len(result)

	if size != 1 {
		t.Fatalf("One level provided so it should return a map with 1 elements, but it has %d", size)
	}
}

func TestNewLevelLoggerWithoutLevelsReturnsLevelLoggerWithWARNINGLevel(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{WARNING})

	_, ok := result[WARNING]

	if !ok {
		t.Fatalf("No levels provided so it should return a map with WARNING level")
	}
}

func TestNewLevelLoggerWithWARNINGLevelReturnsLevelLoggerWithLenght1(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{WARNING})
	size := len(result)

	if size != 1 {
		t.Fatalf("One level provided so it should return a map with 1 elements, but it has %d", size)
	}
}

func TestNewLevelLoggerWithERRORLevelReturnsLevelLoggerWithERRORLevel(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{ERROR})

	_, ok := result[ERROR]

	if !ok {
		t.Fatalf("No levels provided so it should return a map with ERROR level")
	}
}

func TestNewLevelLoggerWithERRORLevelReturnsLevelLoggerWithLenght1(t *testing.T) {

	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, []Level{TRACE})
	size := len(result)

	if size != 1 {
		t.Fatalf("One level provided so it should return a map with 1 elements, but it has %d", size)
	}
}

func TestNewLevelLoggerWithMultipleLevelsReturnsLevelLoggerWithTheSameLenghtAsTheLevels(t *testing.T) {

	levels := []Level{TRACE, WARNING}
	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, levels)

	expected := len(levels)
	size := len(result)

	if size != expected {
		t.Fatalf("%d levels provided so it should return a map with %d elements, but it has %d", expected, expected, size)
	}
}

func TestNewLevelLoggerWithMultipleLevelsReturnsLevelLoggerWithTheLevelsProvided(t *testing.T) {

	levels := []Level{TRACE, WARNING}
	result := NewLevelLogger(os.Stdout, defaultLoggerFlag, levels)

	for _, v := range levels {
		_, ok := result[v]

		if !ok {
			t.Fatalf("No levels provided so it should return a map with %d level", v)
		}

	}
}

func TestContaisLevelWithLevelPresentReturnsTrue(t *testing.T) {

	levels := []Level{TRACE}
	ll := NewLevelLogger(os.Stdout, defaultLoggerFlag, levels)
	result := ll.ContainsLevel(TRACE)

	if !result {
		t.Fatalf("Level is present in LevelLogger so it should return true")
	}
}

func TestContaisLevelWithLevelNotPresentReturnsFalse(t *testing.T) {

	levels := []Level{TRACE}
	ll := NewLevelLogger(os.Stdout, defaultLoggerFlag, levels)
	result := ll.ContainsLevel(ERROR)

	if result {
		t.Fatalf("Level is not present in LevelLogger so it should return false")
	}
}
