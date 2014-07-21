package main

type Pin struct {
  Status int64
  Id int64
}

func NewPin(id int64, status int64) *Pin {
  return &Pin{
    Id: id,
    Status: status,
  }
}
