
# Superclouds SDK for Go

The Superclouds SDK for Go provides a convenient way to interact with the Superclouds API. It supports various services and operations, allowing developers to integrate Superclouds functionality into their Go applications easily.

## Getting Started

### Installation

To get started, install the Superclouds SDK for Go using `go get`:

```sh
go get github.com/yourusername/superclouds
```

### Configuration

The SDK can be configured using environment variables or parameters. The configuration includes the base URL for the Superclouds API, SSL certificate and key paths, and the API token for authorization.

#### Environment Variables

Set the following environment variables:

- `SUPER_CERT`: The path to the SSL certificate file.
- `SUPER_KEY`: The path to the SSL key file.
- `SUPER_TOKEN`: The bearer token for API authorization.

Example:

```sh
export SUPER_CERT="/path/to/cert.pem"
export SUPER_KEY="/path/to/key.pem"
export SUPER_TOKEN="your-api-token"
```

#### Parameters

Alternatively, you can configure the SDK using parameters:

```go
cfg, err := superclouds.NewConfigWithParams(certPath, keyPath, superToken)
if err != nil {
    log.Fatalf("Failed to create config: %v", err)
}
```

### Usage

Here are some examples of how to use the SDK.

#### Initializing the SDK

```go
package main

import (
    "context"
    "log"

    "github.com/yourusername/superclouds"
)

func main() {
    // Using environment variables for config
    cfg, err := superclouds.NewConfig()
    if err != nil {
        log.Fatalf("Failed to create config: %v", err)
    }

    // Or using parameters for config
    // certPath := "/path/to/cert.pem"
    // keyPath := "/path/to/key.pem"
    // superToken := "your-api-token"
    // cfg, err := superclouds.NewConfigWithParams(certPath, keyPath, superToken)
    // if err != nil {
    //     log.Fatalf("Failed to create config: %v", err)
    // }

    usersClient := superclouds.NewUsersClient(cfg)

    // Create User
    newUser, err := usersClient.CreateUser(context.TODO(), &superclouds.CreateUserInput{
        Email: "new.user@example.com",
    })
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }
    log.Printf("Created User: %v", newUser)

    // Other operations...
}
```

#### Creating a User

```go
newUser, err := usersClient.CreateUser(context.TODO(), &superclouds.CreateUserInput{
    Email: "new.user@example.com",
})
if err != nil {
    log.Fatalf("Failed to create user: %v", err)
}
log.Printf("Created User: %v", newUser)
```

#### Listing Users

```go
usersOutput, err := usersClient.ListUsers(context.TODO(), &superclouds.ListUsersInput{
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
err = usersClient.DeleteUser(context.TODO(), &superclouds.DeleteUserInput{
    Email: "delete.user@example.com",
})
if err != nil {
    log.Fatalf("Failed to delete user: %v", err)
}
log.Println("Deleted User")
```

#### Updating a User

```go
updatedUser, err := usersClient.UpdateUser(context.TODO(), &superclouds.UpdateUserInput{
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
err = usersClient.UpdateUserRole(context.TODO(), &superclouds.UpdateUserRoleInput{
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
err = usersClient.ChangePassword(context.TODO(), &superclouds.ChangePasswordInput{
    CurrentPassword: "oldpassword",
    NewPassword:     "newpassword",
    ConfirmPassword: "newpassword",
})
if err != nil {
    log.Fatalf("Failed to change password: %v", err)
}
log.Println("Changed Password")
```

## Users Package

For more detailed examples and usage of the `users` package, see the [Users README](./superclouds/users/README.md).
