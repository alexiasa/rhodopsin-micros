var db = connect('localhost:27017/rhodopsin');

// TODO: remove this sample data load for production use or replace with a restore...
db.ips.insertMany([{'ipaddr': '47.244.9.129'},{'ipaddr': '210.222.113.153'},{'ipaddr': '185.244.25.177'},{'ipaddr': '91.226.210.84'},{'ipaddr': '36.66.220.173'},{'ipaddr': '88.99.126.189'},{'ipaddr': '94.46.14.140'},{'ipaddr': '94.46.14.140'},{'ipaddr': '88.99.126.189'},{'ipaddr': '61.6.247.172'},{'ipaddr': '61.6.247.172'},{'ipaddr': '201.131.244.36'},{'ipaddr': '164.52.12.162'},{'ipaddr': '164.52.12.162'},{'ipaddr': '103.74.69.173'},{'ipaddr': '66.117.9.22'},{'ipaddr': '185.244.25.177'},{'ipaddr': '185.244.25.177'},{'ipaddr': '185.244.25.177'},{'ipaddr': '185.244.25.177'},{'ipaddr': '185.244.25.177'},{'ipaddr': '185.244.25.177'},{'ipaddr': '27.72.17.135'},{'ipaddr': '27.72.17.135'},{'ipaddr': '27.72.17.135'}])

