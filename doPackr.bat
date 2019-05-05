@echo off
setlocal enabledelayedexpansion

set basePath=%CD%
set templatesPath=%basePath%\templates
set datasetsPath=%basePath%\datasets
cd %GOPATH%\src\github.com\gobuffalo\packr\packr
go run main.go -i %templatesPath%
go run main.go -i %datasetsPath%

cd %templatesPath%
for /f %%i in ('dir /b *-packr.go') do (set templatesPackrFile=%%i)
call::deleteLineContainsDotGo %templatesPackrFile%

cd %datasetsPath%
for /f %%i in ('dir /b *-packr.go') do (set datasetsPackrFile=%%i)
call::deleteLineContainsDotGo %datasetsPackrFile%

goto:end


:deleteLineContainsDotGo
for /f "delims=" %%i in (%~1) do (
    set tmp=%%i
    if "!tmp:.go=!"=="!tmp!" (echo %%i>>%~1.doPackr.bak)
)
move %~1.doPackr.bak %~1
goto:eof


:end
