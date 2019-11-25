.PHONY: all

all:
	go build \
		-ldflags "-X \"github.com/nwpc-oper/nmic-message-client/cmd/app.Version=${VERSION}\" \
        -X \"github.com/nwpc-oper/nmic-message-client/cmd/app.BuildTime=${BUILD_TIME}\" \
        -X \"github.com/nwpc-oper/nmic-message-client/cmd/app.GitCommit=${GIT_COMMIT}\" " \
		-o ${BIN_PATH}/nmic_message_client \
		main.go