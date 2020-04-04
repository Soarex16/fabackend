package domain

import (
	"github.com/google/uuid"
	"time"
)

// Achievement - достижения пользователя за выполненные упражнения
type Achievement struct {
	ID          uuid.UUID `json:"id,omitempty"`
	UserID      uuid.UUID `json:"userId,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	IconColor   string    `json:"iconColor,omitempty"`
	Description string    `json:"description,omitempty"`
	Title       string    `json:"title,omitempty"`
	// Вес достижения (нужно для календаря)
	Price int64 `json:"price,omitempty"`
	// Имя иконки для отображения в карточке
	IconName string `json:"iconName,omitempty"`
}
