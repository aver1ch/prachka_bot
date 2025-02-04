package errs

import "errors"

var (
	ErrSendMessage         = errors.New("не удалось отправить сообщение\n")
	ErrRoomNumber          = errors.New("отправленное сообщение не является номером комнаты\n")
	ErrConnectionToDB      = errors.New("не удалось подключиться к базе данных\n")
	ErrInsertingDataFromDB = errors.New("не удалось записать данные в базу данных\n")
	ErrPullingDataFromDB   = errors.New("не удалось подтянуть данные из базы данных\n")
	ErrAlreadyAutorized    = errors.New("пользователь пытался верифицироваться, хотя уже верифицирован")
	ErrAuthorizationError  = errors.New("неавторизованный пользователь пытается выполнить действие")
	ErrCallbackQuery       = errors.New("ошибка обработки кнопки")
)
