package processing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"laundryBot/internal/errs"
	"log"
	"os"
)

type QueueItem struct {
	Username string `json:"username"`
	Time     Time   `json:"time"`
	ChatID   int64  `json:"chatID"`
	Room     int    `json:"room"`
}

type ServiceStatus struct {
	TimeUntilRelease Time        `json:"timeUntilRelease"`
	Status           bool        `json:"status"`
	Queue            []QueueItem `json:"queue"`
}

func ReadInfoFromJSON(service string) (ServiceStatus, error) {

	if service == "orderDry" { // костыль
		service = "Cушка"
	} else if service == "orderLaundry+Стиралка 1+quick" || service == "orderLaundry+Стиралка 1+long" {
		service = "Стиралка 1"
	} else if service == "orderLaundry+Стиралка 2+quick" || service == "orderLaundry+Стиралка 2+long" {
		service = "Стиралка 2"
	} else {
		service = "Стиралка 3"
	}

	file, err := os.Open("/prachka_bot/status.json")
	if err != nil {
		log.Println("Ошибка открытия файла status.json:", err)
		return ServiceStatus{}, fmt.Errorf("%w: %w", err, errs.ErrReadStatusFile)
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка чтения данных из файла:", err)
		return ServiceStatus{}, fmt.Errorf("ошибка при чтении данных: %w", err)
	}

	var services map[string]ServiceStatus
	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	err = decoder.Decode(&services)
	if err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		return ServiceStatus{}, fmt.Errorf("ошибка при декодировании JSON: %w", err)
	}

	log.Print(services[service], "вот это сервис")

	serviceStatus, exists := services[service]
	if !exists {
		return ServiceStatus{}, fmt.Errorf("сервис %s не найден", service)
	}

	serviceStatus.TimeUntilRelease = calculateTimeUntilRelease(serviceStatus.Queue)

	log.Printf("Читаем в JSON: %+v", services)

	return serviceStatus, nil
}

func WriteInfoToJSON(service string, status ServiceStatus) error {
	file, err := os.OpenFile("/prachka_bot/status.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println("Ошибка открытия файла status.json для записи:", err)
		return fmt.Errorf("%w: %w", err, errs.ErrWriteStatusFile)
	}
	defer file.Close()

	var services map[string]ServiceStatus
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&services)
	if err != nil && err.Error() != "EOF" {
		log.Println("Ошибка декодирования JSON при чтении:", err)
		return fmt.Errorf("ошибка при декодировании JSON: %w", err)
	}

	if services == nil {
		services = make(map[string]ServiceStatus)
	}
	services[service] = status

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Println("Ошибка при перемещении курсора в начало файла:", err)
		return fmt.Errorf("ошибка при перемещении курсора в начало файла: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(services)
	if err != nil {
		log.Println("Ошибка записи в JSON:", err)
		return fmt.Errorf("ошибка записи в JSON: %w", err)
	}

	log.Printf("Записываем в JSON: %+v", services)
	return nil
}
