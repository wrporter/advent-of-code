## Day 12

### Part 1

Oh nice, I don't even have to parse the JSON! I can just use a regular expression. Muahhahaha!

### Part 2

Okay, I just had to drop my idea of doing **everything** in Go, because now I had to parse the JSON. In just the time (about 10 minutes) it took me to look at how I might go about this in a decently readable way in Go, I wrote the solution in JavaScript. Maybe I'll look at this again in Go, I just didn't think it was worth the time.

Ah, okay, that really wasn't so bad. I wasn't sure if Go would allow me to determine the difference between a generic object and an array, but it turns out that their object for JSON is a `map[string]interface{}`. My code is almost exactly the same as the JavaScript version.
