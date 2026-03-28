package models

type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"todo_id"`
	Title       string `gorm:"type:text;not null" json:"title"`
	Description string `gorm:"type:text;not null" json:"description"`
	UserID      uint   `gorm:"index" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"-"`
}

type TodoBuilder struct {
	todo *Todo
}

func NewTodoBuilder() *TodoBuilder {
	return &TodoBuilder{
		todo: &Todo{},
	}
}

func (b *TodoBuilder) WithTitle(title string) *TodoBuilder {
	b.todo.Title = title
	return b
}

func (b *TodoBuilder) WithDescription(description string) *TodoBuilder {
	b.todo.Description = description
	return b
}

func (b *TodoBuilder) WithUserID(userID uint) *TodoBuilder {
	b.todo.UserID = userID
	return b
}

func (b *TodoBuilder) Build() *Todo {
	return b.todo
}
