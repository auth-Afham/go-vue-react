# Go-Vue-React Stack

This project is a full-stack web application that integrates:
- **Go** (Backend)
- **React** (Frontend 1)
- **Vue** (Frontend 2)

## Project Structure
```
/go-vue-react
  ├── go-vue-backend      # Go backend service
  ├── react-frontend      # React frontend application
  ├── vue-frontend        # Vue frontend application
```

## Backend - Go
### Setup & Run
1. Navigate to the backend folder:
   ```sh
   cd go-vue-backend
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the backend server:
   ```sh
   go run main.go
   ```

## Frontend - React
### Setup & Run
1. Navigate to the React frontend folder:
   ```sh
   cd react-frontend
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Run the React development server:
   ```sh
   npm start
   ```

## Frontend - Vue
### Setup & Run
1. Navigate to the Vue frontend folder:
   ```sh
   cd vue-frontend
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Run the Vue development server:
   ```sh
   npm run serve
   ```

## API Routes (Backend)
- `GET /api/health` - Check API status
- `POST /api/data` - Submit data

## Environment Variables
Create a `.env` file for each service:
- **Backend (`go-vue-backend/.env`)**
- **React (`react-frontend/.env`)**
- **Vue (`vue-frontend/.env`)**

## Contributing
Feel free to fork the repository, create a new branch, and submit a pull request.

## License
MIT License

