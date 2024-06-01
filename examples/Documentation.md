# Project Documentation Manual
## 1. Introduction
### Project Name
Vedi-DBQ Automation

### Brief Overview
Vedi-DBQ Automation is an advanced web application built to manage and automate the Disability Benefits Questionnaires (DBQ) process. The application enables users to submit, monitor, and manage DBQs efficiently. The system leverages role-based access controls to ensure secure operations and detailed task assignments among veterans, doctors, and the VeDi team.

### Technologies Used
- **Node.js**
- **Express.js**
- **MongoDB (Mongoose ORM)**
- **JSON Web Tokens (JWT)**
- **TypeScript**
- **Axios**
- **Bcrypt**
- **UUID**

---

## 2. API Reference
### Application Route
- **Update Application Status**
  - **Endpoint URL:** `/application/:id`
  - **Method:** PUT
  - **Request Parameters:**
    - `id` - MongoDB Object ID of the application (URL Parameter)
    - `status` - New status of the application (Request Body)
  - **Example Request:**
    ```json
    PUT /application/60c72b2f5f1b2c001cf74611
    {
      "status": "Approved"
    }
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Application updated",
      "data": { /* Application Object */ }
    }
    ```

- **Get Applications by User**
  - **Endpoint URL:** `/user/:id`
  - **Method:** GET
  - **Request Parameters:**
    - `id` - MongoDB Object ID of the user (URL Parameter)
  - **Example Request:**
    ```json
    GET /user/60c72b2f5f1b2c001cf74611
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Applications retrieved",
      "data": [/* Applications Array */]
    }
    ```

- **Get User Progress**
  - **Endpoint URL:** `/application/:id/progress`
  - **Method:** GET
  - **Request Parameters:**
    - `id` - MongoDB Object ID of the application (URL Parameter)
  - **Example Request:**
    ```json
    GET /application/60c72b2f5f1b2c001cf74611/progress
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "User progress retrieved",
      "data": { /* Progress Data */ }
    }
    ```

- **Get User Responses**
  - **Endpoint URL:** `/application/:id/responses`
  - **Method:** GET
  - **Request Parameters:**
    - `id` - MongoDB Object ID of the application (URL Parameter)
  - **Example Request:**
    ```json
    GET /application/60c72b2f5f1b2c001cf74611/responses
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "User progress retrieved",
      "data": { /* Responses Data */ }
    }
    ```

- **Get Next Recommended Action**
  - **Endpoint URL:** `/application/next-action`
  - **Method:** GET
  - **Request Parameters:**
    - None
  - **Example Request:**
    ```json
    GET /application/next-action
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "",
      "data": { /* Action Data */ }
    }
    ```

- **Update Application State**
  - **Endpoint URL:** `/application/:id/state`
  - **Method:** PUT
  - **Request Parameters:**
    - `id` - MongoDB Object ID of the application (URL Parameter)
    - `status` - New state ID of the application (Request Body)
  - **Example Request:**
    ```json
    PUT /application/60c72b2f5f1b2c001cf74611/state
    {
      "status": "ApprovedByDoctor"
    }
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Application status updated"
    }
    ```

---

### Auth Route
- **Login**
  - **Endpoint URL:** `/auth/login`
  - **Method:** POST
  - **Request Parameters:**
    - `email` - User's email address (Request Body)
    - `password` - User's password (Request Body)
  - **Example Request:**
    ```json
    POST /auth/login
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Logged in successfully",
      "data": {
        "user": { /* User Object */ },
        "accessToken": "JWT_ACCESS_TOKEN"
      }
    }
    ```

- **Forgot Password**
  - **Endpoint URL:** `/auth/forgot-password`
  - **Method:** POST
  - **Request Parameters:**
    - `email` - User's email address (Request Body)
    - `redirection_url` - URL to redirect after password reset (Request Body)
  - **Example Request:**
    ```json
    POST /auth/forgot-password
    {
      "email": "user@example.com",
      "redirection_url": "https://example.com/reset-password"
    }
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Password reset link sent successfully"
    }
    ```

- **Reset Password**
  - **Endpoint URL:** `/auth/reset-password`
  - **Method:** POST
  - **Request Parameters:**
    - `resetToken` - Token for password reset (Request Body)
    - `newPassword` - New password (Request Body)
  - **Example Request:**
    ```json
    POST /auth/reset-password
    {
      "resetToken": "UUID_RESET_TOKEN",
      "newPassword": "newPassword123"
    }
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Password updated successfully"
    }
    ```

---

### Intake Route
- **Submit Intake Form**
  - **Endpoint URL:** `/intake/submit`
  - **Method:** POST
  - **Request Parameters:**
    - `formId` - ID of the form being submitted (Request Body)
    - `userId` - ID of the user submitting the form (Request Body)
    - `fields` - Field values of the form (Request Body)
    - `isDraft` - Whether the form is submitted as a draft (Request Body)
  - **Example Request:**
    ```json
    POST /intake/submit
    {
      "formId": 1,
      "userId": "60c72b2f5f1b2c001cf74611",
      "fields": [/* Fields Array */],
      "isDraft": false
    }
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "DBQ Submitted",
      "data": { /* Submission Data */ }
    }
    ```

- **Delete Intake Form**
  - **Endpoint URL:** `/intake/:intakeId`
  - **Method:** DELETE
  - **Request Parameters:**
    - `intakeId` - ID of the intake form (URL Parameter)
  - **Example Request:**
    ```json
    DELETE /intake/60c72b2f5f1b2c001cf74611
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Intake form deleted"
    }
    ```

### DBQ Route
- **Get DBQ Gen Response**
  - **Endpoint URL:** `/dbq`
  - **Method:** GET
  - **Request Parameters:**
    - `formId` - ID of the form (Query Parameter, optional)
    - `transactionId` - Transaction ID (Query Parameter, optional)
    - `userId` - User ID (Query Parameter, optional)
  - **Example Request:**
    ```json
    GET /dbq?formId=1&transactionId=60c72b2f5f1b2c001cf74611
    ```
  - **Example Response:**
    ```json
    {
      "success": true,
      "message": "Data retrieved",
      "data": [/* DBQ Gen Responses Array */]
    }
    ```

---

## 3. Class and Function Documentation
### Application Route
- **Function: `router.put("/:id", ... )`**
  - **Description:** Endpoint to update the status of an application by its ID.
  - **Parameters:**
    - `req.params.id` (string) - Application ID
    - `req.body.status` (string) - New Status
  - **Return Type:** JSON Response

- **Function: `getMissingFormsAndUpdateDb`**
  - **Description:** Helper function to get missing forms for a user and update the DB.
  - **Parameters:**
    - `userId` (mongoose.Types.ObjectId) - User ID
    - `formId` (number) - Form ID
    - `wantedStatusId` (string) - Desired status ID
  - **Return Type:** Promise<number[]>

### Auth Route
- **Function: `router.post("/login", ... )`**
  - **Description:** Endpoint for user login.
  - **Parameters:**
    - `req.body.email` (string) - User's email
    - `req.body.password` (string) - User's password
  - **Return Type:** JSON Response

- **Function: `router.post("/forgot-password", ... )`**
  - **Description:** Endpoint to generate unique URL for resetting password.
  - **Parameters:**
    - `req.body.email` (string) - User's email
    - `req.body.redirection_url` (string) - Redirection URL
  - **Return Type:** JSON Response

### Intake Route
- **Function: `router.delete("/:intakeId", ... )`**
  - **Description:** Endpoint to delete an intake form.
  - **Parameters:**
    - `req.params.intakeId` (string) - Intake form ID
  - **Return Type:** JSON Response

- **Function: `generateDbqResponse`**
  - **Description:** Helper function to generate DBQ response.
  - **Parameters:**
    - `formId` (number) - Form ID
    - `userId` (mongoose.Types.ObjectId) - User ID
    - `userIntakeResponse` (UserIntakeResponseType) - User Intake Response Object
    - `fields` (any) - Fields data
  - **Return Type:** Promise<object>

### DBQ Route
- **Function: `router.get("/", ... )`**
  - **Description:** Endpoint to get DBQ Gen Response.
  - **Parameters:**
    - `req.query.formId` (string, optional) - Form ID
    - `req.query.transactionId` (string, optional) - Transaction ID
    - `req.query.userId` (string, optional) - User ID
  - **Return Type:** JSON Response

- **Function: `highlightChangedFields`**
  - **Description:** Helper function to highlight changed fields between two versions.
  - **Parameters:**
    - `newVersion` (object) - New version of the DBQ
    - `oldVersion` (object) - Old version of the DBQ
  - **Return Type:** object

---

## 4. Configuration and Deployment
### Configuration File Explanations
- **`config.ts`**
  - Contains JWT configuration.
  - Example:
    ```typescript
    export const JwtConfig = {
      accessTokenExpiresIn: 3600, // 1 hour
    };
    ```

### Deployment Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/vedi-dbq-automation.git
   ```
2. Navigate to the project directory:
   ```bash
   cd vedi-dbq-automation
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Set environment variables:
   - Create a `.env` file and add necessary environment variables.
     ```plaintext
     PORT=3000
     DB_URI=mongodb://localhost:27017/yourdb
     AUTH_SECRET=your_secret_key
     DBQ_GEN_API_URL=http://dbq-gen-api-url
     ```
5. Start the server:
   ```bash
   npm run start
   ```

---

## 5. FAQs and Troubleshooting
### Common Issues and Their Resolutions

- **Issue:** Application not found.
  - **Resolution:** Ensure the application ID passed in the request is a valid MongoDB Object ID and exists in the database.

- **Issue:** Incorrect or missing request parameters.
  - **Resolution:** Ensure all required request parameters are provided and correctly formatted.

- **Issue:** User not authenticated.
  - **Resolution:** Ensure the JWT token is passed in the request headers under `Authorization` and is valid.

- **Issue:** Internal server error.
  - **Resolution:** Check the server logs for detailed error messages and stack traces to identify the issue.

- **Issue:** Reset token expired.
  - **Resolution:** Request a new password reset link as the token is only valid for a limited time.

---

This documentation manual provides an organized and comprehensive overview of the Vedi-DBQ Automation project. It covers the API endpoints, class and function details, configuration, deployment, and troubleshooting steps, ensuring a user-friendly and professional experience for developers and users alike.