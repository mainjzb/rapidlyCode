@echo off
route add 192.168.30.1
rundll32.exe %~dp0/cmroute.dll,SetRoutes /STATIC_FILE_NAME  %~dp0/link.rul /DONT_REQUIRE_URL /IPHLPAPI_ACCESS_DENIED_OK
pause