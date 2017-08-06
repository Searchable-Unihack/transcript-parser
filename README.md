# transcript-parser

Transcript Parser takes the .json reply from Watson's transcription service, and outputs each block of text into an individual record in an array, with format:

```
[
  {
    "id": "3043092d-83bf-47ca-8466-586af288e869",
    "source": "audio",
    "text": "so even though ",
    "time": 48.95,
    "weight": 1.604
  }
]
```

To use this, provide a name for the output json, and a collection of input json files.  The name of the input file is used to fill the "id" field of the record.  For example:

./transcript-parser output.json *.json
