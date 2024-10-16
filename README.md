#### I. Running the service
1. Start test using `make test`.
2. Tidy the dependencies using `go mod tidy`.
3. Build the image using `docker compose build`.
4. Run the image using `docker compose up -d`.
5. Test the api using postman or else.
6. Remember to add **X-HMAC** in the header of each API call.
