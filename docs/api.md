## Base URL

The base URL for all endpoints is `/api/license`.

## Endpoints

### Create License

- **URL**: `/create`
- **Method**: `POST`
- **Description**: Creates a new license.
- **Request Body**:
  ```json
  {
    "user_email": "example@example.com",
    "product_id": "12345"
  }
  ```
- **Success Response**:
  - **Code**: `201 Created`
  - **Content**:
    ```json
    {
      "user_email": "example@example.com",
      "product_id": "12345",
      "license_key": "ABC123DEF456GHIJ"
    }
    ```
- **Error Response**:
  - **Code**: `400 Bad Request`
  - **Content**:
    ```json
    {
      "error": "Invalid request body"
    }
    ```
  - **Code**: `500 Internal Server Error`
  - **Content**:
    ```json
    {
      "error": "Internal server error"
    }
    ```

### Validate License

- **URL**: `/validate`
- **Method**: `POST`
- **Description**: Validates a license.
- **Request Body**:
  ```json
  {
    "license_key": "ABC123DEF456GHIJ",
    "product_id": "12345"
  }
  ```
- **Success Response**:
  - **Code**: `200 OK`
  - **Content**:
    ```json
    {
      "user_email": "example@example.com",
      "product_id": "12345",
      "license_key": "ABC123DEF456GHIJ"
    }
    ```
- **Error Response**:
  - **Code**: `400 Bad Request`
  - **Content**:
    ```json
    {
      "error": "Invalid request body"
    }
    ```
  - **Code**: `404 Not Found`
  - **Content**:
    ```json
    {
      "error": "License not found"
    }
    ```
  - **Code**: `500 Internal Server Error`
  - **Content**:
    ```json
    {
      "error": "Internal server error"
    }
    ```

