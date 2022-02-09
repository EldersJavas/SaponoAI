del /F /S /Q SaponoAI.exe
del /F /S /Q .\win64\SaponoAI.exe
go build -buildmode=exe
powershell Copy-Item SaponoAI.exe ./win64/
cd win64
SaponoAI.exe

