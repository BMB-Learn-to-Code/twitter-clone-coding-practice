package env

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	const envKey = "TEST_ENV_STRING"
	const fallbackValue = "fallback"

	// Test when the environment variable is not set
	os.Unsetenv(envKey)
	if val := GetString(envKey, fallbackValue); val != fallbackValue {
		t.Errorf("Expected fallback value '%s', got '%s'", fallbackValue, val)
	}

	// Test when the environment variable is set
	expectedValue := "test_value"
	os.Setenv(envKey, expectedValue)
	if val := GetString(envKey, fallbackValue); val != expectedValue {
		t.Errorf("Expected value '%s', got '%s'", expectedValue, val)
	}
}

func TestGetInt(t *testing.T) {
	const envKey = "TEST_ENV_INT"
	const fallbackValue uint32 = 42

	// Test when the environment variable is not set
	os.Unsetenv(envKey)
	if val := GetInt(envKey, fallbackValue); val != fallbackValue {
		t.Errorf("Expected fallback value '%d', got '%d'", fallbackValue, val)
	}

	// Test when the environment variable is set to a valid integer
	expectedValue := uint32(100)
	os.Setenv(envKey, "100")
	if val := GetInt(envKey, fallbackValue); val != expectedValue {
		t.Errorf("Expected value '%d', got '%d'", expectedValue, val)
	}
}

func TestGetBool(t *testing.T) {
	const envKey = "TEST_ENV_BOOL"
	const fallbackValue = true

	// Test when the environment variable is not set
	os.Unsetenv(envKey)
	if val := GetBool(envKey, fallbackValue); val != fallbackValue {
		t.Errorf("Expected fallback value '%t', got '%t'", fallbackValue, val)
	}

	// Test when the environment variable is set to a valid boolean
	expectedValue := false
	os.Setenv(envKey, "false")
	if val := GetBool(envKey, fallbackValue); val != expectedValue {
		t.Errorf("Expected value '%t', got '%t'", expectedValue, val)
	}
}
