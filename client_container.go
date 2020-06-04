package main

// ClientContainer is a struct that has all components of a client container
type ClientContainer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Network string `json:"net"`
	AdbPort string `json:"adb"`
	VNCPort string `json:"vnc"`
}
