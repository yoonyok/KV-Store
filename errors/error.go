package errors

import "fmt"

// Thrown when client reads a value from non-leader, tells client to request again to leader
// e: leader's address
type NonLeaderReadError string

func (e NonLeaderReadError) Error() string {
	return fmt.Sprintf("Read value from non-leader store. Please request again to leader address [%s]", string(e))
}

// Thrown when client reads from a key that does not exist
// e: key
type KeyDoesNotExistError string

func (e KeyDoesNotExistError) Error() string {
	return fmt.Sprintf("Key [%v] does not exist in database ", e)
}

// Thrown when a store is disconnected
// e: address of disconnected entity
type DisconnectedError string

func (e DisconnectedError) Error() string {
	return fmt.Sprintf("[%s] is disconnected", string(e))
}
