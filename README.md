# E-commerce Microservices Application

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/ecommerce.git
    ```

2. Navigate to each service directory and build Docker images:
    ```sh
    cd auth-service
    docker build -t your-dockerhub-username/auth-service .

    cd ../product-service
    docker build -t your-dockerhub-username/product-service .

    cd ../order-service
    docker build -t your-dockerhub-username/order-service .
    ```

3. Push the Docker images to your Docker Hub repository:
    ```sh
    docker push your-dockerhub-username/auth-service
    docker push your-dockerhub-username/product-service
    docker push your-dockerhub-username/order-service
    ```

4. Apply the Kubernetes deployment:
    ```sh
    kubectl apply -f deployment.yaml
    ```

## Endpoints

### Authentication Service
- `POST /register`: Register a new user.
- `POST /login`: Login a user.

### Product Service
- `POST /products`: Create a new product.
- `GET /products/{id}`: Get product details.
- `PUT /products/{id}`: Update product details.
- `DELETE /products/{id}`: Delete a product.

### Order Service
- `POST /orders`: Create a new order.
- `GET /orders/{id}`: Get order details.
