package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/superclouds/super-sdk-go-v1/superclouds"
	"net/http"
	"net/url"
)

// UsersClient provides methods to interact with the users endpoint of the Superclouds API.
type UsersClient struct {
	config *superclouds.Config
}

// NewUsersClient creates a new UsersClient instance with the provided configuration.
//
// Parameters:
// - cfg: The configuration instance created using NewConfig or NewConfigWithParams.
//
// Example usage:
//
//	usersClient := superclouds.NewUsersClient(cfg)
func NewUsersClient(cfg *superclouds.Config) *UsersClient {
	return &UsersClient{config: cfg}
}

// SuperAPIResponse represents the structure of the response from the Superclouds API.
type SuperAPIResponse struct {
	Data    []User `json:"data"`
	Message string `json:"message"`
	Page    int    `json:"page"`
	Pages   int    `json:"pages"`
	Size    int    `json:"size"`
	Status  int    `json:"status"`
	Total   int    `json:"total"`
}

// ListUsersInput defines the input parameters for the ListUsers method.
type ListUsersInput struct {
	Size       int    `json:"size"`
	Page       int    `json:"page"`
	SearchTerm string `json:"s"`
}

// ListUsersOutput defines the output structure for the ListUsers method.
type ListUsersOutput struct {
	Users []User `json:"data"`
}

// User represents a user in the Superclouds system.
type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

// CreateUserInput defines the input parameters for the CreateUser method.
type CreateUserInput struct {
	Email string `json:"email"`
}

// DeleteUserInput defines the input parameters for the DeleteUser method.
type DeleteUserInput struct {
	Email string `json:"email"`
}

// UpdateUserInput defines the input parameters for the UpdateUser method.
type UpdateUserInput struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Contact   string `json:"contact,omitempty"`
}

// UserOutput defines the output structure for user-related methods.
type UserOutput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	// Add more fields as needed
}

// ListRolesOutput defines the output structure for the ListRoles method.
type ListRolesOutput struct {
	Roles []string `json:"roles"`
}

// UpdateUserRoleInput defines the input parameters for the UpdateUserRole method.
type UpdateUserRoleInput struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// ChangePasswordInput defines the input parameters for the ChangePassword method.
type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// ListUsers retrieves a paginated list of users.
//
// Parameters:
// - ctx: The context for the request.
// - input: The input parameters for the request.
//
// Returns:
// - ListUsersOutput: The list of users and pagination details.
// - error: Any error encountered during the request.
//
// Example usage:
//
//	usersOutput, err := usersClient.ListUsers(context.TODO(), &users.ListUsersInput{
//	    Size: 10,
//	    Page: 1,
//	    SearchTerm: "search_term",
//	})
//	if err != nil {
//	    log.Fatalf("Failed to list users: %v", err)
//	}
//	log.Printf("Users: %v", usersOutput.Users)
func (c *UsersClient) ListUsers(ctx context.Context, input *ListUsersInput) (*ListUsersOutput, error) {
	baseURL, err := url.Parse(c.config.SuperURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %v", err)
	}

	baseURL.Path += "/users"

	params := url.Values{}
	if input.Size > 0 {
		params.Add("size", fmt.Sprintf("%d", input.Size))
	}
	if input.Page > 0 {
		params.Add("page", fmt.Sprintf("%d", input.Page))
	}
	if input.SearchTerm != "" {
		params.Add("s", input.SearchTerm)
	}
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse SuperAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &ListUsersOutput{
		Users: apiResponse.Data,
	}, nil
}

// CreateUser creates a new user within the organization.
//
// Parameters:
// - ctx: The context for the request.
// - input: The input parameters for the request.
//
// Returns:
// - UserOutput: The created user's details.
// - error: Any error encountered during the request.
//
// Example usage:
//
//	newUser, err := usersClient.CreateUser(context.TODO(), &users.CreateUserInput{
//	    Email: "new.user@example.com",
//	})
//	if err != nil {
//	    log.Fatalf("Failed to create user: %v", err)
//	}
//	log.Printf("Created User: %v", newUser)
func (c *UsersClient) CreateUser(ctx context.Context, input *CreateUserInput) (*UserOutput, error) {
	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/users", c.config.SuperURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	var output UserOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &output, nil
}

// DeleteUser removes a user from the organization.
//
// Parameters:
// - ctx: The context for the request.
// - input: The input parameters for the request.
//
// Returns:
// - error: Any error encountered during the request.
//
// Example usage:
//
//	err := usersClient.DeleteUser(context.TODO(), &users.DeleteUserInput{
//	    Email: "delete.user@example.com",
//	})
//	if err != nil {
//	    log.Fatalf("Failed to delete user: %v", err)
//	}
//	log.Println("Deleted User")
func (c *UsersClient) DeleteUser(ctx context.Context, input *DeleteUserInput) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s/users?email=%s", c.config.SuperURL, input.Email), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete user: %s", resp.Status)
	}

	return nil
}

// UpdateUser updates the details of the authenticated user.
//
// Parameters:
// - ctx: The context for the request.
// - input: The input parameters for the request.
//
// Returns:
// - UserOutput: The updated user's details.
// - error: Any error encountered during the request.
//
// Example usage:
//
//	updatedUser, err := usersClient.UpdateUser(context.TODO(), &users.UpdateUserInput{
//	    FirstName: "John",
//	    LastName:  "Doe",
//	    Contact:   "999XXXX999",
//	})
//	if err != nil {
//	    log.Fatalf("Failed to update user: %v", err)
//	}
//	log.Printf("Updated User: %v", updatedUser)
func (c *UsersClient) UpdateUser(ctx context.Context, input *UpdateUserInput) (*UserOutput, error) {
	reqBody, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("%s/user", c.config.SuperURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	var output UserOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &output, nil
}

// GetUser retrieves detailed information about the authenticated user.
//
// Parameters:
// - ctx: The context for the request.
//
// Returns:
// - UserOutput: The authenticated user's details.
// - error: Any error encountered during the request.
//
// Example usage:
//
//	user, err := usersClient.GetUser(context.TODO())
//	if err != nil {
//	    log.Fatalf("Failed to get user: %v", err)
//	}
//	log.Printf("Authenticated User: %v", user)
func (c *UsersClient) GetUser(ctx context.Context) (*UserOutput, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/user", c.config.SuperURL), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	var output UserOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &output, nil
}

// ListRoles retrieves a list of available roles within the system.
//
// Parameters:
// - ctx: The context for the request.
//
// Returns:
// - ListRolesOutput: The list of roles.
// - error: Any error encountered during the request.
//
// Example usage:
//
//	roles, err := usersClient.ListRoles(context.TODO())
//	if err != nil {
//	    log.Fatalf("Failed to list roles: %v", err)
//	}
//	log.Printf("Available Roles: %v", roles)
func (c *UsersClient) ListRoles(ctx context.Context) (*ListRolesOutput, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/users/roles", c.config.SuperURL), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	var output ListRolesOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &output, nil
}

// UpdateUserRole updates the role of a user within the organization.
//
// Parameters:
// - ctx: The context for the request.
// - input: The input parameters for the request.
//
// Returns:
// - error: Any error encountered during the request.
//
// Example usage:
//
//	err := usersClient.UpdateUserRole(context.TODO(), &users.UpdateUserRoleInput{
//	    Email: "user@example.com",
//	    Role:  "MODIFY",
//	})
//	if err != nil {
//	    log.Fatalf("Failed to update user role: %v", err)
//	}
//	log.Println("Updated User Role")
func (c *UsersClient) UpdateUserRole(ctx context.Context, input *UpdateUserRoleInput) error {
	reqBody, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("%s/users/role", c.config.SuperURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update user role: %s", resp.Status)
	}

	return nil
}

// ChangePassword allows the authenticated user to change their password.
//
// Parameters:
// - ctx: The context for the request.
// - input: The input parameters for the request.
//
// Returns:
// - error: Any error encountered during the request.
//
// Example usage:
//
//	err := usersClient.ChangePassword(context.TODO(), &users.ChangePasswordInput{
//	    CurrentPassword: "oldpassword",
//	    NewPassword:     "newpassword",
//	    ConfirmPassword: "newpassword",
//	})
//	if err != nil {
//	    log.Fatalf("Failed to change password: %v", err)
//	}
//	log.Println("Changed Password")
func (c *UsersClient) ChangePassword(ctx context.Context, input *ChangePasswordInput) error {
	reqBody, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("%s/change-password", c.config.SuperURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.SuperToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SuperToken)
	}

	resp, err := c.config.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to change password: %s", resp.Status)
	}

	return nil
}
