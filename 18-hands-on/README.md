
### start mongo
brew services start mongodb-community@6.0

### authenticate on new terminal
mongosh --port 27017  --authenticationDatabase "admin" -u "ifeanyi" -p
