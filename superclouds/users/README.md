
#### Example : Creating a User

```go
newUser, err := usersClient.CreateUser(context.TODO(), &users.CreateUserInput{
    Email: "new.user@example.com",
})
if err != nil {
    log.Fatalf("Failed to create user: %v", err)
}
log.Printf("Created User: %v", newUser)
```

#### Example : Listing Users

```go
usersOutput, err := usersClient.ListUsers(context.TODO(), &users.ListUsersInput{
    Size:       10,
    Page:       1,
    SearchTerm: "search_term",
})
if err != nil {
    log.Fatalf("Failed to list users: %v", err)
}
log.Printf("Users: %v", usersOutput.Users)
```

#### Deleting a User

```go
err = usersClient.DeleteUser(context.TODO(), &users.DeleteUserInput{
    Email: "delete.user@example.com",
})
if err != nil {
    log.Fatalf("Failed to delete user: %v", err)
}
log.Println("Deleted User")
```

#### Updating a User

```go
updatedUser, err := usersClient.UpdateUser(context.TODO(), &users.UpdateUserInput{
    FirstName: "John",
    LastName:  "Doe",
    Contact:   "999XXXX999",
})
if err != nil {
    log.Fatalf("Failed to update user: %v", err)
}
log.Printf("Updated User: %v", updatedUser)
```

#### Retrieving Authenticated User Details

```go
user, err := usersClient.GetUser(context.TODO())
if err != nil {
    log.Fatalf("Failed to get user: %v", err)
}
log.Printf("Authenticated User: %v", user)
```

#### Listing Roles

```go
roles, err := usersClient.ListRoles(context.TODO())
if err != nil {
    log.Fatalf("Failed to list roles: %v", err)
}
log.Printf("Available Roles: %v", roles)
```

#### Updating User Role

```go
err = usersClient.UpdateUserRole(context.TODO(), &users.UpdateUserRoleInput{
    Email: "user@example.com",
    Role:  "MODIFY",
})
if err != nil {
    log.Fatalf("Failed to update user role: %v", err)
}
log.Println("Updated User Role")
```

#### Changing Password

```go
err = usersClient.ChangePassword(context.TODO(), &users.ChangePasswordInput{
    CurrentPassword: "oldpassword",
    NewPassword:     "newpassword",
    ConfirmPassword: "newpassword",
})
if err != nil {
    log.Fatalf("Failed to change password: %v", err)
}
log.Println("Changed Password")
```
