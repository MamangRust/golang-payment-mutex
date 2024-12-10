# Payment Gateway Project

Welcome to the **Payment Gateway Project**! This system simulates a payment gateway handling operations like card transactions, top-ups, withdrawals, and transfers. The project involves multiple entities such as users, cards, merchants, and financial transactions. To ensure data consistency and avoid concurrency issues, Go's `sync.Mutex` is utilized.

## Project Overview

### Key Structs

#### Card
Represents a credit/debit card with the following attributes:
- `CardID`: Unique identifier for the card.
- `UserID`: Identifier of the card owner.
- `CardNumber`: The card number (e.g., `1234-5678-9876-5432`).
- `CardType`: Type of card (e.g., `Visa`, `MasterCard`).
- `ExpireDate`: Expiration date of the card.
- `CVV`: Card verification value.
- `CardProvider`: The provider of the card (e.g., `Visa`, `MasterCard`).

#### Merchant
Represents a merchant in the payment gateway:
- `MerchantID`: Unique identifier for the merchant.
- `Name`: Name of the merchant.
- `ApiKey`: API key associated with the merchant.
- `UserID`: Identifier of the merchant's owner.
- `Status`: Status of the merchant (e.g., `Active`, `Inactive`).

#### Saldo
Represents the balance details of a card:
- `SaldoID`: Unique identifier for the balance entry.
- `CardNumber`: The associated card number.
- `TotalBalance`: Current balance on the card.
- `WithdrawAmount`: The amount withdrawn.
- `WithdrawTime`: Timestamp when the withdrawal occurred.

#### Topup
Represents a top-up operation:
- `TopupID`: Unique identifier for the top-up.
- `CardNumber`: The card number being topped up.
- `TopupNo`: A unique top-up number (e.g., `TOPUP0001`).
- `TopupAmount`: Amount added to the card.
- `TopupMethod`: Method used for the top-up (e.g., `Bank Transfer`, `Credit Card`).
- `TopupTime`: Timestamp when the top-up was performed.

#### Transaction
Represents a transaction record:
- `TransactionID`: Unique identifier for the transaction.
- `CardNumber`: The card number used in the transaction.
- `Amount`: Amount of the transaction.
- `PaymentMethod`: The payment method (e.g., `Credit Card`, `Bank Transfer`).
- `MerchantID`: The merchant ID where the transaction occurred.
- `TransactionTime`: Timestamp when the transaction was made.

#### Transfer
Represents a fund transfer between cards:
- `TransferID`: Unique identifier for the transfer.
- `TransferFrom`: Card number from which the funds are transferred.
- `TransferTo`: Card number to which the funds are transferred.
- `TransferAmount`: The amount transferred.
- `TransferTime`: Timestamp of the transfer.

#### User
Represents a user in the system:
- `UserID`: Unique user identifier.
- `FirstName`: User's first name.
- `LastName`: User's last name.
- `Email`: User's email address.
- `Password`: User's password for authentication.

#### Withdraw
Represents a withdrawal operation:
- `WithdrawID`: Unique identifier for the withdrawal.
- `CardNumber`: The card number from which the withdrawal is made.
- `WithdrawAmount`: The amount withdrawn.
- `WithdrawTime`: Timestamp of the withdrawal.


## Usage

To use the Payment Gateway Project, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/MamangRust/golang-payment-mutex.git
    ```

2. Navigate to the project directory:
    ```bash
    cd payment-gateway
    ```
3. Build the project:
    ```bash
    go build
    ```

4. Run the project:
   ```bash
   ./payment-gateway
   ```
---------------------