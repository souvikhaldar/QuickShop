#!/bin/bash
cd $jio
git pull origin master
psql -U postgres -d quickshop -f setup.sql
cd $jio/cmd/http_server
go install
sudo systemctl restart quickshop.service
journalctl -u quickshop.service -f