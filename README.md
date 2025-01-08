# Golang Server for Vue App and Google Maps Integration

This is a Golang server that provides backend support for a Vue application. The server includes functionality for managing display settings, retrieving GPS information, and integrating with the Google Maps API.

## Features

### 1. Save Settings to Display Fields in the Vue App

The server allows saving customizable settings to control the visibility of fields in the Vue application. These settings determine which data fields are displayed on the frontend.

### 2. Get Settings to Display Fields in the Vue App

Retrieve previously saved settings to dynamically render fields in the Vue application. This ensures a consistent and user-configurable display experience.

### 3. Get Google Maps API URL for Map Rendering

The server provides a secure endpoint to fetch the Google Maps API URL with the API key. This approach helps prevent exposing sensitive credentials directly in the frontend code.

### 4. Get GPS Info of Devices from an External Website

Fetch GPS information of devices associated with a user's account from an external API. This data is utilized for map rendering and device tracking in the Vue app.

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/codelur/golang_gps.git
   Install Dependencies Ensure you have Go installed.
   ```

Set Environment Variables Create a .env file in the project root and configure the following variables:

API_KEY=your-google-maps-api-key
EXTERNAL_API_URL=your-external-website-url
Run the Server Start the Golang server:

go run main.go

API Endpoints

1. Save Settings
   Endpoint: POST /api/savesettings
   Description: Save field visibility settings for the Vue app.

2. Get Settings
   Endpoint: GET /api/getsettings
   Description: Retrieve field visibility settings.

3. Get Google Maps API URL
   Endpoint: GET /api/map
   Description: Fetch the Google Maps API URL with the API key.

4. Get GPS Info
   Endpoint: GET /api/getDevicesInfo
   Description: Retrieve GPS data of devices from an external API.

   Usage

   Integrate this server with a Vue frontend to manage dynamic field rendering and Google Maps integration.
   Use the endpoints to store settings, retrieve settings, and fetch GPS data securely.

   License

   This project is licensed under the MIT License.

Contributions

Contributions are welcome! Feel free to fork this repository, make changes, and submit a pull request.

Contact

For inquiries or support, please contact codelur@gmail.com.
