# Energia API Mini Project

Energia API is a simple API designed to manage the usage of electrical devices in a household. Additionally, this API provides device usage recommendations based on current weather information.

---

## Available Features

### Core Features
- **Device Management**: Provides features to manage and estimate the cost of devices used.
- **Weather Forecast**: Provides weather forecasts.
- **Usage Recommendations**: Provides recommendations for device usage based on the weather.

### Technical Features
- **AI-Based Usage Recommendations**: Utilizes the OpenAI API to provide recommendations.
- **Secure Authentication**: JWT integration for secure authentication.
- **Email Notifications**: Notifications sent via SMTP.

---

## Entity Relationship Diagram (ERD)

Below is the ERD used in this project:

![ERD](images/Energia_ERD.png)

---

## High-Level Architecture (HLA)

Below is the High-Level Architecture (HLA) of the Energia API:

![HLA](images/Energia_HLA.png)

---

## Host for API Usage

The Energia API has been hosted on the cloud, allowing you to use it without any additional installation. You can access the API via the following host:

http://52.65.161.24

---

## API Documentation

To access the API documentation, follow these steps:

### 1. Clone this repository

Clone the repository into your local directory using the following command:

```bash  
git clone <repository_url>
```

### 2. Create a .env file
Create a .env file in the root directory of the project and fill it with the following configuration:

```plaintext
DATABASE_HOST=""
DATABASE_PORT=""
DATABASE_USER=""
DATABASE_PASSWORD=""
DATABASE_NAME=""
JWT_SECRET_KEY=""
OPENWEATHER_API_KEY=""
OPENAI_API_KEY=""
MAIL_USER=""
MAIL_PASSWORD=""
MAIL_HOST=""
MAIL_PORT=""
```

### 3. Run the application
Run the following command to start the application:

```bash
go run main.go
```

### 4. Access API Documentation
Open your browser and access the API documentation via the following URL:

```plaintext
http://{{host}}/swagger/index.html#
```

Note: Replace {{host}} with the appropriate host address, for example, localhost:8080 if the application is running locally.
