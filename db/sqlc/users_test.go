package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"multitenant.com/app/db/util"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	searchedUser, err := testQueries.GetUser(context.Background(), user.ID)

	// assertions
	require.NoError(t, err)
	require.NotEmpty(t, searchedUser)
	require.Equal(t, user.ID, searchedUser.ID)
}

func TestGetUsers(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  10,
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	// assertions
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.NotZero(t, len(users))
	require.Len(t, users, 10)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	var fullName = sql.NullString{String: "Laxmi Narsimha Reddy Jammula", Valid: true}

	arg := UpdateUserParams{
		ID:       user.ID,
		FullName: fullName,
		RoleID:   sql.NullInt32{Int32: 1, Valid: true},
	}
	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)

	// assertions
	require.NoError(t, err)
	require.Equal(t, updatedUser.FullName, arg.FullName)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	// assertions
	require.NoError(t, err)

	//negative test
	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func createRandomUser(t *testing.T) User {
	randomName := util.RandomString(6)
	randomDomain := util.RandomString(4)
	randomEmail := randomName + "@" + randomDomain + ".com"
	arg := CreateUserParams{
		Email:    randomEmail,
		UserName: randomName,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)

	// assertions
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.UserName, user.UserName)
	require.NotZero(t, user.ID)
	require.NotEmpty(t, user.CreatedTimestamp)
	require.NotZero(t, user.UpdatedTimestamp)

	return user
}
