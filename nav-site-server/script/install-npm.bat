::, please run the *.bat in the document

@echo off
cd ../nav-site-web
echo "current path: %cd%"
call npm run build
echo "npm run build end..."

cd ../nav-site-server/server
echo "current path: %cd%"
rmdir /S/Q static

xcopy ..\..\nav-site-web\dist static\  /S /F
echo "web build end"