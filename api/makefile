BINARY_NAME=control-panel.exe
 
build:
	go build -o ${BINARY_NAME} .

	signtool sign /f "mycert.pfx" /fd SHA256 /p $(password) /t http://timestamp.digicert.com ${BINARY_NAME}
 
run:
	go build -o ${BINARY_NAME} .
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
