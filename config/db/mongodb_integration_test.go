package db

import (
	"testing"
)

func TestNewMongoDBConnection(t *testing.T) {
	// Given
	// When
	_, err := NewMongoDBClient()
	// Then
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err)
	}
}

func TestNewMongoDBCollection(t *testing.T) {
	// Given
	// When
	_, err := NewMongoDBCollection("test", "test")
	// Then
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err)
	}
}
