package db

import (
	"context"
	"testing"
	"time"

	"github.com/email2vimalraj/streamapp/util"
	"github.com/stretchr/testify/require"
)

func createRandomStreamWithUser(t *testing.T, user User) Stream {
	arg := CreateStreamParams{
		StreamName: util.RandomString(6),
		StreamLink: util.RandomString(10),
		Username:   user.Username,
	}

	stream, err := testQueries.CreateStream(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stream)

	require.Equal(t, arg.StreamName, stream.StreamName)
	require.Equal(t, arg.StreamLink, stream.StreamLink)
	require.Equal(t, arg.Username, stream.Username)

	require.NotZero(t, stream.ID)
	require.NotZero(t, stream.CreatedAt)

	return stream
}

func createRandomStream(t *testing.T) Stream {
	user := createRandomUser(t)

	return createRandomStreamWithUser(t, user)
}

func TestCreateStream(t *testing.T) {
	createRandomStream(t)
}

func TestGetStream(t *testing.T) {
	stream1 := createRandomStream(t)
	stream2, err := testQueries.GetStream(context.Background(), stream1.StreamName)
	require.NoError(t, err)
	require.NotEmpty(t, stream2)

	require.Equal(t, stream1.Username, stream2.Username)
	require.Equal(t, stream1.StreamName, stream2.StreamName)
	require.Equal(t, stream1.StreamLink, stream2.StreamLink)
	require.WithinDuration(t, stream1.CreatedAt, stream2.CreatedAt, time.Second)
}

func TestListStreams(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomStreamWithUser(t, user)
	}

	arg := ListStreamsParams{
		Username: user.Username,
		Limit:    5,
		Offset:   5,
	}

	streams, err := testQueries.ListStreams(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, streams, 5)

	for _, stream := range streams {
		require.NotEmpty(t, stream)
	}
}
