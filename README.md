<div id="top">

<!-- HEADER STYLE: CLASSIC -->
<div align="center">

# GO-BID

<em></em>

<!-- BADGES -->
<img src="https://img.shields.io/github/license/ecbDeveloper/go-bid?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
<img src="https://img.shields.io/github/last-commit/ecbDeveloper/go-bid?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/ecbDeveloper/go-bid?style=default&color=0080ff" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/ecbDeveloper/go-bid?style=default&color=0080ff" alt="repo-language-count">

<!-- default option, no dependency badges. -->


<!-- default option, no dependency badges. -->

## Overview
A real-time auction API built in **Golang**, leveraging **WebSocket** for instant bid communication between participants and **channels** to handle concurrency safely and efficiently.

### Features

- Create and list auctions.
- Join auctions via WebSocket.
- Place bids in real-time.
- Receive automatic broadcast updates of new bids.
- Enforced business rules (e.g., product owner cannot place bids).

### Technologies

- [Golang](https://golang.org/)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [Go-Chi](https://github.com/go-chi/chi) (HTTP router)
- Native Go channels
- UUIDs via [google/uuid](https://github.com/google/uuid)

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
    - [Project Index](#project-index)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Testing](#testing)

## Project Structure

```sh
└── go-bid/
    ├── cmd
    │   ├── go-bid
    │   └── terndotenv
    ├── compose.yml
    ├── go.mod
    ├── go.sum
    └── internal
        ├── api
        ├── db
        ├── services
        ├── shared
        └── usecase
```

### Project Index

<details open>
	<summary><b><code>GO-BID/</code></b></summary>
	<!-- cmd Submodule -->
	<details>
		<summary><b>cmd</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ cmd</b></code>
			<!-- terndotenv Submodule -->
			<details>
				<summary><b>terndotenv</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ cmd.terndotenv</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/cmd/terndotenv/main.go'>main.go</a></b></td>
							<td style='padding: 8px;'>- The <code>main.go</code> file executes database migrations<br>- It loads environment variables, then uses the <code>tern</code> command-line tool to apply migrations located in the <code>./internal/db/migrations/</code> directory, guided by a configuration file<br>- Successful execution confirms migration completion; failure results in an error message and program termination<br>- The process ensures the database schema is up-to-date.</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- go-bid Submodule -->
			<details>
				<summary><b>go-bid</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ cmd.go-bid</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/cmd/go-bid/main.go'>main.go</a></b></td>
							<td style='padding: 8px;'>- The <code>main.go</code> file bootstraps the Go-Bid application<br>- It initializes database connections, session management using PostgreSQL, and instantiates core application services (user, product, bid)<br>- The file then configures and starts a Chi router-based HTTP server, binding application routes to handle incoming requests, including websocket connections for real-time auction features.</td>
						</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<!-- internal Submodule -->
	<details>
		<summary><b>internal</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ internal</b></code>
			<!-- shared Submodule -->
			<details>
				<summary><b>shared</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.shared</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/shared/validator.go'>validator.go</a></b></td>
							<td style='padding: 8px;'>- Validator.go` provides a reusable validation layer for the application<br>- It defines an interface and associated functions for validating various data types, including strings (checking for blank values, email format, and character length)<br>- The package centralizes validation logic, promoting consistency and maintainability across the project by providing a common error handling mechanism<br>- This improves code organization and reduces redundancy.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/shared/json_utils.go'>json_utils.go</a></b></td>
							<td style='padding: 8px;'>- Json_utils.go` provides reusable functions for encoding and decoding JSON data within the application<br>- It offers functions to encode arbitrary data structures into JSON responses, and to decode JSON requests, optionally validating the data against a defined schema<br>- These utilities streamline JSON handling across the application, ensuring consistent data serialization and input validation.</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- services Submodule -->
			<details>
				<summary><b>services</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.services</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/services/auction_services.go'>auction_services.go</a></b></td>
							<td style='padding: 8px;'>- The <code>auction_services.go</code> file implements a real-time auction system using websockets<br>- It manages auction rooms, client connections, and bid processing<br>- The system handles bid placement, broadcasts updates to all participants, and manages client registration and disconnections, ensuring a synchronized auction experience for all users<br>- A core component is the <code>AuctionRoom</code> struct which facilitates communication and bid processing within each auction.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/services/users_service.go'>users_service.go</a></b></td>
							<td style='padding: 8px;'>- The <code>users_service.go</code> file implements a user service within a larger Go application<br>- It provides functionalities for user account creation, handling potential duplicate email or username errors, and user authentication, securely comparing provided passwords against stored hashes<br>- The service interacts with a database using SQL queries to manage user data.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/services/products_service.go'>products_service.go</a></b></td>
							<td style='padding: 8px;'>- Products_service.go<code> provides a service layer for product-related operations within the </code>go-bid<code> application<br>- It offers functionalities to create, retrieve, and list products, interacting with a database via the </code>sqlc` package<br>- The service handles database errors gracefully, returning custom error messages where appropriate, and uses pagination for efficient data retrieval.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/services/bids_service.go'>bids_service.go</a></b></td>
							<td style='padding: 8px;'>- The <code>bids_service.go</code> file implements a service layer for managing bids within an auction system<br>- It handles placing new bids, validating bid amounts against existing highest bids and base prices, and retrieving bid history for specific products<br>- The service interacts with a database via a defined query interface, ensuring data persistence and retrieval<br>- Error handling is incorporated to manage invalid bids and database access issues.</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- api Submodule -->
			<details>
				<summary><b>api</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.api</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/api.go'>api.go</a></b></td>
							<td style='padding: 8px;'>- Api.go` defines the API application structure, acting as the central hub for handling requests<br>- It integrates various services (user, product, bid, and auction lobby) and session management, using Chi router for request routing and Gorilla WebSockets for real-time communication<br>- The API facilitates interactions between clients and the applications core functionalities.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/products_handler.go'>products_handler.go</a></b></td>
							<td style='padding: 8px;'>- Products_handler.go<code> implements HTTP handlers for product-related API endpoints within the </code>go-bid` application<br>- It handles creating new product auctions, launching asynchronous auction processes, and retrieving paginated lists of products<br>- The handlers utilize services and use cases to manage business logic, returning JSON responses indicating success or failure<br>- Error handling ensures appropriate HTTP status codes are returned to clients.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/auth.go'>auth.go</a></b></td>
							<td style='padding: 8px;'>- Auth.go<code> provides authentication middleware and CSRF token handling for the </code>go-bid` API<br>- The middleware verifies user sessions, rejecting unauthorized requests<br>- A dedicated handler generates and returns CSRF tokens, crucial for protecting against cross-site request forgery attacks<br>- These functions ensure secure access control within the applications API layer.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/auction_handlers.go'>auction_handlers.go</a></b></td>
							<td style='padding: 8px;'>- Auction_handlers.go<code> manages user subscription to auction rooms<br>- It handles incoming requests, validates user and product IDs, and prevents sellers from bidding on their own products<br>- Upon successful validation, it upgrades the connection to a websocket, registers the user in the appropriate auction room, and manages the clients read and write event loops via the </code>AuctionLobby` and associated services.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/users_handlers.go'>users_handlers.go</a></b></td>
							<td style='padding: 8px;'>- Users_handlers.go<code> provides HTTP handlers for user authentication and account management within the </code>go-bid<code> application<br>- It handles user signup, login, and logout requests, interacting with the </code>UserService` to manage user data and sessions<br>- Error handling ensures appropriate HTTP status codes and informative error messages are returned to clients<br>- The handlers utilize JSON for data encoding and decoding.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/routes.go'>routes.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/api/bids_handlers.go'>bids_handlers.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- db Submodule -->
			<details>
				<summary><b>db</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.db</b></code>
					<!-- queries Submodule -->
					<details>
						<summary><b>queries</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.db.queries</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/queries/products.sql'>products.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/queries/bids.sql'>bids.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/queries/users.sql'>users.sql</a></b></td>
									<td style='padding: 8px;'>- The <code>users.sql</code> file defines SQL queries for managing user data within the database<br>- It provides functions to create new users, retrieve user information by ID, and retrieve user information by email address<br>- These queries are integral to the applications user authentication and data access layer, supporting core user management functionalities.</td>
								</tr>
							</table>
						</blockquote>
					</details>
					<!-- sqlc Submodule -->
					<details>
						<summary><b>sqlc</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.db.sqlc</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/sqlc/products.sql.go'>products.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/sqlc/sqlc.yaml'>sqlc.yaml</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/sqlc/db.go'>db.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/sqlc/bids.sql.go'>bids.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/sqlc/users.sql.go'>users.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/sqlc/models.go'>models.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
					<!-- migrations Submodule -->
					<details>
						<summary><b>migrations</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.db.migrations</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/migrations/002_create_sessions_table.sql'>002_create_sessions_table.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/migrations/003_create_products_table.sql'>003_create_products_table.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/migrations/004_create_bids_table.sql'>004_create_bids_table.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/migrations/001_create_users_table.sql'>001_create_users_table.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/db/migrations/tern.conf'>tern.conf</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
			<!-- usecase Submodule -->
			<details>
				<summary><b>usecase</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.usecase</b></code>
					<!-- user Submodule -->
					<details>
						<summary><b>user</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.usecase.user</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/usecase/user/login_user.go'>login_user.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/usecase/user/create_user.go'>create_user.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
					<!-- product Submodule -->
					<details>
						<summary><b>product</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.usecase.product</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-bid/blob/master/internal/usecase/product/create_product.go'>create_product.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** Go
- **Package Manager:** Go modules

### Installation

Build go-bid from the source and intsall dependencies:

1. **Clone the repository:**

    ```sh
    ❯ git clone https://github.com/ecbDeveloper/go-bid
    ```

2. **Navigate to the project directory:**

    ```sh
    ❯ cd go-bid
    ```

3. **Install the dependencies:**

	```sh
	❯ go mod tidy
	❯ go build
	```

### Usage

Run the project with:

**Using [go modules](https://golang.org/):**
```sh
go run {entrypoint}
```

### Testing

Go-bid uses the {__test_framework__} test framework. Run the test suite with:

**Using [go modules](https://golang.org/):**
```sh
go test ./...
```

<div align="center">

[![][back-to-top]](#top)

</div>


[back-to-top]: https://img.shields.io/badge/-BACK_TO_TOP-151515?style=flat-square

