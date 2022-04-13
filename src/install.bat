mkdir GoDBMS

cd CLI
go build
move CLI.exe ..\GoDBMS

cd ..

cd DBMS
go build
move .\GoDBMS.exe ..\GoDBMS