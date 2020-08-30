package custom_types

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

type Time = time.Time

func New() Time {
	return time.Now()
}

// Lets redefine the base ID type to use an id as int
//func MarshalID(id int) graphql.Marshaler {
//	return graphql.WriterFunc(func(w io.Writer) {
//		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
//	})
//}
//
//// And the same for the unmarshaler
//func UnmarshalID(v interface{}) (int, error) {
//	id, ok := v.(string)
//	if !ok {
//		return 0, fmt.Errorf("ids must be strings")
//	}
//	i, e := strconv.Atoi(id)
//	return int(i), e
//}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return Time{}, errors.New("error unmarshalling timestamp")
}