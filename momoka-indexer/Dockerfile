# Use the official Node.js 18 image as a parent image
FROM node:18

# Set the working directory to /app
WORKDIR /app

# Copy the package.json, yarn.lock, and prisma directories to the container
COPY package.json yarn.lock ./
COPY prisma ./prisma/

# Install dependencies
RUN yarn install

# Copy the rest of the application files to the container
COPY . .

# Set the environment variable for the Prisma database URL
ENV DATABASE_URL=""

# Start the application
CMD ["yarn", "start"]
