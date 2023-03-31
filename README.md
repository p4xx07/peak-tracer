# Peak Tracer
Peak Tracer is a command-line tool that helps you find the loudest parts of an audio file and export them as JSON or concatenate them into a new file.

# Requirements
Peak Tracer requires FFmpeg to be installed on your system.

# Installation
To install Peak Tracer, you need to have Go installed on your system. Then, run the following command:

```go
go get -u github.com/Paxx-RnD/peak-tracer
```

This will download the source code and build the binary for you. You should now be able to run the peak-tracer command in your terminal.

# Usage
Peak Tracer can be used with the following command line options:


# Parameters

```shell
  -i string
        The path to the input audio file (required)
  -o string
        The path to the output JSON file (required)
  -after float
        The number of seconds after each peak to include (default 1)
  -before float
        The number of seconds before each peak to include (default 1)
  -concat string
        The path to the output file for concatenation
  -samples int
        The number of samples per peak (default 48000)
  -target int
        The target duration of the output file in seconds (default 60)
```
        
# Finding peaks
To find the peaks in a media file, you can use the following command:

```shell
peak-tracer -i input.mp3 -o output.json
```

This will analyze the input file and export the peaks as a JSON file.

# Concatenating peaks
To concatenate the loudest parts of a media file into a new file, you can use the following command:

```shell
peak-tracer -i input.mp3 -concat output.mp3
```

This will analyze the input file and create a new file that contains the concatenated peaks.

# Customizing the behavior
You can customize the behavior of Peak Tracer with the following options:

```shell
-after: The number of seconds after each peak to include (default: 1).
-before: The number of seconds before each peak to include (default: 1).
-samples: The number of samples per peak (default: 48000).
-target: The target duration of the output file in seconds (default: 60).
```

# License
Peak Tracer is licensed under the MIT License. See the LICENSE file for more information.

# Credits
Peak Tracer was created by Paxx R&D.
