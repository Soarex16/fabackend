package domain

// Achievement - достижения пользователя за выполненные упражнения
type Achievement struct {
	ID string `json:"id,omitempty"`
	// Пользователь, который получил это достижение
	UserID string `json:"userId,omitempty"`
	// Дата в формате unix timestamp
	Date int64 `json:"date,omitempty"`
	// Вес достижения (нужно для календаря)
	Price int64 `json:"price,omitempty"`
	// Имя иконки для отображения в карточке
	IconName    string `json:"iconName,omitempty"`
	IconColor   string `json:"iconColor,omitempty"`
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
}
