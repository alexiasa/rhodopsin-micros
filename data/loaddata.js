var db = connect('db:27017/rhodopsin');

mongoimport --jsonArray --db test --collection docs --file example2.json