# wave

## nosilence

### Parameters

`-delay` duration

Delay to pause after level felt down threshold (default 1s).

`-fall` float

Level EMA fall speed (default 0.0005).

`-log-interval` duration

Current level EMA value logging interval (default 1s).

`-multiply` float

Multiply level (default 1).

`-power` float

Power level (default 1).

`-rise` float

Level EMA rise speed (default 0.9).

`-threshold` float

Level threshold to pause (default 0.007).

Two  positional parameters are input and output wav-files respectfully.
Two optional positional parameters are the input and output wav-files,
respectively (both may be omitted or replaced with `-` to use standard IO).

### Example

```shell
arecord -f FLOAT_LE -r 22050 -c 1 | nosilence --threshold 0.03 - - | ffmpeg -f wav -i pipe: -codec:a libmp3lame -b:a 96k -f mp3 rec.mp3
```
