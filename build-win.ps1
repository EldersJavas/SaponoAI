go build -buildmode=exe
Copy-Item SaponoAI.exe ./win64/
cd win64
Start-Process SaponoAI.exe
