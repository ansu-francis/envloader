package envloader

import (
	"testing"
	"os"
)

func TestLoadFileNotFound(t *testing.T) {
	err := LoadEnv("somefilethatwillneverexistever.env")
	if err == nil {
		t.Error("File wasn't found but Load didn't return an error")
	}
}

func TestReadPlainEnv(t *testing.T) {
	envFileName := "fixtures/plain.env"
	expectedValues := map[string]string{
		"OPTION_A": "1",
		"OPTION_B": "2",
		"OPTION_C": "3",
		"OPTION_D": "4",
		"OPTION_E": "5",
		"OPTION_H": "1",
	}

	err := LoadEnv(envFileName)
	if err != nil {
		t.Error("Error reading file")
	}

	for key, value := range expectedValues {
		if os.Getenv(key) != value {
			t.Error("Read got one of the keys wrong")
		}
	}
}

func TestActualEnvVarsAreOverwritten(t *testing.T) {
	os.Clearenv()
	os.Setenv("OPTION_A", "actualenv")
	_ = LoadEnv("fixtures/plain.env")

	if os.Getenv("OPTION_A") == "actualenv" {
		t.Error("An ENV var set earlier is not overwritten")
	}
}
