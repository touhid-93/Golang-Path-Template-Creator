package consts

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
)

var (
	GlobalMutex        = &sync.Mutex{}
	SlugForbiddenArray = [256]rune{
		constants.ExclamationChar:    constants.One,
		'`':                          constants.One,
		'@':                          constants.One,
		'#':                          constants.One,
		'%':                          constants.One,
		'$':                          constants.One,
		'^':                          constants.One,
		'&':                          constants.One,
		'*':                          constants.One,
		'(':                          constants.One,
		')':                          constants.One,
		'{':                          constants.One,
		'}':                          constants.One,
		'[':                          constants.One,
		']':                          constants.One,
		'\\':                         constants.One,
		'/':                          constants.One,
		',':                          constants.One,
		'.':                          constants.One,
		constants.TabByte:            constants.One,
		constants.LineFeedUnixByte:   constants.One,
		constants.TabVByte:           constants.One,
		constants.FormFeedByte:       constants.One,
		constants.CarriageReturnByte: constants.One,
		constants.SpaceByte:          constants.One,
		0x85:                         constants.One, // reference : https://bit.ly/2JWdIoj
		0xA0:                         constants.One, // reference : https://bit.ly/2JWdIoj
	}
)
