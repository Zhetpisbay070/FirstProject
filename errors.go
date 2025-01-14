package entity

import "errors"

var ProductDoesNotExistError = errors.New("product does not exist")
var ErrOrderCannotBeCancelled = errors.New("you cannot cancel your order at this stage")
