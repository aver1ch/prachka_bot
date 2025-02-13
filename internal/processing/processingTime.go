/*

имя пользователя, во сколько, на сколько и какую машинку
отзывы
киллер фича

*/

package processing

import (
	"log"
	"time"
)

const quickCycleDurationMin = 30
const otherCycleTimeMin = 50
const timeToChangeUsersMin = 5

type Time struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

func calculateEndTime(currentTime time.Time, bookingTime int) time.Time {
	return currentTime.Add(time.Duration(bookingTime) * time.Minute)
}

func calculateTimeUntilRelease(queue []QueueItem) Time {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("Ошибка загрузки часового пояса Moscow: %v", err)
	}

	currentTime := time.Now().In(loc)
	log.Print("Current time in Moscow: ", currentTime)

	var latestEndTime time.Time

	for _, item := range queue {
		bookingDuration := item.Time.Minutes

		endTime := calculateEndTime(currentTime, bookingDuration)

		log.Printf("User %s, booked for %d minutes, will finish at: %s", item.Username, item.Time.Minutes, endTime.Format("15:04:05"))

		if endTime.After(latestEndTime) {
			latestEndTime = endTime
		}
	}

	log.Printf("Max end time (when service will be free): %s", latestEndTime.Format("15:04:05"))

	return Time{
		Hours:   latestEndTime.Hour(),
		Minutes: latestEndTime.Minute(),
		Seconds: latestEndTime.Second(),
	}
}
