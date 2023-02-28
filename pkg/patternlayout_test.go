package pkg

import (
	"fmt"
	"testing"
)

func TestFormatRecordWithEmptyShouldBeEmpty(t *testing.T) {
	format := PatternLayout("")
	result := format.FormatRecord(Record{Message: "test"})

	if result != string(format) {
		t.Fatalf("Empty Formating String should result in empty, but it was %s", result)
	}
}

func TestFormatRecordWithoutPatternsShouldBeFormatingString(t *testing.T) {
	format := PatternLayout("no patterns")
	result := format.FormatRecord(Record{Message: "test"})

	if result != string(format) {
		t.Fatalf("Formating String without patterns should result in Formating String, but it was %s", result)
	}
}

func TestFormatRecordWithOnlyDatePatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(DatePattern)
	result := format.FormatRecord(Record{Message: "test"})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyDatePatternShouldHaveSize10(t *testing.T) {
	format := PatternLayout(DatePattern)
	result := format.FormatRecord(Record{Message: "test"})
	expected := 10
	size := len(result)

	if size != expected {
		t.Fatalf("Formating String only contains Date Pattern so the size should be %d but was %d (%s)", expected, size, result)
	}
}

func TestFormatRecordWithOnlyFileNamePatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(FileNamePattern)
	result := format.FormatRecord(Record{Message: "test"})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyLevelPatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(LevelPattern)
	result := format.FormatRecord(Record{Message: "test", Level: TRACE})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyLevelPatternShouldContainLevel(t *testing.T) {
	format := PatternLayout(LevelPattern)
	result := format.FormatRecord(Record{Message: "test", Level: TRACE})
	expected := "TRACE"

	if result != expected {
		t.Fatalf("Formating String (%s) should result in %s but was %s", format, expected, result)
	}
}

func TestFormatRecordWithOnlyLinePatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(LinePattern)
	result := format.FormatRecord(Record{Message: "test"})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyMethodPatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(MethodPattern)
	result := format.FormatRecord(Record{Message: "test"})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyMessagePatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(MessagePattern)
	result := format.FormatRecord(Record{Message: "test"})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyMessagePatternShouldContainMessage(t *testing.T) {
	format := PatternLayout(MessagePattern)
	expected := "test"
	result := format.FormatRecord(Record{Message: expected})

	if result != expected {
		t.Fatalf("Formating String (%s) should result in %s but was %s", format, expected, result)
	}
}

func TestFormatRecordWithOnlyTimePatternShouldNotContainPattern(t *testing.T) {
	format := PatternLayout(TimePattern)
	result := format.FormatRecord(Record{Message: "test"})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithOnlyTimePatternShouldHaveSize8(t *testing.T) {
	format := PatternLayout(TimePattern)
	result := format.FormatRecord(Record{Message: "test"})
	expected := 8
	size := len(result)

	if size != expected {
		t.Fatalf("Formating String only contains Time Pattern so the size should be %d but was %d (%s)", expected, size, result)
	}
}

func TestFormatRecordWithAllPatternShouldNotContainPatterns(t *testing.T) {
	format := PatternLayout(DatePattern + FileNamePattern + LevelPattern + LinePattern + MethodPattern + MessagePattern + TimePattern)
	result := format.FormatRecord(Record{Message: "test", Level: TRACE})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithLevelAndMessagePatternShouldNotContainPatterns(t *testing.T) {
	format := PatternLayout(fmt.Sprintf("[%s] %s", LevelPattern, MessagePattern))
	result := format.FormatRecord(Record{Message: "test", Level: TRACE})

	if result == string(format) {
		t.Fatalf("Formating String (%s) should not contain the Pattern (%s)", format, result)
	}
}

func TestFormatRecordWithLevelAndMessagePatternShouldResultInLevelAndMessage(t *testing.T) {
	format := PatternLayout(fmt.Sprintf("[%s] %s", LevelPattern, MessagePattern))
	result := format.FormatRecord(Record{Message: "test", Level: TRACE})
	expected := "[TRACE] test"

	if result != expected {
		t.Fatalf("Formating String (%s) should result in '%s' but was '%s'", format, expected, result)
	}
}
