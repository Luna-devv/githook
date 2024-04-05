package utils

type Colors struct {
	Default int
	Success int
	Error   int
}

func GetColors() Colors {
	return Colors{
		Default: 0x895af6,
		Success: 0x59f68e,
		Error:   0xf65959,
	}
}
