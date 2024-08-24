enum AccountPermission {
  Base          = 0b00000,
  Read          = 0b00001,
  History       = 0b00010,
  Withdraw      = 0b00100,
  Send          = 0b01000,
  Subscription  = 0b10000,
  All           = 0b11111
}