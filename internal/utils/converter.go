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

func ConvertStringToPointerUUID(s string) *uuid.UUID {
	parsedUUID, err := uuid.Parse(s)
	if err != nil {
		return nil
	}
	return &parsedUUID
}

func CreateUuid() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}

func StringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
