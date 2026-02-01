package sample

// User はドメインモデルの例。
type User struct {
	ID     string
	Name   string
	Email  string
	Active bool
}

// NewActiveUser は固定パターンのファクトリメソッド例。
func NewActiveUser() User {
	return User{
		ID:     "u-001",
		Name:   "Alice",
		Email:  "alice@example.com",
		Active: true,
	}
}

// UserBuilder は差分指定で生成できるビルダー。
type UserBuilder struct {
	user User
}

// NewUserBuilder は標準値で初期化したビルダーを返す。
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: User{
		ID:     "u-default",
		Name:   "Default",
		Email:  "default@example.com",
		Active: true,
	}}
}

func (b *UserBuilder) WithID(id string) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithInactive() *UserBuilder {
	b.user.Active = false
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}
