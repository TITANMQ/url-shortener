package utils

import "testing"


func TestMessage(t *testing.T) {
	input := Message(true, "Test message")

	if input["status"] != true {
		t.Errorf(`input["status"] returned %v when expected %v`, input["status"], true)
	}

	if input["message"] != "Test message"{
		t.Errorf(`input["message"] returned %s when expected %s`, input["message"], "Test message")
	}

}

// func TestRespond(t *testing.T) {
	
// }
