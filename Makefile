CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	bash ${CURRENT_DIR}/script/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
server:
	go run cmd/main.go

pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge

create:
	migrate create -ext sql -dir migrations -seq create_user_table

up-version:
	migrate -source file:./migrations/ -database 'postgres://postgres:compos1995@localhost:5432/userdb?sslmode=disable' up

down-version:
	migrate -source file:./migrations/ -database 'postgres://postgres:compos1995@localhost:5432/userdb?sslmode=disable' down

# run:
# 	go run cmd/main.go

