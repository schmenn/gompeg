# gompeg
a *wrapper* for ffmpeg

## installation

`go get -u github.com/Schmenn/gompeg`

## usage

### converting files

```go
file, err := gompeg.Open("input.mp4")
if err != nil {
	// handle error
}

// Set output file
file.SetOutputFile("output.webm")

// Execute
err = file.Do()
if err != nil {
	// handle error
}
```

### cut video

```go
file, err := gompeg.Open("input.mp4")
if err != nil {
	// handle error
}

// Set output file
file.SetOutputFile("output.webm")

// Set start time to 5 seconds and 250 milliseconds
file.SetStartTime(time.Second*5 + time.Millisecond*250)

// Set end time to 10 and a half seconds
file.SetEndTime(time.Second*10 + time.Millisecond*500)

// Execute
err = file.Do()
if err != nil {
	// handle error
}
```
