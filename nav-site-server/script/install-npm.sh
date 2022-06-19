#!/bin/sh
cd ../../nav-site-web
echo "current path: %cd%"
npm run build
echo "npm run build end..."

cd ../nav-site-server/server
echo "current path: %cd%"
rm -rf static/

\cp ../../nav-site-web/dist static/
echo "web build end"