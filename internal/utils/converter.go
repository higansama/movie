package utils

import "github.com/google/uuid"

// create uuid to string function
func ConvertUUIDToString(uuid *uuid.UUID) *string {
	if uuid == nil {
		return nil
	}
	str := uuid.String()
	return &str
}

func CreateUuid() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}

func StringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
