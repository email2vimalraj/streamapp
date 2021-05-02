package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/email2vimalraj/streamapp/util"
	"github.com/stretchr/testify/require"
)

func createRandomIviteeWithUserAndStream(t *testing.T, user User, stream Stream) Invitee {
	arg := CreateInviteeParams{
		Inviter:  user.Username,
		StreamID: stream.ID,
		FullName: util.RandomString(6),
		Email:    sql.NullString{String: util.RandomEmail(), Valid: true},
	}

	invitee, err := testQueries.CreateInvitee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invitee)

	require.Equal(t, arg.Inviter, invitee.Inviter)
	require.Equal(t, arg.StreamID, invitee.StreamID)
	require.Equal(t, arg.FullName, invitee.FullName)
	require.Equal(t, arg.Email, invitee.Email)

	require.Empty(t, invitee.MobileNumber)
	require.NotZero(t, invitee.ID)
	require.NotZero(t, stream.CreatedAt)

	return invitee
}

func createRandomIvitee(t *testing.T) Invitee {
	user := createRandomUser(t)
	stream := createRandomStreamWithUser(t, user)

	return createRandomIviteeWithUserAndStream(t, user, stream)
}

func TestCreateInvitee(t *testing.T) {
	createRandomIvitee(t)
}

func TestListInvitees(t *testing.T) {
	user := createRandomUser(t)
	stream := createRandomStreamWithUser(t, user)

	for i := 0; i < 10; i++ {
		createRandomIviteeWithUserAndStream(t, user, stream)
	}

	arg := ListInviteesParams{
		Inviter:  user.Username,
		StreamID: stream.ID,
		Limit:    5,
		Offset:   5,
	}

	invitees, err := testQueries.ListInvitees(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, invitees, 5)

	for _, invitee := range invitees {
		require.NotEmpty(t, invitee)
	}
}
