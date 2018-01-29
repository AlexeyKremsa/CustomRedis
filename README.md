**API**

- ../
GET  
Returns a simple response to check if server is alive

-  ../str
POST
Set string key and value
Example:
```json
{
	   "key": "k1",
	   "strValue": "v1",
	   "expirationSec": 955
}
```

-  ../strnx
POST
Set string key and value in if key doesn't exist
Example:
```json
{
	   "key": "k1",
	   "strValue": "v1",
	   "expirationSec": 955
}
```

- ../str
GET
Parameters: key
Returns value stored by specified key
Example: ../str?key=k1

- ../list
POST
Set list key and value
Example:
```json
{
	   "key": "k1",
	   "listValue": ["v1", "v2", "v3"],
	   "expirationSec": 955
}
```

- ../list
GET
Parameters: key
Get list by key
Example:
..list?key=k1

- ../listinsert
POST
Add elements to the end of the list. Returns list length.
:
```json
{
   "key": "k1",
   "listValue": ["q1", "q2", "q3"]
}
```

- ../listpop
GET
Parameters: key
Removes and returns the last element of the list stored at key
Example:
../listpop?key=k1


- ../listindex
GET
Parameters: key, index
Returns the element at index in the list stored at key
Example:
../listindex?key=k1&index=1

- ../map
POST
Set map key and value
Example:
```json
{
   "key": "k1",
   "mapValue": { "m1" : "v1"},
   "expirationSec": 110
}
```

- ../map
GET  
Parameters: key  
Get map by key  
Example:  
../map?key=k1  

- ../del
DELETE
Parameters: key  
Delete value by key
Example:
../del?key=k1

- ../keys
GET  
Get all keys  
Example:  
../keys  




