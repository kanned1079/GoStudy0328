package main

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"GoStudy0328/SimpleTest/1_ExampleApi/user"
	"testing"
)

func TestUnscoped(t *testing.T) {
	dao.InitDB()
	user.HaveBeenDeleted(0, "dhc@ikanned.com")
}
