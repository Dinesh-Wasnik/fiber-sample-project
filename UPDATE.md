# Update 
# March 5 ,2026
- Added automatically database creation from env file setting.
- Copy setting from .env.sample file to .env file 
- Then uncomment below code in main.go file .
	```
        config.Connect()
        config.Migration()

    ```
- Added demomoel for example.