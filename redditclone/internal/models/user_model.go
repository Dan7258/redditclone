//package models
//
//import "sync"
//
//type UserMemory struct {
//	mu      sync.Mutex
//	Storage map[int]*User
//	counter int
//}
//
//type User struct {
//	Id       int
//	Username string `json:"username"`
//	Password string `json:"password,omitempty"`
//}
//
//func NewUserMemory() *UserMemory {
//	return &UserMemory{
//		Storage: make(map[int]*User),
//		counter: 1,
//		mu:      sync.Mutex{},
//	}
//}
//
//func (d *UserMemory) AddUser(user *User) {
//	d.mu.Lock()
//	defer d.mu.Unlock()
//	user.Id = d.counter
//	d.Storage[d.counter] = user
//
//	d.counter++
//}
//
//func (d *UserMemory) RemoveUser(user User) {
//	d.mu.Lock()
//	defer d.mu.Unlock()
//	delete(d.Storage, user.Id)
//}
//
//func (d *UserMemory) GetUser(login string) (*User, bool) {
//	d.mu.Lock()
//	defer d.mu.Unlock()
//	for _, user := range d.Storage {
//		if user.Username == login {
//			return user, true
//		}
//	}
//	return nil, false
//}
