@SET WINDOWS="windows"
@SET PACKAGEPIN="github.com/vincent87720/AutoPunch/punchin"
@SET PACKAGEPOUT="github.com/vincent87720/AutoPunch/punchout"

::build
go build -o bin/punchin.exe %PACKAGEPIN%
go build -o bin/punchout.exe %PACKAGEPOUT%