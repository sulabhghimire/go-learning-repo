Command: protoc --go_out=. <path-to-proto-files>

--go_out=. means return file in same folder where command is run

If we have imports from other proto files then the commands look like

protoc -I=<folder-that-contains-all-proto-files> --go_out=. <path-to-proto-files>

multiple files can be done by seperating with comma
