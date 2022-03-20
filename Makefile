MAKE=go build -o holidayapi
SERVER=server.go
YEAR2021=holiday_2021.go
YEAR2022=holiday_2022.go
YEAR2023=holiday_2023.go

holidayapi:
	$(MAKE) $(SERVER) $(YEAR2021) $(YEAR2022) $(YEAR2023)

test:
	go test

clean:
	rm -rf holidayapi
