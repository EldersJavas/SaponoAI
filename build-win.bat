del /F /S /Q SaponoAI.exe
del /F /S /Q .\win64\SaponoAI.exe
go build -buildmode=exe
powershell Copy-Item SaponoAI.exe ./win64/
powershell Copy-Item ./design/*.png ./win64/res/
powershell Copy-Item ./design/*.bmp ./win64/res/
cd win64
SaponoAI.exe

