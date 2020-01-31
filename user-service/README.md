# user-service
Service to manage user entities.
```
type User struct {
	Id    string
	Name  string
	Age   int
	Phone string
}
```
Supports add, update and get operations on user entities.
```
	// POST    /users/                             create new user
	// GET     /users/:id/                         retrieves the given user by id
	// PUT     /users/:id/                         updates user after it is created
```

phone field is standardized using standardize API exposed by phone-standardization service before persistence.
