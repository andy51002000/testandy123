# Use Node.js as the base image
FROM node:21 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Vue.js application
RUN npm run build

# Stage 2: Use nginx as the base image for serving static files
FROM nginx:alpine

# Copy the built app from the previous stage to the nginx directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 8080

# Start nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]
