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
phone field is standardized using standardize API exposed by phone-standardization service before persistence.
