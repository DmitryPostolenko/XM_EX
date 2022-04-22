package ipapi

import "testing"

func TestGetIpData(t *testing.T) {
	_, err := GetIpData()
	if err != nil {
		t.Fatalf("Error getting or decoding data from ipapi.co: %s", err)
	}
}
