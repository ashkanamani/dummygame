package message

import "github.com/ashkanamani/dummygame/internal/entity"

func MyInfoMessage(acc entity.Account) string {
	return `🏯Welcome %s 
Can I do anything for you?`
}
