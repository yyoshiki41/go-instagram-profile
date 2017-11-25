package instagram

import "testing"

func TestGetProfile(t *testing.T) {
	userID := "528817151"
	userName := "nasa"
	profile, err := GetProfile(userName)
	if err != nil {
		t.Fatalf("Failed to get profile: %s", err)
	}

	if profile.User.ID != userID {
		t.Fatalf("UserID: expected %s, but got %s", userID, profile.User.ID)
	}
	if profile.User.Username != "nasa" {
		t.Fatalf("UserName: expected %s, but got %s", userID, profile.User.ID)
	}
}
