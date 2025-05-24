package message

import "time"

type Message struct {
	AckId        string    `json:"ack_id"`
	Acknowledged bool      `json:"acknowledged"`
	Topic        string    `json:"topic"`
	Data         string    `json:"message"`
	Timestamp    time.Time `json:"timestamp"`
}
