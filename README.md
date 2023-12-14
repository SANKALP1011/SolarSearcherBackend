# SOLAR SEARCHECR BACKEND

- This is the backend for the solar searcher api built using golang and fiber.

# Tech Stack

- Golang
- Fibre (Golang rest api framework)

# Deployement

- Render

# Live URL

https://solar-geek-api.onrender.com

# Folder Strcuture

- Main.go (Main server file)

  ```go
      - main.go
  ```

- Controller

  ```go
     - space.controller.go
  ```

- Error ( Custom error file )

  ```go
     - customHandler.error.go
  ```

- Helper ( functions that would be used in the controllers )

  ```go
     ## This helper converts the user location into coordinates.
     - geoCordinate.helper.go
     ## This helper helps in performing the analysis over the weather conditions.
     - weatherAnalysis.helper.go
  ```

- Model ( Contains the model that would help in mapping through api response )

  ```go
     - spaceData.model.go
  ```

- Services ( This would be used for making api calls to the third party space api's )

  ```go
     - spaceApis.service.go
  ```

- Routes

  ```go
    - space.router.go
  ```

- Utils ( This contains the base url for the various third parties api that are being used in our backend )

  ```go
     - apis.utils.go
  ```

# Endpoints

```go

app.Get("/v1/getIssLocation", Controller.GetISSLocation)
app.Get("/v1/getMarsRoverPic", Controller.GetMarsRoverImage)
app.Get("/v1/getApodApiImage", Controller.GetApodImages)
app.Get("/v1/performWeatherAnalysis", Controller.PerformWeatherAnalysis)

```
