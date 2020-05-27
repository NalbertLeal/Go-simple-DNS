package manipulatePacket

import (
	"errors"
	"strconv"
)

func QueryTypeToInt16(qt string) (int16, error) {
	switch qt {
	case "A":
		return 1, nil
	default:
		number, err := strconv.Atoi(qt)
		if err != nil {
			return 0, errors.New("> Error: Query type isn't a number.\n" + err.Error())
		}
		return int16(number), nil
	}
}

func QueryTypeFromInt16(qt int16) string {
	switch qt {
	case 1:
		return "A"
	default:
		return strconv.Itoa(int(qt))
	}
}
