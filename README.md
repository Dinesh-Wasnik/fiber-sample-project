# fiber-demo-service v0.2
- Go version  1.26.0
- Fiber version /v3 v3.1.0
- protopuf
- GRPC 
- Swag package added for swagger

# Prerequisite knowledge
- Docker 
- GRPC
- Protopuf

# How to use 
- Create docker network, ```docker network create game-local```, can give any name to network.
- If newtwork name changed, then need to mention it in compose file.
- Create logs folder with 777 permission at your system, path should match to this path ```../../logs/``` 
 or you can create logs folder as per your choice and update path this path - ```../../logs/${APP_NAME}:/app/logs```.
- The application folder is automatically created by docker inside logs folder.
- After docker create folder  in host machine , give that folder 775 permission.

