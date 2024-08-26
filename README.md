# Logbook (logbk)

A journaling application.

## Inspiration

The journal in Zed.

## MVP

A CLI that can create journal entries in markdown based on the date and time, similar to the journal in Zed.

**Example:**
`logbk --new-entry`

The folder structure should be `yyyy/mm/dd.md`, and the file should open with the cursor below a new timestamp.

## Research

- [X] Reading from a config file (Viper)
- [X] File path handling
- [X] Walking a directory
- [X] Time and Date
- [X] Opening a program in Go
- [X] Writing to a file
  - [ ] Now do it without overwriting
- [ ] Make neovim open with cursor on the last line
