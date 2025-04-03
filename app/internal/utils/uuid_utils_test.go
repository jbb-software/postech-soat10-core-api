package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestStringToUuid(t *testing.T) {
	validUUID := "123e4567-e89b-12d3-a456-426614174000"
	invalidUUID := "invalid-uuid"

	t.Run("Valid UUID", func(t *testing.T) {
		parsedUUID := StringToUuid(validUUID)
		expectedUUID, err := uuid.Parse(validUUID)
		if err != nil {
			t.Fatalf("Failed to parse valid UUID in test setup: %v", err)
		}

		if parsedUUID != expectedUUID {
			t.Errorf("StringToUuid(%q) = %v, expected %v", validUUID, parsedUUID, expectedUUID)
		}
	})

	t.Run("Invalid UUID", func(t *testing.T) {
		parsedUUID := StringToUuid(invalidUUID)
		expectedUUID := uuid.Nil // Expecting zero UUID for invalid input

		if parsedUUID != expectedUUID {
			t.Errorf("StringToUuid(%q) = %v, expected %v", invalidUUID, parsedUUID, expectedUUID)
		}
	})

	t.Run("Empty UUID", func(t *testing.T) {
		parsedUUID := StringToUuid("")
		expectedUUID := uuid.Nil

		if parsedUUID != expectedUUID {
			t.Errorf("StringToUuid(%q) = %v, expected %v", "", parsedUUID, expectedUUID)
		}
	})
}
