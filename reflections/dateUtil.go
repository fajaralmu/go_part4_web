package reflections

import "time"

//GetTimeGreeting generate greeting message
func GetTimeGreeting() string {
	//var time string = "Morning"
	currentTime := time.Now()

	var hour int = currentTime.Hour()
	var timeStr string
	if hour >= 3 && hour < 11 {
		timeStr = "Morning"
	} else if hour >= 11 && hour < 18 {
		timeStr = "Afternoon"
	} else {
		timeStr = "Evening"
	}

	return timeStr
}
