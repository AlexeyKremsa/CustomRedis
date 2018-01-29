###API

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
../list?key=k1




