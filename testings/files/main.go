package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := writeUserStream(context.Background())
	fmt.Printf("err: %v\n", err)
}

func writeToFile() {
	if err := os.WriteFile("./data", []byte{100, 123}, 0o600); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

var _ fmt.Stringer = User{}

type User struct {
	Name string
	ID   int
}

// String implements fmt.Stringer.
func (u User) String() string {
	return fmt.Sprintf("%s:%d", u.Name, u.ID)
}

func parseUser(line string) (User, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return User{}, fmt.Errorf("record(%s) was not in the correct format", line)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return User{}, fmt.Errorf("record(%s) had non-numeric ID", line)
	}

	return User{Name: strings.TrimSpace(parts[0]), ID: id}, nil
}

type UserWithError struct {
	err  error
	user User
}

func decodeUsers(ctx context.Context, reader io.Reader) <-chan UserWithError {
	ch := make(chan UserWithError, 1)

	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			if ctx.Err() != nil {
				ch <- UserWithError{err: ctx.Err()}
				return
			}

			user, err := parseUser(scanner.Text())
			if err != nil {
				ch <- UserWithError{err: err}
				return
			}
			ch <- UserWithError{user: user}
		}
	}()

	return ch
}

func readUsersStream(ctx context.Context) ([]User, error) {
	f, err := os.Open("./users")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	users := make([]User, 0)
	for user := range decodeUsers(ctx, f) {
		if user.err != nil {
			return nil, err
		}
		users = append(users, user.user)
	}

	return users, nil
}

func writeUsers(ctx context.Context, writer io.Writer, users []User) error {
	for _, user := range users {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		_, err := fmt.Fprintf(writer, "%s\n", user)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeUserStream(ctx context.Context) error {
	users, err := readUsersStream(ctx)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("users_new", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	writeUsers(ctx, f, users)
	return nil
}
