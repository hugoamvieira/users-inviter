# Intercom Users Inviter

This program takes a list of users and it calculates which ones you should invite based on
parameters on your configuration file.
This program is ready (designed) to be made concurrent (reading from file and parsing entries at the same time), but it is not doing so right now due to time constraints.

This program takes a JSON configuration file (normally under config/config.json) with 4 keys:
- "latitude": The latitude to calculate distance with the users'.
- "longitude": The longitude to calculate distance the users'.
- "distance_threshold_km": The max distance, in km, that a user must be from the aforementioned latitude and longitude.
- "users_file_path": The filepath for the users JSON file.

## Run it:
1. Run `dep ensure`
2. Have a[n incorrect ;)] JSON file with users somewhere on your system, specified as follows: 
```json
{"latitude": string, "user_id": int, "name": string, "longitude": string}
{"latitude": string, "user_id": int, "name": string, "longitude": string}
{"latitude": string, "user_id": int, "name": string, "longitude": string}
...
```
3. Specify your config on `config/config.json` (you can use the default)
```json
{
	"latitude": float,
	"longitude": float,
	"distance_threshold_km": float,
	"users_file_path": string
}
```

4. Run `go run .`
