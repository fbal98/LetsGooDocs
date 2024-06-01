# Incoming Orders API Server Documentation

## Introduction

- **Project Name**: Incoming Orders API Server
- **Brief Overview**: This project provides an API to manage incoming orders for a store. It includes functionalities such as creating, updating, and deleting orders, managing products, categories, notifications, and handling income reports.
- **Technologies Used**: Node.js, Express.js, Prisma, TypeScript, APNs (Apple Push Notification service)

## API Reference

### Categories API

#### Get Categories
- **Endpoint URL**: `/categories`
- **Method**: GET
- **Request Parameters**: None
- **Example Request**: 
    ```
    GET /categories
    ```
- **Example Response**:
    ```json
    {
        "data": [
            {
                "id": "category_id_1",
                "name": "Category 1",
                "_count": {
                    "products": 10
                }
            },
            ...
        ]
    }
    ```

#### Get Products by Category
- **Endpoint URL**: `/categories/:id/products`
- **Method**: GET
- **Request Parameters**: 
    - `id`: Category ID
- **Example Request**: 
    ```
    GET /categories/123/products
    ```
- **Example Response**:
    ```json
    {
        "data": {
            "id": "123",
            "name": "Category 1",
            "products": [
                {
                    "id": "product_id_1",
                    "name": "Product 1",
                    ...
                },
                ...
            ]
        }
    }
    ```

#### Get Category Count
- **Endpoint URL**: `/categories/count`
- **Method**: GET
- **Request Parameters**: None
- **Example Request**: 
    ```
    GET /categories/count
    ```
- **Example Response**:
    ```json
    {
        "categoriesNumber": 5
    }
    ```

#### Create Category
- **Endpoint URL**: `/categories`
- **Method**: POST
- **Request Parameters**:
    - `name`: Category name
    - `storeId`: Store ID
- **Example Request**: 
    ```
    POST /categories
    {
        "name": "New Category",
        "storeId": "store_id_1"
    }
    ```
- **Example Response**:
    ```json
    {
        "data": {
            "id": "new_category_id",
            "name": "New Category",
            "storeId": "store_id_1"
        }
    }
    ```

#### Update Category
- **Endpoint URL**: `/categories/:id`
- **Method**: PATCH
- **Request Parameters**:
    - `id`: Category ID
    - `name`: New category name
    - `storeId`: Store ID
- **Example Request**: 
    ```
    PATCH /categories/123
    {
        "name": "Updated Category",
        "storeId": "store_id_1"
    }
    ```
- **Example Response**:
    ```json
    {
        "msg": "The category has been updated",
        "data": {
            "id": "123",
            "name": "Updated Category",
            "storeId": "store_id_1"
        }
    }
    ```

#### Delete Category
- **Endpoint URL**: `/categories/:id`
- **Method**: DELETE
- **Request Parameters**:
    - `id`: Category ID
- **Example Request**: 
    ```
    DELETE /categories/123
    ```
- **Example Response**:
    ```json
    {
        "msg": "Category has been deleted"
    }
    ```

### Notifications API

#### Register Device for Notifications
- **Endpoint URL**: `/notifications/register`
- **Method**: POST
- **Request Parameters**: 
    - `token`: Device token
    - `userId`: User ID
    - `deviceId`: Device ID
    - `deviceType`: Device Type
- **Example Request**:
    ```
    POST /notifications/register
    {
        "token": "device_token",
        "userId": "user_id",
        "deviceId": "device_id",
        "deviceType": "ios"
    }
    ```
- **Example Response**:
    ```json
    {
        "message": "device registered"
    }
    ```

#### Send iOS Push Notification
- **Endpoint URL**: `/notifications/ios`
- **Method**: POST
- **Request Parameters**: None
- **Example Request**:
    ```
    POST /notifications/ios
    ```
- **Example Response**:
    ```json
    {
        "message": "push notification sent"
    }
    ```

### Orders API

#### Create Order
- **Endpoint URL**: `/orders`
- **Method**: POST
- **Request Parameters**: 
    - `customerId`: Customer ID
    - `storeId`: Store ID
    - `address`: Address
    - `phoneNumber`: Phone Number
    - `orderType`: Order Type
    - `orderItems`: Order items
- **Example Request**:
    ```
    POST /orders
    {
        "customerId": "customer_id_1",
        "storeId": "store_id_1",
        "address": "123 street",
        "phoneNumber": "1234567890",
        "orderType": "DINE_IN",
        "orderItems": [
            {
                "productId": "product_id_1",
                "quantity": 2,
                "modifierItemIds": [1, 2]
            }
        ]
    }
    ```
- **Example Response**:
    ```json
    {
        "data": {
            "id": "new_order_id",
            ...
        },
        "msg": "Created order successfully."
    }
    ```

#### Re-Order by Order ID
- **Endpoint URL**: `/orders/:orderId`
- **Method**: POST
- **Request Parameters**: 
    - `orderId`: Order ID
- **Example Request**:
    ```
    POST /orders/order_id_1
    ```
- **Example Response**:
    ```json
    {
        "data": {
            "id": "reordered_order_id",
            ...
        },
        "msg": "Re-created order successfully."
    }
    ```

#### Get All Orders
- **Endpoint URL**: `/orders`
- **Method**: GET
- **Request Parameters**: None
- **Example Request**:
    ```
    GET /orders
    ```
- **Example Response**:
    ```json
    {
        "data": [
            {
                "id": "order_id_1",
                "status": "PENDING",
                ...
            },
            ...
        ]
    }
    ```

#### Get Order by ID
- **Endpoint URL**: `/orders/:id`
- **Method**: GET
- **Request Parameters**: 
    - `id`: Order ID
- **Example Request**:
    ```
    GET /orders/order_id_1
    ```
- **Example Response**:
    ```json
    {
        "data": {
            "id": "order_id_1",
            "status": "PENDING",
            ...
        }
    }
    ```

## Class and Function Documentation

### CategoriesRouter
- **Description**: Router for managing categories.
- **Class/Function Name**: `router`
- **Parameters and Return Types**:
    - No parameters for the router itself.
- **Sample Code Snippet**:
    ```typescript
    import { Request, Response, Router } from "express";
    import prisma from "../../db/prisma";

    const router = Router();

    import { StatusCodes } from "http-status-codes";
    import { NotFoundError, BadRequestError } from "../../errors";

    // Define routes here...

    export { router as categoriesRouter };
    ```

### NotificationsRouter
- **Description**: Router for handling notification-related endpoints.
- **Class/Function Name**: `router`
- **Parameters and Return Types**:
    - No parameters for the router itself.
- **Sample Code Snippet**:
    ```typescript
    import { Router } from "express";
    const router = Router();

    const sendIosPushNotification = async () => {
        // iOS Push Notification logic...
    };

    router.post("/register", async (req: Request, res: Response) => {
        // Registration logic...
    });

    router.post("/ios", async (req: Request, res: Response) => {
        // iOS Push Notification logic...
    });

    export { router as notificationsRouter };
    ```

### OrdersRouter
- **Description**: Router for managing orders.
- **Class/Function Name**: `router`
- **Parameters and Return Types**:
    - No parameters for the router itself.
- **Sample Code Snippet**:
    ```typescript
    import { Request, Response, Router } from "express";
    import prisma from "../../db/prisma";

    const router = Router();

    interface OrderItemInput {
        productId: string;
        quantity: number;
        modifierItemIds: number[];
    }

    router.post("/", async (req: Request, res: Response) => {
        // Create order logic...
    });

    router.post("/:orderId", async (req: Request, res: Response) => {
        // Re-order logic...
    });

    router.get("/", async (req: Request, res: Response) => {
        // Get orders logic...
    });

    export { router as ordersRouter };
    ```

## Configuration and Deployment

### Configuration File Explanation

#### `.env`
- **Description**: Environment variables needed for the project.
- **Content**:
    ```
    APN_KEY_ID=your_apn_key_id
    APN_TEAM_ID=your_apn_team_id
    ```

#### `prisma/schema.prisma`
- **Description**: Prisma schema file defining the database schema.
- **Content**:
    ```prisma
    datasource db {
        provider = "postgresql"
        url      = env("DATABASE_URL")
    }

    generator client {
        provider = "prisma-client-js"
    }

    model Category {
        id       String   @id @default(uuid())
        name     String
        products Product[]
    }

    // Other models...
    ```

### Deployment Instructions
- **Step-by-Step**:
    1. **Install Dependencies**: 
        ```bash
        npm install
        ```
    2. **Setup Environment Variables**: Create a `.env` file and add the required environment variables.
    3. **Run Migrations**: 
        ```bash
        npx prisma migrate dev
        ```
    4. **Start the Server**: 
        ```bash
        npm run dev
        ```
    5. **Deploy to Production**:
        - Ensure the environment variables are set.
        - Build the project: 
            ```bash
            npm run build
            ```
        - Start the server:
            ```bash
            npm start
            ```

## FAQs and Troubleshooting

### Common Issues and Their Resolutions

#### Issue: "Category not found" when fetching products by category ID
- **Resolution**: Ensure the category ID provided in the request URL is correct and exists in the database.

#### Issue: "Unable to send push notification"
- **Resolution**: Ensure the APN credentials in the environment variables are correct and the APN key file path is correct.

#### Issue: "Invalid item data" when creating or updating items
- **Resolution**: Validate that all required fields (`variantName`, `sku`, `price`, `quantity`) are provided and correctly formatted in the request body.

### Additional Tips
- Always check the server logs for detailed error messages.
- Ensure your database is correctly configured and accessible by the application.

This documentation provides a comprehensive overview of the Incoming Orders API Server, detailing the available API endpoints, their usage, and the necessary configurations. For further queries or issues, refer to the provided troubleshooting section.