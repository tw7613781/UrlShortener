# UrlShortener

It's a exercise project for book **The Way To Go**

## It has two main functions:    

### 1. Add  
given a long URL, it returns a short version  

### 2. Redirect  
whenever a shortened URL is requested, it redirects the user to the original, long URL

## Involved with different version   

### Version 1   
A map and a struct are used, together with a Mutex from the sync package and a struct

### Version 2   
The data is made persistent because written to a file in gob-format.

### Version 3   
The application is rewritten with goroutines and channels (see chapter 14) 

### Version 4   
What to change if we want a json-version?

### Version 5   
A distributed version is made with the rpc protocol.