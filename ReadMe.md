Very basic go api to retrieve api keys from files. Simply save the keys you need you need as a json file in the same
directory as you pull this api.

Format required for json file
{
"secretKey": "INSERT_KEY_VALUE_HERE"
}

Then, go run . this project and localhost:8080/secretKey/fileName will retrieve your secret key as a json object
