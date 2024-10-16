migrate_up=go run main.go migrate --direction=up

model/mock/mock_author_repository.go:
	mockgen -destination=model/mock/mock_author_repository.go -package=mock github.com/rhtyx/bayarind-service.git/model AuthorRepository

mockgen:
	@mockgen -destination=model/mock/mock_author_repository.go -package=mock github.com/rhtyx/bayarind-service.git/model AuthorRepository
	@mockgen -destination=model/mock/mock_book_repository.go -package=mock github.com/rhtyx/bayarind-service.git/model BookRepository
	@mockgen -destination=model/mock/mock_user_repository.go -package=mock github.com/rhtyx/bayarind-service.git/model UserRepository
	@mockgen -destination=model/mock/mock_session_repository.go -package=mock github.com/rhtyx/bayarind-service.git/model SessionRepository
	@mockgen -destination=model/mock/mock_jwt.go -package=mock github.com/rhtyx/bayarind-service.git/token JWTService

migrate:
	go run main.go migrate --direction=$(DIRECTION)

cert:
	@openssl genrsa -out cert/id_rsa.pri 4096
	@openssl rsa -in cert/id_rsa.pri -pubout -out cert/id_rsa.pub

run:
	go run main.go server

test:
	go test ./... -v -cover

.PHONY: cert model/mock/mock_author_repository.go mockgen